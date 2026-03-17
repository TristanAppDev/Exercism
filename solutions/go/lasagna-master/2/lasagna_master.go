package lasagna

import "strings"

// PreparationTime returns the total preparation time of the Lasagna recipe.
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// Quantities returns the quantities of the ingredients in the Lasagna recipe.
func Quantities(layers []string) (int, float64) {
	noodleGramsPerLayer := 50
	sauceLitersPerLayer := 0.2
	noodleLayerCount := 0
	sauceLayerCount := 0

	for _, layer := range layers {
		if strings.Contains(layer, "noodles") {
			noodleLayerCount++
		} else if strings.Contains(layer, "sauce") {
			sauceLayerCount++
		}
	}

	return noodleGramsPerLayer * noodleLayerCount, sauceLitersPerLayer * float64(sauceLayerCount)
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
