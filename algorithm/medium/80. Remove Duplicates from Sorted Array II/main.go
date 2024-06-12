func removeDuplicates(nums []int) int {
    k := 1

    if len(nums) < 2 {
        return len(nums)
    }

    for i := 2; i < len(nums) ; i ++ {
        if nums[i] == nums[k] {
            if nums[k] != nums[k-1] {
                k += 1
                nums[k] = nums[i]
            } 
            continue
        } else {
            k += 1
            nums[k] = nums[i]
        }
    }

    return k+1
}
