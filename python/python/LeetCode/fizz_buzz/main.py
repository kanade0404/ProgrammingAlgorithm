from typing import List
class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        result = []
        for i in range(1, n + 1):
            v = str(i)
            if i % 3 == 0:
                v = "Fizz"
            if i % 5 == 0:
                v = "Buzz" if v == "" or v != "Fizz" else v + "Buzz"
            result.append(v)
        return result
