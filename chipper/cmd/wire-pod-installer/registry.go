package main

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func UpdateRegistry(is InstallSettings) {
	keyPath := `Software\Microsoft\Windows\CurrentVersion\Uninstall\wire-pod`
	appName := "wire-pod"
	displayIcon := filepath.Join(is.Where, `\chipper\icons\ico\pod256x256.ico`)
	displayVersion := "1.0.0"
	publisher := "github.com/kercre123"
	uninstallString := filepath.Join(is.Where, `\uninstall.exe`)
	installLocation := filepath.Join(is.Where, `\chipper\chipper.exe`)
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		k, _, err = registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating registry key: %v\n", err)
			return
		}
	}
	defer k.Close()

	err = k.SetStringValue("DisplayName", appName)
	if err != nil {
		fmt.Printf("Error setting DisplayName: %v\n", err)
		return
	}
	k.SetStringValue("DisplayIcon", displayIcon)
	k.SetStringValue("DisplayVersion", displayVersion)
	k.SetStringValue("Publisher", publisher)
	k.SetStringValue("UninstallString", uninstallString)
	k.SetStringValue("InstallLocation", installLocation)
	fmt.Println("Registry entries successfully created")
}