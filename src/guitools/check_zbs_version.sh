#!/bin/bash
zbspath=/export/App/zbs4
cd $zbspath
array=( zbs-agent zbs-gateway zbs-scheduler zbs-server zbs-worker zbs-storage )
for i in "${array[@]}"
do
    modulename=$i
    echo "=========checking ${modulename}=================="
    cd ${modulename} && ls -l |grep ${modulename} |awk '{print $11}' && ./${modulename} -v && cd ..
    sleep 1s
done
