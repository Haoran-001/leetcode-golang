package main

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	i := 0

	for candies > 0 {
		ans[i%num_people] += min(i+1, candies)
		candies -= min(i+1, candies)
		i++
	}

	return ans
}
