func canJump(nums []int) bool {
    max_idx := 0
    for i,v := range(nums) {
        if i <= max_idx {
            if i + v > max_idx {
                max_idx = i + v
            }
        }
    }

    fmt.Println(max_idx, len(nums))

    return max_idx + 1 >= len(nums)
}
