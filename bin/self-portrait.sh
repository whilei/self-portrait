#!/bin/bash

# execute 50% of the time
p=$(( ( RANDOM % 2 )  + 1 ))
q=1

if [ "$p" -eq "$q" ]
then
# sleep randomly bewteen 1s and 12m
s=$(( ( RANDOM % 720 )  + 1 ))
sleep $s
d="$(date +%y%m%d)-$(date +%H%M).png"

# osx
# imagesnap ~/Pictures/self-portraits/$d

# thought about capturing screen as well, no go.
# screencapture -x ~/Pictures/self-screenshots/sc$d

# linux
/usr/bin/fswebcam --no-subtitle --no-banner --no-timestamp -d /dev/video0 -r 640x480 --png 1 $HOME/Pictures/self-portraits/$d

fi
