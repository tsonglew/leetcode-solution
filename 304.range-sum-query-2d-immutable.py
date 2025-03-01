class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        self.cache = {}
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                if i == 0 and j == 0:
                    self.cache[(i, j)] = matrix[i][j]
                elif i == 0:
                    self.cache[(i, j)] = matrix[i][j] + self.cache[(i, j - 1)]
                elif j == 0:
                    self.cache[(i, j)] = matrix[i][j] + self.cache[(i - 1, j)]
                else:
                    self.cache[(i, j)] = (
                        matrix[i][j] + self.cache[(i - 1, j)] + self.cache[(i, j - 1)] - self.cache[(i - 1, j - 1)]
                    )

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        res = self.cache[(row2, col2)]
        if row1 == 0 and col1 == 0:
            pass
        elif row1 == 0:
            res -= self.cache[(row2, col1 - 1)]
        elif col1 == 0:
            res -= self.cache[(row1 - 1, col2)]
        else:
            res = res - self.cache[(row2, col1 - 1)] - self.cache[(row1 - 1, col2)] + self.cache[(row1 - 1, col1 - 1)]
        return res


# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)
