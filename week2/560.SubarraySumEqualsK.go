import "fmt"

func subarraySum(nums []int, k int) int {
	var s []int
	s = append(s, 0)
	for i := 1; i <= len(nums); i++ {
		s = append(s, s[i-1]+nums[i-1])
	}
	fmt.Println(s) // 0 1 2 3
	c := make(map[int]int)
	c[0] = 1
	ans := 0
	for i := 1; i <= len(nums); i++ {
		ans += c[s[i]-k]
		c[s[i]]++
	}
	fmt.Println(c)
	return ans
}
