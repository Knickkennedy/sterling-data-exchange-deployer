package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	value, exists := os.LookupEnv("CONFIG_DIR")
	if !exists {
		log.Fatal("CONFIG_DIR not set correctly.")
	}
	command := exec.Command("ansible-playbook", "playbooks/deploy_sterling_b2bi.yml", "--extra-vars", "config_dir="+value)

	fmt.Println(command)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.Run()

}
