package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {

	var dp []int
	temp := int(math.Abs(float64(jumps[0] - jumps[1])))
	dp = append(dp, 0)
	dp = append(dp, temp)
	for i := 2; i < len(jumps); i++ {
		temp_1 := dp[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		temp_2 := dp[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		if temp_1 < temp_2 {
			dp = append(dp, temp_1)
		} else {
			dp = append(dp, temp_2)
		}

	}
	
	return dp[len(jumps)-1]

}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}
