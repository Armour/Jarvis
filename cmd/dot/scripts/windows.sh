# Install scoop
set-executionpolicy unrestricted -s cu -Force
iex (new-object net.webclient).downloadstring('https://get.scoop.sh')

# Install basic tools
scoop install sudo 7zip coreutils curl git grep openssh sed wget vim grep concfg pshazz nodejs yarn less go

# Use material theme
concfg import material

# Install extra tools
scoop bucket add extras 
scoop bucket add nerd-fonts
scoop install vscode anaconda3 hyper kitematic now-cli p4merge postman shadowsocks slack telegram gcloud
sudo scoop install FiraCode-NF SourceCodePro-NF
