package lasagna

import "strings"

const noodleGramsPerLayer int = 50
const sauceLitersPerLayer float64 = 0.2

// PreparationTime returns the total preparation time of the Lasagna recipe.
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// Quantities returns the quantities of the ingredients in the Lasagna recipe.
func Quantities(layers []string) (int, float64) {
	noodleCounter := 0
	sauceCounter := 0.0

	for _, layer := range layers {
		if strings.Contains(layer, "noodles") {
			noodleCounter += noodleGramsPerLayer
		} else if strings.Contains(layer, "sauce") {
			sauceCounter += sauceLitersPerLayer
		}
	}

	return noodleCounter, sauceCounter
}

// AddSecretIngredient adds a secret ingredient to the Lasagna recipe.
func AddSecretIngredient(incridientsFriend []string, incridientsMe []string) {
	incridientsMe[len(incridientsMe)-1] = incridientsFriend[len(incridientsFriend)-1]
}

// ScaleRecipe returns a scaled version of the Lasagna recipe.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		newQuantities[i] = quantity / 2 * float64(portions)
	}
	return newQuantities
}
