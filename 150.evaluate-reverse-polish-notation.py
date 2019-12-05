class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        if len(tokens) == 0:
            return 0
        i = 0
        while i < len(tokens) and len(tokens) > 1:
            if tokens[i] == "+":
                tokens[i-2] += tokens[i-1]
            elif tokens[i] == "-":
                tokens[i-2] -= tokens[i-1]
            elif tokens[i] == "*":
                tokens[i-2] *= tokens[i-1]
            elif tokens[i] == "/":
                tokens[i-2] = int(tokens[i-2]/tokens[i-1])
            else:
                tokens[i] = int(tokens[i])
                i += 1
                continue
            i = i-1
            tokens = tokens[:i] + tokens[i+2:]
        return tokens[0]