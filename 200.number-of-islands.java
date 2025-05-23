class Solution {

    static class Pair {
        public int l;
        public int r;

        public Pair(int l, int r) {
            this.l = l;
            this.r = r;
        }
    }

    public int numIslands(char[][] grid) {
        Queue<Pair> queue = new LinkedList<>();
        boolean[][] visited = new boolean[grid.length][grid[0].length];
        int ans = 0;
        int[] dx = new int[]{0, 1, 0, -1};
        int[] dy = new int[]{-1, 0, 1, 0};
        
        for (int i = 0; i < visited.length; i++) {
            for (int j = 0; j < visited[0].length; j++) {
                if (visited[i][j]) {
                    continue;
                }
                if (grid[i][j] != '1') {
                    visited[i][j] = true;
                    continue;
                }
                ans ++;
                queue.add(new Pair(i, j));
                while (!queue.isEmpty()) {
                    Pair p = queue.poll();
                    if (visited[p.l][p.r]) {
                        continue;
                    }
                    if (grid[p.l][p.r] != '1') {
                        continue;
                    }
                    visited[p.l][p.r] = true;
                    for (int di = 0; di < 4; di++) {
                        int nl = p.l + dx[di];
                        int nr = p.r + dy[di];

                        if (nl >= 0 && nl < grid.length && nr >= 0 && nr < grid[0].length) {
                            queue.add(new Pair(nl, nr));
                        }
                    }
                }
            }
        }
        return ans;

    }
}

