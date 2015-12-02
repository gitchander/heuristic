package main

func min(as ...int) int {
	var m int
	if n := len(as); n > 0 {
		m = as[0]
		for i := 1; i < n; i++ {
			if as[i] < m {
				m = as[i]
			}
		}
	}
	return m
}

func max(as ...int) int {
	var m int
	if n := len(as); n > 0 {
		m = as[0]
		for i := 1; i < n; i++ {
			if as[i] > m {
				m = as[i]
			}
		}
	}
	return m
}
