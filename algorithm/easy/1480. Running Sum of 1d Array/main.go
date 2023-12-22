func runningSum(nums []int) []int {
    var res []int
    t := 0
    for _, n := range nums {
        res = append(res, n + t)
        t = n + t
    }

    return res
}
