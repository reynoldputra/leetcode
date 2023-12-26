func removeDuplicates(nums []int) int {
    i := 101
    c := 0
    
    for _, v := range nums {
        if v != i {
            nums[c] = v
            i = v
            c += 1
        }
    }

    return c
}
