
package main
import(
	"fmt"
	"sort"
)
func main(){
	nums1 := []int{0,0,0}
	nums2 := []int{1,5,6}
	mergev3(nums1, 0, nums2, 3)
	fmt.Println(nums1)
}

//时间O((n+m)log(n+m))
//空间O(log(m+n))
func mergev1(nums1 []int, m int, nums2 []int, n int)  {
	for i:= m ;i < m+n; i++{
		nums1[i] = nums2[i-m]
	}
	sort.Slice(nums1, func(i, j int) bool{
		return nums1[i] < nums1[j]
	})
}

//时间O(m+n)，空间O(m+n)))
func mergev2(nums1 []int, m int, nums2 []int, n int){
	result := make([]int,0, m+n)
	i, j := 0,0
	for i<m && j< n{
		if nums1[i]<nums2[j]{
			result = append(result, nums1[i])
			i++
		}else{
			result = append(result, nums2[j])
			j++
		}
	}
	 if i<m{
		result = append(result, nums1[i:m]...)
	 }
	 if j<n{
		result = append(result, nums2[j:n]...)
	 }
	copy(nums1, result)
}

//时间O(m+n)
//空间O(1)
func mergev3(nums1[]int, m int, nums2[]int, n int){
    if n==0 {return}
    if m==0 {copy(nums1, nums2)}
	i,j,k := m-1,n-1,m+n-1
	for i>=0 && j>=0 {
		if nums1[i]>nums2[j]{
			nums1[k] = nums1[i]
			k--
			i--
		}else{
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	for i >= 0{
		nums1[k]=  nums1[i]
		k--
		i--
	}
	for j >= 0{
		nums1[k] = nums2[j]
		k--
		j--
	}
}