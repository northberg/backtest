package backtest

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func NewRunnable(instance *BotInstance, version *BotVersion, port int) *BotRunner {
	return &BotRunner{
		Version:    version,
		Owner:      instance,
		Cmd:        nil,
		Log:        &bytes.Buffer{},
		Port:       port,
		Terminated: false,
		Error:      nil,
	}
}

type BotRunner struct {
	Version    *BotVersion
	Owner      *BotInstance
	Cmd        *exec.Cmd
	Log        *bytes.Buffer
	Port       int
	Terminated bool
	Error      error
}

func (r *BotRunner) handleError(err error, status string) {
	r.updateStatus(status)
	if err == nil {
		err = errors.New(status)
	}
	r.Error = err
	r.Owner.Error = r.Error.Error()
	r.Owner.Log = r.Log.String()
	r.Cmd = nil
	r.Terminated = true
}

func (r *BotRunner) handleExit() {
	r.Owner.Log = r.Log.String()
	r.Cmd = nil
	r.Terminated = true
}

func (r *BotRunner) updateStatus(status string) {
	r.Owner.Status = status
	fmt.Printf("[%s] %s\n", r.Owner.Name, status)
}

func (r *BotRunner) Launch(dst string) {

	if r.Port == 0 {
		log.Fatalln("port is not set")
	}

	// check if no process exists
	if r.Cmd != nil {
		log.Fatalf("bot was already running: %s\n", r.Owner.Id)
		return
	}

	owner := r.Owner
	botId := r.Version.Id
	runDir := dst + "/nbb-" + botId

	// pull or clone bot repository
	if _, err := os.Stat(runDir); os.IsNotExist(err) {

		// update status
		r.updateStatus("Cloning repository")

		// run clone
		cmd := exec.Command("git", "clone", "git@github.com:northberg/nbb-"+botId)
		cmd.Stderr = r.Log
		cmd.Stdout = r.Log
		cmd.Dir = dst
		err = cmd.Run()

		// handle errors
		if err != nil {
			r.handleError(err, "Could not clone")
			return
		}

	} else {

		// update status
		r.updateStatus("Pulling repository")

		// run pull
		cmd := exec.Command("git", "pull")
		cmd.Stderr = r.Log
		cmd.Stdout = r.Log
		cmd.Dir = runDir
		err = cmd.Run()

		// handle errors
		if err != nil {
			r.handleError(err, "Could not pull")
			return
		}

	}

	// Extract the latest commit hash
	cmd := exec.Command("git", "log -1 --pretty=%H:=:=:%B")
	out, err := cmd.Output()
	if err != nil {
		r.handleError(err, "Could not retrieve version")
	}
	outParts := strings.Split(string(out), ":=:=:")
	r.Version.Commit = outParts[0]
	r.Version.Description = outParts[1]

	// remove existing build if present
	if _, err = os.Stat(runDir); !os.IsNotExist(err) {
		err = os.RemoveAll(runDir + "/build")
		if err != nil {
			r.handleError(err, "Could not remove old version")
			return
		}
	}

	// create directory to put executable in
	err = os.Mkdir(runDir+"/build", 0777)
	if err != nil {
		r.handleError(err, "Could not create directory")
		return
	}

	// build new version
	r.updateStatus("Compiling")
	var compileBuffer bytes.Buffer
	cmd = exec.Command("go", "build", "-o", "./build/bot", "./cmd/bot")
	cmd.Dir = runDir
	cmd.Stderr = &compileBuffer
	if err = cmd.Run(); err != nil {
		r.handleError(err, "Compilation failed")
		return
	}

	// setup process
	r.updateStatus("Starting")
	cmd = exec.Command("./bot")
	cmd.Dir = runDir + "/build"
	cmd.Env = append(os.Environ(), "PORT="+strconv.Itoa(r.Port))

	// link output buffers
	cmd.Stdout = r.Log
	cmd.Stderr = r.Log

	// save
	r.Cmd = cmd

	// start process
	if err = cmd.Start(); err != nil {
		_ = cmd.Wait()
		r.handleError(err, "Could not start")
		return
	}
	fmt.Printf("Starting \"%s\"\n", owner.Name)

	// keep track of process termination state
	go func() {
		err = cmd.Wait()
		if err != nil {
			switch err.(type) {
			default:
				fmt.Printf("unexpected bot exit: %s\n", err.Error())
			case *exec.ExitError:
				break
			}
		}
		fmt.Printf("\"%s\" exited", owner.Name)
		r.handleExit()
	}()

	// check if bot is online, try until status is returned
	r.updateStatus("Waiting for server")
	for {
		time.Sleep(time.Second)
		ok := r.Heartbeat()
		if ok {
			break
		}
		if r.Terminated {
			r.handleError(err, "Crash on startup")
			return
		}
	}

	fmt.Printf("\"%s\" is running\n", owner.Name)
}

func (r *BotRunner) Heartbeat() bool {
	_, err := http.Get("http://localhost:" + strconv.Itoa(r.Port) + "/heartbeat")
	fmt.Print(err)
	return err == nil
}

func (r *BotRunner) Kill() {
	if r.Terminated {
		return
	}
	for {
		err := r.Cmd.Process.Kill()
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < 10; i++ {
			time.Sleep(150 * time.Millisecond)
			if r.Terminated {
				return
			}
		}
		fmt.Printf("failed to kill process \"%s\", trying again...\n", r.Owner.Name)
	}
}
