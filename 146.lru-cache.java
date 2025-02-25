class LRUCache {

    private Integer capacity;
    private HashMap<Integer, LinkedNode> m;
    private LinkedNode head;
    private LinkedNode end;

    static class LinkedNode {
        public Integer key;
        public Integer val;
        public LinkedNode prev;
        public LinkedNode next;

        public LinkedNode(Integer key, Integer val) {
            this.key = key;
            this.val = val;
            this.prev = null;
            this.next = null;
        }
    }

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.m = new HashMap<>();
        this.head = new LinkedNode(-1, -1);
        this.end = new LinkedNode(-1, -1);
        this.head.next = this.end;
        this.end.prev = this.head;
    }

    public int get(int key) {
        // System.out.printf("get key:%d\n", key);

        if (!this.m.containsKey(key)) {
            return -1;
        }
        this.moveFront(key);
        // this.iterate();
        return this.m.get(key).val;
    }

    public void put(int key, int value) {
        // System.out.printf("put key:%d, value:%d\n", key, value);
        if (this.m.containsKey(key)) {
            this.m.get(key).val = value;
            this.moveFront(key);
            return;
        }
        if (this.capacity.equals(this.m.size())) {
            LinkedNode endPrev = this.end.prev;
            LinkedNode end2ndPrev = endPrev.prev;
            end2ndPrev.next = this.end;
            this.end.prev = end2ndPrev;

            // System.out.printf("remove val: %d\n", endPrev.key);
            this.m.remove(endPrev.key);
        }
        LinkedNode newNode = new LinkedNode(key, value);
        LinkedNode first = this.head.next;
        newNode.prev = this.head;
        newNode.next = first;
        first.prev = newNode;
        this.head.next = newNode;
        this.m.put(key, newNode);
        // this.iterate();
    }

    private void moveFront(Integer key) {
        // System.out.printf("move %d to front\n", key);
        LinkedNode n = this.m.get(key);
        LinkedNode prev = n.prev;
        LinkedNode next = n.next;
        prev.next = next;
        next.prev = prev;
        this.head.next.prev = n;
        n.next = this.head.next;
        n.prev = this.head;
        this.head.next = n;
    }

    private void iterate() {
        LinkedNode i = this.head;
        while (i != null) {
            System.out.printf("%d ", i.key);
            i = i.next;
        }
        System.out.printf("\n");
        i = this.end;
        while (i != null) {
            System.out.printf("%d ", i.key);
            i = i.prev;
        }
        System.out.printf("\n");
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */