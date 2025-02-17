class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<Integer, List<String>> m = new HashMap<>();
        for (int i = 0; i < strs.length; i++) {
            int hashCode = this.calcHashCode(strs[i]);
            if (!m.containsKey(hashCode)) {
                m.put(hashCode, new ArrayList<String>());
            }
            m.get(hashCode).add(strs[i]);
        }
        List<List<String>> res = new ArrayList<>();
        m.forEach((key, value) -> {
            res.add(value);
        });
        return res;
    }

    public int calcHashCode(String s) {
        int[] arr = new int[26];
        for (int i = 0; i < s.length(); i++) {
            arr[s.charAt(i) - 'a']++;
        }
        return Arrays.hashCode(arr);
    }
}