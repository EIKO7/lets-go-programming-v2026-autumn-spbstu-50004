package main

import "fmt"

func updateRange(L, R int, op string, val int) (int, int) {
	switch op {
	case "<=":
		if val < R {
			R = val
		}
	case ">=":
		if val > L {
			L = val
		}
	case "<":
		if val-1 < R {
			R = val - 1
		}
	case ">":
		if val+1 > L {
			L = val + 1
		}
	}
	return L, R
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		L, R := 15, 30

		for j := 0; j < k; j++ {
			var op string
			var val int
			fmt.Scan(&op, &val)

			L, R = updateRange(L, R, op, val)
		}
	}
}
