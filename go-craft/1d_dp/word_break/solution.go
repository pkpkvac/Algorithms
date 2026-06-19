package word_break

func wordBreak(s string, wordDict []string) bool {
    if len(s) == 0 {
        return true  // empty string can be segmented
    }
    
    // Try every word in dict
    for _, word := range wordDict {
        if strings.HasPrefix(s, word) {
            // If s starts with this word, recurse on the rest
            if wordBreak(s[len(word):], wordDict) {
                return true
            }
        }
    }
    return false
}

func wordBreakMemo(s string, wordDict []string, memo map[string]bool) bool {
    // Check cache first
    if result, exists := memo[s]; exists {
        return result
    }
    
    // Base case: empty string can be segmented
    if len(s) == 0 {
        return true
    }
    
    // Try each word in dictionary
    for _, word := range wordDict {
        if strings.HasPrefix(s, word) {
            if wordBreakMemo(s[len(word):], wordDict, memo) {
                memo[s] = true  // Cache success
                return true
            }
        }
    }
    
    // No valid segmentation found
    memo[s] = false  // Cache failure
    return false
}

func wordBreakTabulation(s string, wordDict []string) bool {
    n := len(s)
    dp := make([]bool, n+1)
    dp[0] = true  // empty string can be segmented
    
    // For each position i in the string
    for i := 1; i <= n; i++ {
        // Try every word in dictionary
        for _, word := range wordDict {
            wordLen := len(word)
            // Can we use this word ending at position i?
            if i >= wordLen && dp[i-wordLen] {
                // Check if s[i-wordLen:i] matches word
                if s[i-wordLen:i] == word {
                    dp[i] = true
                    break  // Found one valid way, move to next i
                }
            }
        }
    }
    
    return dp[n]  // Can we segment entire string?
}

// dp[0] = true  (base case)
// dp[1] = false (can't segment "l")
// dp[2] = false (can't segment "le")
// dp[3] = false (can't segment "lee")
// dp[4] = true  (dp[0]=true AND s[0:4]="leet" ✓)
// dp[5] = false (can't segment "leetc")
// dp[6] = false (can't segment "leetco")
// dp[7] = false (can't segment "leetcod")
// dp[8] = true  (dp[4]=true AND s[4:8]="code" ✓)

// Answer: dp[8] = true