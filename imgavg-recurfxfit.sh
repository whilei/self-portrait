#!/usr/bin/env bash

# Best one for cropped faces.
selfies_dir=$1
out_file=$2

selfies=($(find "$selfies_dir" -type f))

echo "Output average file will be $out_file"

i=0
len=${#selfies[@]}
for file in "${selfies[@]}"; do
	if [[ $file =~ .png ]]; then
		echo "$i/$len - averaging $file into $out_file"
		if [ $i -eq 0 ]; then
			convert "$file" -background transparent -gravity center -resize 640x480 "$out_file"
		else
			convert "$file" -background transparent -gravity center -resize 640x480 resized_tmp_file.png
			convert resized_tmp_file.png "$out_file" -fx "(u+$i*v)/$((i + 1))" "$out_file"
			rm resized_tmp_file.png
		fi
		i=$((i + 1))
	fi
done

unset selfies_dir
unset out_file
unset selfies
unset i
unset len