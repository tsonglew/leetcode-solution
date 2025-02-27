class BrowserHistory {

    static class HistoryNode {
        public String url;
        public HistoryNode prev;
        public HistoryNode next;

        public HistoryNode(String url) {
            this.url = url;
            this.prev = null;
            this.next = null;
        }
    }

    private HistoryNode ptr;

    public BrowserHistory(String homepage) {
        this.insert(homepage);
    }

    public void visit(String url) {
        this.insert(url);
    }

    public String back(int steps) {
        for (int i = 0; i < steps && this.ptr.prev != null; i++) {
            this.ptr = this.ptr.prev;
        }
        return this.ptr.url;
    }

    public String forward(int steps) {
        for (int i = 0; i < steps && this.ptr.next != null; i++) {
            this.ptr = this.ptr.next;
        }
        return this.ptr.url;
    }

    private void insert(String url) {
        HistoryNode newNode = new HistoryNode(url);
        if (this.ptr != null) {
            this.ptr.next = newNode;
            newNode.prev = this.ptr;
        }
        this.ptr = newNode;
    }
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * BrowserHistory obj = new BrowserHistory(homepage);
 * obj.visit(url);
 * String param_2 = obj.back(steps);
 * String param_3 = obj.forward(steps);
 */