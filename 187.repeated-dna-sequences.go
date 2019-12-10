func findRepeatedDnaSequences(s string) []string {
    m := make(map[string]bool)
    for i := 0; i <= len(s)-10; i++ {
        if _, ok := m[s[i:i+10]]; ok {
            m[s[i:i+10]] = true
        } else {
            m[s[i:i+10]] = false
        }
    }
    result := []string{}
    for k, b := range m {
        if b {
            result = append(result, k)   
        }
    }
    return result
}