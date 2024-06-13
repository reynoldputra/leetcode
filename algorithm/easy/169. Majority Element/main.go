func majorityElement(nums []int) int {
    i := 0
    j := 0

    for _,v := range(nums) {
        if i == v {
            j++
        } else {
            if j == 0 {
                j ++
                i = v
            } else {
                j --
            }
        }
    }

    return i
}
