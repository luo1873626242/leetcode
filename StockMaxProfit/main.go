package  main 

import(
	"math"
	"fmt"
)

func main(){
	prices := []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
    maxProfit := int(0) 
	minPrice := math.MaxInt32
	for _, price := range prices{
		if price < minPrice{
			minPrice = price
		}else if price - minPrice > maxProfit{
			maxProfit = price - minPrice
		}
	}
    return maxProfit
}