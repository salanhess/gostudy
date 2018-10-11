#!/bin/bash
set -e
set -o nounset
#Only available in bash
#status 3,4,11 > already deleted ,status 1 > del ,status 2 > try detach,del ,other status > need zbs admin operation
vol=$1
#tid=ebs-admin
sip=10.226.134.236:9298
#sip=172.19.58.198:80
#MariaDB [zbs_global]> GRANT ALL PRIVILEGES ON zbs_global.* TO 'zbs_global'@'172.19.58.194' IDENTIFIED BY 'zbs_global';
#export sqlstr="-h192.168.241.101 -uroot -pCPtXq@Nv0k -P3306 zbs_global"
export sqlstr=" -h10.226.134.236 -uzbs_global_zbs2 -pzbs_global_zbs2 -P3306 zbs_global_zbs2"
Maxtry=20
logpath=log/operation.log
#ln -sf /export/jcloud-zbs/bin/zbs-cli /usr/bin/zbs-cli
echo ==== mysql ${sqlstr} ======
echo `date "+%Y-%m-%d-%H:%M:%S"`, start del volume by volumeid |tee -ai $logpath

volstatus=`mysql ${sqlstr} -e"select status from volume where id='${vol}'" |grep -v status |awk '{print $1}'`
if [ -z $volstatus ]; then
    echo "${vol} not exist!" |tee -ai $logpath
    exit 0
fi

tid=`mysql ${sqlstr} -e"select tenant_id from volume where id='${vol}'" |grep -v tenant_id |awk '{print $1}'`
export CC_SERVER_URL="http://$sip/zbs-server"
export CC_TENANT_ID="$tid"

volaction=`mysql ${sqlstr} -e"select action from volume where id='${vol}'" |grep -v action |awk '{print $1}'`
if [ $volaction ]; then
    echo "${vol} is in status: $volstatus ,with action:$volaction ,can't del!" |tee -ai $logpath
    echo "try to resume volume status and action..." |tee -ai $logpath
    echo curl -X POST http://${sip}/zbs-server?Action=ResumeVolume -H  "Content-Type: application/json" -H "User-Agent: ZbsClient" -d "{\"id\":\"${vol}\",\"tenant_id\":\"$tid\"}"
    curl -X POST http://${sip}/zbs-server?Action=ResumeVolume -H  "Content-Type: application/json" -H "User-Agent: ZbsClient" -d "{\"id\":\"${vol}\",\"tenant_id\":\"$tid\"}"
    exit 0
fi

echo [Info]${vol} status is $volstatus
case $volstatus in
    0)
    echo "${vol} is creating,will delete from DB" |tee -ai $logpath
    mysql ${sqlstr} -e"delete from volume where id='${vol}' and status in (0)"
    ;;
    3|4|11)
    echo "${vol} already deleted" |tee -ai $logpath
    ;;
    1|7)
    echo "${vol} is availalbe(1) or error_creating(7),will del" |tee -ai $logpath
    zbs-cli volume-delete ${vol} |tee -ai $logpath
    zbs-cli volume-describe ${vol} |tee -ai $logpath
    ;;
    2)
    echo "${vol} is in-use,will detach,then del(Support multi-detach NOW)" |tee -ai $logpath
    attachid=`mysql ${sqlstr} -e"select *  from volume_attachment where volume_id='${vol}' and status in (32,35)" |grep vol-tach |awk '{print $1}'`
    #attachid=`mysql  -h$sqlip -uzbs_global -pzbs_global zbs_global -e"select *  from volume_attachment where volume_id='${vol}' and status!=33" |grep vol-tach |awk '{print $1}'`
    if [ -z "$attachid" ] ;then
       echo " volume ${vol} attach status not avaible,need reset process" |tee -ai $logpath
    else
       echo "${vol} attachid is $attachid" |tee -ai $logpath
       echo delete attached volume by cli |tee -ai $logpath
       zbs-cli volume-detach ${vol} $attachid |tee -ai $logpath

       volstatus=
       volattachstatus=	
       num=0
       while [[ $volattachstatus != "detached" && $volstatus != "available" ]]
       do
     if [ $num -eq $Maxtry ]
     then
            echo "Max try num access!" |tee -ai $logpath
            zbs-cli attachment-describe $attachid
        break              #Abandon the while lopp.
         fi
          attachid=`mysql ${sqlstr} -e"select *  from volume_attachment where volume_id='${vol}' and status in (32,35)" |grep vol-tach |awk '{print $1}'`
	  if [ -n "$attachid" ] ;then
              echo "${vol} multi-attachid is $attachid" |tee -ai $logpath
              zbs-cli volume-detach ${vol} $attachid |tee -ai $logpath
              sleep 3s
          fi
          volattachstatus=`zbs-cli attachment-describe $attachid |  grep status  | awk -F '|' '{print $3}' | sed 's/ //g'`
          volstatus=`zbs-cli volume-describe $vol |  grep status  | awk -F '|' '{print $3}' | sed 's/ //g'`
          echo "wait vol detach. now $vol is $volstatus, volattachstatus is $volattachstatus" |tee -ai $logpath
          sleep 3s
      num=`expr $num + 1`
       done
       echo "begin to del"
       zbs-cli volume-delete ${vol} |tee -ai $logpath
       zbs-cli volume-describe ${vol} |tee -ai $logpath
       sleep 1s
    fi
    ;;
    *)
    echo "${vol} is in status $volstatus ,pls contact administrator to del" |tee -ai $logpath
    ;;
esac

exit 0

