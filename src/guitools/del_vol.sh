#!/bin/bash
tid=$1
ip_port=192.168.180.116:9698
logpath=log/operation.log
export CC_SERVER_URL="http://$ip_port/zbs-server"
export CC_TENANT_ID="$tid"
echo `date "+%Y-%m-%d-%H:%M:%S"`, start volume del |tee -ai $logpath
#for line in `cat vol.list`
#mysql zbs_global -e"select id from volume where status =1 and tenant_id='test'" |grep vol-
for line in `mysql zbs_global -e"select id from volume where status =1 and tenant_id='$tid'" |grep vol-`
do
    echo "delete ${line}" |tee -ai $logpath
    echo zbs-cli volume-delete ${line} |tee -ai $logpath
done
echo `date "+%Y-%m-%d-%H:%M:%S"`, end volume del |tee -ai $logpath
