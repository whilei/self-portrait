#!/usr/bin/env bash

selfies_dir=$1
out_file=$2

echo "Using simple 'convert -evaluate-sequence mean'."
echo "Output average file will be $out_file"

convert "$selfies_dir/*.png" \
	-evaluate-sequence mean \
	"$out_file"

unset selfies_dir
unset out_file
