package main

import "fmt"

func main(){
	nums := []int{3, 2, 4}
	target := 6
	var i, j int
	flag :=true
	for i = 0; i < len(nums) - 1 && flag; i++{
		for j = i + 1; j < len(nums); j++{
			if nums[i] + nums[j] == target{
				flag = false
				break
			}
		}
	}
	fmt.Println(i - 1,j)
	return
}
