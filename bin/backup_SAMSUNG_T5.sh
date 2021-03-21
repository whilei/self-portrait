#!/usr/bin/env bash

set -e

declare -a bks
bks=(SAMSUNG_T5 SAMSUNG_T5_3)
for bk in "${bks[@]}"
do
    if [[ ! -f ~/$bk/CANARY.md ]]; then
        echo "Unmounted: $bk"
        exit 0;
    fi

    rsync -a ~/Pictures/self-portraits/ ~/$bk/self-portraits
    # rsync -a ~/Pictures/self-portraits2/ ~/$bk/self-portraits2

    echo "Backup OK: $bk $(date)"
done
