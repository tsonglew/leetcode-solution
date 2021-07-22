class Solution:
    def isNumber(self, s: str) -> bool:
        return bool(re.search("^(\+|-)?((\d+\.?\d+)|(\d+\.?)|(\.?\d+))((e|E)(\+|-)?\d+)?$", s.strip()))
