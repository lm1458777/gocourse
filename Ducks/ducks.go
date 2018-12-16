package main

import (
	"fmt"
	"github.com/lm1458777/Ducks/behaviors"
)

type Duck struct {
	name          string
	flyBehavior   behaviors.Flyer
	quackBehavior behaviors.Quacker
}

func (duck *Duck) Fly() {
	duck.flyBehavior.Fly()
}

func (duck *Duck) Quack() {
	duck.quackBehavior()
}

func PlayWithDuck(duck *Duck) {
	fmt.Println()
	fmt.Println("I'm", duck.name)
	duck.Fly()
	duck.Quack()
	duck.Fly()
}

func main() {
	d1 := Duck{"Duck #1", &behaviors.FlyWithWings{}, behaviors.QuackQuackBehavior}
	PlayWithDuck(&d1)

	d2 := Duck{"Duck #2", &behaviors.FlyWithRocketEngine{}, behaviors.SqueekQuackBehavior}
	PlayWithDuck(&d2)

	d3 := Duck{"Duck #3", &behaviors.FlyNoWay{}, behaviors.MuteQuackBehavior}
	PlayWithDuck(&d3)

	d3.flyBehavior = &behaviors.FlyWithRocketEngine{}
	d3.quackBehavior = behaviors.SqueekQuackBehavior
	PlayWithDuck(&d3)
	d3.quackBehavior = behaviors.MuteQuackBehavior
	PlayWithDuck(&d3)
}
