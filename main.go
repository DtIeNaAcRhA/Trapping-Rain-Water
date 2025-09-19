package main

import "fmt"

func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	result := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]

	for left < right {
		if height[left] > height[right] {
			right--
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				result += min(maxLeft, maxRight) - height[right]
			}

		} else {
			left++
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				result += min(maxLeft, maxRight) - height[left]
			}

		}
	}
	return result

}
