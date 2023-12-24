func merge(nums1 []int, m int, nums2 []int, n int)  {
    var arr []int
    arr = append(nums1[:m], nums2[:n]...)
    sort.Ints(arr)
    fmt.Println(arr)
}
