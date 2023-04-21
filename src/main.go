package main

import (
	// "fmt"
	"path/filepath"
	"strconv"
	"systemMonitor/src/arguments"
	"systemMonitor/src/collection/cpuUsage"
	"systemMonitor/src/collection/diskUsage"
	"systemMonitor/src/collection/networkUsage"
	"systemMonitor/src/collection/ramUsage"
	"systemMonitor/src/levelDB"
	"systemMonitor/src/readJson"
	"systemMonitor/src/teleBot"
)

// function to compare current usage with limit
// variable: current usage, limit
// send message to telegram if current usage is larger than limit
func compareUsage(levelDBPath string, currentUsage float64, limit float64, serverName string, telegramToken string, telegramChatID string, usage float64, usageType string) {
	// check if currentUsage is larger than limit
	if currentUsage >= limit {
		// fmt.Println(currentUsage, limit)
		switch usageType {
		case "CPU":
			// read the value of the key, if the value is False, send message to telegram, otherwise do nothing
			value := levelDB.GetLevelDBValue(levelDBPath, "CPU")
			if value == "False" {
				msg := "CPU超限! 服务器: " + serverName + " 的CPU使用率超过 " + strconv.FormatFloat(limit, 'f', 2, 64) + " 的限制。目前CPU的使用率是 " + strconv.FormatFloat(usage, 'f', 2, 64) + " 。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "CPU", "True")
			}
		case "Disk":
			value := levelDB.GetLevelDBValue(levelDBPath, "Disk")
			if value == "False" {
				msg := "硬盘超限! 服务器: " + serverName + " 的硬盘使用率为 " + strconv.FormatFloat(currentUsage, 'f', 2, 64) + "%" + " 已超过 " + strconv.FormatFloat(limit, 'f', 2, 64) + "%" + " 的限制。目前硬盘已使用了 " + strconv.FormatFloat(usage, 'f', 2, 64) + " GB。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "Disk", "True")
			}
		case "RAM":
			value := levelDB.GetLevelDBValue(levelDBPath, "RAM")
			if value == "False" {
				msg := "内存超限! 服务器: " + serverName + " 的内存使用率为 " + strconv.FormatFloat(currentUsage, 'f', 2, 64) + "%" + " 已超过 " + strconv.FormatFloat(limit, 'f', 2, 64) + "%" + " 的限制。目前内存已使用了 " + strconv.FormatFloat(usage, 'f', 2, 64) + " MB。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "RAM", "True")
			}
		case "Network":
			value := levelDB.GetLevelDBValue(levelDBPath, "Network")
			if value == "False" {
				msg := "带宽超限! 服务器: " + serverName + " 的带宽使用量超过 " + strconv.FormatFloat(limit, 'f', 2, 64) + " GB" + " 的限制。目前带宽的使用量是 " + strconv.FormatFloat(usage, 'f', 2, 64) + " GB。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "Network", "True")
			}
		}
	} else {
		switch usageType {
		case "CPU":
			// fmt.Println("CPU is smaller than limit")
			value := levelDB.GetLevelDBValue(levelDBPath, "CPU")
			if value == "True" {
				msg := "CPU恢复! 服务器: " + serverName + " 的CPU使用率已低于 " + strconv.FormatFloat(limit, 'f', 2, 64) + " 的限制。目前CPU的使用率是 " + strconv.FormatFloat(usage, 'f', 2, 64) + " 。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "CPU", "False")
			}
		case "Disk":
			// fmt.Println("Disk is smaller than limit")
			value := levelDB.GetLevelDBValue(levelDBPath, "Disk")
			if value == "True" {
				msg := "硬盘恢复! 服务器: " + serverName + " 的硬盘使用率是 " + strconv.FormatFloat(currentUsage, 'f', 2, 64) + "%" + " 已低于 " + strconv.FormatFloat(limit, 'f', 2, 64) + "%" + " 的限制。目前硬盘使用量是 " + strconv.FormatFloat(usage, 'f', 2, 64) + " GB。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "Disk", "False")
			}
		case "RAM":
			// fmt.Println("RAM is smaller than limit")
			value := levelDB.GetLevelDBValue(levelDBPath, "RAM")
			if value == "True" {
				msg := "内存恢复! 服务器: " + serverName + " 的内存使用率是 " + strconv.FormatFloat(currentUsage, 'f', 2, 64) + "%" + " 已低于 " + strconv.FormatFloat(limit, 'f', 2, 64) + "%" + " 的限制。目前内存使用量是 " + strconv.FormatFloat(usage, 'f', 2, 64) + " MB。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "RAM", "False")
			}
		case "Network":
			// fmt.Println("Network is smaller than limit")
			value := levelDB.GetLevelDBValue(levelDBPath, "Network")
			if value == "True" {
				msg := "带宽恢复! 服务器: " + serverName + " 的带宽使用量已低于 " + strconv.FormatFloat(limit, 'f', 2, 64) + " GB" + " 的限制。目前带宽的使用率是 " + strconv.FormatFloat(usage, 'f', 2, 64) + " GB。"
				teleBot.SendMessage(telegramToken, telegramChatID, msg)
				levelDB.UpdateLevelDBValue(levelDBPath, "Network", "False")
			}
		}
	}

}

// main function
func main() {

	configPath := arguments.GetCommandArguments()
	// fmt.Println("config file path:", configPath)
	levelDBPath := filepath.Join(filepath.Dir(configPath), "levelDB")
	// fmt.Println("levelDB path:", levelDBPath)

	// check if levelDB exist, if not, create one and write inital file to levelDB
	if !levelDB.CheckLevelDB(levelDBPath) {
		levelDB.CreateLevelDB(levelDBPath)
		// write initial value to levelDB
		levelDB.WriteLevelDB(levelDBPath, "CPU", "False")
		levelDB.WriteLevelDB(levelDBPath, "Disk", "False")
		levelDB.WriteLevelDB(levelDBPath, "RAM", "False")
		levelDB.WriteLevelDB(levelDBPath, "Network", "False")
	}

	// read config file
	config := readJson.ReadConfig(configPath)
	serverName := config.Name
	telegramToken := config.Telegram.Token
	telegramChatID := config.Telegram.ChatID
	diskName := config.Disk
	cpuUsageLimit := config.Status.CPU
	diskUsageLimit := config.Status.DiskPercentage
	networkUsageLimit := config.Status.Network
	ramUsageLimit := config.Status.MemoryPercentage

	// print the config file
	// fmt.Println("server name:", serverName)
	// fmt.Println("telegram token:", telegramToken)
	// fmt.Println("telegram chat id:", telegramChatID)
	// fmt.Println("disk name:", diskName)
	// fmt.Println("cpu usage limit:", cpuUsageLimit)
	// fmt.Println("disk usage limit:", diskUsageLimit)
	// fmt.Println("network usage limit:", networkUsageLimit)
	// fmt.Println("ram usage limit:", ramUsageLimit)

	// fmt.Println("------------------------")

	// get current disk usage
	currentDiskUsage, currentDiskUsagePercentage := diskUsage.GetDiskUsage(diskName)
	// fmt.Println("current disk usage:", currentDiskUsage)
	// fmt.Println("current disk usage percentage:", currentDiskUsagePercentage)

	// get current cpu usage
	currentCPUUsage := cpuUsage.GetCPUUsage()
	// fmt.Println("current cpu usage:", currentCPUUsage)

	// get current ram usage
	currentRAMUsage, currentRAMUsagePercentage := ramUsage.GetRamUsage()
	// fmt.Println("current ram usage:", currentRAMUsage)
	// fmt.Println("current ram usage percentage:", currentRAMUsagePercentage)

	// get current network usage
	currentNetworkUsage := networkUsage.GetMonthlyBandwidthUsage()
	// fmt.Println("current network usage:", currentNetworkUsage)

	// check if currentCPUUsage is larger than config.Status.CPU
	compareUsage(levelDBPath, currentCPUUsage, cpuUsageLimit, serverName, telegramToken, telegramChatID, currentCPUUsage, "CPU")
	// check if currentDiskUsagePercentage is larger than config.Status.DiskPercentage
	compareUsage(levelDBPath, currentDiskUsagePercentage, diskUsageLimit, serverName, telegramToken, telegramChatID, currentDiskUsage, "Disk")
	// check if currentRAMUsage is larger than config.Status.MemoryPercentage
	compareUsage(levelDBPath, currentRAMUsagePercentage, ramUsageLimit, serverName, telegramToken, telegramChatID, currentRAMUsage, "RAM")
	// check if currentNetworkUsage is larger than config.Status.Network
	compareUsage(levelDBPath, currentNetworkUsage, networkUsageLimit, serverName, telegramToken, telegramChatID, currentNetworkUsage, "Network")

}
