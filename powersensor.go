package lmsensors

import (
	"path/filepath"
	"strconv"
	"time"
)

var _ Sensor = &PowerSensor{}

// A PowerSensor is a Sensor that detects average electrical power consumption
// in watts.
type PowerSensor struct {
	//Path of the sensor
	Path string

	// The name of the sensor.
	Name string

	// The average electrical power consumption, in watts, indicated
	// by the sensor.
	Average float64

	// The interval of time over which the average electrical power consumption
	// is collected.
	AverageInterval time.Duration

	// Whether or not this sensor has a battery.
	Battery bool

	// The model number of the sensor.
	ModelNumber string

	// Miscellaneous OEM information about the sensor.
	OEMInfo string

	// The serial number of the sensor.
	SerialNumber string
}

func (s *PowerSensor) name() string        { return s.Name }
func (s *PowerSensor) setName(name string) { s.Name = name }
func (s *PowerSensor) GetPath() string     { return s.Path }
func (s *PowerSensor) setPath(path string) { s.Path = path }
func (s *PowerSensor) parse(raw map[string]SensorInfo) error {
	for k, v := range raw {
		s.setPath(filepath.Dir(v.Path))
		switch k {
		case "average":
			f, err := strconv.ParseFloat(v.Value, 64)
			if err != nil {
				return err
			}

			// Raw temperature values are scaled by one million
			f /= 1000000
			s.Average = f
		case "average_interval":
			// Time values in milliseconds
			d, err := time.ParseDuration(v.Value + "ms")
			if err != nil {
				return err
			}

			s.AverageInterval = d
		case "is_battery":
			s.Battery = v.Value != "0"
		case "model_number":
			s.ModelNumber = v.Value
		case "oem_info":
			s.OEMInfo = v.Value
		case "serial_number":
			s.SerialNumber = v.Value
		}
	}

	return nil
}
