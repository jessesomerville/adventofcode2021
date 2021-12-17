package main

import "log"

// SumOp returns the sum of nums.
func SumOp(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// ProductOp returns the product of nums.
func ProductOp(nums []int) int {
	product := 1
	for _, n := range nums {
		product *= n
	}
	return product
}

// MinOp returns returns the minimum value in nums.
func MinOp(nums []int) int {
	min := int(^uint(0) >> 1)
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

// MaxOp returns the maximum value in nums.
func MaxOp(nums []int) int {
	max := -1
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

// GtOp returns 1 if the first value is greater than the second value, otherwise returns 0.
func GtOp(nums []int) int {
	if len(nums) != 2 {
		log.Fatalf("[GtOp] Length of nums = %d, but should = 2", len(nums))
	}
	if nums[0] > nums[1] {
		return 1
	}
	return 0
}

// LtOp returns 1 if the first value is less than the second value, otherwise returns 0.
func LtOp(nums []int) int {
	if len(nums) != 2 {
		log.Fatalf("[LtOp] Length of nums = %d, but should = 2", len(nums))
	}
	if nums[0] < nums[1] {
		return 1
	}
	return 0
}

// EqOp returns 1 if the first value is equal to the second value, otherwise returns 0.
func EqOp(nums []int) int {
	if len(nums) != 2 {
		log.Fatalf("[EqOp] Length of nums = %d, but should = 2", len(nums))
	}
	if nums[0] == nums[1] {
		return 1
	}
	return 0
}
