package board

import "fmt"

func BoardSize(size int) string {
	var board string

	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {

			if (i+j)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	return fmt.Sprintf("%s\n", board)

}
