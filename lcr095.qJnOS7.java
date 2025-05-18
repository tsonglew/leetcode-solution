class Solution {
    int[][] length;

    public int longestCommonSubsequence(String text1, String text2) {
        this.length = new int[text1.length()][text2.length()];
        for (int i = 0; i < text1.length(); i++) {
            for (int j = 0; j < text2.length(); j++) {
                int val = 0;
                if (text1.charAt(i) == text2.charAt(j)) {
                    val = 1;
                    this.length[i][j] = val;
                }

                if (i - 1 >= 0 && j - 1 >= 0) {
                    this.length[i][j] = this.length[i-1][j-1]+val;
                }

                if (i - 1 >= 0) {
                    this.length[i][j] = Math.max(this.length[i-1][j], this.length[i][j]);
                }

                if (j - 1 >= 0) {
                    this.length[i][j] = Math.max(this.length[i][j-1], this.length[i][j]);
                }
                
            }
        }
        // this.iter();
        return this.length[this.length.length-1][this.length[0].length-1];

    }

    private void iter() {
        for (int i = 0; i < this.length.length; i++) {
            for (int j = 0; j < this.length[0].length; j++) {
                System.out.printf("%d ", this.length[i][j]);
            }
            System.out.println();
        }
    }

}
