package lmsensors

import (
	"path/filepath"
	"strconv"
)

var _ Sensor = &FanSensor{}

// A FanSensor is a Sensor that detects fan speeds in rotations per minute.
type FanSensor struct {
	//Path of the sensor
	Path string

	//path of the sensor input
	InputPath string

	// The name of the sensor.
	Name string

	// Whether or not the fan speed is below the minimum threshold.
	Alarm bool

	// Whether or not the fan will sound an audible alarm when fan speed is
	// below the minimum threshold.
	Beep bool

	// The input fan speed, in rotations per minute, indicated by the sensor.
	Input int

	// The low threshold fan speed, in rotations per minute, indicated by the
	// sensor.
	Minimum int
}

func (s *FanSensor) name() string             { return s.Name }
func (s *FanSensor) setName(name string)      { s.Name = name }
func (s *FanSensor) GetPath() string          { return s.Path }
func (s *FanSensor) GetInputPath() string     { return s.InputPath }
func (s *FanSensor) setPath(path string)      { s.Path = path }
func (s *FanSensor) setInputPath(path string) { s.InputPath = path }

func (s *FanSensor) parse(raw map[string]SensorInfo) error {
	for k, v := range raw {
		s.setPath(filepath.Dir(v.Path))
		switch k {
		case "input", "min":
			i, err := strconv.Atoi(v.Value)
			if err != nil {
				return err
			}

			switch k {
			case "input":
				s.Input = i
				s.setInputPath(v.Path)
			case "min":
				s.Minimum = i
			}
		case "alarm":
			s.Alarm = v.Value != "0"
		case "beep":
			s.Beep = v.Value != "0"
		}
	}

	return nil
}
