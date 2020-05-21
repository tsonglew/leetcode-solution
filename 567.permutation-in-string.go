func checkInclusion(s1 string, s2 string) bool {
    cm := make(map[byte]int)
    wc := 0
    for i := range s1 {
        if _, ok := cm[s1[i]]; !ok {
            cm[s1[i]] = 0
        }
        cm[s1[i]]++
        wc++
    }
    i, j := 0, 0
    for j < len(s2) {
        if v, ok := cm[s2[j]]; ok && v > 0 {
            cm[s2[j]]--
            j++
            wc--
            if wc == 0 {
                return true
            }
        } else {
            n := 0
            for j < len(s2) {
                if _, ok := cm[s2[j]]; ok { 
                    break 
                }
                n++
                j++
            }
            if n > 0 {
               for i < j && i < len(s2) {
                if _, ok := cm[s2[i]]; ok {
                    cm[s2[i]] ++
                    wc++
                }
                i++
               }
            } else {
                cm[s2[i]] ++
                wc++
                i++
            }            
        }
    }
    return false
}
