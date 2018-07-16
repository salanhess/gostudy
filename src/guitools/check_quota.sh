#!/bin/bash
tid=$1
ip_port=192.168.241.104:9698
logpath=log/operation.log
export CC_SERVER_URL="http://$ip_port/zbs-server"
export CC_TENANT_ID="$tid"
echo `date "+%Y-%m-%d-%H:%M:%S"`, start check quota |tee -ai $logpath
mysql zbs_global -e"select * from quota where tenant_id='$tid'"
echo "=======current available and in-use Volumes list:========="
zbs-cli volume-list |egrep "available|in-use"
echo "=======current available and in-use Snapshots list:========="
zbs-cli snapshot-list |egrep "available|in-use"
echo `date "+%Y-%m-%d-%H:%M:%S"`, end check quota |tee -ai $logpath

echo "======check current DB use volume storage status:======"
mysql zbs_global -e"select sum(size)/1024/1024/1024, count(*) ,tenant_id  from volume where status in (1,2) group by tenant_id order by count(*) DESC"

echo "[======check current DB zbs-tenant id whilelist(ONLY support following tenant_id,az_name,volume_type  create vol)======]"
mysql zbs_global -e"pager less -SFX;select * from ebs_tenant"

echo "=======Full quotalist as following:========="
mysql zbs_global -e"select * from quota"
