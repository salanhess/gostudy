#!/bin/bash

printf "*************************************\n"  
echo " cat file whiel read line"  
cat vollist.txt |while read line  
do  
  echo sh del_volid.sh $line
  sh del_volid.sh $line
done  
