class Solution:
    def constructArr(self, a: List[int]) -> List[int]:
        left_mul = [1]*len(a)
        right_mul = left_mul.copy()
        for i in range(len(a)):
            left_mul[i] = (left_mul[i-1]*a[i-1] if i > 0 else 1)
            right_mul[len(a)-1-i] = (right_mul[len(a)-i]*a[len(a)-i] if i>0 else 1)
        for i in range(len(left_mul)):
            right_mul[i] *= left_mul[i]
        return right_mul
