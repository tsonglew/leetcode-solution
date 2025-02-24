class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> res = new ArrayList<>();
        HashMap<String, Integer> splitWordsMap = new HashMap<>();
        for (int i = 0; i < words.length; i++) {
            splitWordsMap.put(words[i], splitWordsMap.getOrDefault(words[i], 0) + 1);
        }

        // System.out.println("==== splitWordsMap ====");
        // splitWordsMap.forEach((key, value) -> {
        // System.out.printf("%s:%d\n", key, value);
        // });
        // System.out.println("=======================");

        for (int i = 0; i <= s.length() - words.length * words[0].length(); i++) {
            String subs = s.substring(i, i + words[0].length() * words.length);
            HashMap<String, Integer> sMap = new HashMap<>();
            for (int j = 0; j < (int) (subs.length() / words[0].length()); j++) {
                String subss = subs.substring(j * words[0].length(), j * words[0].length() + words[0].length());
                sMap.put(subss, sMap.getOrDefault(subss, 0) + 1);
            }
            // System.out.println("==== sMap ====");
            // sMap.forEach((key, value) -> {
            // System.out.printf("%s:%d\n", key, value);
            // });
            // System.out.println("=======================");

            if (splitWordsMap.size() != sMap.size()) {
                continue;
            }
            boolean next = false;
            for (Map.Entry<String, Integer> e : splitWordsMap.entrySet()) {
                if (!e.getValue().equals(sMap.getOrDefault(e.getKey(), -1))) {
                    next = true;
                    break;
                }
            }
            if (next) {
                continue;
            }
            res.add(i);
        }
        return res;
    }
}