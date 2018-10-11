#!/bin/bash

echo "[===check 101 102 OPENAPI UC and account charge===]"
ping -c 3 uc-inner-api-test1.jcloud.com
pin=iass10
region=cn-north-1
ucurl=uc-inner-api-test1.jcloud.com
curl -H 'token: test' -XPOST "http://$ucurl/usercenter/getRegionAzMapping?pin=${pin}&region=${region}" |python -mjson.tool  
curl -H "token: test" -XPOST "http://$ucurl/usercenter/getBalance?pin=${pin}" |python -mjson.tool 

echo "pls config ets host and parameter according to jss cfg"
cd ../jss
echo "[===check hd  env s3 method===]"
sh use_s3.sh hd put uploadabc.txt
sh use_s3.sh hd list uploadabc.txt
sh use_s3.sh hd del uploadabc.txt

echo "[===check yf env s3 method===]"
sh use_s3.sh yf put uploadabc.txt
sh use_s3.sh yf list uploadabc.txt
sh use_s3.sh yf del uploadabc.txt


#echo "===check 241.103 104 test env ak/sk==="
sh test.sh hddev1
echo "[===check pre-public 198  env ak/sk===]"
sh test.sh yf

echo "[===check pre-public 198  env ak/sk===]"
sh test.sh yf

#echo "===check 241.101 test env ak/sk==="
#sh test.sh hd
#echo "===check 241.110 DEV-DEV env ak/sk==="
#sh test.sh hddev2
#sh test.sh hddev3

