package behavior

import "fmt"

func Wander(value int, next func()) {
	if value > 10 {
		next()
		return
	}

	fmt.Println("Wandering")
}
