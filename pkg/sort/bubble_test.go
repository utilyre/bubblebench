package sort

import (
	"math/rand"
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testSort(t, BubbleSort)
}

func TestFakeSort(t *testing.T) {
	testSort(t, FakeSort)
}

func testSort(t *testing.T, sort func(s []int)) {
	got := []int{-8, 5, 10, 2}
	want := []int{-8, 2, 5, 10}
	sort(got)
	if !slices.Equal(got, want) {
		t.Errorf("result = %v; want %v", got, want)
	}

	got = []int{2}
	want = []int{2}
	sort(got)
	if !slices.Equal(got, want) {
		t.Errorf("result = %v; want %v", got, want)
	}

	got = []int{}
	want = []int{}
	sort(got)
	if !slices.Equal(got, want) {
		t.Errorf("result = %v; want %v", got, want)
	}
}

func BenchmarkBubbleSort_n100(b *testing.B)     { benchmarkSort(b, BubbleSort, 100) }
func BenchmarkBubbleSort_n1000(b *testing.B)    { benchmarkSort(b, BubbleSort, 1_000) }
func BenchmarkBubbleSort_n10000(b *testing.B)   { benchmarkSort(b, BubbleSort, 10_000) }
func BenchmarkBubbleSort_n100000(b *testing.B)  { benchmarkSort(b, BubbleSort, 100_000) }
func BenchmarkBubbleSort_n1000000(b *testing.B) { benchmarkSort(b, BubbleSort, 1_000_000) }

func BenchmarkFakeSort_n100(b *testing.B)     { benchmarkSort(b, FakeSort, 100) }
func BenchmarkFakeSort_n1000(b *testing.B)    { benchmarkSort(b, FakeSort, 1_000) }
func BenchmarkFakeSort_n10000(b *testing.B)   { benchmarkSort(b, FakeSort, 10_000) }
func BenchmarkFakeSort_n100000(b *testing.B)  { benchmarkSort(b, FakeSort, 100_000) }
func BenchmarkFakeSort_n1000000(b *testing.B) { benchmarkSort(b, FakeSort, 1_000_000) }

func benchmarkSort(b *testing.B, sort func([]int), n int) {
	for range b.N {
		b.StopTimer()
		s := randomSlice(n)
		b.StartTimer()

		sort(s)
	}
}

func randomSlice(n int) []int {
	s := make([]int, n)
	for i := range n {
		s[i] = rand.Int()
	}
	return s
}
