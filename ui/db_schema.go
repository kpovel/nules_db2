package main

import (
	"fmt"
	"time"
)

type Server_Status string

const (
	enabled  Server_Status = "enabled"
	disabled Server_Status = "disabled"
)

type MQTT_Server struct {
	id_server uint
	url       string
	status    Server_Status
}

type Designation_Category string

const (
	Excellent Designation_Category = "Excellent"
	Fine      Designation_Category = "Fine"
	Moderate  Designation_Category = "Moderate"
	Poor      Designation_Category = "Poor"
	Very_Poor Designation_Category = "Very Poor"
	Severe    Designation_Category = "Severe"
)

type Category struct {
	ID_Category uint
	Designation Designation_Category
}

type Units string

const (
	Percent   Units = "%"
	MgPerM3   Units = "mg/m3"
	HPa       Units = "hPa"
	Celsius   Units = "Celsius"
	PPM       Units = "ppm"
	PPB       Units = "ppb"
	AQI       Units = "aqi"
	MkgPerM3  Units = "mkg/m3"
	N3vPerGod Units = "n3v/god"
)

type Measured_Unit struct {
	ID_Measured_Unit uint
	Title            string
	Unit             Units
}

type Optimal_Value struct {
	ID_Category      uint
	ID_Measured_Unit uint
	Bottom_Border    *int16
	Upper_Border     *int16
}

type Point struct {
	X float64
	Y float64
}

func (p *Point) Scan(value interface{}) error {
	strValue := string(value.([]byte))

	n, err := fmt.Sscanf(strValue, `(%f,%f)`, &p.X, &p.Y)
	if err != nil {
		return err
	}
	if n != 2 {
		return fmt.Errorf("did not scan enough items from point string: %d", n)
	}
	return nil
}

type Station struct {
	ID_Station    uint
	City          string
	Name          string
	Status        Server_Status
	ID_SaveEcoBot string
	ID_Server     *uint
	Coordinates   Point
}

type Measurement struct {
	ID_Measurement   uint
	Time             time.Time
	Value            float64
	ID_Station       uint
	ID_Measured_Unit uint
}

type MQTT_Unit struct {
	ID_Station       uint
	ID_Measured_Unit uint
	Message          string
	Order            *uint
}

type User_Data struct {
	ID_User  uint
	Login    string
	Password string
}

type Favorite struct {
	ID_User    uint
	ID_Station uint
}
