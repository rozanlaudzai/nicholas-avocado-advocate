package payment

import (
	"fmt"

	nicholasavocadoadvocate "github.com/rozanlaudzai/nicholas-avocado-advocate"
)

func Pay(fruit nicholasavocadoadvocate.Fruit) (bool, string) {
	// doc aja sih ini
	return true, fmt.Sprintf("You successfully purchased a %v.", fruit.Name)
}
