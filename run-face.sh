#!/usr/bin/env bash

go run face-detector/face-detect.go \
  --dirIn "$(pwd)/data/examples/originals-2103/" \
  --dirOut "$(pwd)/data/examples/originals-2103/faces/" \
  --cache-nofacelist "$(pwd)/data/knownnofaces"

  # --harrcascade /home/ia/dev/go-opencv/go-opencv/samples/haarcascade_frontalface_alt.xml \
