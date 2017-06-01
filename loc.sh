#!/usr/bin/env bash
ROOT=github.com/theodesp/go-shuffled-queue


for pkg in $(go list $ROOT...); do
    wc -l $(go list -f '{{range .GoFiles}}{{$.Dir}}/{{.}} {{end}}' $pkg) | \
        tail -1 | awk '{ print $1 " '$pkg'" }'
done | sort -nr