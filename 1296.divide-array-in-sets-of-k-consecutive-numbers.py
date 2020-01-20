class Solution:
    def isPossibleDivide(self, nums: List[int], k: int) -> bool:
        from collections import OrderedDict
        
        od = OrderedDict()
        for i in sorted(nums):
            if i in od:
                od[i]+=1
            else:
                od[i] = 1
        
        while len(od.keys()) > 0:
            cnt = k
            current = list(od.keys())[0]
            while cnt > 0:
                if current in od and od[current] > 0:
                    od[current]-=1
                else:
                    return False
                if od[current] == 0:
                    del od[current]
                current += 1
                cnt-=1
        return True
