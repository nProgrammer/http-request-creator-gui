#!/bin/bash

mkdir $HOME/.http-request-creator-gui/
mkdir $HOME/.http-request-creator-gui/bin
mkdir $HOME/.http-request-creator-gui/tools

go build

mv request-creator-gui $HOME/.http-request-creator-gui/bin/request-creator-gui
cp ./installer.sh $HOME/.http-request-creator/tools

echo "Now add to the .bashrc or .zshrc command: export PATH=\$PATH:\$HOME/.http-request-creator-gui/bin" #FIXME - check grammar
echo "To use http-request-creator-gui use command: request-creator-gui" #FIXME - check grammar%  