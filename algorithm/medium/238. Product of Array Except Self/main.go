func productExceptSelf(nums []int) []int {
    pref := make([]int, len(nums))
    suff := make([]int, len(nums))
    res := make([]int, len(nums))
    
    for i, v := range(nums) {
        if i == 0 {
            pref[i] = v
            continue
        }

        pref[i] = pref[i - 1] * v
    }

    for i, _ := range(nums) {
        idx := len(nums) - 1 - i
        v := nums[idx]

        if i == 0 {
            suff[idx] = v
            continue
        }

        suff[idx] = suff[idx + 1] * v
    }

    for i, _ := range(res) {
        if i == 0 {
            res[i] = suff[i + 1]
            continue
        }

        if i >= len(res) - 1 {
            res[i] = pref[i - 1]
            continue
        }

        res[i] = pref[i - 1] * suff[i+1]
    }
    
    return res
}
