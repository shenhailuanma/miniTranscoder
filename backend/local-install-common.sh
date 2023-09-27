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
# [must]      $1 : server name
# example:
#       local-install-common.sh mtserver

if [ $# -lt 1 ]; then
    log_error "Input params number error, current number is $#"
    return 1
fi

TargetBin=$1

log_info "TargetBin: ${TargetBin}"

BACKUP_TIME=`date "+%F_%H-%M-%S"`

rm -rf ${TargetBin}

log_info "Stop the service ${TargetBin}.service"
systemctl stop ${TargetBin}.service

log_info "Backup ${TargetBin}"
tar zxvf ${TargetBin}.tar.gz

cp -f /opt/bin/${TargetBin} /opt/bin/${TargetBin}.${BACKUP_TIME}

cp -f /tmp/${TargetBin} /opt/bin/${TargetBin}

log_info "Start the service ${TargetBin}.service"
systemctl start ${TargetBin}.service