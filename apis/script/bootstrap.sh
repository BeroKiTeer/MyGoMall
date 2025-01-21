#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=apis
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}