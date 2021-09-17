package main

import "fmt"

type patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}

//department текущая служба
type department interface {
    execute(*patient)
    setNext(department)
}

//cashier касса
type cashier struct {
    next department
}

//medical терапевт
type medical struct {
    next department
}

//doctor специализированный врач
type doctor struct {
    next department
}

//reception стойка регистрации
type reception struct {
    next department
}

func main() {

    cashier := &cashier{}

    //Set next for medical department
    medical := &medical{}
    medical.setNext(cashier)

    //Set next for doctor department
    doctor := &doctor{}
    doctor.setNext(medical)

    //Set next for reception department
    reception := &reception{}
    reception.setNext(doctor)
    
    patient := &patient{name: "abc"}
    //Patient visiting
    reception.execute(patient)
}

//execute произвести действие над пациентом
func (d *doctor) execute(p *patient) {
    if p.doctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.doctorCheckUpDone = true
    d.next.execute(p)
}

//execute произвести действие над пациентом
func (m *medical) execute(p *patient) {
    if p.medicineDone {
        fmt.Println("Medicine already given to patient")
        m.next.execute(p)
        return
    }
    fmt.Println("Medical giving medicine to patient")
    p.medicineDone = true
    m.next.execute(p)
}

//execute произвести действие над пациентом
func (c *cashier) execute(p *patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

//execute произвести действие над пациентом
func (r *reception) execute(p *patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.registrationDone = true
    r.next.execute(p)
}

//setNext отправить дальше
func (c *cashier) setNext(next department) {
    c.next = next
}

//setNext отправить дальше
func (m *medical) setNext(next department) {
    m.next = next
}

//setNext отправить дальше
func (d *doctor) setNext(next department) {
    d.next = next
}

//setNext отправить дальше
func (r *reception) setNext(next department) {
    r.next = next
}