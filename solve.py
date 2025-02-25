def maxScore(cards):
    n = len(cards)
    dp = [[0] * n for _ in range(n)]

    # Base case: When only one card is left
    for i in range(n):
        dp[i][i] = cards[i]

    # Fill DP table for all possible subarray lengths
    for length in range(2, n+1):  # Subarray sizes from 2 to n
        for i in range(n - length + 1):
            j = i + length - 1  # Ending index
            # Compute the best score the first player can get
            take_left = cards[i] + min(dp[i+2][j] if i+2 <= j else 0, dp[i+1][j-1] if i+1 <= j-1 else 0)
            take_right = cards[j] + min(dp[i+1][j-1] if i+1 <= j-1 else 0, dp[i][j-2] if i <= j-2 else 0)
            dp[i][j] = max(take_left, take_right)

    return dp[0][n-1]  # The max score Player 1 can collect from the full deck

# Example usage:
cards = list([59, 324, 915, 608, 779, 958, 814, 387, 454, 505])
print(maxScore(cards))
