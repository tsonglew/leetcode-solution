class Solution {
    private int m;
    private int n;
    private HashMap<Integer, ArrayList<Integer>> to;
    private int[] dx;
    private int[] dy;
    private int[] degree;
    private int ans;
    private int[] dist;

    public int longestIncreasingPath(int[][] matrix) {
        this.m = matrix.length;
        this.n = matrix[0].length;
        this.to = new HashMap<>();
        this.dx = new int[]{0, 0, -1, 1};
        this.dy = new int[]{-1, 1, 0, 0};
        this.degree = new int[this.m*this.n];
        this.ans = 1;
        this.dist = new int[this.m*this.n];
        this.setUpTopo(matrix);
        // System.out.println("finish setupt topo");
        this.bfs();
        return this.ans;
    }

    private int hash(int i, int j) {
        return i * this.m + j;
    }

    private boolean valid(int i, int j) {
        return i >= 0 && i < this.m && j >= 0 && j < this.n;
    }

    private void setUpTopo(int[][] matrix) {
        for (int i = 0; i < this.m; i++) {
            for (int j = 0; j < this.n; j++) {
                int hashij = this.hash(i, j);
                this.to.put(hashij, new ArrayList<>());
                for (int d = 0; d < this.dx.length; d++) {
                    int ni = i + dx[d];
                    int nj = j + dy[d];
                    if (!this.valid(ni, nj)) {
                        continue;
                    }
                    if (matrix[ni][nj] > matrix[i][j]) {
                        this.to.get(hashij).add(this.hash(ni, nj));
                        this.degree[hashij] ++;
                    }
                }  
            }
        }
    }

    private void bfs() {
        Queue<Integer> q = new LinkedList<>();

        for (int i = 0; i < this.degree.length; i++) {
            if (this.degree[i] == 0) {
                continue;
            }
            q.add(this.degree[i]);
            int sum = 1;
            while (!q.isEmpty()) {
                Integer ele = q.poll();
                // System.out.printf("ele: %d\n", ele);
                this.dist[ele] = Math.max(this.dist[ele], sum + 1);
                this.ans = Math.max(this.dist[ele], this.ans);
                int row = (int) Math.floor(ele / this.m);
                int col = ele % this.m;
                for (Integer nexthash: this.to.get(ele)) {
                    q.add(nexthash);
                }                
                sum++;
                
            }
        }
    }
}
