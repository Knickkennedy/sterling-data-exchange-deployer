#!/bin/bash
SCRIPT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )

echo "-------------------------------------------------------------------------------"
echo "Entering Sterling Data Exchange Deployer command line in a container."
echo 'Use the "exit" command to leave the container and return to the hosting server.'
echo "-------------------------------------------------------------------------------"

export PS1='\[\e]0;\w\a\]\n[\#] \[\e[32m\u@SDE Deployer:\[\e[33m\]\w \e[m\$ ';
export PATH=$PATH:/sde-deployer

/bin/bash

exit 0