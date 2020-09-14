package p27

func ExampleScreen_DisplayTemp() {
	s := &Screen{}
	s.DisplayTemp(20.3)
	// output:
	// temperature： 20.3
}

func ExampleScreen_DisplayPressure() {
	s := &Screen{}
	s.DisplayPressure(20.2)
	// output:
	// pressure： 20.2
}
