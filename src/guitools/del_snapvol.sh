#!/bin/bash
#Only available in bash
#data=data.dat

#usr=ebs-admin
usr=$1
#for line in `cat $data`
ip_port=127.0.0.1:9698
#ip_port=192.168.180.116:9698
logpath=log/operation.log
export CC_SERVER_URL="http://$ip_port/zbs-server"
export CC_TENANT_ID="$usr"
#ln -sf /export/jcloud-zbs/bin/zbs-cli /usr/bin/zbs-cli

echo `date "+%Y-%m-%d-%H:%M:%S"`, get volume and snapshot name from zbs_global.snapshot |tee -ai $logpath
for line in `mysql zbs_global -e"select id from snapshot where tenant_id='$usr' and status in (21)" |grep snapshot`
#for line in `mysql zbs_global -e"select id from snapshot where status=21 and snapshot_name like 'snapbh1218_%'" |grep -v name |sed 's/[[:space:]][[:space:]]*/\//g'`
do
    echo "File:${line}" |tee -ai $logpath
    echo delete snapshot by cli |tee -ai $logpath
    zbs-cli snapshot-delete $usr ${line} |tee -ai $logpath
    sleep 2s
done
echo `date "+%Y-%m-%d-%H:%M:%S"` finished |tee -ai $logpath
