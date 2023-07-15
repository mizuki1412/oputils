#! /bin/sh

SERVICE_NAME=waster

#启动方法 
start(){
  count=`ps -ef |grep -e "/root/waster" |grep -v "grep" |wc -l`
  if [ $count -eq 0 ];then
    echo "Starting $SERVICE_NAME ..."
    /root/waster waster --mem=5 &
  fi
  echo "$SERVICE_NAME started ..."
}

#停止方法    
stop(){    
  ps -ef |grep -e "/root/waster" |grep -v "grep" |awk '{print $2}'|xargs kill -9
}

# 检查是否运行中
active(){
  start
}

    
case "$1" in
  active)
  active
  ;;
  start)
  start
  ;;
  stop)
  stop
  ;;
  restart)
  stop
  start
  ;;
*)    
printf 'Usage: %s {start|stop|restart}\n' "$prog"    
exit 1    
;;    
esac  

