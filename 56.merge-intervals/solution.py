# Definition for an interval.
# class Interval:
#     def __init__(self, s=0, e=0):
#         self.start = s
#         self.end = e

class Solution:
    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        if len(intervals) <= 1:
            return intervals
        result = []
        sorted_intervals = sorted(intervals, key=lambda x: x.start)
        lsi = len(sorted_intervals)
        
        i = 0
        while i < len(sorted_intervals)-1:
            new_interval = Interval(s=sorted_intervals[i].start, e=sorted_intervals[i].end)
            while i < lsi-1 and sorted_intervals[i+1].start <= new_interval.end:
                if sorted_intervals[i+1].end > new_interval.end:
                    new_interval.end = sorted_intervals[i+1].end
                i += 1
            result.append(new_interval)
            i += 1
        if i > 0 and i<lsi and not sorted_intervals[i].end == sorted_intervals[i-1].end:
            result.append(sorted_intervals[i])
        return result
