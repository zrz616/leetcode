func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
    for _, str := range strs {
        prefix = lcp(prefix, str)
        if len(prefix) == 0 {
            break
        }
    }
    return prefix
}

func lcp(prefix, s string) string {
    i := 0
    for ; i < len(prefix) && i < len(s); i++ {
        if prefix[i] != s[i] {
            break
        }
    }
    return prefix[:i]
}
