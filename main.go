package main

import (
	"fmt"
	"math/rand"
	"sort"
)


type GameState struct {
	RedArmy int
	BlueArmy int
}

func Roll() int {
	return int(rand.Uint32()) % 6 + 1
}

func (gs *GameState) Fight(red []int, blue []int) {
	sort.Ints(red)
	sort.Ints(blue)

	for i := 0; i < len(blue); i++ {
		if red[i] > blue[i] {
			gs.BlueArmy--
		} else {
			gs.RedArmy--
		}
	}
}

func (gs *GameState) TwoVOne() {
	red := []int{Roll(), Roll()}
	blue := []int{Roll()}

	gs.Fight(red, blue)
}

func (gs *GameState) TwoVTwo() {
	red := []int{Roll(), Roll()}
	blue := []int{Roll(), Roll()}

	gs.Fight(red, blue)
}

func (gs *GameState) ThreeVTwo() {
	red := []int{Roll(), Roll(), Roll()}
	blue := []int{Roll(), Roll()}

	gs.Fight(red, blue)
}

func (gs *GameState) ThreeVOne() {
	red := []int{Roll(), Roll(), Roll()}
	blue := []int{Roll()}

	gs.Fight(red, blue)
}

func (gs *GameState) Simulate(scenario func(), verbose bool) {
	i := 1
	for {
		if gs.RedArmy <=0 || gs.BlueArmy <= 0 {
			break
		}

		scenario()

		if verbose {
			fmt.Printf("Iteration: %d Red: %d, Blue: %d\n", i, gs.RedArmy, gs.BlueArmy)
		}
		i++
	}

	fmt.Printf("Red: %d, Blue: %d\n", gs.RedArmy, gs.BlueArmy)
	if gs.BlueArmy > gs.RedArmy {
		fmt.Println("Blue won!")
	} else {
		fmt.Println("Red won!")
	}
}

func main(){
	gs1 := GameState{RedArmy: 1000, BlueArmy: 1000}
	gs2 := GameState{RedArmy: 1000, BlueArmy: 1000}
	gs3 := GameState{RedArmy: 1000, BlueArmy: 1000}
	gs4 := GameState{RedArmy: 1000, BlueArmy: 1000}

	fmt.Println("TwoVOne")
	gs1.Simulate(gs1.TwoVOne, false)
	fmt.Println("TwoVTwo")
	gs2.Simulate(gs2.TwoVTwo, false)
	fmt.Println("ThreeVTwo")
	gs3.Simulate(gs3.ThreeVTwo, false)
	fmt.Println("ThreeVOne")
	gs4.Simulate(gs4.ThreeVOne, false)
}

