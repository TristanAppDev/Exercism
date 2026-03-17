package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

const initialHitPoints int = 10
const numberOfRolls int = 4
const numberOfDiceSides int = 6

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score - 10.0)) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := make([]int, numberOfRolls)
	modifier := 0

	for i := 0; i < numberOfRolls; i++ {
		rolls[i] = rand.Intn(numberOfDiceSides-1) + 1
		modifier += rolls[i]
	}

	modifier -= slices.Min(rolls)

	return modifier
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	consitution := Ability()
	return Character{
		Ability(),
		Ability(),
		consitution,
		Ability(),
		Ability(),
		Ability(),
		initialHitPoints + Modifier(consitution),
	}
}
