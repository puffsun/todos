#!/usr/bin/env bash

[ -z "$GOPATH" ] && { echo "Need to set GOPATH"; exit 1; }

if [ -x ${GOPATH}/bin/fresh ]; then
    $GOPATH/bin/fresh
else
    go run ./main.go
fi
