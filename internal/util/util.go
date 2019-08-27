package util

type StringDiff struct {
	Inc []string
	Dec []string
}

func StringArrayDiff(before []string, after []string) StringDiff {
	return StringDiff{
		Inc: StringArraySub(after, before),
		Dec: StringArraySub(before, after),
	}
}

// StringArraySub return a - b
func StringArraySub(a []string, b []string) []string {
	r := []string{}
	m := make(map[string]bool)
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if _, ok := m[v]; !ok {
			r = append(r, v)
		}
	}
	return r
}
