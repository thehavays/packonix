package main

import (
	"fmt"
	"os"
	"os/exec"
)

// ShellAppInstaller manages shell app installations
type ShellAppInstaller struct{}

// Check if an application is installed
func (s *ShellAppInstaller) isInstalled(command string, args ...string) bool {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	return err == nil
}

// Install Docker
func (s *ShellAppInstaller) InstallDocker() {
	if s.isInstalled("docker", "--version") {
		fmt.Println("Docker is already installed.")
		return
	}

	fmt.Println("Installing Docker...")
	script := `
		sudo apt-get update
		sudo apt-get install -y ca-certificates curl
		sudo install -m 0755 -d /etc/apt/keyrings
		curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.asc
		echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list
		sudo apt-get update
		sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
	`
	runShellScript(script, "Docker")
}

// Install Neovim
func (s *ShellAppInstaller) InstallNeovim() {
	if s.isInstalled("nvim", "--version") {
		fmt.Println("Neovim is already installed.")
		return
	}

	fmt.Println("Installing Neovim...")
	script := `
		wget https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz
		tar xzvf nvim-linux64.tar.gz
		sudo mv nvim-linux64 /usr/local/nvim
		sudo ln -s /usr/local/nvim/bin/nvim /usr/bin/nvim
		rm nvim-linux64.tar.gz
		git clone https://github.com/NvChad/starter ~/.config/nvim
	`
	runShellScript(script, "Neovim")
}

// Install fzf
func (s *ShellAppInstaller) InstallFzf() {
	if s.isInstalled("fzf", "--version") {
		fmt.Println("fzf is already installed.")
		return
	}

	fmt.Println("Installing fzf...")
	script := `
		git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
		~/.fzf/install
	`
	runShellScript(script, "fzf")
}

// Install eza
func (s *ShellAppInstaller) InstallEza() {
	if s.isInstalled("eza", "--version") {
		fmt.Println("eza is already installed.")
		return
	}

	fmt.Println("Installing eza...")
	script := `
		sudo apt update
		sudo apt install -y gpg
		sudo mkdir -p /etc/apt/keyrings
		wget -qO- https://raw.githubusercontent.com/eza-community/eza/main/deb.asc | sudo gpg --dearmor -o /etc/apt/keyrings/gierens.gpg
		echo "deb [signed-by=/etc/apt/keyrings/gierens.gpg] http://deb.gierens.de stable main" | sudo tee /etc/apt/sources.list.d/gierens.list
		sudo chmod 644 /etc/apt/keyrings/gierens.gpg /etc/apt/sources.list.d/gierens.list
		sudo apt update
		sudo apt install -y eza
	`
	runShellScript(script, "eza")
}

// Install NVM
func (s *ShellAppInstaller) InstallNvm() {
	if s.isInstalled("nvm", "--version") {
		fmt.Println("nvm is already installed.")
		return
	}

	fmt.Println("Installing nvm...")
	script := `
		curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
		source ~/.bashrc
	`
	runShellScript(script, "nvm")
}

// Install Termius
func (s *ShellAppInstaller) InstallTermius() {
	fmt.Println("Installing Termius...")
	script := `
		wget https://www.termius.com/download/linux/Termius.deb
		sudo apt install ./Termius.deb  --fix-broken -y
		rm -rf Termius.deb
	`
	runShellScript(script, "termius")
}

// Utility function to execute shell scripts
func runShellScript(script, appName string) {
	cmd := exec.Command("bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to install %s: %v\n", appName, err)
	} else {
		fmt.Printf("%s installed successfully.\n", appName)
	}
}
