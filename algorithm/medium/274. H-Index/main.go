func hIndex(citations []int) int {
    arr := make([]int, 1005)

    for _,v := range(citations) {
        arr[v] += 1
    }

    for i,_ := range(arr) {
        if i == 0 {
            continue
        }

        arr[len(arr) - i - 1] += arr[len(arr) - i]

        if len(arr) - i <= arr[len(arr) - i] {
            return len(arr) - i 
        }
    }

    fmt.Println(arr)
    
    return 0
}
