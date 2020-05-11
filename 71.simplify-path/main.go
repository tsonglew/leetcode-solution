package main

import "strings"

func simplifyPath(path string) string {
    dirs := make([]string, 0)
    pts := strings.Split(path, "/")
    for i := range pts {
        if pts[i] == ".." && len(dirs) > 0 {
            dirs = dirs[:len(dirs)-1]
        } else if len(pts[i]) > 0 && (!strings.HasPrefix(pts[i], ".")||len(pts[i])>2) {
            dirs = append(dirs, pts[i])
        }
    }
    return "/" + strings.Join(dirs, "/")
}
