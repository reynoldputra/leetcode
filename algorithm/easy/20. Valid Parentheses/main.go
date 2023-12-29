func isValid(s string) bool {
    if len(s) < 2 {
        return false
    }

    t := ""
    t = fmt.Sprintf("%s%c", t, s[0])

    for i := 1; i < len(s); i++ {
        if s[i] == '{' || s[i] == '(' || s[i] == '[' {
            t = fmt.Sprintf("%s%c", t, s[i])
            continue
        }

        if len(t) < 1 {
            return false
        }

        
        if s[i] == ')'  {
            if t[len(t)-1:] == "("  {
                t = t[:len(t)-1]
                continue
            } else {
                return false
            }
        }

        if s[i] == ']'  {
            if t[len(t)-1:] == "["  {
                t = t[:len(t)-1]
            } else {
                return false
            }
        }

        if s[i] == '}'  {
            if t[len(t)-1:] == "{"  {
                t = t[:len(t)-1]
            } else {
                return false
            }
        }
    }

    if len(t) > 0 {
        return false
    }

    return true
}
