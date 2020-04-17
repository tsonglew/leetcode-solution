func replaceSpace(s string) string {
    return strings.Join(strings.Split(s, " "), "%20")
}
