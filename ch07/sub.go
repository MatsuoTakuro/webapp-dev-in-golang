package ch07

import (
	"errors"
	"fmt"
)

func Sub() {
	// list_7_2()
	list_7_3()
}

type Modem struct{}

func (m *Modem) Dial() {}

func (m *Modem) Hangup() {}

func (m *Modem) Sender() {}

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

type MyErr struct{}

func (me *MyErr) Error() string { return "" }

var _ error = (*MyErr)(nil)

func Apply() error {
	var err *MyErr = nil
	if err != nil {
		return errors.New("err is not nil")
	}
	fmt.Println(err)
	return err
	// return nil
}

func Apply2() error {
	var err error = nil
	if err != nil {
		return errors.New("err is not nil")
	}
	fmt.Println(err)
	return nil
}

func list_7_3() {
	fmt.Println(Apply() == nil)
	fmt.Println(Apply2() == nil)
}
