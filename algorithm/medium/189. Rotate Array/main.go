func rotate(nums []int, k int)  {
    if len(nums) == 1 {
        return
    }
    
    k = k % len(nums)
    t := len(nums) - k
    tr := make([]int, len(nums))
    copy(tr, nums)
    copy(nums[:], tr[t:])
    copy(nums[k:], tr[:t])
    
}
