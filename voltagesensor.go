package lmsensors

import (
	"path/filepath"
	"strconv"
)

var _ Sensor = &VoltageSensor{}

// A VoltageSensor is a Sensor that detects voltage.
type VoltageSensor struct {
	//Path of the sensor
	Path string

	//path of the sensor input
	InputPath string

	// The name of the sensor.
	Name string

	// A label that describes what the sensor is monitoring.  Label may be
	// empty.
	Label string

	// Whether or not the sensor has an alarm triggered.
	Alarm bool

	// Whether or not the sensor will sound an audible alarm when an alarm
	// is triggered.
	Beep bool

	// The input voltage indicated by the sensor.
	Input float64

	// The maximum voltage threshold indicated by the sensor.
	Maximum float64
}

func (s *VoltageSensor) name() string             { return s.Name }
func (s *VoltageSensor) setName(name string)      { s.Name = name }
func (s *VoltageSensor) GetPath() string          { return s.Path }
func (s *VoltageSensor) setPath(path string)      { s.Path = path }
func (s *VoltageSensor) GetInputPath() string     { return s.InputPath }
func (s *VoltageSensor) setInputPath(path string) { s.InputPath = path }

func (s *VoltageSensor) parse(raw map[string]SensorInfo) error {
	for k, v := range raw {
		s.setPath(filepath.Dir(v.Path))
		switch k {
		case "input", "max":
			f, err := strconv.ParseFloat(v.Value, 64)
			if err != nil {
				return err
			}

			// Raw temperature values are scaled by 1000
			f /= 1000

			switch k {
			case "input":
				s.Input = f
				s.setInputPath(v.Path)
			case "max":
				s.Maximum = f
			}
		case "alarm":
			s.Alarm = v.Value != "0"
		case "beep":
			s.Beep = v.Value != "0"
		case "label":
			s.Label = v.Value
		}
	}

	return nil
}
