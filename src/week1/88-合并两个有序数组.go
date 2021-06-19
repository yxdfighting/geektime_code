package week1

func merge(nums1 []int, m int, nums2 []int, n int)  {
	//主体
	i,j := m-1,n-1
	//考虑下标会越界
	for k := m+n-1;k >= 0;k --{
		if (j >= 0) && (i < 0 || nums2[j] > nums1[i]) {
			nums1[k] = nums2[j]
			j--
		}else{
			nums1[k] = nums1[i]
			i--
		}
	}
}
