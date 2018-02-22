#!/bin/bash
set -e
set -o nounset
#Only available in bash
#status 35 Errordetached, 32 volumeattached
tid=$1
#tid=ebs-admin
sip=192.168.180.116:9698
Maxtry=3
export CC_SERVER_URL="http://$sip/zbs-server"
export CC_TENANT_ID="$tid"
logpath=log/operation.log
#ln -sf /export/jcloud-zbs/bin/zbs-cli /usr/bin/zbs-cli
echo `date "+%Y-%m-%d-%H:%M:%S"`, start del volume with attach |tee -ai $logpath
for vol in `mysql zbs_global -e"select id  from volume where tenant_id='$tid' and status=2" |grep vol-`
#for vol in `mysql  -h$sqlip -uzbs_global -pzbs_global zbs_global -e"select id  from volume where tenant_id='$tid' and status=2" |grep vol-`
do
    attachid=`mysql zbs_global -e"select *  from volume_attachment where volume_id='${vol}' and status=32" |grep vol-tach |awk '{print $1}'`
    #attachid=`mysql  -h$sqlip -uzbs_global -pzbs_global zbs_global -e"select *  from volume_attachment where volume_id='${vol}' and status!=33" |grep vol-tach |awk '{print $1}'`
    if [ -z "$attachid" ] ;then
       echo " volume ${vol} attach status not avaible,need reset process" |tee -ai $logpath
    else
       echo "${vol} attachid is $attachid" |tee -ai $logpath
       echo delete attached volume by cli |tee -ai $logpath
       echo zbs-cli volume-detach ${vol} $attachid |tee -ai $logpath

       volstatus=
       num=0
       while [[ $volstatus != "available" ]]
       do
  	 	if [ $num -eq $Maxtry ]
  	 	then
            echo "Max try num access!" |tee -ai $logpath
	    	break       	   #Abandon the while lopp.
        fi
          volstatus=`zbs-cli volume-describe $vol |  grep status  | awk -F '|' '{print $3}' | sed 's/ //g'`
          echo "wait vol detach. now $vol is $volstatus" |tee -ai $logpath
          sleep 3
	  num=`expr $num + 1`
       done
       zbs-cli volume-delete ${vol} |tee -ai $logpath
       sleep 1s
    fi
    #mysql  -h172.19.54.101 -uzbs_global -pzbs_global zbs_global -e"delete from volume where tenant_id='$tid' and status=3"
done
echo `date "+%Y-%m-%d-%H:%M:%S"`, end  del volume with attach |tee -ai $logpath
