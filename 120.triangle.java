class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        for (int level = triangle.size()-2; level >= 0; level--) {
            for (int i = 0; i < triangle.get(level).size(); i++) {
                triangle.get(level).set(
                    i, 
                    triangle.get(level).get(i) + Math.min( 
                        triangle.get(level+1).get(i), 
                        triangle.get(level+1).get(i+1)
                    )
                );
            }
        }
        return triangle.get(0).get(0);
    }
}
