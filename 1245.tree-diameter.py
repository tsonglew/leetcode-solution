class Solution:
    def treeDiameter(self, edges: list[list[int]]) -> int:
        edges_m = self.init_edges_m(edges)
        start = list(edges_m.keys())[0]
        node, distance = self.find_furthest(start, edges_m, set())
        _, distance = self.find_furthest(node, edges_m, set())
        return distance

    def init_edges_m(self, edges: list[list[int]]):
        edges_m = {}
        for edge in edges:
            if edge[0] not in edges_m:
                edges_m[edge[0]] = []
            edges_m[edge[0]].append(edge[1])
            if edge[1] not in edges_m:
                edges_m[edge[1]] = []
            edges_m[edge[1]].append(edge[0])
        return edges_m

    def find_furthest(self, i, edges_m: dict[int, list[int]], visited: set[int]):
        furthest_node = 0
        furthest_distance = 0
        visited.add(i)
        for node in edges_m[i]:
            if node in visited:
                continue
            node, distance = self.find_furthest(node, edges_m, visited)
            if distance > furthest_distance:
                furthest_node = node
                furthest_distance = distance
        return furthest_node, furthest_distance + 1


# Test Cases
if __name__ == "__main__":
    solution = Solution()
    assert solution.treeDiameter([[0, 1], [0, 2]]) == 2
    assert solution.treeDiameter([[0, 1], [1, 2], [2, 3], [1, 4], [4, 5]]) == 4
