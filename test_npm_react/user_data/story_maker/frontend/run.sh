#!/bin/bash
echo "Running file ${BASH_SOURCE[0]} in directory $(dirname ${BASH_SOURCE[0]})"

CURRENT_DIR="$(pwd)"

TARGET_DIR="$(dirname ${BASH_SOURCE[0]})"
echo "Changing to $TARGET_DIR"
cd $TARGET_DIR
COMMAND="npm run dev -- --host"
echo "executing command \"$COMMAND\""
eval "$COMMAND"
echo "Script finished changing back to $CURRENT_DIR"
cd $CURRENT_DIR