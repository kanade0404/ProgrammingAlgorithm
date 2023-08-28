class Solution:
    def myAtoi(self, s: str) -> int:
        tmp = ""
        max_int = 2**31 - 1
        min_int = -2**31
        is_already_plus_or_minus = False
        is_already_digit = False
        is_minus = False
        for i in s:
            if i == ' ':
                if is_already_digit or is_already_plus_or_minus:
                    break
                continue
            elif i == '-' or i == '+':
                if is_already_digit:
                    break
                if is_already_plus_or_minus:
                    break
                is_already_plus_or_minus = True
                if i == '-':
                    is_minus = True
            elif i.isdigit():
                is_already_digit = True
                tmp += i
            else:
                break
        result = int(tmp) if tmp != "" else 0
        if is_minus:
            result *= -1
        if result > max_int:
            return max_int
        if result < min_int:
            return min_int
        return result
