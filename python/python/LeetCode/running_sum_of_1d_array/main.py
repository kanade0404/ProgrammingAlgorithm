from typing import List

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        result = []
        tmp = 0
        for i in nums:
            tmp += i
            result.append(tmp)
        return result
