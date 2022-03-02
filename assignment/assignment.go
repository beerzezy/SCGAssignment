package assignment

import (
	"strconv"
	"strings"
)

func Calculation(argsStr string) (int, error) {

	args := strings.Split(argsStr, " ")
	wide, convWideErr := strconv.Atoi(args[0])
	if convWideErr != nil {
		return 0, convWideErr
	}

	high, convHighErr := strconv.Atoi(args[1])
	if convHighErr != nil {
		return 0, convHighErr
	}

	var areaUnit = []int{}
	for _, i := range args[2:] {
		j, err := strconv.Atoi(i)
		if err != nil {
			return 0, err

		}
		areaUnit = append(areaUnit, j)
	}

	area := make([][]int, wide)
	for i := range area {
		area[i] = make([]int, high)
	}

	for _, unit := range areaUnit {
		for row := 0; row < wide; row++ {
			for cell := 0; cell < high; cell++ {
				area[unit][cell] = 1
				area[row][unit] = 1
			}
		}
	}

	areaNotUse := 0
	for row := 0; row < wide; row++ {
		for cell := 0; cell < high; cell++ {
			if area[row][cell] == 0 {
				areaNotUse++
			}
		}
	}

	return areaNotUse, nil
}
