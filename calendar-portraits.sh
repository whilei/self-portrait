#!/usr/bin/env bash

croppedFacesBaseDir=/Users/ia/Pictures/faceout-test

outBaseDir="./calendar"
mkdir -p "$outBaseDir"

for year in {15..18}; do
    yearbook="$croppedFacesBaseDir/$year"
    for month in {01..12}; do 
        monthbook="$yearbook/$month"
        mkdir -p "$monthbook"

        for f in "$yearbook/face_$year$month"*; do
            mv "$f" "$monthbook/"
        done

        calendarSpot="$outBaseDir/$year-$month.avg-recurfxfit.png"
        source ./imgavg-recurfxfit.sh "$monthbook" "$calendarSpot"
    done
done