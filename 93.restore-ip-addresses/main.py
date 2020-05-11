class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        return self.restore(s, 4)[0]
    
    def restore(self, s, cnt):
        if len(s) == 0:
            return [], 0
        if cnt == 1:
            if len(s) > 3 or len(s) == 0 or (s[0] == '0' and len(s) > 1) or int(s) > 255:
                return [], 0
            else:
                return [s], 1
        if s[0] == '0':
            sub, subn = self.restore(s[1:], cnt-1)
            if subn == 0:
                return [], 0
            else:
                return [f'0.{n}' for n in sub], subn
        i = max(0, len(s)-cnt*3)
        result = []
        while i < len(s)-cnt+1 and int(s[:i+1])<=255:
            sub, subn = self.restore(s[i+1:], cnt-1)
            if subn > 0:
                for n in sub:
                    result.append(f'{s[:i+1]}.{n}')
            i += 1
        return result, len(result)
