package code

// 股票买入的最佳时机
// 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

func maxProfit(prices []int) int {
	minPrice := prices[0] // 最小价格
	maxProfit := 0        // 最大利润
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

// 股票买入的最佳时机 II
// 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易
func maxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

// 股票买入的最佳时机 III
// 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易
func maxProfit3(prices []int) int {
	firstBuy, firstSell := -prices[0], 0
	secondBuy, secondSell := -prices[0], 0
	for _, price := range prices {
		firstBuy = max(firstBuy, -price)
		firstSell = max(firstSell, firstBuy+price)
		secondBuy = max(secondBuy, firstSell-price)
		secondSell = max(secondSell, secondBuy+price)
	}
	return secondSell
}

// 股票买入的最佳时机 IV
// 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易
// 你不可以参与多笔交易（你必须在再次购买前出售掉之前的股票）
func maxProfit4(k int, prices []int) int {
	if k == 0 {
		return 0
	}
	if k > len(prices)/2 {
		return maxProfit2(prices)
	}
	buy := make([]int, k)
	sell := make([]int, k)
	for i := 0; i < k; i++ {
		buy[i] = -prices[0]
		sell[i] = 0
	}
	for _, price := range prices {
		for i := 0; i < k; i++ {
			if i == 0 {
				buy[i] = max(buy[i], -price)
			} else {
				buy[i] = max(buy[i], sell[i-1]-price)
			}
			sell[i] = max(sell[i], buy[i]+price)
		}
	}
	return sell[k-1]
}
