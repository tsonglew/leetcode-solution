class TopVotedCandidate {

    int[] top;
    int[] times;

    public TopVotedCandidate(int[] persons, int[] times) {
        this.times = times;
        this.top = new int[persons.length];
        HashMap<Integer, Integer> topMap = new HashMap<>();
        Integer topPerson = -1;
        for (int i = 0; i < persons.length; i++) {
            int person = persons[i];
            topMap.computeIfAbsent(person, k -> 0);
            topMap.put(person, topMap.get(person) + 1);
            if (topPerson < 0 || topMap.get(person) >= topMap.get(topPerson)) {
                topPerson = person;
            }
            this.top[i] = topPerson;
        }
    }
    
    public int q(int t) {
        int l = 0, r = this.times.length - 1;
        int ans = 0;
        while (l <= r) {
            int mid = (int) Math.floor((l + r) / 2);
            if (this.times[mid] == t) {
                return this.top[mid];
            } else if (this.times[mid] > t) {
                r = mid - 1;
            } else {
                l = mid + 1;
                ans = Math.max(ans, mid);
            }
        }
        return this.top[ans];

    }


    private void printString() {
        for (int i = 0; i < this.top.length; i++) {
            System.out.printf("%d ", this.top[i]);
        }
        System.out.println();
    }
}
