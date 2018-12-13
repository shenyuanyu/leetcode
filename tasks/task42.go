package tasks

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	rain := 0
	first, second := twoHeighNum(height)
	rain += getRain(first, second, height)

	for first > 1 {
		l := getHeighest(0, first, height)
		rain += getRain(l, first, height)
		first = l
	}

	for second < len(height)-2 {
		r := getHeighest(second+1, len(height), height)
		rain += getRain(second, r, height)
		second = r
	}

	return rain
}

func twoHeighNum(height []int) (int, int) {
	first, second := 0, -1
	for i := 1; i < len(height); i++ {
		if height[i] >= height[first] {
			second = first
			first = i
		} else if second == -1 {
			second = i
		} else if height[i] >= height[second] {
			second = i
		}
	}

	if first > second {
		return second, first
	}

	return first, second
}

func getRain(first int, second int, height []int) int {
	lower := height[first]
	if height[first]-height[second] > 0 {
		lower = height[second]
	}

	area := (second - first - 1) * lower
	for i := first + 1; i < second; i++ {
		area -= height[i]
	}

	return area
}

func getHeighest(first int, second int, height []int) int {
	h := first
	for first++; first < second; first++ {
		if height[first] > height[h] {
			h = first
		}
	}
	return h
}
