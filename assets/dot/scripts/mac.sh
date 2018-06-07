# Install brew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# Init brew repo
brew tap buo/cask-upgrade
brew tap heroku/brew
brew tap homebrew/core
brew tap homebrew/cask
brew tap homebrew/cask-fonts
brew tap homebrew/services
brew update
brew upgrade -y

# Install brew stuff
brew install git \
             go \
             autojump \
             docker-clean\
             heroku \
             libpng \
             mono \
             nginx \
             node \
             postgresql \
             redis \
             yarn \
             wget \
             sqlmap \
             taglib \
             tree \
             ccat \
             archey \
             colordiff \
             htop \
             ffmpeg

# Install brew cask stuff
brew cask install qlcolorcode \
                  qlstephen \
                  qlmarkdown \
                  quicklook-json \
                  qlimagesize \
                  qlvideo \
                  quicklookapk \
                  qladdict \
                  jupyter-notebook-ql \
                  betterzip \
                  neteasemusic \
                  teamviewer \
                  google-chrome \
                  slack \
                  lingon-x \
                  shadowsocksx-ng \
                  appzapper \
                  iterm2 \
                  scroll-reverser \
                  dash \
                  unity \
                  postico \
                  rowanj-gitx \
                  alfred \
                  docker \
                  visual-studio-code \
                  lyricsx \
                  qq \
                  kitematic \
                  db-browser-for-sqlite \
                  vox \
                  font-source-code-pro-for-powerline \
                  font-fira-code \
                  google-cloud-sdk \
                  sketch \
                  now \
                  kap \
                  postman \
                  anaconda \
                  telegram \
                  dotnet \
                  dotnet-sdk \
                  p4merge

# Install brew cask stuff (optional)
# brew cask install wireshark \
#                   texmaker \
#                   mactex \
#                   skype \
#                   spotify \
#                   appcode \
#                   clion \
#                   pycharm \
#                   webstorm \
#                   clip-studio-paint \
#                   burp-suite

# Cleanup
brew cleanup
brew cask cleanup

# Install oh-my-zsh
sh -c "$(wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"

# Download zsh plugins
git clone https://github.com/zdharma/fast-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/fast-syntax-highlighting
git clone https://github.com/zsh-users/zsh-autosuggestions.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

# Sync dot files
go get -u github.com/armour/jarvis
jarvis dot sync

# Unblock neteasemusic :P
echo '158.199.142.239 music.163.com p1.music.126.net p2.music.126.net p3.music.126.net p4.music.126.net' | sudo tee -a /etc/hosts

# Enable hold and repeat for vscode on mactex
defaults write com.microsoft.VSCode ApplePressAndHoldEnabled -bool false
