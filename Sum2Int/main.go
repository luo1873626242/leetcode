package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 4, 1, 5, 2, 3}
	ans := twoSumV3(nums, 6)
	fmt.Println(ans)
}

// hash表
func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	hashMap := make(map[int]int, 0)
	for idx, val := range nums {
		if k, exist := hashMap[target-val]; exist {
			return []int{idx, k}
		} else {
			hashMap[val] = idx
		}
	}
	return []int{}
}

// 排序+双指针
type Node struct {
	Val int
	Idx int
}

func twoSumV2(sums []int, target int) []int {
	nodeList := make([]*Node, 0)
	for idx, val := range sums {
		node := &Node{
			Idx: idx,
			Val: val,
		}
		nodeList = append(nodeList, node)
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].Val < nodeList[j].Val
	})
	i, j := 0, len(nodeList)-1
	for i < len(nodeList)-1 && j > i {
		sum := nodeList[i].Val + nodeList[j].Val
		if sum == target {
			return []int{nodeList[i].Idx, nodeList[j].Idx}
		} else if sum > target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

//若需要返回所有可能的二元组，且二元组内的顺序不重要，即（a,b）和(b,a)是一种情况。这里返回的不是下标，而是数值
//上述两种方案，都需要新增一个判断，即当前得出的二元组，在结果集中是否已经出现（hashMap判断）

func twoSumV3(sums []int, target int) [][]int {
	ans := make([][]int, 0)
	hashMap := make(map[int]int, 0) //用于快速查找整数是否存在, val为下标
	ansMap := make(map[string]bool, 0)
	//为防止找到自身，在遍历过程中差map,没有则存入
	for idx, val := range sums {
		if _, ok := hashMap[target-val]; !ok {
			hashMap[val] = idx
		} else {
			a, b := val, target-val
			if a > b {
				tmp := a
				a = b
				b = tmp
			}
			key := fmt.Sprintf("%d_%d", a, b)
			if _, exist := ansMap[key]; !exist {
				ansMap[key] = true
				ans = append(ans, []int{a, b})
			}
		}
	}
	return ans
}
