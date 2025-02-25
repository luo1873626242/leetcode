package main
import(
	"fmt"
)
func main(){
	s := []string{"abba"}
	for _, str := range s{
		fmt.Println(lengthOfLongestSubstring(str))
	}
}
func lengthOfLongestSubstring(s string) int {
	var maxLen = 0
	var start, end = 0, 0
	hashMap := make(map[byte]int)
	for end = 0; end < len(s); end ++{
		if index, ok := hashMap[byte(s[end])]; ok && start <= index{
			 start = index +1 
		}
		hashMap[byte(s[end])] = end 
		if maxLen < end - start + 1 {
			maxLen = end - start + 1
		}
	}
	return maxLen
}