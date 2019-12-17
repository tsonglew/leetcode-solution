func isIsomorphic(s string, t string) bool {
	m1, m2 := make(map[byte]byte), make(map[byte]byte)
	n1, n2 := make([]byte, len(s)), make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}
	for i := 0; i < len(s); i++ {
		n1[i] = m1[s[i]]
		n2[i] = m2[t[i]]
	}
	return string(n1) == t && string(n2) == s
}