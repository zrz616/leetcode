func lengthOfLastWord(s string) int {
    i := len(s) - 1
    for s[i] == ' ' {
        i--
    }
    ans := 0
    for i >= 0 && s[i] != ' ' {
        ans++
        i--
    }
    return ans
}
