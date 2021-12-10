func shortestCompletingWord(licensePlate string, words []string) string {
    licencesMap := makeByteMap(licensePlate)
    result := ""
    for i := range words {
        wordsMap := makeByteMap(words[i])
        flag := true
        print(wordsMap)
        for j, _ := range licencesMap {
            if v, _ := wordsMap[j]; v < licencesMap[j] {
                flag = false
                break
            }
        }
        if flag && (result == "" || len(words[i]) < len(result)){
            result = words[i]
        }
    }
    return result
}

func makeByteMap(licensePlate string) map[byte]int {
    licencesMap := make(map[byte]int, 26)
    for i := range licensePlate {
        if licensePlate[i] < 'A' || licensePlate[i] > 'z' {
           continue
        }
        v := licensePlate[i]
        if v > 'Z' {
            v -= ('a'-'A')
        }
        if _, ok := licencesMap[v]; ok {
            licencesMap[v] ++
        } else {
            licencesMap[v] = 1
        }
    }
    return licencesMap
}
