class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        l = []
        poped_i = 0
        for i in pushed:
            l.append(i)
            while len(l) > 0 and poped_i < len(popped) and l[-1] == popped[poped_i]:
                l.pop()
                poped_i += 1
        return len(l) == 0
            
