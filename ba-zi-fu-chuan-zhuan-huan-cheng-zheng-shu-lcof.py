class Solution:
    def is_digit(self, s):
        return len(s) == 1 and s in set(range(10))

    def first_letter(self, s):
        m = re.search(r'[a-z\s+-]', s, re.I)
        if m is not None:
            return m.start()
        return -1

    def strToInt(self, s: str) -> int:
        s = s.strip()
        sign = 1
        if len(s) == 0:
            return 0
        if s[0] == "-":
            sign = -1
            s = s[1:]
        elif s[0] == "+":
            s = s[1:]

        first_letter = self.first_letter(s)
        if first_letter >= 0:
            s = s[:first_letter]
        r = 0
        try:
            r = sign * int(float(s))
        except:
            pass
        return min(pow(2, 31)-1, max(-pow(2, 31), r))
