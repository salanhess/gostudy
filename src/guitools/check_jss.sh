#!/bin/bash
echo "pls config ets host and parameter according to jss cfg"
cd ../jss
echo "===check 241.103 104 test env ak/sk==="
sh test.sh hddev1
echo "===check pre-public 198  env ak/sk==="
sh test.sh yf
echo "===check 241.101 test env ak/sk==="
sh test.sh hd
#echo "===check 241.110 DEV-DEV env ak/sk==="
#sh test.sh hddev2
#sh test.sh hddev3

