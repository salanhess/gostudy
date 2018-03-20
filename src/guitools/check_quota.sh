#!/bin/bash
tid=$1
ip_port=192.168.180.116:9698
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

