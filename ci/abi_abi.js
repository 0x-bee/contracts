#!/usr/bin/env node

const path = require('path')
const fs = require('fs')

let inFile = process.argv[2]
let outFile = process.argv[3]

let jsonObject = JSON.parse(fs.readFileSync(inFile).toString())

fs.writeFileSync(outFile, JSON.stringify(jsonObject.abi, undefined, 2))
