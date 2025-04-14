class Solution {
    int m;
    int n;
    int[][] direction = { { 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 } };
    HashMap<Integer, ArrayList<Integer>> to;
    HashMap<Integer, Integer> dist;
    int ans;
    int[] indegree;
    Queue<Integer> q;

    private void init(int[][] matrix) {
        this.m = matrix.length;
        this.n = matrix[0].length;
        this.to = new HashMap<>(this.m * this.n);
        this.dist = new HashMap<>(this.m * this.n);
        this.ans = 1;
        this.indegree = new int[this.m * this.n];
        this.q = new LinkedList<>();
    }

    public int longestIncreasingPath(int[][] matrix) {
        this.init(matrix);
        this.setTopo(matrix);
        
        for (int i = 0; i < this.indegree.length; i++) {
            if (this.indegree[i] == 0) {
                this.dist.putIfAbsent(i, 1);
                this.q.add(i);
                
            }
        }
        this.bfs();
        return this.ans;
    }

    private void bfs() {
        while (!this.q.isEmpty()) {
            Integer startNode = this.q.poll();
            this.ans = Math.max(this.ans, this.dist.get(startNode));
            ArrayList<Integer> toList = this.to.get(startNode);
            if (toList == null) {
                continue;
            }
            for (Integer to : toList) {
                this.dist.put(to, Math.max(this.dist.getOrDefault(to, 1), this.dist.get(startNode)+1));
                this.indegree[to]--;
                if (this.indegree[to] == 0) {
                    q.add(to);
                }

            }
        }
    }

    private void setTopo(int[][] matrix) {
        for (int i = 0; i < this.m; i++) {
            for (int j = 0; j < this.n; j++) {
                int hashij = this.hash(i, j);
                this.to.put(hashij, new ArrayList<Integer>());
                for (int[] d : this.direction) {
                    int next_i = i + d[0], next_j = j + d[1];
                    if (!this.valid(next_i, next_j)) {
                        continue;
                    }
                    if (matrix[next_i][next_j] > matrix[i][j]) {
                        this.to.get(hashij).add(this.hash(next_i, next_j));
                    } else if (matrix[next_i][next_j] < matrix[i][j]) {
                        this.indegree[hashij]++;
                    }
                }

            }
        }

    }

    private int hash(int i, int j) {
        return this.n * i + j;
    }

    private boolean valid(int i, int j) {
        return i >= 0 && i < this.m && j >= 0 && j < this.n;
    }

    private void iterList(List<Integer> l) {
        for (Integer i : l) {
            System.out.printf("%d ", i);
        }
        System.out.println();
    }
}
