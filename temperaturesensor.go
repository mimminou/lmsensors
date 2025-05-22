package lmsensors

import (
	"path/filepath"
	"strconv"
)

// A TemperatureSensorType is value that indicates the type of a
// TemperatureSensor.
type TemperatureSensorType int

// All possible TemperatureSensorType constants.
const (
	TemperatureSensorUnknown             TemperatureSensorType = 0
	TemperatureSensorTypePIICeleronDiode TemperatureSensorType = 1
	TemperatureSensorType3904Transistor  TemperatureSensorType = 2
	TemperatureSensorTypeThermalDiode    TemperatureSensorType = 3
	TemperatureSensorTypeThermistor      TemperatureSensorType = 4
	TemperatureSensorTypeAMDAMDSI        TemperatureSensorType = 5
	TemperatureSensorTypeIntelPECI       TemperatureSensorType = 6
)

var _ Sensor = &TemperatureSensor{}

// A TemperatureSensor is a Sensor that detects temperatures in degrees
// Celsius.
type TemperatureSensor struct {
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

	// Whether or not the sensor will sound an audible alarm if an alarm
	// is triggered.
	Beep bool

	// The type of sensor used to report tempearatures.
	Type TemperatureSensorType

	// The input temperature, in degrees Celsius, indicated by the sensor.
	Input float64

	// A high threshold temperature, in degrees Celsius, indicated by the
	// sensor.
	High float64

	// A critical threshold temperature, in degrees Celsius, indicated by the
	// sensor.
	Critical float64

	// Whether or not the temperature is past the critical threshold.
	CriticalAlarm bool
}

func (s *TemperatureSensor) name() string             { return s.Name }
func (s *TemperatureSensor) setName(name string)      { s.Name = name }
func (s *TemperatureSensor) GetPath() string          { return s.Path }
func (s *TemperatureSensor) GetInputPath() string     { return s.InputPath }
func (s *TemperatureSensor) setPath(path string)      { s.Path = path }
func (s *TemperatureSensor) setInputPath(path string) { s.InputPath = path }

func (s *TemperatureSensor) parse(raw map[string]SensorInfo) error {
	for k, v := range raw {
		s.setPath(filepath.Dir(v.Path))
		switch k {
		case "input", "crit", "max":
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
			case "crit":
				s.Critical = f
			case "max":
				s.High = f
			}
		case "alarm":
			s.Alarm = v.Value != "0"
		case "beep":
			s.Beep = v.Value != "0"
		case "type":
			t, err := strconv.Atoi(v.Value)
			if err != nil {
				return err
			}

			s.Type = TemperatureSensorType(t)
		case "crit_alarm":
			s.CriticalAlarm = v.Value != "0"
		case "label":
			s.Label = v.Value
		}
	}

	return nil
}
