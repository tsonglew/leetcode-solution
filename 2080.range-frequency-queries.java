class RangeFreqQuery {

    private HashMap<Integer, List<Integer>> m;

    public RangeFreqQuery(int[] arr) {
        this.m = new HashMap<>();
        for (int i = 0; i < arr.length; i++) {
            if (!this.m.containsKey(arr[i])) {
                this.m.put(arr[i], new ArrayList<>());
            }
            this.m.get(arr[i]).add(i);
        }
    }

    public int query(int left, int right, int value) {
        if (!this.m.containsKey(value)) {
            return 0;
        }
        List<Integer> l = this.m.get(value);
        int li = this.f1(l, left);
        int ri = this.f2(l, right);
        // System.out.printf("leftIndex: %d, rightIndex: %d\n", li, ri);
        return ri - li + 1;
    }

    public int f1(List<Integer> l, int val) {
        // 找最小的 i 使 l[i]>=val
        int left = 0, right = l.size() - 1;
        int ans = l.size();
        while (left <= right) {
            int mid = (left + right) >> 1;
            if (l.get(mid) == val) {
                return mid;
            } else if (l.get(mid) < val) {
                left = mid + 1;
            } else {
                // l[mid] > val
                ans = Math.min(ans, mid);
                right = mid - 1;
            }
        }
        return ans;
    }

    public int f2(List<Integer> l, int val) {
        // 找最大的 i 使 l[i]<=val
        int left = 0, right = l.size() - 1;
        int ans = -1;
        while (left <= right) {
            int mid = (left + right) >> 1;
            if (l.get(mid) == val) {
                return mid;
            } else if (l.get(mid) < val) {
                ans = Math.max(ans, mid);
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return ans;
    }
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * RangeFreqQuery obj = new RangeFreqQuery(arr);
 * int param_1 = obj.query(left,right,value);
 */