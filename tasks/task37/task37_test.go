package task37

import (
	"reflect"
	"testing"
)

func Test_generateMapByByteRange(t *testing.T) {
	type args struct {
		start byte
		end   byte
	}
	tests := []struct {
		name string
		args args
		want map[byte]struct{}
	}{
		// TODO: Add test cases.
		{
			"c1",
			args{'1', '9'},
			map[byte]struct{}{
				'1': struct{}{},
				'2': struct{}{},
				'3': struct{}{},
				'4': struct{}{},
				'5': struct{}{},
				'6': struct{}{},
				'7': struct{}{},
				'8': struct{}{},
				'9': struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMapByByteRange(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMapByByteRange() = %v, want %v", got, tt.want)
			}
			t.Logf("%s case success, generate [%v]", tt.name, tt.want)
		})
	}
}

func Test_deletePositionArray(t *testing.T) {
	// case1
	arr := []position{{0, 1}, {0, 2}, {0, 3}, {0, 4}}
	i := 2

	out := deletePositionArray(arr, i)

	t.Logf("the result is: [%v], the origin is: [%v]", out, arr)
}

func Test_excludeByPossible(t *testing.T) {
	pos := position{0, 1}
	notSolvePosArry := []position{{0, 2}, {0, 4}, {0, 5},{0, 7}}
	notSolvePosPossibleMap := map[position]map[byte]struct{}{
		{0, 1}: {
			'2': struct{}{},
			'3': struct{}{},
			'4': struct{}{},
			'5': struct{}{},
		},
		{0, 2}: {
			'2': struct{}{},
			'3': struct{}{},
		},
		{0, 4}: {
			'2': struct{}{},
			'3': struct{}{},
		},
	}

	excludeByPossible(pos, notSolvePosArry, notSolvePosPossibleMap)

	t.Logf("the possible is: [%v]", notSolvePosPossibleMap[pos])
}

func Test_excludeByNotPossible(t *testing.T) {
	pos := position{0, 1}
	notSolvePos := []position{{0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6},
		{0, 7}, {0, 8}, {0, 9}}
	notSolvePosPossibleMap := map[position]map[byte]struct{}{
		{0,1}: {'1': {}, '2': {}, '3': {}},
		{0,2}: {'1': {}, '3': {}},
		{0,3}: {'4': {}, '3': {}},
		{0,4}: {'5': {}, '3': {}},
		{0,5}: {'7': {}, '3': {}},
		{0,6}: {'8': {}, '3': {}},
		{0,7}: {'6': {}, '3': {}},
		{0,8}: {'9': {}, '3': {}},
		{0,9}: {'3': {}, '6': {}},
	}

	excludeByNotPossible(pos,notSolvePos, notSolvePosPossibleMap)

	t.Logf("this pos possible is: [%v]", notSolvePosPossibleMap[pos])

}
