#!/bin/sh
#
# Convert JSON-exported terminal.sexy colorscheme to rainbou YAML
#

usage() {
	echo "Usage: $0 <filename>"
	exit 1
}

filename="$1"

stat "$filename" > /dev/null 2>&1 || usage

name="$(jq '.name' "$filename" | sed 's/\"//g')"
author="$(jq '.author' "$filename" | sed 's/\"//g')"
background="$(jq '.background' "$filename" | sed 's/\#//')"
foreground="$(jq '.foreground' "$filename" | sed 's/\#//')"

cat <<BLOCK1
metadata:
  name: ${name}
  author: ${author}
colors:
  bg: ${background}
  fg: ${foreground}
  cr: ${foreground}

BLOCK1

colors="$(jq '.color[]' "$filename" | sed 's/\#//g' | tr '\n' ' ')"

i=0
while [ "$i" -ne 8 ]; do
	dark="$(echo "$colors" | cut -d' ' -f$((i + 1)) )"
	light="$(echo "$colors" | cut -d' ' -f$((i + 8 + 1)) )"

	printf "  %02d: %s\n" "$i" "$dark"
	printf "  %02d: %s\n" "$((i + 8))" "$light"

	test "$i" -ne 7 && echo ""

	i="$((i + 1))"
done
