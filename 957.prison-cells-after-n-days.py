class Solution:
    def prisonAfterNDays(self, cells: List[int], n: int) -> List[int]:

        cache = []
        cache_hash = {}

        cells_list = cells
        j = 0
        while j < n:

            cells_list_hash = self.hash_list(cells_list)
            if cells_list_hash in cache_hash:
                break
            cache_hash[cells_list_hash] = j
            cache.append(cells_list)

            cells_list = self.calc_list(cells_list)
            j += 1

        if j == n:
            result = cells_list
        else:
            prev = cache_hash[cells_list_hash]
            cache = cache[prev:]
            # print(len(cache), cache)
            idx = (n - prev) % len(cache)
            result = cache[idx]
        return result

    def hash_list(self, cells):
        return ''.join([str(i) for i in cells])
        
    def calc_list(self, cells_list):
        new_list = [0] * len(cells_list)
        for i in range(len(cells_list)):
            new_list[i] = self.calc(cells_list, i)
        return new_list

    def calc(self, cells: List[int], i: int) -> int:
        if i == 0 or i == len(cells) - 1:
            return 0
        return int((cells[i-1] and cells[i+1]) or (not cells[i-1] and not cells[i+1]))
