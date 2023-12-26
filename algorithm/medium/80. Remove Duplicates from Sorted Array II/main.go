func removeDuplicates(nums []int) int {
    i := 101
    c := 0
    t := 0
    
    for idx, v := range nums {
        fmt.Println(idx, v, i, t, c)
        if v == i {
            if t < 2  {
                nums[c] = v
                c += 1
                t += 1
            }
        } else {
            nums[c] = v
            c += 1
            i = v
            t = 1
        }
    }

    return c
}
