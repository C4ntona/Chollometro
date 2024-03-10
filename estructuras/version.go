package estructuras

import (
	"strconv"
	"time"
)

type Version struct {
	Version   float32 `json:"version"`
	TimeStamp string  `json:"timeStamp"`
}

func NewVersion() *Version {
	return &Version{1.0, getTimeStampActual()}
}

func getTimeStampActual() string {
	return strconv.Itoa(time.Now().Nanosecond())
}

func (v *Version) GetVersion() float32 {
	return v.Version
}

func (v *Version) SetNewVersion() float32 {
	return v.Version + 1
}

func (v *Version) GetTimeStamp() string {
	return v.TimeStamp
}
