class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
        distances = collections.defaultdict(dict)
        for (a, b), v in zip(equations, values):
            distances[a][b] = v
            distances[b][a] = 1 / v
            distances[a][a] = distances[b][b] = 1
        for i in distances:
            for j in distances[i]:
                for k in distances[i]:
                    distances[j][k] = distances[j][i] * distances[i][k]
                    distances[k][j] = 1 / distances[j][k]
        return [distances.get(e[0], {e[1]:-1.0}).get(e[1], -1.0) for e in queries]
            