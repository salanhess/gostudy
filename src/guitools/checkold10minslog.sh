#!/bin/bash
num=4
app=/export/log/zbs${num}/
serverLog=/export/log/zbs4/zbs-server/zbs-server.log
schedulerLog=/export/log/zbs4/zbs-scheduler/zbs-scheduler.log
gatewayLog=/export/log/zbs4/zbs-gateway/zbs-gateway.log
wokerLog=/export/log/zbs4/zbs-worker/zbs-worker.log
agentLog=/export/log/zbs4/zbs-agent/zbs-agent.log
proxyLog=/export/log/zbs4/zbs-proxy/zbs-proxy.log

#Not fix path
storage1Log=/export/log/zbs4/zbs-storage/sdg1.log
storage2Log=/export/log/zbs4/zbs-storage/sdh1.log
clientLog=/export/log/zbs4/zbs-agent/nbd201/zbs-client.log
launcherLog=/export/log/zbs4/zbs-agent/nbd201/launcher.log

array=( zbs-agent zbs-gateway zbs-scheduler zbs-server zbs-worker zbs-proxy )
for i in "${array[@]}"
do
    modulename=$i
    echo "=========[Checking $app/${modulename}/${modulename}.log] =================="
    sh checklog.sh old10 $app/${modulename}/${modulename}.log
    sleep 1s
done
