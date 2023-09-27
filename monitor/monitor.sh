#!/bin/bash

function usage() {
  echo "agora private monitor scripts, version 1.0"
  echo "usage:"
  echo "  ./start.sh [server name]/[all]"
  echo ""
  echo "supported server names are: "
  echo "\tnetdata cadvisor influxdb grafana"
  echo ""
  echo "eg. : ./start.sh netdata"
}

#=======configs========
host=10.10.10.10
influx_host=$host
#=======end config=====


function check_net() {
  net=`docker network ls -q -f name=agora_monitor_network`
  if [[ -z $net ]]; then
    docker network create agora_monitor_network
  fi
}

function start() {
  if [ $# -eq 0 ];then
    echo "no service specified"
    exit 1
  fi
  echo "start service: "$1
  case $1 in
  netdata)
    docker rm -f agora_netdata
    sed -i -e 's/destination.*/destination=$influx_host:4242/' netdata.conf
    docker run --pid host --restart=unless-stopped \
              --cap-add=NET_ADMIN --cap-add sys_ptrace --ulimit core=-1 \
              --security-opt seccomp=unconfined -itd --privileged \
              --name agora_netdata \
              --hostname $host \
              -p 19999:19999 \
              -v /etc/passwd:/host/etc/passwd:ro \
              -v /etc/group:/host/etc/group:ro \
              -v /proc:/host/proc:ro \
              -v /sys:/host/sys:ro \
              -v /var/run/docker.sock:/var/run/docker.sock:ro \
              -v `pwd`/netdata.conf:/etc/netdata/netdata.conf \
              netdata/netdata:v1.33
    ;;
  cadvisor)
    docker rm -f agora_cadvisor
    docker run --pid host --restart=unless-stopped \
              --cap-add=NET_ADMIN --cap-add sys_ptrace --ulimit core=-1 \
              --security-opt seccomp=unconfined -itd --privileged \
              -p 8080:8080 \
              --hostname $host \
              -v /:/rootfs:ro \
              -v /var/run:/var/run \
              -v /sys:/sys:ro \
              -v /dev/kmsg:/dev/kmsg:ro \
              -v /var/lib/docker:/var/lib/docker:ro \
              --name agora_cadvisor \
	            google/cadvisor:v0.32.0 \
              -docker_only \
              -storage_driver=influxdb \
              -storage_driver_db=proxy_monitor \
              -storage_driver_host=${influx_host}
    ;;
  influxdb)
    docker rm -f agora_monitor_influxdb
    chmod 777 -R influxdb
    docker run --pid host --restart=unless-stopped --network agora_monitor_network \
              -p 8086:8086 -p 4242:4242 \
              --cap-add=NET_ADMIN --cap-add sys_ptrace --ulimit core=-1 \
              --security-opt seccomp=unconfined -itd \
              -v `pwd`/influxdb:/var/lib/influxdb \
              -v `pwd`/influxdb.conf:/etc/influxdb/influxdb.conf:ro \
              --name agora_monitor_influxdb influxdb:1.8    
    ;;
  grafana)
    docker rm -f agora_monitor_grafana
    sed -i -e 's/domain.*DOMAIN_NAME./domain = $host/' grafana.ini
    chmod 777 -R grafana
    docker run --pid host --restart=unless-stopped --network agora_monitor_network -p 3000:3000 \
              --cap-add=NET_ADMIN --cap-add sys_ptrace --ulimit core=-1 \
              --security-opt seccomp=unconfined -itd \
              -v `pwd`/grafana:/var/lib/grafana \
              -v `pwd`/grafana.ini:/etc/grafana/grafana.ini:ro \
              --name agora_monitor_grafana grafana/grafana:8.2.0    
    ;;
  ?)
    echo "unsupport service: " $1
    exit 1
    ;;
  esac    
}

all=0
if [[ $# -eq 0 ]];then
  all=1
elif [[ $1 == "all" ]];then
  all=1
fi

if [[ $all -eq 1 ]];then
  check_net
  start "netdata"
  start "cadvisor"
  start "influxdb"
  start "grafana"
  exit 0
fi

if [[ $1 == "influxdb" || $1 == "grafana" ]];then
  check_net
fi
start $1

