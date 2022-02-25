package model

type InitMeasured struct {
	MoistLimit      float64 `json:"moistLimit"`
	WaterLevelLimit float64 `json:"waterLevelLimit"`
}

type LiveMeasure struct {
	Moist float64 `json:"moist"`
	Hum   float64 `json:"hum"`
	Temp  float64 `json:"temp"`
}

type LiveNotify struct {
	Title  string `json:"title"`
	State  string `json:"state"`
	Action string `json:"action"`
}

type LiveControl struct {
	Restart   bool `json:"restart"`
	PumpState bool `json:"pumpState"`
}
