type Vertex struct {
    Word string
    Edges []*Vertex
    checked bool
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    words := make([]*Vertex, len(wordList)+1)
    words[0] = &Vertex{beginWord, []*Vertex{}, false}
    for i := range wordList {
        words[i+1] = &Vertex{wordList[i], []*Vertex{}, false}
    }
    
    for i := range words {
        for j := i+1; j < len(words); j++ {
            if canTransform(words[i].Word, words[j].Word) {
                words[i].Edges = append(words[i].Edges, words[j])
                words[j].Edges = append(words[j].Edges, words[i])
            }
        }
    }
    
    
    for cnt , queue := 0, words[0].Edges; len(queue) > 0; cnt++ {
        newQueue := []*Vertex{}
        for i := range queue {
            queue[i].checked = true
            if queue[i].Word == endWord {
                return cnt+2
            }
            for j := range queue[i].Edges {
                if !queue[i].Edges[j].checked {
                    queue[i].Edges[j].checked = true
                    newQueue = append(newQueue, queue[i].Edges[j])
                }
            }
        }
        queue = newQueue
    }
    return 0
}

func canTransform(a, b string) bool {
    charDiff := 0
    for i := range a {
        if a[i] != b[i] {
            charDiff ++
        }
        if charDiff > 1 {
            return false
        }
    }
    return true
}