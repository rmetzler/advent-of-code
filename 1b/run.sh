cat input.txt | \
    gsed \
    -e 's/one/one1one/g' \
    -e 's/two/two2two/g' \
    -e 's/three/three3three/g' \
    -e 's/four/four4four/g' \
    -e 's/five/five5five/g' \
    -e 's/six/six6six/g' \
    -e 's/seven/seven7seven/g' \
    -e 's/eight/eight8eight/g' \
    -e 's/nine/nine9nine/g' | \
gsed -e 's/[^0-9]//g' | gsed -r 's/^(.)$/\1\1/' | gsed -r 's/^(.)(.*)(.)$/\1\3/' | awk '{s+=$1} END {print s}'
