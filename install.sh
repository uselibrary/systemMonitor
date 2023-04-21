#!/usr/bin/env bash

clear
echo "    ################################################"
echo "    #                                              #"
echo "    #         System Monitor Installation          #"
echo "    #                https://pa.ci                 #"
echo "    #          Version 0.1.0, 2023-04-21           #"
echo "    #                                              #"
echo "    ################################################"

# https://github.com/uselibrary/systemMonitor


# check root user, if not root, exit
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

# check if env is utf-8, if not, change to utf-8 temporarily
if [ "$LANG" != "en_US.UTF-8" ]; then
    export LANG=en_US.UTF-8
fi

# check architecture, if not x86_64, exit
if [ "$(uname -m)" != "x86_64" ]; then
    echo "暂时仅支持x86_64架构CPU，ARM等正在开发中"
    exit
fi

# get linux distribution and architecture
# distribution: 1. redhat/centos/rocky/almalinux 2. debian/ubuntu
if [ -f /etc/redhat-release ]; then
    distribution="redhat"
elif [ -f /etc/os-release ]; then
    distribution="debian"
else
    echo "暂时仅支持RedHat系列和Debian系列发行版"
    exit
fi

# check if vnsta, if not, update and install via package manager
if [ ! -f /usr/bin/vnstat ]; then
    if [ "$distribution" == "redhat" ]; then
        yum update -y
        yum install -y vnstat
    elif [ "$distribution" == "debian" ]; then
        apt update -y
        apt install -y vnstat
    fi
fi

# make dir
mkdir -p /usr/local/systemMonitor
cd /usr/local/systemMonitor
# check th lastest version of systemMonitor by github api
version=$(curl -s https://api.github.com/repos/uselibrary/systemMonitor/releases/latest | grep "tag_name" | cut -d '"' -f 4)
# build download url and download systemMonitor
# example:https://github.com/uselibrary/systemMonitor/releases/download/v0.0.1/systemMonitor.zip
downloadUrl="https://github.com/uselibrary/systemMonitor/releases/download/$version/systemMonitor.zip"
wget -O systemMonitor.zip $downloadUrl
unzip systemMonitor.zip
rm -f systemMonitor.zip
chmod +x systemMonitor
# add systemMonitor to crontab
# example: */10 * * * * /usr/local/systemMonitor/systemMonitor -c /usr/local/systemMonitor/config.json >> /usr/local/systemMonitor/error.log 2>& 1
echo "*/10 * * * * /usr/local/systemMonitor/systemMonitor -c /usr/local/systemMonitor/config.json >> /usr/local/systemMonitor/error.log 2>& 1" >> /var/spool/cron/root

touch error.log

# wget config.json, https://raw.githubusercontent.com/uselibrary/systemMonitor/main/data/config.json
wget -O config.json https://raw.githubusercontent.com/uselibrary/systemMonitor/main/data/config.json

# config json, example:
# {
#     "name": "demo.domain.com",
#     "telegram": {
#         "token": "123456789:ABCD45-VCSIDUIC78VS78RN",
#         "chat_id": "123456789"
#     },
#     "disk": "dev/sda1",
#     "status": {
#         "cpu": 0.5,
#         "diskpercentage": 30,
#         "network": 500,
#         "memorypercentage": 50
#     }
# }

# ask user to input name, telegram token, telegram chat_id, disk
read -p "请输入服务器名称：" name
read -p "请输入Telegram Bot Token：" token
read -p "请输入Telegram Chat ID：" chat_id
read -p "请输入磁盘名称：" disk
read -p "请输入CPU使用率阈值：" cpu
read -p "请输入磁盘使用率阈值：" diskpercentage
read -p "请输入网络使用率阈值：" network
read -p "请输入内存使用率阈值：" memorypercentage

# replace config.json
sed -i "s/demo.domain.com/$name/g" config.json
sed -i "s/123456789:ABCD45-VCSIDUIC78VS78RN/$token/g" config.json
sed -i "s/123456789/$chat_id/g" config.json
sed -i "s/dev\/sda1/$disk/g" config.json
sed -i "s/0.5/$cpu/g" config.json
sed -i "s/30/$diskpercentage/g" config.json
sed -i "s/500/$network/g" config.json
sed -i "s/50/$memorypercentage/g" config.json



