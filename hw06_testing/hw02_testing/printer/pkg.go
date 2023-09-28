package printer

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/hw06_testing/hw02_testing/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
