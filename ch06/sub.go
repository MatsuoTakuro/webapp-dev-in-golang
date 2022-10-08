package ch06

import "fmt"

func Sub() {
	list_6_3()
	list_6_4()
}

type Dog struct{}

func (d *Dog) Bark() string { return "Bow" }

type BullDog struct{ Dog }

type ShibaInu struct{ Dog }

func (d *ShibaInu) Bark() string { return "wan" }

func DogVoice(d *Dog) string { return d.Bark() }

func list_6_3() {
	bd := &BullDog{}
	fmt.Println(bd.Bark())
	si := &ShibaInu{}
	fmt.Println(si.Bark())

	// cannot use bd (variable of type *BullDog) as *Dog value in argument to DogVoice
	// fmt.Println(DogVoice(bd))
	// cannot use si (variable of type *ShibaInu) as *Dog value in argument to DogVoice
	// fmt.Println(DogVoice(si))
}

type Person struct {
	Name string
	Age  int
}

type Japanse struct {
	Person
	MyNumber int
}

func Hello(p Person) {
	fmt.Println("Hello,", p.Name)
}

func list_6_4() {
	j := &Japanse{
		Person: Person{
			Name: "Taro",
			Age:  20,
		},
		MyNumber: 100,
	}

	Hello(j.Person)
}
