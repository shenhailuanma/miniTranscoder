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

log_info "npm run build ..."

npm run build

log_info "tar package  ..."
tar czvf ${TargetBin}.tar.gz dist

sshpass -p $Password scp -P ${RemotePort} ${TargetBin}.tar.gz ${RemoteServer}:/tmp
check_error "Copy ${TargetBin}.tar.gz to remote server:${RemoteServer}"

sshpass -p $Password ssh -p ${RemotePort} ${RemoteServer} "cd /tmp && tar zxvf ${TargetBin}.tar.gz"
check_error "cd /tmp && tar zxvf ${TargetBin}.tar.gz"

sshpass -p $Password ssh -p ${RemotePort} ${RemoteServer} "sudo rsync -av /tmp/dist/ /opt/vod/www/ui/"
check_error "Copy ${TargetBin}.tar.gz to remote server:${RemoteServer}"
