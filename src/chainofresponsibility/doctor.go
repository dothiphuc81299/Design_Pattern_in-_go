package chainofresponsibility

import "fmt"

type Doctor struct {
	next Department
}

// Execute ..
func (d *Doctor) Execute(p *Patient) {
	if !p.isDoctorChecked {
		fmt.Println("patient already checked by doctor ")
		d.next.Execute(p)
		return
	}// else {
	// 	fmt.Println("Doctor is checking patient")
	// 	p.isDoctorChecked = true
	// 	d.next.Execute(p)
	// }

}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}
