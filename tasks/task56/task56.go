package task56

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Sort(IntervalsSort(intervals))

	results := make([][]int, 0, len(intervals)/2)
	results = append(results, intervals[0])
	for i, j := 0, 1; j < len(intervals); {
		if intervals[j][0] > intervals[i][1] {
			results = append(results, intervals[j])
			i = j
			j++
			continue
		} else {
			if intervals[i][1] >= intervals[j][1] {
				j++
				continue
			} else {
				results[len(results)-1][1] = intervals[j][1]
			}
		}
	}

	return results
}

type IntervalsSort [][]int

func (interval IntervalsSort) Len() int {
	return len(interval)
}

func (interval IntervalsSort) Less(i, j int) bool {
	return interval[i][0] < interval[j][0]
}

func (interval IntervalsSort) Swap(i, j int) {
	interval[i], interval[j] = interval[j], interval[i]
}
