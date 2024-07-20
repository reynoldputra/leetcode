func recr(arr1 []int, arr2 []int ) [][]int {
    var result [][]int

    if len(arr2) == 0 {
        result = append(result, arr1)
        fmt.Println("break", arr1)
        return result
    } else {
        for i,_ := range(arr2) {
            fmt.Println(i, arr1, arr2, arr2[:i], arr2[i+1:])
            pop := arr2[i]
            newarr1 := append([]int(nil), arr1...)
            newarr1 = append(newarr1, pop)
            newarr2 := append([]int(nil), arr2[:i]...)
            newarr2 = append(newarr2, arr2[i+1:]...)
            result_recr := recr(newarr1, newarr2)
            result = append(result, result_recr...)
        }
    }

    return result
}

func permute(nums []int) [][]int {
    return recr([]int{}, nums)
}
