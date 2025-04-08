package algorithm_test

import (
	"go-algorithm/algorithm/sort"
	"go-algorithm/utils"
	"runtime"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	list := utils.GenDisOrderIntList(b.N)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	before := memStats.Alloc
	sort.MergeSort(list)
	runtime.ReadMemStats(&memStats)
	after := memStats.Alloc
	b.ReportMetric(float64(after-before)/float64(b.N), "B/op")
}

func BenchmarkInsertionSort(b *testing.B) {
	list := utils.GenDisOrderIntList(b.N)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	before := memStats.Alloc
	sort.InsertionSort(list)
	runtime.ReadMemStats(&memStats)
	after := memStats.Alloc
	b.ReportMetric(float64(after-before)/float64(b.N), "B/op")
}

func BenchmarkSelectionSort(b *testing.B) {
	list := utils.GenDisOrderIntList(b.N)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	before := memStats.Alloc
	sort.SelectionSort(list)
	runtime.ReadMemStats(&memStats)
	after := memStats.Alloc
	b.ReportMetric(float64(after-before)/float64(b.N), "B/op")
}

// isSorted 函数用于判断切片是否有序（升序或降序）
func isSorted(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}
