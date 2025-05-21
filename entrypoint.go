package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	configDir, exists := os.LookupEnv("CONFIG_DIR")
	if !exists {
		log.Fatal("CONFIG_DIR not set correctly.")
	}

	argsSlice := []string{"playbooks/deploy_sterling_b2bi.yml"}
	argsSlice = append(argsSlice, "-e")
	argsSlice = append(argsSlice, "config_dir="+configDir)

	entitlementKey, exists := os.LookupEnv("ENTITLEMENT_KEY")
	if !exists {
		log.Fatal("ENTITLEMENT_KEY not set correctly.")
	}

	argsSlice = append(argsSlice, "-e")
	argsSlice = append(argsSlice, "entitlement_key="+entitlementKey)

	command := exec.Command("ansible-playbook", argsSlice...)

	fmt.Println(command)

	fmt.Println(os.Args)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.Run()

}
