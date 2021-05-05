// KULINA TEST ALGORITHM
// DHIYA ULHAQ RIFQI

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func bulbSwitch(bulbNumber, loop int) {

	bulb := make(map[int]bool)
	for z := 1; z <= bulbNumber; z++ {

		bulb[z] = false
	}

	for j := 1; j <= loop; j++ {

		for i := j; i <= loop; i++ {

			if i%j == 0 {

				if bulb[i] == false {

					bulb[i] = true
				} else {

					bulb[i] = false
				}
			}
		}

	}

	count := 0
	for i, v := range bulb {

		if v == true {

			fmt.Println("LampNumber", i, "it's on")
			count++
		}
	}

	fmt.Println(count, "lamp is on")

}

func PrintNumber(number int) {

	arr := strings.Split(strconv.Itoa(number), "")
	lenArr := len(arr)

	for i := lenArr - 1; i >= 0; i-- {

		k := arr[len(arr)-i-1]
		var j int
		power := int(math.Pow10(i))
		fmt.Sscan(k, &j)
		fmt.Println(j * power)
	}
}

func countingValley(n int, path string) int {

	a := strings.Split(path, "")
	seaLvl := 0
	valley := 0

	for _, v := range a {

		if v == "U" {
			seaLvl += 1

		} else {

			seaLvl -= 1
		}

		if v == "U" && seaLvl == 0 {
			valley += 1
		}

	}

	return valley

}

func pairSocks(n int, arr []int) int {

	checkMap := make(map[int]bool)
	count := 0
	for i := 0; i < n; i++ {

		if checkMap[arr[i]] {
			checkMap[arr[i]] = false
			count++

		} else {

			checkMap[arr[i]] = true
		}
	}

	return count
}

func main() {

	fmt.Println("KULINA TEST ALGORITHM - DHIYA ULHAQ RIFQI")

	fmt.Println("----------------------------")
	fmt.Println("")
	fmt.Println("number 1")

	fmt.Println(pairSocks(9, []int{10, 20, 20, 10, 10, 30, 50, 10, 20}))

	fmt.Println("----------------------------")
	fmt.Println("")
	fmt.Println("number 2")

	fmt.Println(countingValley(8, "UDDDUDUU"))

	fmt.Println("----------------------------")
	fmt.Println("")
	fmt.Println("number 3")

	PrintNumber(1345679)

	fmt.Println("----------------------------")
	fmt.Println("")
	fmt.Println("number 4")

	bulbSwitch(100, 100)

}
