package main

import (
	"fmt"
	"math"
	"sync"
)




func main(){
	arr:=[]int{2,4,6,8,10}
	var wg sync.WaitGroup 
	for i:=0;i<len(arr);i++{
		wg.Add(1)
		go func(value int) {
			fmt.Println(math.Pow(float64(value),2))
			wg.Done()
		}(arr[i])
	}
	wg.Wait()
}