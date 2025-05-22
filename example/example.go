package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mimminou/lmsensors"
)

func main() {

	var cpuTempPath string
	scanner := lmsensors.New()
	stuff, err := scanner.Scan()
	if err != nil {
		panic(err)
	}
	for _, d := range stuff {
		for _, s := range d.Sensors {
			if s, ok := s.(*lmsensors.TemperatureSensor); ok {
				if s.Label == "Tctl" { //catch AMD Ryzen CPU Temp Sensor, change to match your cpu sensor
					cpuTempPath = s.InputPath //cache input path, this is to avoid scanning again which is extremely slow
				}
			}
		}
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// read cpu temp as bytes from filepath
				t, err := os.ReadFile(cpuTempPath)
				if err != nil {
					panic(err)
				}
				//trim space
				s := strings.TrimSpace(string(t))
				temp, err := strconv.ParseFloat(s, 64)
				if err != nil {
					panic(err)
				}
				temp = temp / 1000.0 //this is because the temp is scaled by 1000 in lmsensors
				fmt.Println(temp, "°C")
			}
		}
	}()
	select {}
}
