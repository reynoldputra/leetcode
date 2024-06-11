func removeElement(nums []int, val int) int {
    i := len(nums) - 1 
    j := len(nums) - 1 

    for i >= 0 {
        if nums[i] == val {
            nums[i] = nums[j]
            nums[j] = val
            j--
        } 
        i--
    }

    return j + 1
}
