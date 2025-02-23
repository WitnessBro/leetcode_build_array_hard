package buildarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	//var ghost = [][]int{{2, 0}}       // Input array
	var expectedResult = 6      // The expected answer with correct length.
	res := numOfArrays(2, 3, 1) // Calls your implementation

	assert.Equal(t, expectedResult, res)
}
func numOfArrays(n int, m int, k int) int {
	var MOD = 1_000_000_007
	if k > m {
		return 0
	}
	if k == m && k == 1 {
		return 1
	}
	if k > n {
		return 0
	}
	if n == m && m == k-1 {
		return 1
	}
	possibleArrays := make([][][]int, n+1)
	for i := range possibleArrays {
		possibleArrays[i] = make([][]int, m+1)
		for j := range possibleArrays[i] {
			possibleArrays[i][j] = make([]int, k+1)
		}
	}

	for j := 1; j <= m; j++ {
		possibleArrays[1][j][1] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			for l := 1; l <= k; l++ {
				if possibleArrays[i][j][l] == 0 {
					continue
				}
				for newJ := 1; newJ <= j; newJ++ {
					possibleArrays[i+1][j][l] = (possibleArrays[i+1][j][l] + possibleArrays[i][j][l]) % MOD
				}
				for newJ := j + 1; newJ <= m; newJ++ {
					if l+1 > k {
						continue
					}
					possibleArrays[i+1][newJ][l+1] = (possibleArrays[i+1][newJ][l+1] + possibleArrays[i][j][l]) % MOD
				}
			}
		}
	}
	result := 0
	for j := 1; j <= m; j++ {
		result = (result + possibleArrays[n][j][k]) % MOD
	}
	return result

}
