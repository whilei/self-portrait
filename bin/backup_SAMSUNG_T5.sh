#!/usr/bin/env bash

set -e
set -x

declare -a bks
bks=(SAMSUNG_T5 SAMSUNG_T5_3)
for bk in "${bks[@]}"
do
    if [[ ! -f ~/$bk/CANARY.md ]]; then
        echo "Unmounted: $bk"
        continue
        # exit 0;
    fi

    # rsync -a ~/Pictures/self-portraits/ ~/$bk/self-portraits

    # Development iteration of command above.
    # https://unix.stackexchange.com/a/558240
    # rsync -rlptgoDvz ~/Pictures/self-portraits/ ~/$bk/self-portraits
    rsync -rgoDvz --progress ~/Pictures/self-portraits/ ~/$bk/self-portraits

    # Commented because this directory is a static archive.
    # rsync -a ~/Pictures/self-portraits2/ ~/$bk/self-portraits2

    echo "Backup OK: $bk $(date)"
done
