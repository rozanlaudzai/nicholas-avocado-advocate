package payment

import (
	"fmt"

	nicholasavocadoadvocate "github.com/rozanlaudzai/nicholas-avocado-advocate"
)

func Pay(fruit nicholasavocadoadvocate.Fruit) (bool, string) {
	// doc aja sih ini
	return false, fmt.Sprintf("You successfully purchased a %v with %v acidity.", fruit.Name, fruit.Acidity)
}
