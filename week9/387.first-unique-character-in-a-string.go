func firstUniqChar(s string) int {
    vs := make([]int, 26)
    for _, ch := range s {
        vs[ch-'a']++
    }
    for i, ch := range s {
        if vs[ch-'a'] == 1 {
            return i
        }
    }
    return -1
}
