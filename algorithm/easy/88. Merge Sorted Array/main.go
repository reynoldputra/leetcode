func merge(nums1 []int, m int, nums2 []int, n int)  {
    idx1 := m - 1
    idx2 := n - 1
    i := len(nums1) - 1

    if m == 0 {
        copy(nums1, nums2)
        return
    }

    for i >= 0 {
        if idx2 < 0 || (idx1 >= 0 && nums1[idx1] > nums2[idx2]) {
            nums1[i] = nums1[idx1]
            idx1--
        } else {
            nums1[i] = nums2[idx2]
            idx2--
        }
        i--
    }
}
