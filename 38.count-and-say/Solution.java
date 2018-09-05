import java.util.HashMap;

public class Solution {
    private static HashMap<Integer, String> hm = new HashMap<Integer, String>();
    {
        hm.put(1, "1");
        hm.put(2, "11");
        hm.put(3, "21");
        hm.put(4, "1211");
        hm.put(5, "111221");
    }

    public String countAndSay(int n) {
        if (n <= 0) {
            return null;
        }
        if (hm.containsKey(n)) {
            return hm.get(n);
        }
        for (int i = 6; i <= n; ++i) {
            StringBuilder sb = new StringBuilder();
            String preValue = hm.get(i-1);
            int cnt = 1;
            for (int j = 1; j < preValue.length(); ++j) {
                if (preValue.charAt(j-1) == preValue.charAt(j)) {
                    cnt += 1;
                } else {
                    sb.append(Integer.toString(cnt));
                    sb.append(preValue.charAt(j-1));
                    cnt = 1;
                }
            }
            sb.append(Integer.toString(cnt));
            sb.append(preValue.charAt(preValue.length()-1));
            hm.put(i, sb.toString());
        }
        return hm.get(n);
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.countAndSay(7));
    }
}
