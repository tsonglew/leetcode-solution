class Solution {

    static class RowCol {
        int row;
        int col;
        public RowCol(int row, int col) {
            this.row = row;
            this.col = col;
        }
    }

    ArrayList<RowCol> candidate;
    ArrayList<ArrayList<RowCol>> ans;
    HashMap<Integer, Boolean> usedCol;
    HashMap<Integer, Boolean> usedMinus;
    HashMap<Integer, Boolean> usedPlus;


    public Solution() {
        this.candidate = new ArrayList<>();
    }

    public List<List<String>> solveNQueens(int n) {

        this.usedCol = new HashMap<>();
        this.usedMinus = new HashMap<>();
        this.usedPlus = new HashMap<>();
        this.ans = new ArrayList<>();

        this.dfs(0, n);
        List<List<String>> result = new ArrayList<>();

        for (int i = 0; i < this.ans.size(); i++) {
            ArrayList<String> cache = new ArrayList<>();
            for (int j = 0; j < n; j++) {
                cache.add(this.buildQueenString(this.ans.get(i).get(j), n));
            }
            result.add(cache);
        }
        return result;
    }

    private void dfs(int row, int n) {
        if (row == n) {
            this.ans.add((ArrayList<RowCol>) this.candidate.clone());
        }
        for (int col = 0; col < n; col++) {
            if (this.usedCol.getOrDefault(col, false)) {
                continue;
            }
            int minus = row - col;
            if (this.usedMinus.getOrDefault(minus, false)) {
                continue;
            }
            int plus = row + col;
            if (this.usedPlus.getOrDefault(plus, false)) {
                continue;
            }
            this.usedCol.put(col, true);
            this.usedMinus.put(minus, true);
            this.usedPlus.put(plus, true);
            this.candidate.add(new RowCol(row, col));
            this.dfs(row+1, n);
            this.usedCol.put(col, false);
            this.usedMinus.put(minus, false);
            this.usedPlus.put(plus, false);
            this.candidate.remove(this.candidate.size() - 1);
        }
    }

    private String buildQueenString(RowCol rc, int n) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < n; i++) {
            sb.append(i == rc.col ? "Q" : ".");
        }
        return sb.toString();
    }
}
