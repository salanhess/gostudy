#!/bin/bash
set -e
processname=guitools
chmod +x guitools
echo before restart:
ps aux |grep -v grep|grep $processname
ps aux |grep -v grep|grep $processname |awk '{print $2}' |xargs kill -9
sleep 3s
nohup ./$processname > log/guitools.out 2>&1 &
echo after restart:
ps aux |grep -v grep|grep $processname
