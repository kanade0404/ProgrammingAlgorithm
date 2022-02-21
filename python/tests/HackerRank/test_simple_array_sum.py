from python.HackerRank.simple_array_sum import simpleArraySum


def test_case_1():
    assert simpleArraySum([1, 2, 3, 4, 10, 11]) == 31


def test_case_2():
    assert simpleArraySum([1]) == 1


def test_case_3():
    assert simpleArraySum([1 for _ in range(1, 1001)]) == 1000
