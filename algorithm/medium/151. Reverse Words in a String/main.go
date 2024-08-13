func reverseWords(s string) string {
    result := ""
    tmp := ""
    for i,v := range(s) {
        if i >= len(s) -1 {
            if string(v) != " " {
                if len(result) == 0 {
                    result = tmp + string(v)
                } else {
                    result = tmp + string(v) + " " + result
                }
                break
            }
        }

        if string(v) == " " && tmp != ""  {
            if len(result) == 0 {
                result = tmp 
            } else {
                result = tmp + " " + result 
            }
            tmp = ""
        } else if string(v) != " " {
            tmp += string(v)
        }

        fmt.Println(tmp)
    }

    return result
}
