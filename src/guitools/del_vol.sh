#!/bin/bash
tid=$1
ip_port=172.19.58.198:80
#MariaDB [zbs_global]> GRANT ALL PRIVILEGES ON zbs_global.* TO 'zbs_global'@'172.19.58.194' IDENTIFIED BY 'zbs_global';
export sqlstr="-h172.19.58.198 -uzbs_global -pzbs_global zbs_global"

logpath=log/operation.log
export CC_SERVER_URL="http://$ip_port/zbs-server"
export CC_TENANT_ID="$tid"
echo `date "+%Y-%m-%d-%H:%M:%S"`, start volume del |tee -ai $logpath
#for line in `cat vol.list`
#mysql zbs_global -e"select id from volume where status =1 and tenant_id='test'" |grep vol-
for line in `mysql $sqlstr -e"select id from volume where status =1 and tenant_id='$tid'" |grep vol-`
do
    echo "delete ${line}" |tee -ai $logpath
    zbs-cli volume-delete ${line} |tee -ai $logpath
done
echo `date "+%Y-%m-%d-%H:%M:%S"`, end volume del |tee -ai $logpath
