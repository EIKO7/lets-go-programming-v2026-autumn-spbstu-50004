package main

import "fmt"

func updateRange(left, right int, op string, val int) (int, int) {
	switch op {
	case "<=":
		if val < right {
			right = val
		}
	case ">=":
		if val > left {
			left = val
		}
	case "<":
		if val-1 < right {
			right = val - 1
		}
	case ">":
		if val+1 > left {
			left = val + 1
		}
	}

	return left, right
}

func main() {
	var count int
	if _, err := fmt.Scan(&count); err != nil {
		return
	}

	for range count {
		var operations int
		if _, err := fmt.Scan(&operations); err != nil {
			return
		}

		left, right := 15, 30

		for range operations {
			var op string
			var val int
			if _, err := fmt.Scan(&op, &val); err != nil {
				return
			}

			left, right = updateRange(left, right, op, val)
		}

		if left <= right {
			fmt.Println(left)
		} else {
			fmt.Println(-1)
		}
	}
}
