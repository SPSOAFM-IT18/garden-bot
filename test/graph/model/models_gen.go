// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type IrrigationHistory struct {
	ID             *int     `json:"id"`
	Timestamp      *string  `json:"timestamp"`
	WaterLevel     *float64 `json:"water_level"`
	WaterAmount    *float64 `json:"water_amount"`
	WaterOverdrawn *float64 `json:"water_overdrawn"`
}

type Measurement struct {
	ID             *int     `json:"id"`
	Timestamp      *string  `json:"timestamp"`
	Hum            *float64 `json:"hum"`
	Temp           *float64 `json:"temp"`
	Moist          *float64 `json:"moist"`
	WithIrrigation *bool    `json:"with_irrigation"`
}

type NewMeasurement struct {
	Hum            *float64 `json:"hum"`
	Temp           *float64 `json:"temp"`
	Moist          *float64 `json:"moist"`
	WithIrrigation *bool    `json:"with_irrigation"`
}

type NewSettings struct {
	LimitsTrigger      *bool    `json:"limits_trigger"`
	WaterLevelLimit    *float64 `json:"water_level_limit"`
	WaterAmountLimit   *float64 `json:"water_amount_limit"`
	MoistLimit         *float64 `json:"moist_limit"`
	ScheduledTrigger   *bool    `json:"scheduled_trigger"`
	HourRange          *int     `json:"hour_range"`
	Location           *string  `json:"location"`
	IrrigationDuration *int     `json:"irrigation_duration"`
	ChartType          *bool    `json:"chart_type"`
	Language           *bool    `json:"language"`
	Theme              *bool    `json:"theme"`
	Lat                *float64 `json:"lat"`
	Lon                *float64 `json:"lon"`
}

type Setting struct {
	ID                 *int     `json:"id"`
	LimitsTrigger      *bool    `json:"limits_trigger"`
	WaterLevelLimit    *float64 `json:"water_level_limit"`
	WaterAmountLimit   *float64 `json:"water_amount_limit"`
	MoistLimit         *float64 `json:"moist_limit"`
	ScheduledTrigger   *bool    `json:"scheduled_trigger"`
	HourRange          *int     `json:"hour_range"`
	Location           *string  `json:"location"`
	IrrigationDuration *int     `json:"irrigation_duration"`
	ChartType          *bool    `json:"chart_type"`
	Language           *bool    `json:"language"`
	Theme              *bool    `json:"theme"`
	Lat                *float64 `json:"lat"`
	Lon                *float64 `json:"lon"`
}
