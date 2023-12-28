func majorityElement(nums []int) int {
    // the key is n / 2
    c := 1
    n := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] == n {
            c += 1
        } else {
            c -= 1
            if c == 0 {
                n = nums[i]
                c = 1
            } 
        }
    }

    return n
}
