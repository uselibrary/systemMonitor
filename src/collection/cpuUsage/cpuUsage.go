package cpuUsage

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// function to get cpu usage
// return: cpu usage in float64
func GetCPUUsage() float64 {
	// get cpu usage
	out, err := exec.Command("bash", "-c", "top -bn1 | grep load | awk '{printf \"%.2f\", $(NF-2)}'").Output()
	if err != nil {
		log.Println(err)
		return 0
	}
	cpuUsage, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return cpuUsage
}