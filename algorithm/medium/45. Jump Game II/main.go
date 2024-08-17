func jump(nums []int) int {
    if len(nums) <= 1{
        return 0
    }
    
    r := 0
    l := 0
    c := 0

    for r < len(nums) - 1 {
        max := nums[l] + l
        for i := l; i <= r; i ++ {
            if max < nums[i] + i {
                max = nums[i] + i
            }
        }
        
        l = r + 1
        r = max
        c++
    }

    return c
}
