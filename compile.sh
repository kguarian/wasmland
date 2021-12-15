#/usr/bin/bash

#variables:
#errorCode
cd goimp
#compiles to pkg/wasm_lib.wasm
`./compile.sh`
errorcode=$?
echo $errorcode

cd ../

if (( errorcode != 0 )); then
    echo "$errorCode"
    echo "failed Go build."
    exit 0
fi

#export
mkdir docs
sudo cp -r html/* style pkg docs
sudo cp -r docs/* /var/www/html/lootjs/