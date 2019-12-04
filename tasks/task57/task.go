package task57

func insert(intervals [][]int, newInterval []int) [][]int {
	t := 0

	for i := len(intervals) - 1; ; i-- {
		if i == -1 {
			s := make([][]int, 0, len(intervals)+1)
			intervals = append(s, intervals...)
			return merge(intervals)
		}

		if intervals[i][0] < newInterval[0] {
			s := make([][]int, len(intervals[i+1:]))
			copy(s, intervals[i+1:])
			intervals = append(intervals[:i+1], newInterval)
			intervals = append(intervals, s...)
			t = i
			break
		}
	}

	s := merge(intervals[t:])
	return append(intervals[:t], s...)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

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
