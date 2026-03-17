package algorithms

import "sync"

/*
Example: a = [3 1 4],

b = [ 4 3 ]

	[ 2 5 ]
	[ 6 8 ]

Result: [38 46]
Explanation:

	3(4) + 1(2) + 4(6) = 38
	3(3) + 1(5) + 4(8) = 46
*/
func multiply2DNaive(a, b [][]int) [][]int {
	numRowsA := len(a)
	numColsA := len(a[0])
	numRowsB := len(b)
	numColsB := len(b[0])

	// We cannot multiple two matrices of mismatched transposed sizes
	if numColsA != numRowsB {
		return nil
	}

	// Create resulting matrix from (rowsA x colsB)
	result := make([][]int, numRowsA)
	for i := range numRowsA {
		result[i] = make([]int, numColsB)

		// Fill out the result row we just created
		for j := range len(result[i]) {

			// Multiply the i'th row of a by j'th column of b
			sum := 0
			for k, aVal := range a[i] {
				bVal := b[k][j]
				sum += aVal * bVal
			}

			result[i][j] = sum
		}
	}

	return result
}

func multiplyWorker(i, j int, a, b, result [][]int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0

	for k, aVal := range a[i] {
		bVal := b[k][j]
		sum += aVal * bVal
	}

	result[i][j] = sum
}

func multiply2DConcurrent(a, b [][]int) [][]int {
	numRowsA := len(a)
	numColsA := len(a[0])
	numRowsB := len(b)
	numColsB := len(b[0])

	// We cannot multiple two matrices of mismatched transposed sizes
	if numColsA != numRowsB {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(numRowsA * numColsB)

	// Create resulting matrix from (rowsA x colsB)
	result := make([][]int, numRowsA)
	for i := range numRowsA {
		result[i] = make([]int, numColsB)

		for j := range len(result[i]) {
			go multiplyWorker(i, j, a, b, result, &wg)
		}
	}

	wg.Wait()

	return result
}

func multiplyRowWorker(i int, a, b [][]int, ch chan<- rowResult, wg *sync.WaitGroup) {
	defer wg.Done()

	row := make([]int, len(b[0]))

	// i is the row of result we're on (aka row of a)
	// Iterate over columns of b (j)
	for j := range len(row) {
		row[j] = 0
		for k, aVal := range a[i] {
			bVal := b[k][j]
			row[j] += bVal * aVal
		}
	}
	ch <- rowResult{
		Index: i,
		Row:   row,
	}
}

type rowResult struct {
	Index int
	Row   []int
}

func multiply2DConcurrentByRow(a, b [][]int) [][]int {
	numRowsA := len(a)
	numColsA := len(a[0])
	numRowsB := len(b)
	// numColsB := len(b[0])

	// We cannot multiple two matrices of mismatched transposed sizes
	if numColsA != numRowsB {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(numRowsA)
	ch := make(chan rowResult)

	// Create resulting matrix from (rowsA x colsB)
	result := make([][]int, numRowsA)
	for i := range numRowsA {
		go multiplyRowWorker(i, a, b, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		result[res.Index] = res.Row
	}

	return result
}

/*


 */
