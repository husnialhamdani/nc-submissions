func coinChange(coins []int, amount int) int {
    
    dp := make([]int, amount+1)
    dp[0]=0

    maxVal := amount+1
    for i:=1; i<=amount; i++{
        dp[i]=maxVal
    }

    for i:=1; i<=amount; i++{
        for _, coin := range coins {
            if i-coin >= 0 {
                dp[i] = min(dp[i], dp[i-coin]+1)
            }
        }
    }

    if dp[amount]==maxVal {
        return -1
    }

    return dp[amount]
}
