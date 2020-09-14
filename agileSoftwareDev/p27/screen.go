package p27

import "fmt"

var screen = &Screen{}

type Screen struct {
}

func NewScreen() *Screen {
	return screen
}

func (s *Screen) Create() {
	ob := NewObservable()
	ob.addObserver(NewTemperatureObserver())
	ob.addObserver(NewPressureObserver())
}

// 展示温度
func (s *Screen) DisplayTemp(temp float64) {
	fmt.Println("temperature：", temp)
}

// 展示压力
func (s *Screen) DisplayPressure(temp float64) {
	fmt.Println("pressure：", temp)
}

// 展示
func (s *Screen) DisplayPressureTend(temp float64) {
	fmt.Println("pressure tend：", temp)
}
