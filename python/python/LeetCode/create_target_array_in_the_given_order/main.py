from typing import List
class Solution:
    def createTargetArray(self, nums: List[int], index: List[int]) -> List[int]:
        results = []
        l = len(nums)
        for i in range(0, l):
            n = nums[i]
            idx = index[i]
            if not results[idx - 1:idx]:
                results.append(n)
            elif results[idx-1:idx][0] != n:
                r = results[:idx]
                r.append(n)
                results = r.extend(results[idx:])
        return results
