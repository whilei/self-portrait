#!/usr/bin/env bash

target_schema="$(cat <<EOF

/2103/
/2103/29-0533.png

/2103/faces/
/2103/faces/29-0533.png

/2103/faces/avgs/
/2103/faces/avgs/avg.png

/2103/faces/montages/
/2103/faces/montages/montage_1.png
/2103/faces/montages/montage_2.png

# -- Nice to have:

/2103/avgs/
/2103/avgs/avg.png

/2103/montages/
/2103/montages/montage_1.png
/2103/montages/montage_2.png

EOF
)"

steps="$(cat <<EOF
- Faces extraction
- Faces -> average creation
- Faces -> montage creation

EOF
)"

program_args="$(cat <<EOF
- ARG1: year_month directory (eg. /path/to/2103)

EOF
)"

set -e

process() {
    local indir
    indir="$1"
    # dirout="$indir"
    outdir="/home/ia/dev/iha.studio/static/self-portrait/$(basename $indir)"

    sp-montage-maker -dirIn "$indir" -dirOut "$outdir/montages" --montage-geo "64x48+1+1"
    sp-face-detector -dirIn "$indir" -dirOut "$outdir/faces"

    sp-avg-maker -dirIn "$outdir/faces" -fileOut "$outdir/faces/avgs/avg.png"
    sp-montage-maker -dirIn "$outdir/faces" -dirOut "$outdir/faces/montages" --montage-geo "48x64+1+1"
}

echo "Directories: $@"

for dir in "$@"
do
    echo "Directory: $dir"
    process "$dir"
done


