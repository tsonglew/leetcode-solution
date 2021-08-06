class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        if len(pushed) != len(popped):
            return False
        num_index_map = {}
        for idx, num in enumerate(pushed):
            num_index_map[num] = idx
        popped = [num_index_map[num] for num in popped]
        return self.validate(popped)

    def validate(self, l):
        if len(l) <= 1:
            return True
        less = l[0]
        for i in l[1:]:
            if i < l[0]:
                if i > less:
                    return False
                less = i
        return self.validate(l[1:])
            
