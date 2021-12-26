class Solution:
    def findOcurrences(self, text: str, first: str, second: str) -> List[str]:
        text = text.split()
        return [text[i+2] for i in filter(lambda i: text[i]==first and text[i+1]==second,range(0, len(text)-2))]
