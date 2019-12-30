func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	coin := 0
	if !(A >= G || B >= H || D <= F || C <= E) {
		coin += (min(C, G) - max(A, E)) * (min(D, H) - max(B, F))
	}
	return (D-B) * (C-A) + (H-F)*(G-E) - coin
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
