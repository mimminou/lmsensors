package lmsensors

import "path/filepath"

var _ Sensor = &IntrusionSensor{}

// An IntrusionSensor is a Sensor that detects when the machine's chassis
// has been opened.
type IntrusionSensor struct {
	//Path of the sensor
	Path string

	// The name of the sensor.
	Name string

	// Whether or not the machine's chassis has been opened, and the alarm
	// has been triggered.
	Alarm bool
}

func (s *IntrusionSensor) name() string        { return s.Name }
func (s *IntrusionSensor) setName(name string) { s.Name = name }
func (s *IntrusionSensor) GetPath() string     { return s.Path }
func (s *IntrusionSensor) setPath(path string) { s.Path = path }

func (s *IntrusionSensor) parse(raw map[string]SensorInfo) error {
	for k, v := range raw {
		s.setPath(filepath.Dir(v.Path))
		switch k {
		case "alarm":
			s.Alarm = v.Value != "0"
		}
	}

	return nil
}
