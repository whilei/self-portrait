#!/usr/bin/env bash

set -e

process() {
    local indir
    local outdir

    indir="$1"
    # dirout="$indir"

    outdir="~/iha.studio.dev/static/self-portrait/$(basename $indir)"

    ~/bin/sp-montage-maker -dirIn "$indir" -dirOut "$outdir/montages" --montage-geo "64x48+1+1"
    ~/bin/sp-face-detector -dirIn "$indir" -dirOut "$outdir/faces"

    ~/bin/sp-avg-maker -dirIn "$outdir/faces" -fileOut "$outdir/faces/avgs/avg.png"
    ~/bin/sp-montage-maker -dirIn "$outdir/faces" -dirOut "$outdir/faces/montages" --montage-geo "48x64+1+1"
}

echo "Directories: $@"

for dir in "$@"
do
    echo "Directory: $dir"
    process "$dir"
done

