#!/bin/bash

# Usage: ./self-ish.sh -[|m|M|r] path/to/selfies/dir out/dir
# 
# TODO: argue name of out file

avg_method='evalseq' # default -evaluate-sequence mean
					#'evalseqtp' # set transparent background for -evaluate-sequence mean (possible best for face-detectored)
					#'recurfx' # one-by-one recurs custom -fx strategy per stephan pauker (see comm below)
										 # takes forever. is less likely to break the compbrain.

while getopts ':mMr' opt; do
	case $opt in
		m)
			# default
			shift
			;;
		M)
			echo "Using 'convert -background transparent -evaluate-sequence mean'."
			avg_method='evalseqtp'
			shift
			;;
		r)
			echo "Using one-by-one recursive -fx strategy."
			avg_method='recurfx'
			shift
			;;
		\?)
			echo "Use: ./self-ish.sh -[|m|M|r] path/to/selfies/dir out/dir" 
			exit 1;
			;;
		*)
			echo "Doing star thingey." 
			exit 1
			;;
	esac
done

selfies_dir=$1
out_dir=$2
mkdir -p "$out_dir"

echo "Using selfies from $selfies_dir"

if [[ "$avg_method" == recurfx ]]; then

	# fix recurs avg?
	# http://stephan.paukner.cc/syslog/archives/362-Averaging-an-image-sequence-with-ImageMagick.html
	selfies=($(find "$selfies_dir" -type f))
	out_file="$out_dir/self-ish-recurfx.png"
	echo "Output average file will be $out_file"
	# SAVEIFS=$IFS
	# IFS=$(echo -en "\n\b") # Use newline instead of spaces.
	i=0
	for file in "${selfies[@]}"; do
	  echo "Averaging $file into $out_file"
	  if [ $i -eq 0 ]; then
	    cp "$file" "$out_dir/self-ish-recurfx.png"
	  else
	    # u is first image in list
	    # v is second
	    convert "$file" "$out_file" -fx "(u+$i*v)/$((i+1))" "$out_file"
	  fi
	  i=$((i+1))
	done

elif [[ "$avg_method" == evalseqtp ]]; then
	echo "Output average file will be $out_dir/self-ish-esmean-bgtrans.png"
	# http://blog.patdavid.net/2012/08/imagemagick-average-blending-files.html
	
	# since they all same size (can use for face-detected selfie tho)
	convert -background transparent "$selfies_dir/*.png" \
	-gravity center -extent 640x480 \
	-evaluate-sequence mean \
	"$out_dir/self-ish-esmean-bgtrans.png"
else
	# be advised this reads all files into mem
	echo "Using default 'convert -evaluate-sequence mean'."
	echo "Output average file will be $out_dir/self-ish-esmean.png"
	convert "$selfies_dir/*.png" -evaluate-sequence mean "$out_dir/self-ish-esmean.png"	
fi

# sources
unset selfies_dir
unset selfies
unset out_dir
unset out_file
# IFS=$SAVEIFS # revert ifs
