func maxProfit(prices []int) int {
    profit := 0
    last_price := prices[0]
    buy := prices[0]

    for i,today_price := range(prices) {
        if today_price < last_price {
            profit += last_price - buy
            buy = today_price
        } else if i == len(prices)-1 {
            profit += today_price - buy
        }
        
        last_price = today_price
    }

    return profit
}
