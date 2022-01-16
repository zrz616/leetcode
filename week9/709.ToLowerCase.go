func toLowerCase(s string) string {
    lower := &strings.Builder{}
    lower.Grow(len(s))
    for _, ch := range s {
        if ch >= 'A' && ch <= 'Z' {
            ch |= 32
        }
        lower.WriteRune(ch)
    }
    return lower.String()
}
