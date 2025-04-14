
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Queue;

class Solution {

    static class Pair {
        public Integer left;
        public Integer right;

        public Pair(Integer left, Integer right) {
            this.left = left;
            this.right = right;
        }
    }

    int m;
    int n;
    int[][] direction = { { 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 } };
    HashMap<Integer, ArrayList<Integer>> to;
    HashMap<Integer, Integer> dist;

    private void init(int[][] matrix) {
        this.m = matrix.length;
        this.n = matrix[0].length;
        this.to = new HashMap<>(this.m * this.n);
        this.dist = new HashMap<>(this.m * this.n);
    }

    public int longestIncreasingPath(int[][] matrix) {
        this.init(matrix);
        this.setTopo(matrix);

        for (Map.Entry<Integer, ArrayList<Integer>> e : this.to.entrySet()) {
            if (e.getValue().isEmpty()) {
                continue;
            }

        }
    }

    private void bfs(Integer hashij) {
        int dist = 1;

        this.dist.put(hashij, Math.max(dist, this.dist.getOrDefault(hashij, 0)));
        Queue<Pair> q = new ArrayList<>();
        q.add(new Pair(hashij, this.dist.get(hashij)));
        while (!q.isEmpty()) {
            Pair p = q.poll();
            int hashd = p.left;
            int dist = p.right;
            int i = (int) Math.floor(hashij / this.m);
            int j = hashij % this.m;

        }

    }

    private void setTopo(int[][] matrix) {
        for (int i = 0; i < this.m; i++) {
            for (int j = 0; j < this.n; j++) {
                for (int[] d : this.direction) {
                    int next_i = i + d[0], next_j = j + d[1];
                    if (!this.valid(next_i, next_j)) {
                        continue;
                    }
                    int hashij = this.hash(i, j);
                    this.to.putIfAbsent(hashij, k -> new ArrayList<>());
                    if (matrix[next_i][next_j] > matrix[i][j]) {
                        this.to.get(hashij).add(this.hash(next_i, next_j));
                    }
                }

            }
        }

    }

    private int hash(int i, int j) {
        return this.m * i + j;
    }

    private boolean valid(int i, int j) {
        return i >= 0 && i < this.m && j >= 0 && j < this.n;
    }
}
