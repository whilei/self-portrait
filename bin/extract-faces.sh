#!/usr/bin/env bash

# extract_faces takes one arg: day offset as a integer (eg. 0, -1)
extract_faces() {
    offset="$1"
    [[ -z $offset ]] && { echo "extract_faces: day offset cannot be empty"; exit 1; }

    ye=$(date -d "${offset}day" +%y)
    mo=$(date -d "${offset}day" +%m)
    day=$(date -d "${offset}day" +%d)

    indir="$HOME/Pictures/self-portraits/$ye$mo/"
    outdir="$HOME/Pictures/self-portraits/$ye$mo/faces/"

    ia_face_extractor --inDir="${indir}" --outDir="${outdir}" \

}

extract_faces -1
extract_faces 0




