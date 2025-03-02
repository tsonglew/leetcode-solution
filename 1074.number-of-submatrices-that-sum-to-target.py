class Solution:
    def numSubmatrixSumTarget(self, matrix: List[List[int]], target: int) -> int:
        res = 0
        for row1 in range(len(matrix)):
            temp = [0] * len(matrix[0])
            for row2 in range(row1, len(matrix)):
                m = {0: 1}
                preSum = 0
                for col in range(len(matrix[0])):
                    # print("***(%d,0),(%d,%d)"%(row1,row2,col))
                    temp[col] += matrix[row2][col]
                    preSum += temp[col]
                    res += m.get(preSum - target, 0)
                    m[preSum] = m.get(preSum, 0) + 1
                    # print(res, temp, m)
        return res
