package behavior

import "fmt"

func Avoid(value int, next func()) {
	if value > 6 {
		next()
		return
	}

	fmt.Println("Avoiding")
}
