package utils

import (
	"log"
	"os/exec"
)

func ExecuteMain() {
	//ExecuteCmd("cp","demo.jar","d://tmp")
	ExecuteCmd("ps","aux")
}

func ExecuteCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("pid:",cmd.Process.Pid)
	log.Println(string(stdout))
}
