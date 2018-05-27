#!/usr/bin/env bash

croppedFacesBaseDir=/Users/ia/Pictures/faceout-test

outBaseDir="$1"
mkdir -p "$outBaseDir"

for year in {15..18}; do
    yearbook="$croppedFacesBaseDir/$year"
    for month in {01..12}; do 
        calendarSpot="$outBaseDir/$year-$month.avg-recurfxfit.png"
        if [ -f "$calendarSpot" ]; then
            echo "Found pre-existing average: $calendarSpot - skipping"
            continue
        fi
        monthbook="$yearbook/$month"
        mkdir -p "$monthbook"

        for f in "$yearbook/face_$year$month"*; do
            mv "$f" "$monthbook/"
        done

        source ./imgavg-recurfxfit.sh "$monthbook" "$calendarSpot"
    done
done