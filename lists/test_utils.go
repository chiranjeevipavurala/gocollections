package lists

type IntComparator struct{}

func (c *IntComparator) Compare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

type StringComparator struct{}

func (c *StringComparator) Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
