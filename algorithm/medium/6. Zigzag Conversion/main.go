func convert(s string, numRows int) string {
    res := ""
    for j := 0; j < numRows; j++ {
        i := j
        for i < len(s) {
            if j == 0 || j == numRows - 1 {
                res += string(s[i])
                if numRows == 1 {
                    i += 1
                } else {
                    i += (numRows - 1) *2
                }
                continue
            } 

            res += string(s[i])

            // bottom
            bottom := (numRows - j - 1) * 2
            i = i + bottom
            if i < len(s) {
                res += string(s[i])
            } else {
                break
            }
            
            // top
            top := j * 2
            if top < 0 {
                top = 0
            }
            i = i + top
            if i >= len(s) {
                break
            }
        }
    }

    return res
}
