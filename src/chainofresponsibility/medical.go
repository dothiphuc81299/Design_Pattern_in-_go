package chainofresponsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if !p.isMedicineProvied {
		fmt.Println("patient already provided medicine")
		m.next.Execute(p)
		return
	} // else {
	// 	fmt.Println("we are providing medicine to patient")
	// 	p.isMedicineProvied = true
	// 	m.next.Execute(p)
	// }
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}
