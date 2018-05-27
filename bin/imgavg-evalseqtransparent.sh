#!/usr/bin/env bash

selfies_dir=$1
out_file=$2

echo "Output average file will be $out_file"
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
	"$out_file"

unset selfies_dir
unset out_file
