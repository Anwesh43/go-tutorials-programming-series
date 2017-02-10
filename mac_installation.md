To install brew
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/in

To install go with brew
brew install go --cross-compile-common

Setting up the path
~/.bash_profile
export GOPATH="/Users/anweshmishra/go_libraries"
export PATH=$PATH:$GOPATH/bin

source ~/.bash_profile
