package diskUsage

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// function to get disk usage
// variable: disk name, such as /dev/vda1 or /dev/sda1
// return: used space, percentage
func GetDiskUsage(disk string) (float64, float64) {
	// get disk usage, use df -B M to get the result in MB, and use grep to get the result of the disk
	// command: df -B M | grep disk | awk '{print $3, $5}'
	command := "df -B M | grep " + disk + " | awk '{print $3, $5}'"
	//out, err := exec.Command("bash", "-c", "df -B M | grep /dev/sda1 | awk '{print $3, $5}'").Output()
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	// parse output
	split := strings.Split(strings.TrimSpace(string(out)), " ")
	usedSpace, err := strconv.ParseFloat(strings.TrimSuffix(split[0], "M"), 64)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	// convert to GB, and keep 2 decimal places
	usedSpace = float64(int(usedSpace/1024*100)) / 100
	usagePercentage, err := strconv.ParseFloat(strings.TrimSuffix(split[1], "%"), 64)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	return usedSpace, usagePercentage
}
