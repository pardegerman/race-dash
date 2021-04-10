#!/bin/sh
set -e

# Create config directories if needed
mkdir -p .kivy
mkdir -p .config/racecapture

# Configure kivy
cp -n /racecapture/ft5406_kivy_config.ini .kivy/config.ini
crudini --set .kivy/config.ini input btns hidinput,/dev/input/event1
crudini --set .kivy/config.ini input mouse mouse,disable_on_activity
crudini --set .kivy/config.ini graphics show_cursor 0
crudini --del .kivy/config.ini modules cursor

HOMEDIR=$(pwd)
EXECDIR=$(dirname $1)
cd ${EXECDIR} && HOME=${HOMEDIR} $@
