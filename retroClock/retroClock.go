package main

import (
	"fmt"
	"time"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	colon := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine, colon,
	}

	for {
		hour, min, sec := time.Now().Hour(), time.Now().Minute(), time.Now().Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			fmt.Print(line, "    ")
			for digit := range clock {
				fmt.Print(clock[digit][line], " ")
			}
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}

/*
Output:

0    ██  █ █     ███ ███     ██  ███
1     █  █ █  █    █   █  █   █  █ █
2     █  ███     ███   █      █  ███
3     █    █  █  █     █  █   █  █ █
4    ███   █     ███   █     ███ ███
*/
