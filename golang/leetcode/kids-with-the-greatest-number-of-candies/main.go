package kids_with_the_greatest_number_of_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for i := range candies {
		if max < candies[i] {
			max = candies[i]
		}
	}
	var results []bool
	for i := range candies {
		if candies[i]+extraCandies >= max {
			results = append(results, true)
		} else {
			results = append(results, false)
		}
	}
	return results
}
