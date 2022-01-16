func numJewelsInStones(jewels string, stones string) (ans int) {
    for _, ch := range stones {
        for _, js := range jewels {
            if ch == js {
                ans++
                break
            }
        }
    }
    return
}
