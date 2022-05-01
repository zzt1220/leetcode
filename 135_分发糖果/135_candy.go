package _35_分发糖果

/*
1、贪心算法
*/

func candy(ratings []int) int {
	candyNum := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyNum[i] = candyNum[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			if candyNum[i]+1 > candyNum[i-1] {
				candyNum[i-1] = candyNum[i] + 1
			}
		}
	}

	total := len(ratings)
	for i := 0; i < len(candyNum); i++ {
		total += candyNum[i]
	}
	return total
}
