package main

import (
	"os"
	"os/exec"

	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	configDir, exists := os.LookupEnv("CONFIG_DIR")
	if !exists {
		log.Info("CONFIG_DIR not set correctly. Defaulting to example configs.",
			zap.String("CONFIG_DIR", "examples/configs"))
		configDir = "examples/configs"
	}

	argsSlice := []string{"playbooks/deploy_sterling_data_exchange.yml"}
	argsSlice = append(argsSlice, "-e")
	argsSlice = append(argsSlice, "config_dir="+configDir)

	entitlementKey, exists := os.LookupEnv("ENTITLEMENT_KEY")
	if !exists {
		log.Info("ENTITLEMENT_KEY not set correctly. Installation will continue but fail if you do not provide.",
			zap.String("ENTITLEMENT_KEY", entitlementKey))
	} else {
		argsSlice = append(argsSlice, "-e")
		argsSlice = append(argsSlice, "entitlement_key="+entitlementKey)
	}

	command := exec.Command("ansible-playbook", argsSlice...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.Run()
	// output, err := command.CombinedOutput()
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// fmt.Println(string(output))

}
