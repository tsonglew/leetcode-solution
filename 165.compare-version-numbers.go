func compareVersion(version1 string, version2 string) int {
    version1s := strings.Split(version1, ".")
    version2s := strings.Split(version2, ".")
    i := 0
    for ; i < len(version1s); i++ {
        num1, _ := strconv.Atoi(version1s[i])
        if i >= len(version2s) {
			if num1 == 0 {
				continue
			}
            return 1
        }
        num2, _ := strconv.Atoi(version2s[i])
        if num1 > num2  {
            return 1
        } else if num1 < num2 {
            return -1
        }
    }
    for j := i; j < len(version2s); j++ {
        num, _ := strconv.Atoi(version2s[j])
        if num != 0 {
            return -1
        }
    }
    return 0
}