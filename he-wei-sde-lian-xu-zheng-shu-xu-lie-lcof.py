class Solution:
    def findContinuousSequence(self, target: int) -> List[List[int]]:
        i, j = 1, 2
        result = []
        while j <= target // 2 + 2:
            sm = (i+j)*(j-i+1)/2
            if sm == target:
                result.append(list(range(i,j+1)))
                j+=1
            elif sm < target:
                j+=1
            else:
                i+=1
        return result
