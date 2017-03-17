// Package pascal implements Pascal's triangle
package pascal

const testVersion = 1

// Triangle computes Pascal's triangle up to a given number of rows
func Triangle(size int) [][]int {
	var all [][]int
	var row []int

	for i := 0; i < size; i++ {
		if i == 0 {
			all = append(all, []int{1})
		}
		if i == 1 {
			all = append(all, []int{1, 1})
		}

		if i > 1 {

			row = []int{}
			row = append(row, 1)

			for j := 0; j < len(all[i-1])-1; j++ {
				row = append(row, all[i-1][j]+all[i-1][j+1])
			}

			row = append(row, 1)
			all = append(all, row)
		}
	}

	return all
}
