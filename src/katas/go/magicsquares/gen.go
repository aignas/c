package magicsquares

// Matrix is a number grid, where the first index is a column index.
type Matrix [][]int

// Generate gives all magic squares of a given size.
func Generate(size int) []Matrix {
	if size < 3 {
		return nil
	}

	var result []Matrix
	for m := range genAll(size) {
		result = append(result, m)
	}

	return result
}

func genAll(N int) chan Matrix {
	ch := make(chan Matrix, 1)

	go func() {
		numbers := make([]int, N*N)
		for i := range numbers {
			numbers[i] = i + 1
		}

		permutations(N, len(numbers), numbers, ch)
		close(ch)
	}()

	return ch
}

// permutations generate all permutations of A where it can form a magic square.
//
// taken and adapted from wikipedia.
func permutations(N, k int, A []int, out chan Matrix) {
	if k == 1 {
		if check(N, A) {
			out <- newMatrix(N, A)
		}
		return
	}

	// Generate permutations with kth unaltered
	// Initially k == length(A)
	permutations(N, k-1, A, out)

	// Generate permutations for kth swapped with each k-1 initial
	for i := 0; i < k-1; i++ {
		// Swap choice dependent on parity of k (even or odd)
		j := 0
		if k%2 == 0 {
			j = i
		}
		A[j], A[k-1] = A[k-1], A[j]

		permutations(N, k-1, A, out)
	}
}

func newMatrix(size int, numbers []int) Matrix {
	var (
		i int
		m Matrix
		b = make([]int, len(numbers))
	)
	copy(b, numbers)
	for {
		right := size * (i + 1)
		if right > len(numbers) {
			break
		}

		m = append(m, b[size*i:size*(i+1)])
		i++
	}
	return m
}

// Check returns true if the given matrix is a magic square.
//
// Magic square is a matrix, which has distinct elements and the sum of rows,
// columns and diagonals is the same.
func Check(m Matrix) bool {
	N := len(m)

	if N < 3 {
		return false
	}

	r := m[0][:0]

	for i := range m {
		r = append(r, m[i]...)
	}

	return check(N, r)
}

func check(N int, m []int) bool {
	var (
		want = (1 + N*N) * N / 2
		sums = make([]int, N+2)
	)
	if N*N != len(m) || N < 3 {
		return false
	}

	for i := 0; i < N; i++ {
		var sum int
		for j := 0; j < N; j++ {
			sum += m[i*N+j]
			sums[j] += m[i*N+j]
		}

		// check the row sum early
		if sum != want {
			return false
		}

		sums[N] += m[i*N+i]
		sums[N+1] += m[(i+1)*(N-1)]
	}

	for _, s := range sums {
		if s != want {
			return false
		}
	}

	return true
}
