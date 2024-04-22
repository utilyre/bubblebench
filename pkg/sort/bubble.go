package sort

func BubbleSort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := range len(s) - i {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func FakeSort(s []int) {
	for i := range len(s) {
		for j := range i {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
