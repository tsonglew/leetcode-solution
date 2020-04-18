class Solution:
    def getTriggerTime(self, increase: List[List[int]], requirements: List[List[int]]) -> List[int]:
        class Requirement:
            def __init__(self, req):
                self.req = req
                self.d = -1
            def ok(self):
                return sum(self.req) == 0
        
        reqs = [Requirement(i) for i in requirements]
        sorted_reqs = [
            sorted(reqs, key=lambda x: x.req[0]),
            sorted(reqs, key=lambda x: x.req[1]),
            sorted(reqs, key=lambda x: x.req[2]),
        ]
        acc = [0, 0, 0]

        for i in range(len(increase)):
            for idx in range(3):
                while True:
                    if len(sorted_reqs[idx])==0:
                        break
                    if sorted_reqs[idx][0].req[idx] <= acc[idx]:
                        sorted_reqs[idx][0].req[idx] = 0
                        if sorted_reqs[idx][0].ok():
                            sorted_reqs[idx][0].d = i
                        sorted_reqs[idx].pop(0)
                    else:
                        break
            for a in range(3):
                acc[a] += increase[i][a]
        for idx in range(3):
            while True:
                if len(sorted_reqs[idx])==0:
                    break
                if sorted_reqs[idx][0].req[idx] <= acc[idx]:
                    sorted_reqs[idx][0].req[idx] = 0
                    if sorted_reqs[idx][0].ok():
                        sorted_reqs[idx][0].d = len(increase)
                    sorted_reqs[idx].pop(0)
                else:
                    break
        return [r.d for r in reqs]
                
           
