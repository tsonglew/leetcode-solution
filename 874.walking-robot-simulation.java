class Solution {
    public int robotSim(int[] commands, int[][] obstacles) {
        int x = 0;
        int y = 0;
        int res = 0;
        HashMap<Integer, HashSet<Integer>> obs = new HashMap<>();
        for (int i = 0; i < obstacles.length; i++) {
            for (int j = 0; j < 2; j++) {
                if (!obs.containsKey(obstacles[i][0])) {
                    obs.put(obstacles[i][0], new HashSet<Integer>());
                }
                obs.get(obstacles[i][0]).add(obstacles[i][1]);
            }
        }
        int nextX = 0;
        int nextY = 1;
        for (int i = 0; i < commands.length; i++) {
            int command = commands[i];
            if (command != -1 && command != -2) {
                int steps = command;
                while (steps > 0 && !this.isObstacle(obs, x + nextX, y + nextY)) {
                    x += nextX;
                    y += nextY;
                    steps--;
                }
            } else {
                int nX = nextX;
                nextX = this.nextX(command, nextX, nextY);
                nextY = this.nextY(command, nX, nextY);
                // System.out.printf("nextX: %d, nextY: %d\n", nextX, nextY);
            }
            // System.out.printf("x: %d, y: %d\n", x, y);
            res = Math.max(res, (int) (Math.pow(x, 2) + Math.pow(y, 2)));
        }
        return res;
    }

    public boolean isObstacle(HashMap<Integer, HashSet<Integer>> obs, int x, int y) {
        if (obs.containsKey(x) && obs.get(x).contains(y)) {
            return true;
        }
        return false;
    }

    public int nextX(int command, int nextX, int nextY) {
        if (command == -2) {
            return -nextY;
        } else if (command == -1) {
            return nextY;
        }
        return nextX;
    }

    public int nextY(int command, int nextX, int nextY) {
        if (command == -2) {
            return nextX;
        } else if (command == -1) {
            return -nextX;
        }
        return nextY;
    }
}