package nicholasavocadoadvocate

import "fmt"

func ExplainFruit(fruit Fruit) string {
	return fmt.Sprintf("This fruit named %v, has %v of acidity", fruit.Name, fruit.Acidity)
}

func Lick(fruit Fruit) string {
	return fmt.Sprintf("%v is very good!", fruit.Name)
}
