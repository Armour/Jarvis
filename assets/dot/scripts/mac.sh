# Install brew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# Init brew repo
brew tap buo/cask-upgrade
brew tap heroku/brew
brew tap homebrew/core
brew tap homebrew/cask
brew tap homebrew/cask-fonts
brew tap homebrew/cask-versions
brew tap homebrew/services
brew update
brew upgrade -y

# Install brew stuff
brew install archey \
            autojump \
            asciinema \
            ccat \
            cmake \
            colordiff \
            dep \
            docker-clean \
            ffmpeg \
            git \
            go \
            heroku \
            htop \
            libpng \
            mono \
            nginx \
            node \
            postgresql \
            qt \
            redis \
            sqlmap \
            taglib \
            tree \
            wget \
            yarn

# Install brew cask stuff
brew cask install adobe-creative-cloud \
                  alfred \
                  anaconda \
                  android-studio \
                  appcode \
                  appzapper \
                  betterzip \
                  burp-suite \
                  clion \
                  dash \
                  db-browser-for-sqlite \
                  discord \
                  djay-pro \
                  docker-edge \
                  dotnet \
                  dotnet-sdk \
                  font-fira-code \
                  font-source-code-pro-for-powerline \
                  gitter \
                  goland \
                  google-chrome \
                  google-cloud-sdk \
                  iina \
                  intellij-idea \
                  iterm2 \
                  java8 \
                  jupyter-notebook-ql \
                  kap \
                  kitematic \
                  league-of-legends \
                  lingon-x \
                  lyricsx \
                  mactex \
                  minecraft \
                  neteasemusic \
                  now \
                  p4v \
                  postico \
                  postman \
                  pycharm \
                  qladdict \
                  qlcolorcode \
                  qlimagesize \
                  qlmarkdown \
                  qlstephen \
                  qlvideo \
                  qq \
                  qt-creator \
                  quicklook-json \
                  quicklookapk \
                  rowanj-gitx \
                  scroll-reverser \
                  shadowsocksx-ng \
                  shuttle \
                  simpholders \
                  sip \
                  sketch \
                  slack \
                  teamviewer \
                  telegram \
                  texmaker \
                  unity \
                  visual-studio-code \
                  vox \
                  webstorm \
                  wechat \
                  wireshark

# Cleanup
brew cleanup
brew cask cleanup

# Install oh-my-zsh
sh -c "$(wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"

# Download zsh plugins
git clone https://github.com/zdharma/fast-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/fast-syntax-highlighting
git clone https://github.com/zsh-users/zsh-autosuggestions.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

# Unblock neteasemusic :P
echo '158.199.142.239 music.163.com p1.music.126.net p2.music.126.net p3.music.126.net p4.music.126.net' | sudo tee -a /etc/hosts

# Enable hold and repeat for vscode on Mac
defaults write com.microsoft.VSCode ApplePressAndHoldEnabled -bool false

# Adobe creative cloud
echo "To complete the installation of Cask adobe-creative-cloud, you must also run the installer at '/usr/local/Caskroom/adobe-creative-cloud/latest/Creative Cloud Installer.app'"

# Sync dot files
go get -u github.com/armour/jarvis
echo "Please run 'jarvis dot sync' to sync your dot files."
