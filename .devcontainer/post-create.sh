#!/bin/bash

# bash auto completion
sudo apt-get update && sudo apt-get install bash-completion
wget https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash -O ~/.git-completion.bash
echo "source ~/.git-completion.bash" >> ~/.bashrc
