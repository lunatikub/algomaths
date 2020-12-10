#!/bin/bash

IN=$1
OUT=$2

[ ! -f "${IN}" ] && \
    echo "File '${IN}' doesn't exist..." 2>&1 \
    && exit 1

grep -v "'s$" ${IN} |
    grep -v "[éèàùâ]" |
    tr '[:upper:]' '[:lower:]' |
    sort | uniq > ${OUT}

wc -l ${OUT}

exit 0
