#!/bin/sh
set -e

# Create config directories if needed
mkdir -p .kivy
mkdir -p .config/racecapture

# Configure for the correct touchscreen driver
if lsusb | grep -q 04d8:0c02; then
  cp -n ar1100_kivy_config.ini .kivy/config.ini
fi
if lsmod | grep -q rpi_ft5406; then
  cp -n ft5406_kivy_config.ini .kivy/config.ini
fi

HOMEDIR=$(pwd)
EXECDIR=$(dirname $1)
cd ${EXECDIR} && HOME=${HOMEDIR} $@
