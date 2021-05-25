package observer

import "fmt"

type Item struct {
	observerList []Observer
	Name         string
	Instock      bool
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Println("item %s is now in stock", i.Name)
	i.Instock = true
	i.NotifyAll()
}

func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) DeRegister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.Name)
	}
}

func removeFromSlice(observerList []Observer, observerRemove Observer) []Observer {
	l := len(observerList)
	for i, observer := range observerList {
		if observerRemove.GetID() == observer.GetID() {
			observerList[l-1], observerList[i] = observerList[i], observerList[l-1]
			return observerList[:l-1]
		}
	}
	return observerList
}
