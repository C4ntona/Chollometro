package estructuras

type Version struct {
	version   int
	timeStamp string
}

func (v *Version) getVersion() int {
	return v.version
}

func (v *Version) getTimeStamp() string {
	return v.timeStamp
}
