# systemMonitor

systemMonitor是一个基于Go语言的自托管轻量型服务器性能检测与告警软件。可以检测CPU、硬盘内存和流量的使用量，并可在超过(或恢复)设置的限制后，通过电报机器人发送通知。

A Self-hosted Lightweight System Monitor by Golang


# 效果演示
![notice](https://raw.githubusercontent.com/uselibrary/systemMonitor/main/assets/8cb18054e876f2fa70706.jpg)


# 安装方法

## 需要准备的事项

1. 事先准备Telegram Bot (即电报机器人)的密钥(token)和对话ID(chat id)，教程可以参考这里：[自建电报机器人/telegram bot实现消息推送](https://pa.ci/119.html)
2. 需要监控的硬盘的名称，如`dev/sda1`。具体可使用`df -h`查看，如下所示，`/`目录对应着`/dev/vda1`，则此服务器的硬盘是`/dev/vda1`
```
Filesystem      Size  Used Avail Use% Mounted on
udev            2.0G     0  2.0G   0% /dev
tmpfs           394M  484K  393M   1% /run
/dev/vda1        39G  2.8G   34G   8% /
tmpfs           2.0G     0  2.0G   0% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           394M     0  394M   0% /run/user/1000
```
3. 阈值设置细则。配置文件为`/usr/local/systemMonitor/config.json`，安装过程中将自动生成，也可以随后手动修改。安装过程中会要求输入：
```
请输入服务器名称: # 名称用于区别服务器，推送消息的时候也会使用
请输入Telegram Bot Token: # 即上述的密钥(token)
请输入Telegram Chat ID: # 即上述的对话ID(chat id)
请输入磁盘名称: # 即上述的硬盘的名称
请输入CPU使用率阈值: # 单个CPU的理论最大值是1.00，如果无负载则是0.00，一般推荐0.20
请输入磁盘使用率阈值: # 硬盘使用率，如果输入50，当硬盘使用率超过50%则会告警
请输入网络使用率阈值: # 月带宽使用率，单位为GB，如果以前没有安装过vnstat，则统计信息从安装此软件开始计算。
请输入内存使用率阈值: # 内存使用率，如果输入50，当内存使用率超过50%则会告警
```
以下是`config.json`文件的示例
```
{
    "name": "demo.domain.com",
    "telegram": {
        "token": "123456789:ABCD45-VCSIDUIC78VS78RN",
        "chat_id": "123456789"
    },
    "disk": "dev/sda1",
    "status": {
        "cpu": 0.2,
        "diskpercentage": 50,
        "network": 500,
        "memorypercentage": 50
    }
}
```
4. 此程序由`crontab`定时运行，每10分钟检查一次。修改检查频率，可以手动编辑`/var/spool/cron/crontabs/root`文件。

### 一键安装
```
wget --no-check-certificate -O install.sh https://raw.githubusercontent.com/uselibrary/systemMonitor/main/install.sh && chmod +x install.sh && bash install.sh
```

如果需要卸载，则运行
```
wget --no-check-certificate -O uninstall.sh https://raw.githubusercontent.com/uselibrary/systemMonitor/main/uninstall.sh && chmod +x uninstall.sh && bash uninstall.sh
```

