// [[0,0],[1,1],[2,2],[3,3],[4,4],[5,5],[6,6]]

impl Solution {
    pub fn min_rectangles_to_cover_points(mut points: Vec<Vec<i32>>, w: i32) -> i32 {
        let mut r = 0;
        let mut start = -1 - w;

        points.sort_by(|a, b| a[0].cmp(&b[0]));
        for p in (&points).iter() {
            if p[0] >= start && p[0] <= start + w {
                continue;
            }
            start = p[0];
            r += 1;
        }
        return r;
    }
}
