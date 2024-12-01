cat input.txt | gsed -e 's/[^0-9]//g' | gsed -r 's/^(.)$/\1\1/' | gsed -r 's/^(.)(.*)(.)$/\1\3/' | awk '{s+=$1} END {print s}'
