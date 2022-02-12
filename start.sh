#!/usr/bin/env bash
cd ./src

ARGS=$(getopt -a --options r:h --long "rom,help" -- "$@")

eval set -- "$ARGS"

while true; do
  case "$1" in
    -r|--rom)
      rom="$2"
      ./chip8 -rom rom/$rom
      shift 2;;
    -h|--help)
      echo "Please use -r/--rom <local> to execute"
      break;;
    --)
      break;;
    esac
done
