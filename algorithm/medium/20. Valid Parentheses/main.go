func isValid(s string) bool {
    tmp := ""

    for _,v := range(s) {
        if(len(tmp) > 0) {
            if string(v) == "(" || string(v) == "{" || string(v) == "[" {
                tmp += string(v)
            } else {
                if string(v) == ")" && string(tmp[len(tmp)-1]) != "(" {
                    return false
                }
                if string(v) == "]" && string(tmp[len(tmp)-1]) != "[" {
                    return false
                }
                if string(v) == "}" && string(tmp[len(tmp)-1]) != "{" {
                    return false
                }
                tmp = tmp[:len(tmp)-1]
            } 
        } else {
            if string(v) == "(" || string(v) == "{" || string(v) == "[" {
                tmp += string(v)
            } else {
                return false
            }
        }
    }

    if len(tmp) > 0 {
        return false
    }

    return true
}
