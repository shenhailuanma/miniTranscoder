#!/bin/bash

function log_info()
{
    datestring=`date "+%Y-%m-%d"`
    timestring=`date "+%H:%M:%S"`
    echo -e "\033[32m[Info ][$datestring $timestring]" "$1 \033[0m"
}

function log_warn()
{
    datestring=`date "+%Y-%m-%d"`
    timestring=`date "+%H:%M:%S"`
    echo -e "\033[33m[Warn ][$datestring $timestring]" "$1 \033[0m"
}

function log_error()
{
    datestring=`date "+%Y-%m-%d"`
    timestring=`date "+%H:%M:%S"`
    echo -e "\033[31m[Error][$datestring $timestring]" "$1 \033[0m"
}

# params:
#   $1 : message
function check_error() {
    if [ $? -eq 0 ];then
        log_info "[OK] $1"
    else
        log_error "[Failed] $1"
        exit 1
    fi
}



# Main
# Inputs:
# [must]      $1 : server, user@server-address
# [must]      $2 : server, user@server-address
# [must]      $3 : server password
# [optional]  $4 : server ssh port
# example:
#       upload-base.sh ssh root@192.168.1.162 'password' 22

if [ $# -lt 3 ]; then
    log_error "Input params number error, current number is $#"
    return 1
fi

TargetBin=$1
RemoteServer=$2
Password=$3
RemotePort=$4

if [ ! -n "$RemotePort" ] ;then
    RemotePort=22
fi

log_info "TargetBin: ${TargetBin}"
log_info "RemoteServer: ${RemoteServer}"
log_info "RemotePort: ${RemotePort}"

BACKUP_TIME=`date "+%F_%H-%M-%S"`

rm -rf ${TargetBin}
rm -rf ${TargetBin}.tar.gz

log_info "go build ..."

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w " -o ${TargetBin} main.go

log_info "tar package  ..."
tar czvf ${TargetBin}.tar.gz ${TargetBin}

sshpass -p $Password scp -P ${RemotePort} ${TargetBin}.tar.gz ${RemoteServer}:/tmp
check_error "Copy ${TargetBin}.tar.gz to remote server:${RemoteServer}"


#log_info "Stop the service ${TargetBin}.service"
#sshpass -p $Password ssh -p ${RemotePort} ${RemoteServer} "sudo systemctl stop ${TargetBin}.service"
#
#log_info "Backup ${TargetBin}"
#sshpass -p $Password ssh -p ${RemotePort} ${RemoteServer} "cd /tmp && tar zxvf ${TargetBin}.tar.gz"
#sshpass -p $Password ssh -p ${RemotePort} ${RemoteServer} "sudo cp -f /opt/bin/${TargetBin} /opt/bin/${TargetBin}.${BACKUP_TIME}"
#sshpass -p $Password ssh -p ${RemotePort} ${RemoteServer} "sudo cp -f /tmp/${TargetBin} /opt/bin/${TargetBin}"
#
#log_info "Start the service ${TargetBin}.service"
#sshpass -p $Password ssh -p ${RemotePort} ${RemoteServer} "sudo systemctl start ${TargetBin}.service"
#check_error "Start the service ${TargetBin}.service"