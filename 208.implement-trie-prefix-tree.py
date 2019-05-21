class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.dict = dict()

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        r = self.dict
        for i in word:
            if i not in r.keys():
                r[i] = dict()
            r = r[i]
        r['$'] = None

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        r = self.dict
        for i in word:
            if i not in r:
                return False
            r = r[i]
        return '$' in r

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        r = self.dict
        for i in prefix:
            if i not in r:
                return False
            r = r[i]
        return True


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)
