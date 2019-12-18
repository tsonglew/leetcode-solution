type Node struct {
	Next []*Node
	C byte
}

type WordDictionary struct {
    Head *Node
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{&Node{Next: []*Node{}}}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	word = word + "$"
    for i, p := 0, this.Head; i < len(word); i++{
		cp := p
		for j := range p.Next {
			if p.Next[j].C == word[i] {
				p = p.Next[j]
				break
			}
		}
		if p == cp {
			nNode := &Node{Next: []*Node{}, C: word[i]}
			p.Next = append(p.Next, nNode)
			p = nNode
		}
	}
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	word = word + "$"
	for i := range this.Head.Next {
		if this.Head.Next[i].Search(word) {
			return true
		}
	}
	return false
}

func (n *Node) Search(s string) bool {
	if s[0] == '.' || s[0] == n.C {
		if len(s) == 1 {
			return true
		}
		for i := range n.Next {
			if n.Next[i].Search(s[1:]) {
				return true
			}
		}
	}
	return false
}