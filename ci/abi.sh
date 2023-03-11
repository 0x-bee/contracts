#!/bin/sh

cd abi
for ITEM in $(ls *.json)
do
    abi-types-generator $ITEM --output ../ts/src --provider=web3
    ../ci/abi_web3.js $ITEM ../ts/src
done
