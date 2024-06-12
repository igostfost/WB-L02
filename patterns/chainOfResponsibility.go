package main

import "fmt"

/*
Цепочка обязанностей — это поведенческий паттерн,
позволяющий передавать запрос по цепочке потенциальных обработчиков,
пока один из них не обработает запрос.
*/

// Patient - объукт пациента который пришел в поликлинику и по цепочке пойдет по отделениям реализуя паттерн
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	paymentDone       bool
}

// Department - интерфейс отедла поликлиники
type Department interface {
	execute(*Patient)
	setNext(Department)
}

// Reception - регистратура
type Reception struct {
	next Department
}

func (r *Reception) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Reception registration already done")
		r.next.execute(patient)
		return
	}
	fmt.Println("Reception registration done")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

// Doctor - Доктор
type Doctor struct {
	next Department
}

func (d *Doctor) execute(patient *Patient) {
	if patient.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(patient)
		return
	}
	fmt.Println("Doctor checkup done")
	patient.doctorCheckUpDone = true
	d.next.execute(patient)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

// Cashbox - касса
type Cashbox struct {
	next Department
}

func (c *Cashbox) execute(patient *Patient) {
	if patient.paymentDone {
		fmt.Println("Cashbox payment already done")
		return
	}
	fmt.Println("Cashbox payment done")
	patient.paymentDone = true
}

func (c *Cashbox) setNext(next Department) {
	c.next = next
}

// Клиентсикй код
func main() {

	patient := &Patient{name: "alex"}

	register := &Reception{}
	doctor := &Doctor{}
	cashbox := &Cashbox{}

	// Настраиваем цепочку обязанностей
	register.setNext(doctor)
	doctor.setNext(cashbox)

	// Запускаем цепочку
	register.execute(patient)
}
