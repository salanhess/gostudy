#!/bin/bash
set -e
set -o nounset
#Only available in bash
#status 3,4,11 > already deleted ,status 1 > del ,status 2 > try detach,del ,other status > need zbs admin operation
vol=$1
#tid=ebs-admin
sip=192.168.241.104:9698
#sip=172.19.58.198:80
#MariaDB [zbs_global]> GRANT ALL PRIVILEGES ON zbs_global.* TO 'zbs_global'@'172.19.58.194' IDENTIFIED BY 'zbs_global';
#export sqlstr="-h172.19.58.198 -uzbs_global -pzbs_global zbs_global"
export sqlstr="zbs_global"
Maxtry=12
logpath=log/operation.log
#ln -sf /export/jcloud-zbs/bin/zbs-cli /usr/bin/zbs-cli
echo ==== mysql ${sqlstr} ======
echo `date "+%Y-%m-%d-%H:%M:%S"`, start check volume status by volumeid |tee -ai $logpath

volstatus=`mysql ${sqlstr} -e"select status from volume where id='${vol}'" |grep -v status |awk '{print $1}'`
if [ -z $volstatus ]; then
    echo "${vol} not exist!" |tee -ai $logpath
    exit 0
fi

tid=`mysql ${sqlstr} -e"select tenant_id from volume where id='${vol}'" |grep -v tenant_id |awk '{print $1}'`
export CC_SERVER_URL="http://$sip/zbs-server"
export CC_TENANT_ID="$tid"

volaction=`mysql ${sqlstr} -e"select action from volume where id='${vol}'" |grep -v action |awk '{print $1}'`
zbs-cli volume-describe ${vol}

mysql ${sqlstr} -e"select * from volume_attachment where volume_id='${vol}' and status !=33 \G"
