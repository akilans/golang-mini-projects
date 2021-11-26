package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// define port
const PORT string = ":8080"

// define system info type

type systemInfo struct {
	HostName             string  `json:"hostname"`
	TotalMemory          uint64  `json:"total_memory"`
	FreeMemory           uint64  `json:"free_memory"`
	MemoryUsedPercentage float64 `json:"memory_used_percentage"`
	Architecture         string  `json:"architecture"`
	OperationSystem      string  `json:"os"`
	NumberOfCpuCores     int     `json:"number_of_cpu_cores"`
	CpuUsedPercentage    float64 `json:"cpu_used_percentage"`
}

// main function starts here
func main() {

	// http://localhost:8080
	http.HandleFunc("/", getSysInfo)

	fmt.Printf("App is listening on %v\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	// stop the app is any error to start the server
	if err != nil {
		log.Fatal(err)
	}
}

func getSysInfo(w http.ResponseWriter, r *http.Request) {

	totalMemory, freeMemory, usedMemoryPercentage := getMemInfo()
	hostName, architecture, operationSystem := getHostInfo()
	cpuNumCores, cpuPercentage := getCpuInfo()
	// all system info as systemInfo struct
	sysInfo := systemInfo{
		hostName,
		totalMemory,
		freeMemory,
		usedMemoryPercentage,
		architecture,
		operationSystem,
		cpuNumCores,
		cpuPercentage,
	}

	// converting into bytes for writing as response
	sysInfoByte, err := json.Marshal(sysInfo)

	// capture error
	checkError(err)

	//write into response body
	w.Write(sysInfoByte)

}
