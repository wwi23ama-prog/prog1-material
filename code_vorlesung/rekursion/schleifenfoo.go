package rekursion

func CountA_v1(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			count++
		}
	}
	return count
}

func CountA_v2(s string) int {
	count := 0
	for s != "" {
		if s[0] == 'A' {
			count++
		}
		s = s[1:]
	}
	return count
}

func CountA_v3(s string) int {
	if s == "" {
		return 0
	}
	if s[0] == 'A' {
		return CountA_v3(s[1:]) + 1
	}
	return CountA_v3(s[1:])
}

func CountA_v4(s string, count int) int {
	if s == "" {
		return count
	}
	if s[0] == 'A' {
		count++
	}
	return CountA_v4(s[1:], count)
}

func CountA_v5(s string, count int, i int) int {
	if i >= len(s) {
		return count
	}
	if s[i] == 'A' {
		count++
	}
	return CountA_v5(s, count, i+1)
}

func CountA_v6(s string, i int) int {
	if i >= len(s) {
		return 0
	}
	count := 0
	if s[i] == 'A' {
		count = 1
	}
	return count + CountA_v6(s, i+1)
}

func CountA_v7(s string, i int) int {
	if i >= len(s) {
		return 0
	}
	if s[i] == 'A' {
		return CountA_v7(s, i+1) + 1
	}
	return CountA_v7(s, i+1)
}
