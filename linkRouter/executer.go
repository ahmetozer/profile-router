package linkrouter

import (
	"log"
	"os"

	"github.com/ahmetozer/profile-router/config"
)

// Executer
// Execute Google Chrome
func Executer(Config *config.Settings, args *[]string) {

	//var cred = &syscall.Credential{UID, GUID, []uint32{}}
	//var sysproc = &syscall.SysProcAttr{Credential: cred, Noctty: true}
	var attr = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			nil,
			nil,
		},
	}
	process, err := os.StartProcess(Config.ChromePath, *args, &attr)
	if err == nil {

		// It is not clear from docs, but Realease actually detaches the process
		err = process.Release()
		if err != nil {
			log.Fatalf(err.Error())
		}

	} else {
		log.Fatalf(err.Error())
	}
	// cmd := exec.Command(Config.ChromePath, *args...)
	// err := cmd.Run()

	// var stdout, stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// return stdout.String(), stderr.String(), cmd.ProcessState.ExitCode()

	// Then app exists
	os.Exit(0)
}
