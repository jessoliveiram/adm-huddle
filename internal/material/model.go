package material

import "time"

type Measure struct {
	Temperature        float32
	Humidity           float32
	MacAddress         string
	TimeOfMeasurements time.Time
}

type Material struct {
	Id       float32
	Nome     float32
	Lote     string
	Validade time.Time
}
