package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Define application lists
var aptApps = []string{
	"adb", "fastboot", "curl", "vim", "git", "net-tools", "bat", "subversion", "vlc", "build-essential",
	"repo", "meld", "ckermit", "openssh-server", "openjdk-11-jdk", "openjdk-17-jdk", "openjdk-21-jdk",
}
var snapApps = []string{
	"android-studio", "brave", "code", "gh", "go", "intellij-idea", "kubectl", "microk8s", "pycharm-community", "thunderbird", "tmux",
}
var snapClassicApps = map[string]bool{
	"code":              true, // VS Code
	"intellij-idea":     true, // IntelliJ IDEA
	"go":                true, // Go
	"android-studio":    true, // Android Studio
	"microk8s":          true, // MicroK8s
	"thunderbird":       true, // Thunderbird
	"pycharm-community": true, // Pycharm
	"kubectl":           true, // kubectl
	"tmux":              true, // tmux
}
var shellApps = []string{
	"docker", "nvim", "fzf", "eza", "nvm",
}

func main() {
	fmt.Println("Welcome to the App Installer!")

	// Allow the user to select from the lists
	selectedApt := multiSelect("Select APT packages to install:", aptApps)
	selectedSnap := multiSelect("Select SNAP packages to install:", snapApps)
	selectedShell := multiSelect("Select shell scripts to execute:", shellApps)

	// Show summary and confirm installation
	if !confirmSelection(selectedApt, selectedSnap, selectedShell) {
		fmt.Println("Installation process canceled.")
		return
	}

	// Install apps
	if len(selectedApt) > 0 {
		installApps("apt", selectedApt)
	} else {
		fmt.Println("No APT packages selected. Skipping.")
	}

	if len(selectedSnap) > 0 {
		installApps("snap", selectedSnap)
	} else {
		fmt.Println("No SNAP packages selected. Skipping.")
	}

	if len(selectedShell) > 0 {
		runShellScripts(selectedShell)
	} else {
		fmt.Println("No shell scripts selected. Skipping.")
	}

	fmt.Println("\nInstallation completed!")
}

// Multi-select prompt using survey
func multiSelect(prompt string, options []string) []string {
	var selected []string
	surveyPrompt := &survey.MultiSelect{
		Message: prompt,
		Options: options,
	}
	err := survey.AskOne(surveyPrompt, &selected)
	if err != nil {
		fmt.Printf("Error in selection: %v\n", err)
	}
	return selected
}

// Confirm selection and show summary
func confirmSelection(aptApps, snapApps, shellApps []string) bool {
	fmt.Println("\nSummary of selected items:")
	if len(aptApps) > 0 {
		fmt.Printf("APT packages: %s\n", strings.Join(aptApps, ", "))
	} else {
		fmt.Println("APT packages: None")
	}
	if len(snapApps) > 0 {
		fmt.Printf("SNAP packages: %s\n", strings.Join(snapApps, ", "))
	} else {
		fmt.Println("SNAP packages: None")
	}
	if len(shellApps) > 0 {
		fmt.Printf("Shell scripts: %s\n", strings.Join(shellApps, ", "))
	} else {
		fmt.Println("Shell scripts: None")
	}

	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: "Do you want to proceed with the installation?",
		Default: false,
	}
	err := survey.AskOne(confirmPrompt, &confirm)
	if err != nil {
		fmt.Printf("Error in confirmation: %v\n", err)
		return false
	}
	return confirm
}

// Install apps via apt or snap
func installApps(manager string, apps []string) {
	for _, app := range apps {
		fmt.Printf("Installing %s via %s...\n", app, manager)
		var cmd *exec.Cmd
		if manager == "snap" && snapClassicApps[app] {
			// Use --classic for Snap apps that need it
			cmd = exec.Command("sudo", manager, "install", app, "--classic")
		} else if manager == "snap" {
			// Regular Snap installation
			cmd = exec.Command("sudo", manager, "install", app)
		} else {
			// APT installation
			cmd = exec.Command("sudo", manager, "install", app, "-y")
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to install %s: %v\n", app, err)
		} else {
			fmt.Printf("%s installed successfully.\n", app)
		}
	}
}

func runShellScripts(scripts []string) {
	installer := &ShellAppInstaller{}
	for _, script := range scripts {
		fmt.Printf("Running shell script for %s...\n", script)
		switch script {
		case "docker":
			installer.InstallDocker()
		case "nvim":
			installer.InstallNeovim()
		case "fzf":
			installer.InstallFzf()
		case "eza":
			installer.InstallEza()
		case "nvm":
			installer.InstallNvm()
		default:
			fmt.Printf("No installation script defined for %s.\n", script)
		}
	}
}
