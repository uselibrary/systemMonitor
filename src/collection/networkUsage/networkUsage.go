package networkUsage

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// function to get monthly bandwidth usage via vnstat
// return: total bandwidth usage, float64
func GetMonthlyBandwidthUsage() float64 {

	// get monthly bandwidth usage via vnstat
	monthlyBandwidthUsage, err := exec.Command("vnstat", "-m", "--oneline").Output()
	if err != nil {
		log.Fatal(err)
	}

	// the output of vnstat is like this:
	// 1;ens18;2023-04-19;267.21 MiB;719.30 KiB;267.91 MiB;40.27 kbit/s;2023-04;9.57 GiB;1.28 GiB;10.85 GiB;57.87 kbit/s;10.73 GiB;1.32 GiB;12.05 GiB

	// split the output by ";", the 11th element is the total bandwidth usage
	monthlyBandwidthUsageSplit := strings.Split(string(monthlyBandwidthUsage), ";")

	// the total bandwidth usage may end with " KiB", " MiB", " GiB" and " TiB", use the last 3 characters to determine the unit
	// then convert the total bandwidth usage to float64
	var monthlyBandwidthUsageFloat64 float64
	if monthlyBandwidthUsageSplit[10][len(monthlyBandwidthUsageSplit[10])-3:] == "KiB" {
		monthlyBandwidthUsageFloat64, err = strconv.ParseFloat(monthlyBandwidthUsageSplit[10][:len(monthlyBandwidthUsageSplit[10])-4], 64)
		if err != nil {
			log.Fatal(err)
		}
		// convert KiB to GiB
		monthlyBandwidthUsageFloat64 = monthlyBandwidthUsageFloat64 / 1024 / 1024
	}
	if monthlyBandwidthUsageSplit[10][len(monthlyBandwidthUsageSplit[10])-3:] == "MiB" {
		monthlyBandwidthUsageFloat64, err = strconv.ParseFloat(monthlyBandwidthUsageSplit[10][:len(monthlyBandwidthUsageSplit[10])-4], 64)
		if err != nil {
			log.Fatal(err)
		}
		monthlyBandwidthUsageFloat64 = monthlyBandwidthUsageFloat64 / 1024
	}
	if monthlyBandwidthUsageSplit[10][len(monthlyBandwidthUsageSplit[10])-3:] == "GiB" {
		monthlyBandwidthUsageFloat64, err = strconv.ParseFloat(monthlyBandwidthUsageSplit[10][:len(monthlyBandwidthUsageSplit[10])-4], 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	if monthlyBandwidthUsageSplit[10][len(monthlyBandwidthUsageSplit[10])-3:] == "TiB" {
		monthlyBandwidthUsageFloat64, err = strconv.ParseFloat(monthlyBandwidthUsageSplit[10][:len(monthlyBandwidthUsageSplit[10])-4], 64)
		if err != nil {
			log.Fatal(err)
		}
		monthlyBandwidthUsageFloat64 = monthlyBandwidthUsageFloat64 * 1024
	}

	// keep 2 decimal places
	monthlyBandwidthUsageFloat64 = float64(int(monthlyBandwidthUsageFloat64*100)) / 100

	return monthlyBandwidthUsageFloat64
}
