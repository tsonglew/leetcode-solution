package main

import "strings"

func main() {

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	dirs := make([]string, 0)
	for i := range paths {
		if paths[i] == ".." {
			dirs = dirs[:max(0, len(dirs)-1)]
		} else if paths[i] == "" || paths[i] == "." {
			continue
		} else {
			dirs = append(dirs, paths[i])
		}
	}
	return "/" + strings.Join(dirs, "/")
}
