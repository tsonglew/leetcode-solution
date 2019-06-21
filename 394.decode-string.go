func decodeString(s string) string {
    num := 0
    r := ""
    for i := 0; i < len(s); i++ {
        if s[i] <= '9' && s[i] >= '0' {
            num *= 10
            num += int(s[i]) -48
        } else if s[i] == '[' {
            j := i +1
            for bracs := 1; bracs >0; j++ {
                if s[j] == '['{
                    bracs++
                }else if s[j] == ']'{
                    bracs--
                }
            }
            dr := decodeString(s[i+1:j-1])
            i=j-1
            if num == 0 {
                r += dr
            } else {
                r += strings.Repeat(dr, num)
            }
            num = 0
        } else {
            r += string(s[i])
        }
    }
    return r
}
