#!/bin/bash
set -e
#chmod 0600 id_rsa_jenkins_bot
#ssh -i id_rsa_jenkins_bot 192.168.180.118 'cd /export/App/zbs1/zbs-scheduler && sh recycle_deletedvol.sh'
cd /export/App/zbs4/zbs-scheduler

export SCHEDULER_PORT=21400
mysqlstr=" zbs_global"
#mysqlstr="-h192.168.172.93 -uzbs_global -pzbs_global zbs_global"
echo `date "+%Y-%m-%d-%H:%M:%S"`, start recycle volume
sleep 1s
result=`mysql $mysqlstr -e"select id  from volume where status in (4,9)"`

echo ========before recycle volume,zbs-storage in memory===========
curl -s '127.0.0.1:21401/common?action=supernetnode' | grep disk  |grep host

if [ -z "$result" ] ;then
   echo "No deleted volume need to recycle"
   exit 0
fi

for tbl in `mysql $mysqlstr -e"select id  from volume where status in (4,9)" |grep - | sed 's/\r//g'`
do
   volstr=`mysql $mysqlstr -e "select id,status,tenant_id from volume where id = '$tbl'" |grep -`
   echo ./zbs-recycle ====== $volstr ======
   ./zbs-recycle $tbl
   #sleep 3s
done

echo ========After recycle volume,zbs-storage in memory===========
curl -s '127.0.0.1:21401/common?action=supernetnode' | grep disk |grep host
echo `date "+%Y-%m-%d-%H:%M:%S"`, end recycle volume


