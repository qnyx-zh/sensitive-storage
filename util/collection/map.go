package collection

func MapIsEmpty(t map[string]any) bool {
	return !(len(t) > 0)
}
