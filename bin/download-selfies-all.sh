#!/usr/bin/env bash

# Download all available selfies using wget.

outDir="selfless"
if [ $# -gt 0 ]; then
    outDir="$1"
fi
mkdir -p $outDir

wget -P "$outDir" --recursive --no-parent http://isaacardis.com/selfies/
