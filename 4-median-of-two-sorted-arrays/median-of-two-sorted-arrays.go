func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums3 := make([]int, len(nums1)+len(nums2))
	inf := math.MaxInt
	nums1 = append(nums1, inf)
	nums2 = append(nums2, inf)
	i := 0
	j := 0
	k := 0
	for nums1[i] != inf && nums2[j] != inf {
		if nums1[i] <= nums2[j] {
			nums3[k] = nums1[i]
			i++
			k++
		} else {
			nums3[k] = nums2[j]
			j++
			k++
		}
	}
	for nums1[i] != inf {
		nums3[k] = nums1[i]
		k++
		i++
	}
	for nums2[j] != inf {
		nums3[k] = nums2[j]
		k++
		j++
	}
    var median float64
    if len(nums3)%2!=0{
        mid:=len(nums3)/2
        median=float64(nums3[mid])
    }else{
        mid:=len(nums3)/2
        median=(float64(nums3[mid]+nums3[mid-1])/float64(2))
    }
    return median
}