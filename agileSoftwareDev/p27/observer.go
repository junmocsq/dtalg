package p27

type Observer interface {
	Update(float642 float64)
}

type TemperatureObserver struct {
}

func NewTemperatureObserver() *TemperatureObserver {
	return &TemperatureObserver{}
}

func (t *TemperatureObserver) Update(temp float64) {
	NewScreen().DisplayTemp(temp)
}

type PressureObserver struct {
}

func NewPressureObserver() *PressureObserver {
	return &PressureObserver{}
}

func (t *PressureObserver) Update(temp float64) {
	NewScreen().DisplayPressure(temp)
}

type PressureTendObserver struct {
}

func NewPressureTendObserver() *PressureTendObserver {
	return &PressureTendObserver{}
}

func (t *PressureTendObserver) Update(temp float64) {
	NewScreen().DisplayPressureTend(temp)
}
