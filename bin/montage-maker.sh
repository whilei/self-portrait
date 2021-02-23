#!/usr/bin/env bash

montage $(find data/examples/faces -type f | tail -64) -tile 8x8 -geometry 32x32+1+1 montage-faces-64.jpg
montage $(find data/examples/originals -type f | tail -64) -tile 8x8 -geometry 32x24+1+1 montage-originals-64.jpg


