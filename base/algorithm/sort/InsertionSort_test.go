package sort

import (
	"fmt"
	"testing"
	"time"
)

//var Data = [...]int{8,2,9,10,5,1,12,7,6}

func Test_ISort(t *testing.T){
	start := time.Now().UnixNano()

	for i:=1;i<len(Data);i++{
		for j:=i;j>0;j--{
			if Data[j] < Data[j-1]{
				Data[j], Data[j-1] = Data[j-1], Data[j]
			}else {
				break
			}
		}
	}

	end :=time.Now().UnixNano()
	fmt.Println(end-start)
	fmt.Println(Data)
}