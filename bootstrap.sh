#!/bin/bash

export CWD=${PWD##*/}

for F in *
do
    if [ -d $F ]; then
        cd $F
        sed -i 's/vokal-go-seed/'$CWD'/g' *.go
        cd ..
    else
        sed -i 's/vokal-go-seed/'$CWD'/g' *.go
    fi
done

exit 0
