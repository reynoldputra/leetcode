func summaryRanges(nums []int) []string {
    var res []string

    if len(nums) == 0 {
        return res
    }

    prev := nums[0]
    prevIdx := 0

    if len(nums) <= 1 {
        str := fmt.Sprintf("%d", prev)
        res = append(res, str)
        return res
    }

    for i, v := range nums {
        if i == 0 {
            continue
        }
        
        if len(nums) - 1 == i {
            if prev + (i-prevIdx) == v {
                str := fmt.Sprintf("%d->%d", prev, v)

                res = append(res, str)
            } else {
                str := fmt.Sprintf("%d", prev)
    
                if i != prevIdx + 1 {
                    str = fmt.Sprintf("%d->%d", prev, nums[i-1])
                }
                
                lastStr := fmt.Sprintf("%d", v)
                res = append(res, str, lastStr)
                
            }
            return res
        }

        if prev + (i-prevIdx) != v {
            str := fmt.Sprintf("%d", prev)

            if i != prevIdx + 1 {
                str = fmt.Sprintf("%d->%d", prev, nums[i-1])
            }
            
            res = append(res, str)
            prev = v
            prevIdx = i
        }
    }

    return res
}
