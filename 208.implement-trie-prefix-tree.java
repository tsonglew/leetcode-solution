class Trie {

    class TrieNode {
        public Boolean isWord;
        public HashMap<Character, TrieNode> m;

        public TrieNode() {
            this.isWord = false;
            this.m = new HashMap<>();
        }
    }

    public TrieNode head;

    public Trie() {
        this.head = new TrieNode();
    }

    public void insert(String word) {
        TrieNode head = this.head;
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            if (!head.m.containsKey(c)) {
                head.m.put(c, new TrieNode());
            }
            head = head.m.get(c);
        }
        head.isWord = true;
    }

    public boolean search(String word) {
        TrieNode head = this.head;
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            if (!head.m.containsKey(c)) {
                return false;
            }
            head = head.m.get(c);
        }
        return head.isWord;
    }

    public boolean startsWith(String prefix) {
        TrieNode head = this.head;
        for (int i = 0; i < prefix.length(); i++) {
            char c = prefix.charAt(i);
            if (!head.m.containsKey(c)) {
                return false;
            }
            head = head.m.get(c);
        }
        return this.isPrefix(head);
    }

    private boolean isPrefix(TrieNode head) {
        if (head.isWord) {
            return true;
        }
        for (Map.Entry<Character, TrieNode> entry : head.m.entrySet()) {
            if (this.isPrefix(entry.getValue())) {
                return true;
            }
        }
        return false;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */