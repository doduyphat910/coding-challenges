func maxProfit(prices []int) int {
	var (
		min      = prices[0]
		maxPrice int
	)

	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > maxPrice {
			maxPrice = prices[i] - min
		}
	}

	return maxPrice
}

// func maxProfit(prices []int) int {
//     sortedPrices := make([]int, len(prices))
//     copy(sortedPrices, prices)
//     sort.Slice(sortedPrices, func(i, j int) bool {
//         return sortedPrices[i] < sortedPrices[j]
//     })

//     priceByKeyMap := make(map[int]int, len(prices))
//     for i := range prices {
//         priceByKeyMap[prices[i]] = i
//     }

//     var result [][]int
//     for i := range sortedPrices{
//         for j := len(sortedPrices) - 1; j >= 0; j-- {
//             if priceByKeyMap[sortedPrices[i]] < priceByKeyMap[sortedPrices[j]] && i < j {
//                 result = append(result,[]int{priceByKeyMap[sortedPrices[i]],priceByKeyMap[sortedPrices[j]]})
//             }
//         }
//     }

//     var max = 0
//     for i := range result {
//         if prices[result[i][1]] - prices[result[i][0]] > max {
//             max = prices[result[i][1]] - prices[result[i][0]]
//         }
//     }

//     return max
// }