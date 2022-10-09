package ch07

import "fmt"

func Sub() {
	list_7_2()
}

type Modem struct{}

func (m *Modem) Dial() {

}

func (m *Modem) Hangup() {

}

func (m *Modem) Sender() {

}

func (m *Modem) Recv() {
	fmt.Println("Modem implements Receiver interface.")
}

type Receiver interface {
	Recv()
}

// check if type implements an interface.
var _ Receiver = (*Modem)(nil)

func list_7_2() {
	m := &Modem{}
	m.Recv()
}
