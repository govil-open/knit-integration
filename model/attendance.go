package model

import (
	"knit-integration/model/enum"
	"time"
)

type Attendance struct {
	Date           time.Time   `json:"date"`
	ShiftStartTime time.Time   `json:"shiftStartTime"`
	ShiftEndTime   time.Time   `json:"shiftEndTime"`
	TotalHours     float64     `json:"totalHours"`
	BreakHours     float64     `json:"breakHours"`
	EffectiveHours float64     `json:"effectiveHours"`
	FirstInTime    time.Time   `json:"firstInTime"`
	LastOutTime    time.Time   `json:"lastOutTime"`
	Status         enum.Status `json:"status"`
}
