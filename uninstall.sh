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

# uninstall systemMonitor
echo "正在卸载systemMonitor..."
rm -rf /usr/local/systemMonitor

echo "正在删除定时任务..."
# remove crontab
# the installtion using "echo "*/10 * * * * /usr/local/systemMonitor/systemMonitor -c /usr/local/systemMonitor/config.json >> /usr/local/systemMonitor/error.log 2>& 1" >> /var/spool/cron/root"
# so we need to remove the line which contains "systemMonitor"
sed -i '/systemMonitor/d' /var/spool/cron/root

echo "卸载完成！"