#!/bin/bash
set -e
#echo [Sample1]test.sh old10 /export/log/zbs4/zbs-worker/zbs-worker.log
#echo [Sample2]"test.sh 11:20;11:30 /export/log/zbs4/zbs-worker/zbs-worker.log"
#echo "=========================================================================="
case "$1" in
    ([oO][lL][dD]10)
		oldtimeY=`date --date='10 minutes ago' '+ %Y' | sed 's/ //g'`
		oldtimeM=`date --date='10 minutes ago' '+ %m' | sed 's/ //g'`
		oldtimeD=`date --date='10 minutes ago' '+ %d' | sed 's/ //g'`
		oldtime=`date --date='10 minutes ago' '+ %H:%M' | sed 's/ //g'`
		nowtime=`date  '+ %H:%M' | sed 's/ //g'`
		;;
    ([oO][lL][dD]30)
		oldtimeY=`date --date='30 minutes ago' '+ %Y' | sed 's/ //g'`
		oldtimeM=`date --date='30 minutes ago' '+ %m' | sed 's/ //g'`
		oldtimeD=`date --date='30 minutes ago' '+ %d' | sed 's/ //g'`
		oldtime=`date --date='30 minutes ago' '+ %H:%M' | sed 's/ //g'`
		nowtime=`date  '+ %H:%M' | sed 's/ //g'`
		;;		
    (*)
		oldtimeY=`date '+ %Y' | sed 's/ //g'`
		oldtimeM=`date '+ %m' | sed 's/ //g'`
		oldtimeD=`date '+ %d' | sed 's/ //g'`
		oldtime=${1:0:5}
		nowtime=${1:6}
                echo "[Info]use specific timesrange ${oldtime} to ${nowtime}"
		;;
esac

#workerlog=/export/log/zbs4/zbs-worker/zbs-worker.log
workerlog=$2

#echo [Current Case]sed -n ${oldtimeY}\/${oldtimeM}\/${oldtimeD} ${oldtime} '-----------' ${nowtime}  $workerlog
#sed -n "/${oldtimeY}\/${oldtimeM}\/${oldtimeD} ${oldtime}/,/${oldtimeY}\/${oldtimeM}\/${oldtimeD} ${nowtime}/p" $workerlog
#sed -n "/${oldtimeY}\/${oldtimeM}\/${oldtimeD} ${oldtime}/,/${oldtimeY}\/${oldtimeM}\/${oldtimeD} ${nowtime}/p" $workerlog |grep -i Error
sed -n "/${oldtimeY}\/${oldtimeM}\/${oldtimeD} ${oldtime}/,/${oldtimeY}\/${oldtimeM}\/${oldtimeD} ${nowtime}/p" $workerlog |grep -i Error |perl -pe 's/(Error)/\e[1;31m$1\e[0m/g'


