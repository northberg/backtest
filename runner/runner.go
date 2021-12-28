package runner

import (
	"bytes"
	"errors"
	bt "github.com/northberg/backtest"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func NewRunnable(instance *bt.BotInstance, version *bt.BotVersion, port int) *BotRunner {
	return &BotRunner{
		Version:    version,
		Owner:      instance,
		Cmd:        nil,
		Output:     &bytes.Buffer{},
		Port:       port,
		Terminated: false,
		Error:      nil,
	}
}

type BotMode = string

const (
	SimMode  = BotMode("sim")
	LiveMode = BotMode("live")
)

type BotRunner struct {
	Version    *bt.BotVersion
	Owner      *bt.BotInstance
	logger     *bt.BotRunLog
	Cmd        *exec.Cmd
	Output     *bytes.Buffer
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
	r.logger.Error = r.Error.Error()
	r.logger.Output = r.Output.String()
	r.logger.StopTime = time.Now().UTC().Unix()
	r.Cmd = nil
	r.Terminated = true
}

func (r *BotRunner) handleExit() {
	r.logger.Output = r.Output.String()
	r.logger.StopTime = time.Now().UTC().Unix()
	r.Cmd = nil
	r.Terminated = true
}

func (r *BotRunner) updateStatus(status string) {
	r.Owner.Status = status
	log.Printf("[%s] %s\n", r.Owner.Id, status)
}

func (r *BotRunner) Launch(dst string, mode BotMode) {

	if r.Port == 0 {
		log.Fatalln("invalid port")
	}

	// check if no process exists
	if r.Cmd != nil {
		log.Fatalf("[%s] bot was already running\n", r.Owner.Id)
		return
	}

	r.logger = r.Owner.NewRun()

	runName := r.Owner.Id
	botId := r.Version.Id
	runDir := dst + "/nbb-" + botId

	// pull or clone bot repository
	if _, err := os.Stat(runDir); os.IsNotExist(err) {

		// update status
		r.updateStatus("Cloning repository")

		// run clone
		cmd := exec.Command("git", "clone", "git@github.com:northberg/nbb-"+botId)
		cmd.Stderr = r.Output
		cmd.Stdout = r.Output
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
		cmd.Stderr = r.Output
		cmd.Stdout = r.Output
		cmd.Dir = runDir
		err = cmd.Run()

		// handle errors
		if err != nil {
			r.handleError(err, "Could not pull")
			return
		}

	}

	// Extract the latest commit hash
	cmd := exec.Command("git", "log", "-1", "--pretty=%H:=:=:%B")
	cmd.Dir = runDir
	out, err := cmd.Output()
	if err != nil {
		r.handleError(err, "Could not retrieve version")
		return
	}
	outParts := strings.Split(string(out), ":=:=:")
	r.Version.Commit = outParts[0]
	r.Version.Description = strings.ReplaceAll(outParts[1], "\n", " ")

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
	cmd = exec.Command("go", "build", "-o", "./build/bot", "./cmd/bot")
	cmd.Dir = runDir
	cmd.Stderr = r.Output
	cmd.Stdout = r.Output
	if err = cmd.Run(); err != nil {
		r.handleError(err, "Compilation failed")
		return
	}

	// setup process
	r.updateStatus("Starting")
	cmd = exec.Command("./bot", mode)
	cmd.Dir = runDir + "/build"
	cmd.Env = append(os.Environ(), "PORT="+strconv.Itoa(r.Port))
	cmd.Stdout = r.Output
	cmd.Stderr = r.Output

	// save
	r.Cmd = cmd

	// start process
	if err = cmd.Start(); err != nil {
		_ = cmd.Wait()
		r.handleError(err, "Could not start")
		return
	}

	// keep track of process termination state
	go func() {
		err = cmd.Wait()
		if err != nil {
			switch err.(type) {
			default:
				log.Printf("unexpected bot exit: %s\n", err.Error())
			case *exec.ExitError:
				break
			}
		}
		log.Printf("[%s] Finished\n", runName)
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

	log.Printf("[%s] Started\n", runName)
}

func (r *BotRunner) Heartbeat() bool {
	_, err := http.Get("http://localhost:" + strconv.Itoa(r.Port) + "/heartbeat")
	return err == nil
}

func (r *BotRunner) Kill() {
	if r.Terminated {
		return
	}
	for {
		err := r.Cmd.Process.Kill()
		if err != nil {
			log.Println(err)
		}
		for i := 0; i < 10; i++ {
			time.Sleep(150 * time.Millisecond)
			if r.Terminated {
				return
			}
		}
		log.Printf("[%s] Failed to kill\n", r.Owner.Id)
	}
}
