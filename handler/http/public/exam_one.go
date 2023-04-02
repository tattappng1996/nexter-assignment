package public

import (
	"net/http"
	"nexter-assignment/utils"

	"github.com/labstack/echo/v4"
)

// GET ../exam/one
func (h *handler) GetExamOne(c echo.Context) error {
	// Design API to calculate value x,y,z in data set [1, X, 8, 17, Y, Z, 78, 113]

	// what we know now is some distance on dataset number that already exist
	//  [8 -> 17] is 9
	//  [78, 113] is 35
	//  First, as you can see they have a "relationship" just by observed
	//		[1(+s1), X(+s2), 8(+9), 17(+s4), Y(+s5), Z(+s6), 78(+35), 113]
	//  We have to find "a formula" by declare (s2, X) backward until they're matched
	//	And we know X should be in range between (2->7)
	//	the values that matched for (s2, X) is "5, 3" because
	//		[1(+s1), 3(+5), 8(+9),...]
	//	After I fill this I just guess a formula to find s1
	//		[1(+2), 3(+5), 8(+9),...] <- those values are matched
	//	So maybe the relationship about "s(n)" might good to find all answer
	//		[1(+2), 3(+5), 8(+9),...]
	//			[|2-5|, |5-9|]
	//	Everything is clear now
	//	seq1 =	[1(+2), 3(+5), 8(+9),..., 78(+35), 113]
	//	seq2 =		 [+3,   +4,		...,		+?]
	//	"Formula" is seq1(n) = seq1(n-1) + (seq2(n)+1)
	//	let's fill the missing number
	//		[1(+2), 3(+5), 8(+9), 17(+14), 31(+20), 51(+27), 78(+35), 113(+...)]
	//			 [+3,   +4,		+5,		+6		+7,		 +8,	   +9...]
	// The Answer is [1, 3, 8, 17, 31, 51, 78, 113]

	inputs := []int{1, 0, 8, 17, 0, 0, 78, 113}
	sequences1 := []int{0, 0, 9, 0, 0, 0, 35, -1}
	sequences2 := []int{0, 0, 0, 0, 0, 0, 0, -1}

	// Find X, Sequences1[0,1], Sequences2[0,1,2]
	predictX := inputs[0] + 1
	for {
		sequences1[1] = inputs[2] - predictX
		inputs[1] = inputs[2] - sequences1[1]
		sequences1[0] = inputs[1] - inputs[0]
		sequences2[0] = sequences1[1] - sequences1[0]
		sequences2[1] = sequences2[0] + 1
		if inputs[1]-sequences1[0] == inputs[0] &&
			inputs[1]+sequences1[1] == inputs[2] &&
			sequences1[0]+sequences2[0] == sequences1[1] &&
			sequences1[1]+sequences2[1] == sequences1[2] {
			sequences2[0] = sequences1[1] - sequences1[0]
			sequences2[1] = sequences2[0] + 1
			sequences2[2] = sequences2[1] + 1
			break
		}

		predictX++
		sequences1[1] = 0
		inputs[1] = 0
		sequences1[0] = 0
		sequences2[0] = 0
		sequences2[1] = 0
	}

	for i := 0; i < len(inputs); i++ {
		if i == len(inputs)-1 {
			break
		}

		inputs[i+1] = sequences1[i] + inputs[i]
		sequences2[i+1] = sequences2[i] + 1
		sequences1[i+1] = sequences1[i] + sequences2[i]
	}

	return utils.SuccessResponseMessage(http.StatusOK, map[string]interface{}{
		"input":  []interface{}{1, "X", 8, 17, "Y", "Z", 78, 113},
		"output": inputs,
	}, c)
}
