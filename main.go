package main

import ("fmt")



func main() {
	
	var arr [4]int = [4]int{1,2,3,4}
	arr2 := [4]int {5,6,7,8}

	fmt.Println(arr,arr2)

	nums := []int{1,2,3,4,4}
	fmt.Println(nums)
	nums2 := []int{99,99,99}
	nums = append(nums, nums2...)
	fmt.Println(nums)

	user := map[any]any{1:25, // [string]any  any and interface{} both same thing
		"name":"JAY",
	}

	for k,v := range user {
		fmt.Println(k,v,"|")
	}

	fmt.Println(user,user["age"])



	// fmt.Println(arr,arr[0],arr2,arr2[len(arr2)])


}