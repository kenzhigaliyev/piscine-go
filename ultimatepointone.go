package piscine

func UltimatePointOne(n ***int) {
	*n = *n + 1
	**n = *n
	***n = **n
}
