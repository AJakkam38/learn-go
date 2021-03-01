package main
import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"os"
)

const maxTurns = 5

func main() {
	rand.Seed(time.Now().UnixNano())
	guess, err := strconv.Atoi(os.Args[1])

	if guess < 0 || err != nil {
		fmt.Println("Negative numbers are not allowed")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		num := rand.Intn(guess + 1)

		if num == guess {
			fmt.Println("You Won!!!")
			return
		}
	}

	fmt.Println("You Lost!")
}
