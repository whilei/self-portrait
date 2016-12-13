#!/bin/bash

# Usage: ./self-ish.sh -[|m|M|r|R] path/to/selfies/dir out/dir
#
# TODO: argue name of out file
# TODO: average by time period (month, year, decade)
# TODO: keep track of which files were used to make a given average,
# so work can be not redundant and build always something new from what already exists

# -m
avg_method='evalseq' # default -evaluate-sequence mean
	#-M 'evalseqtp' # set transparent background for -evaluate-sequence mean (possible best for face-detectored)
	#-r 'recurfx' # one-by-one recurs custom -fx strategy per stephan pauker (see comm below)
	  # takes forever. is less likely to break the compbrain.
	#-R 'recurfxfit' # same as one-by-one but uses a tmp file to resize before
		#	may be best for averaging cropped faces

while getopts ':mMrR' opt; do
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
		R)
			echo "Using one-by-one recursive -fx strategy with fit"
			avg_method='recurfxfit'
			shift
			;;
		\?)
			echo "Use: ./self-ish.sh -[|m|M|r|R] path/to/selfies/dir out/dir"
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
		if [[ $file =~ .png ]]; then
		  echo "Averaging $file into $out_file"
		  if [ $i -eq 0 ]; then
		    cp "$file" "$out_dir/self-ish-recurfx.png"
		  else
		    # u is first image in list
		    # v is second
		    convert "$file" "$out_file" -fx "(u+$i*v)/$((i+1))" "$out_file"
		  fi
		  i=$((i+1))
		fi
	done

# Best one for cropped faces.
elif [[ "$avg_method" == recurfxfit ]]; then

	selfies=($(find "$selfies_dir" -type f))

	out_file="$out_dir/self-ish-recurfxfit.png"
	echo "Output average file will be $out_file"

	i=0
	for file in "${selfies[@]}"; do
		if [[ $file =~ .png ]]; then
			echo "Averaging $file into $out_file"
			if [ $i -eq 0 ]; then
				convert "$file" -background transparent -gravity center -resize 640x480 "$out_file"
			else
			  convert "$file" -background transparent -gravity center -resize 640x480 resized_tmp_file.png
			  convert resized_tmp_file.png "$out_file" -fx "(u+$i*v)/$((i+1))" "$out_file"
			  rm resized_tmp_file.png
			fi
			i=$((i+1))
		fi
	done

# Be advised this reads all files into mem.
# Memory an issue? For 10k+ files, likely.
# Try: -limit memory 16mb -limit map 32mb
elif [[ "$avg_method" == evalseqtp ]]; then
	echo "Output average file will be $out_dir/self-ish-esmean-bgtrans.png"
	# http://blog.patdavid.net/2012/08/imagemagick-average-blending-files.html

	# background, resize, extent, and gravity will make face-cropped selfies consistently sized.
	# since all faces are smaller than original image, faces will be resized to fit within
	# resize will not FILL the space, they will FIT INTO the requested size
	convert "$selfies_dir/*.png" \
	-background transparent \
	-gravity center \
	-resize 640x480 \
	-extent 640x480 \
	-evaluate-sequence mean \
	"$out_dir/self-ish-esmean-bgtrans.png"
else

	echo "Using default 'convert -evaluate-sequence mean'."
	echo "Output average file will be $out_dir/self-ish-esmean.png"

	convert "$selfies_dir/*.png" \
	-evaluate-sequence mean \
	"$out_dir/self-ish-esmean.png"
fi

# sources
unset selfies_dir
unset selfies
unset out_dir
unset out_file
# IFS=$SAVEIFS # revert ifs
