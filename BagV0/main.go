package main
import (
	"fmt"
)
func main(){
	N, C := 5,8
	w := []int{4,1,2,3,4}
	v := []int{5,2,4,4,5}
	result := bag(N,C,w,v)
	fmt.Println(result)
}


func bag(N, C int, w []int, v []int) int{
	if N==0 || C==0 {
		return 0
	}
	if len(w) != N || len(v) != N{
		return -1
	}
	dp := make([][]int, N+1)
	for i:= 0; i<= N; i++{
		dp[i] = make([]int, C+1)
		for j := 0; j<= C ; j++{
			dp[i][j] = 0
		}
	}

	for i := 1; i<= N; i++{
		for j := 1; j<=C; j++{
			if j < w[i-1] {
				dp[i][j] = dp[i-1][j]
			}else{
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}
	return dp[N][C]
}


func max(a,b int)int{
	if a >b{
		return a
	}
	return b
}