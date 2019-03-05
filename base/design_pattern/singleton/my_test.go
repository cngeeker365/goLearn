package singleton

import (
	"fmt"
	"testing"
)

func TestLa(t *testing.T){
	var a = 1.1234567899776767
	var b = 1.12345678
	var c = 1.1234567899776767543

	t.Log(a==b, a==c, b==c)

}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1,1,1,1,2,2,3,4,4,6}

	count := 0
	for i,_ := range nums{
		if (nums[i] != nums[count]){
			count++
			nums[count]=nums[i]
		}
	}

	fmt.Println(nums)
	fmt.Println(count+1)
	fmt.Println(nums[:count+1])
}

func TestTwoSum(t *testing.T) {
	nums := []int{3,2,4}
	//for i:=0;i<len(nums);i++{
	//	for j:=len(nums)-1;j>i;j--{
	//		if nums[i]+nums[j]==6{
	//			fmt.Println(i,j)
	//		}
	//	}
	//}

	tables := make(map[int]int)
	fmt.Println(tables)
	for index,value := range nums {
		tmp := 6 - value
		fmt.Printf("%v--%v--%v   %v \n\n",index,value, tmp, tables[tmp])
		fmt.Println(tables)

		if i,ok := tables[tmp]; ok {
			fmt.Println(i,index)
		}
		tables[value] = index
	}
}