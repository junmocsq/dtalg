package p27

type Observable struct {
	arr []Observer
}

func NewObservable() *Observable {
	return &Observable{}
}
func (o *Observable) addObserver(ob Observer) {
	o.arr = append(o.arr, ob)
}

func (o *Observable) notifyObservers(temp float64) {
	for _, v := range o.arr {
		v.Update(temp)
	}
}
