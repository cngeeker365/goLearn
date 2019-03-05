package sort

import (
	"fmt"
	"testing"
	"time"
)

var Data = [...]int{8,2,9,10,5,1,12,7,6}

/**
冒泡排序（Bubble Sort）：
	* 原理：
		1. 比较相邻的元素，若第一个＞第二个，交互
		2. 对每一对相邻元素做相同的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数
		3. 针对所有元素重复以上的步骤，除了最后一个
		4. 持续每次对越来越少的元素重复上面的步骤，知道没有任何一对数字需要比较
	* 时间复杂度：
		1. 最好情况（数据为有序）：比较次数与移动次数：Cmin=n-1, Mmin=0, 时间复杂度 O(n)
		2. 最坏情况（数据为逆序）：比较次数与移动次数：Cmax=n(n-1)/2, Mmax=3n(n-1)/2，时间复杂度 O(n^2)
		3. 平均时间复杂度 O(n^2)
	* 算法稳定性：
		冒泡排序是一种稳定排序算法。
 */

func Test_BSort(t *testing.T){
	start := time.Now().UnixNano()
	for i:=0; i< len(Data); i++{
		for j:=i+1;j<len(Data);j++{
			if Data[j] < Data[i]{
				Data[i], Data[j] = Data[j], Data[i]
			}
		}
		//fmt.Println(Data, "\t",  time.Now().UnixNano()-start)
	}
	end :=time.Now().UnixNano()
	fmt.Println(end-start)
	//fmt.Println(Data)
}


func Test_BSort2(t *testing.T){
	start := time.Now().UnixNano()
	for i:=0; i< len(Data); i++{
		for j:=1;j<len(Data)-i;j++{
			if Data[j-1] > Data[j]{
				Data[j-1], Data[j] = Data[j], Data[j-1]
			}
		}
		fmt.Println(Data, "\t",  time.Now().UnixNano()-start)
	}
	end :=time.Now().UnixNano()
	fmt.Println(end-start)
	fmt.Println(Data)
}
