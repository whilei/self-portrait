#!/usr/bin/env bash

selfies_dir=$1
out_file=$2

echo "Using selfies from $selfies_dir"

# fix recurs avg?
# http://stephan.paukner.cc/syslog/archives/362-Averaging-an-image-sequence-with-ImageMagick.html
selfies=($(find "$selfies_dir" -type f))
echo "Output average file will be $out_file"
# SAVEIFS=$IFS
# IFS=$(echo -en "\n\b") # Use newline instead of spaces.
i=0
len=${#selfies[@]}
for file in "${selfies[@]}"; do
	if [[ $file =~ .png ]]; then
		echo "$i/$len - averaging $file into $out_file"
		if [ $i -eq 0 ]; then
			cp "$file" "$out_file"
		else
			# u is first image in list
			# v is second
			convert "$file" "$out_file" -fx "(u+$i*v)/$((i + 1))" "$out_file"
		fi
		i=$((i + 1))
	fi
done

unset selfies_dir
unset out_file
unset selfies
unset i
unset len