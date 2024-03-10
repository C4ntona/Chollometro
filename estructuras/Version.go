package estructuras

import (
	"strconv"
	"time"
)

type Version struct {
	version   float32
	timeStamp string
}

func NewVersion() *Version {
	return &Version{1.0, getTimeStampActual()}
}

func getTimeStampActual() string {
	return strconv.Itoa(time.Now().Nanosecond())
}

func (v *Version) getVersion() float32 {
	return v.version
}

func (v *Version) setNewVersion() float32 {
	return v.version + 1
}

func (v *Version) getTimeStamp() string {
	return v.timeStamp
}
