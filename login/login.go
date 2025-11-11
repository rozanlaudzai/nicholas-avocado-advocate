package login

import (
	"fmt"

	nicholasavocadoadvocate "github.com/rozanlaudzai/nicholas-avocado-advocate"
)

func Login(fruit nicholasavocadoadvocate.Fruit) (bool, string) {
	return true, fmt.Sprintf("%v has successfully logged in.", fruit.Name)
}
