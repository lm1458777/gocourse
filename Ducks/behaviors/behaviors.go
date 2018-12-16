package behaviors

import "fmt"

type Flyer interface {
	Fly()
}

type counter struct {
	flyCounter int
}

type FlyWithWings counter

func (f *FlyWithWings) Fly() {
	f.flyCounter++
	fmt.Printf("I'm flying with wings. Flight number: %v\n", f.flyCounter)
}

type FlyWithRocketEngine counter

func (f *FlyWithRocketEngine) Fly() {
	f.flyCounter++
	fmt.Printf("I'm flying with rocket engine. Flight number: %v\n", f.flyCounter)
}

type FlyNoWay struct {
}

func (f FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

type Quacker func()

func QuackQuackBehavior() {
	fmt.Println("Quack! Quack!")
}

func SqueekQuackBehavior() {
	fmt.Println("Squeek!")
}

func MuteQuackBehavior() {
}
