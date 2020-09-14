package p27

type TemperatureSensor interface {
	Read() float64
}

type temperatureSensor1 struct {
	te, bps, bpt   Observable
	itsLastReading float64
}

func (t *temperatureSensor1) Read() float64 {
	return 0
}

type temperatureSensor2 struct {
}

func (t *temperatureSensor2) Read() float64 {

	return 0
}
