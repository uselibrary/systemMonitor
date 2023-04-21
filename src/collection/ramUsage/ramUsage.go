package ramUsage

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// function to get ram usage
// return: used ram (M) and usage percentage

func GetRamUsage() (float64, float64) {
	// get ram usage
	out, err := exec.Command("bash", "-c", "free -m | grep Mem | awk '{print $3, $2}'").Output()
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	// parse output
	split := strings.Split(strings.TrimSpace(string(out)), " ")
	usedRam, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	totalRam, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	usagePercentage := usedRam / totalRam * 100
	// keep 2 decimal places
	usagePercentage = float64(int(usagePercentage*100)) / 100
	return usedRam, usagePercentage
}
