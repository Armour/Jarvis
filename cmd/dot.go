package cmd

import (
	"os/user"

	"github.com/spf13/cobra"
)

var (
	username string
	mac      bool
	linux    bool
	windows  bool
)

var dotCmd = &cobra.Command{
	Use:   "dot",
	Short: "Manage global dot files.",
	Long:  `Manage global dot files.`,
	Run: func(cmd *cobra.Command, args []string) {
		templatePath := "../dotfiles"
		requireMap := map[string]interface{}{
			"mac":     mac,
			"linux":   linux,
			"windows": windows,
		}
		replaceMap := map[string]interface{}{
			"mac":      mac,
			"linux":    linux,
			"windows":  windows,
			"username": username,
		}
		usr, err := user.Current()
		if err != nil {
			exitOnError(err)
		}
		generateFile(templatePath, usr.HomeDir, requireMap, replaceMap)
		if mac {
			// runCommand("/usr/bin/ruby", "-e", "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"")
			runCommand("brew", "tap", "buo/cask-upgrade", "caskroom/cask", "caskroom/fonts", "heroku/brew", "homebrew/core", "homebrew/python", "homebrew/science", "homebrew/services", "homebrew/versions")
			runCommand("brew", "update")
			runCommand("brew", "upgrade")
			// Install brew stuff
			runCommand("brew", "reinstall", "autojump", "docker-clean", "heroku", "libpng", "mono", "nginx", "node", "postgresql", "redis", "yarn", "wget", "python", "python3", "sqlmap", "taglib", "tree", "ccat", "archey", "colordiff", "stormssh", "htop", "ffmpeg")
			// Install brew cask stuff
			runCommand("brew", "cask", "reinstall", "qlcolorcode", "qlstephen", "qlmarkdown", "quicklook-json", "qlimagesize", "qlvideo", "quicklookapk", "qladdict", "betterzip", "neteasemusic", "teamviewer", "google-chrome", "slack", "lingon-x", "shadowsocksx-ng", "appzapper", "iterm2", "scroll-reverser", "dash", "unity", "postico", "rowanj-gitx", "alfred docker", "visual-studio-code", "alarm-clock", "lyricsx", "qq", "nylas-mail", "kitematic", "db-browser-for-sqlite", "jupyter-notebook-ql", "oversight", "vlc", "vox", "font-source-code-pro-for-powerline", "font-fira-code", "google-cloud-sdk", "usage", "p4merge", "sketch")
			// install brew cask stuff (optional)
			// brew cask install wireshark texmaker mactex skype spotify appcode clion pycharm webstorm clip-studio-paint telegram insomnia kaleidoscope now burp-suite
			runCommand("brew", "cleanup")
			runCommand("brew", "cask", "cleanup")
		}
		if linux {
			// TODO: linux command here
		}
		if windows {
			// TODO: windows command here
		}
		// Download zsh plugins
		runCommand("git", "clone", "https://github.com/zdharma/fast-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/fast-syntax-highlighting")
		runCommand("git", "clone", "https://github.com/zsh-users/zsh-autosuggestions.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions")
		// Run pip install
		runCommand("pip3", "install", "jupyter", "numpy", "pillow", "torch", "torchvision", "requests", "pylint", "opencv-python", "tensorflow", "you-get")
	},
}

func init() {
	rootCmd.AddCommand(dotCmd)

	dotCmd.Flags().StringVarP(&username, "username", "u", "armour", "The name for current user.")
	dotCmd.Flags().BoolVar(&mac, "mac", false, "The flag to enable mac environment.")
	dotCmd.Flags().BoolVar(&linux, "linux", false, "The flag to enable linux environment.")
	dotCmd.Flags().BoolVar(&windows, "windows", false, "The flag to enable windows environment.")
}
