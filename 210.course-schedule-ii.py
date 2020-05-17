class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        courses = [[] for i in range(numCourses)]
        preNum = [0 for i in range(numCourses)]
        for p in prerequisites:
            courses[p[1]].append(p[0])
            preNum[p[0]] += 1
        q = []
        result = []
        for i, v in enumerate(preNum):
            if v == 0:
                q.append(i)
        for i in q:
            result.append(i)
            for j in courses[i]:
                preNum[j] -= 1
                if preNum[j] == 0:
                    q.append(j)

        return result if len(result) == numCourses else []
