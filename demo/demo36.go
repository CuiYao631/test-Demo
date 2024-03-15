package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

func setAppID(appName, appID string) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\`+appName, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	if err := k.SetStringValue("AppID", appID); err != nil {
		return err
	}

	return nil
}

func main() {
	appName := "MyApp.exe"
	appID := "{00000000-0000-0000-0000-000000000001}"

	if err := setAppID(appName, appID); err != nil {
		fmt.Println("Failed to set AppID:", err)
		return
	}

	fmt.Println("AppID set successfully for", appName)
}
