package genericHardware

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const powerSupplyPath = "/sys/class/power_supply"

func discoverBattery() {
	folder, err := os.Open(powerSupplyPath)
	if err != nil {
		fmt.Println("Error opening power_supply folder:", err)
		return
	}
	defer folder.Close()

	files, err := folder.Readdirnames(0)
	if err != nil {
		fmt.Println("Error reading power_supply folder:", err)
		return
	}

	for _, file := range files {
		devicePath := filepath.Join(powerSupplyPath, file)

		powerType, err := readPowerSupplyFile(devicePath, "type")
		if err != nil {
			fmt.Println("Error reading power supply type:", err)
			continue
		}

		if powerType == "Battery" {
			fmt.Println("Battery found:", file)
		}
	}
}

// Helper function to read a file from the power supply device path and return its content as a string.
func readPowerSupplyFile(devicePath string, name string) (string, error) {
	data, err := os.ReadFile(filepath.Join(devicePath, name))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
