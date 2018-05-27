#!/usr/bin/env bash

basedir=/Users/ia/Pictures/self-portraits

i=0
yy="00"
mm="00"
for p in "$basedir"/*.png; do
    if [ i -eq 0 ]; then
        i=$((i + 1))
        yy=${p:34:2}
        mm=${p:36:2}
    fi
done