class Solution:
    def removeInvalidParentheses(self, s: str) -> List[str]:
        def isValid(s):
            stack = []
            for i in s:
                if i == '(':
                    stack.append(i)
                elif i == ')':
                    if len(stack) > 0 and stack[-1] == '(':
                        stack = stack[:-1]
                    else:
                        return False
            return len(stack) == 0
            
        def removeRecurse(s, leftCnt, rightCnt):
            if leftCnt == 0 and rightCnt == 0:
                return set([s])
            if leftCnt + rightCnt > len(s):
                return set()
            result = set()
            for i,c in enumerate(s):
                if leftCnt > 0 and c == '(':
                    removed = removeRecurse(s[i+1:], leftCnt-1, rightCnt)
                elif rightCnt > 0 and c == ')':
                    removed = removeRecurse(s[i+1:], leftCnt, rightCnt-1)
                else:
                    continue
                if len(removed) > 0:
                    for j in removed:
                        result.add(s[:i]+j)
            return result
               
        stack = []
        leftCnt = 0
        rightCnt = 0
        for i in s:
            if i == '(':
                stack.append(i)
            elif i == ')':
                if len(stack) > 0 and stack[-1] == '(':
                    stack = stack[:-1]
                else:
                    stack.append(i)
        for i in stack:
            if i == '(':
                leftCnt += 1
            else:
                rightCnt += 1
        
        removed = removeRecurse(s, leftCnt, rightCnt)
        return [s for s in removed if isValid(s)]
        
        
                