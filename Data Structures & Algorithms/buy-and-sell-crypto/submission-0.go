func maxProfit(prices []int) int {
    l, r := 0, 1
    maxP := 0

    for r < len(prices) {
        if prices[l] < prices[r] {

            profit := prices[r]-prices[l]
            if profit > maxP{
                maxP = profit
            }
        } else {
            l = r
        }
        r++
    }
    return maxP
}

/*
        l r
prices=[7,1,5,3,6,4]

*/