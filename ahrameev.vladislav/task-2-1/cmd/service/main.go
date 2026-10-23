package main

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
}
