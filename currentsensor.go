package lmsensors

import (
	"path/filepath"
	"strconv"
)

var _ Sensor = &CurrentSensor{}

// A CurrentSensor is a Sensor that detects current in Amperes.
type CurrentSensor struct {
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

	// The input current, in Amperes, indicated by the sensor.
	Input float64

	// The maximum current threshold, in Amperes, indicated by the sensor.
	Maximum float64

	// The critical current threshold, in Amperes, indicated by the sensor.
	Critical float64
}

func (s *CurrentSensor) name() string             { return s.Name }
func (s *CurrentSensor) setName(name string)      { s.Name = name }
func (s *CurrentSensor) GetPath() string          { return s.Path }
func (s *CurrentSensor) setPath(path string)      { s.Path = path }
func (s *CurrentSensor) GetInputPath() string     { return s.InputPath }
func (s *CurrentSensor) setInputPath(path string) { s.InputPath = path }

func (s *CurrentSensor) parse(raw map[string]SensorInfo) error {
	for k, v := range raw {
		s.setPath(filepath.Dir(v.Path))
		switch k {
		case "crit", "input", "max":
			f, err := strconv.ParseFloat(v.Value, 64)
			if err != nil {
				return err
			}

			// Raw current values are scaled by 1000
			f /= 1000

			switch k {
			case "crit":
				s.Critical = f
			case "input":
				s.Input = f
				s.setInputPath(v.Path)
			case "max":
				s.Maximum = f
			}
		case "alarm":
			s.Alarm = v.Value != "0"
		case "label":
			s.Label = v.Value
		}
	}

	return nil
}
