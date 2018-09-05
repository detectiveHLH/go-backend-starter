<h1 align="center">go-backend-starter</h1>

<div align="center"> 
Based on gin, gorm RESTFUL style backend framework, using dep as the project's package manager
</div>

[简体中文](./README.md) | English

## Install golang
On a Mac, install golang and install it using the following command
```
brew install go
```
## Configuring environment variables
Configuring environment variables with following command
```
vim ~/.bash_profile
```
Add following code
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```
Then execute the following code to make the changes take effect
```
source ~/.bash_profile
```
## Build a local workspace
```
mkdir $HOME/go
```
In the above go directory, create three new directories, namely bin, pkg, src. Bin is the executable file after the compilation, pkg stores the third-party package, src stores the project source code.

Your project should be placed under $HOME/go/src.

## IDE
It is recommended to use JetBrains' [GoLand](https://www.jetbrains.com/go/?fromMenu) and set the GOPATH of the project to $HOME/go in Preferences.

## Installation package manager
```
brew install dep
```

## Clone project
```
git clone git@github.com:detectiveHLH/go-backend-starter.git
```

## Installation dependency package
```
dep ensure
```

## Start project
```
cd $HOME/go/src/go-backend-starter && go run main.go
```