func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    
    m, n := len(nums1), len(nums2)
    left, right := 0, m
    
    for left <= right {
        i := (right-left)/2+left
        j := (m + n + 1) / 2 - i 
        
		maxLeft1, minRight1, maxLeft2, minRight2 := math.MinInt32, math.MaxInt32, math.MinInt32, math.MaxInt32
		if i > 0 {
			maxLeft1 = nums1[i-1]
		}
		if i < m {
			minRight1 = nums1[i]
		}

		if j > 0 {
			maxLeft2 = nums2[j-1]
		}
		if j < n {
			minRight2 = nums2[j]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 0 {
				return float64(max(maxLeft1, maxLeft2) + min(minRight1, minRight2)) / 2.0
			} else {
				return float64(max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			right = i-1
		} else {
			left = i+1
		}
    }
    return 0.0
}