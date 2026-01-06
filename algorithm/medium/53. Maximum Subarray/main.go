func maxSubArray(nums []int) int {
    max := nums[0]
    last := nums[0]
    
    for i:= 1; i < len(nums); i++ {
        curr_val := nums[i]
        sum := last + curr_val

        if last < 0 {
            last = curr_val
        } else {
            last = sum
        }

        if sum > max {
            max = sum
        }

        if curr_val > max {
            max = curr_val
        }

        fmt.Println(max, last)
    }

    if last > max {
        max = last
    }

    return max
}