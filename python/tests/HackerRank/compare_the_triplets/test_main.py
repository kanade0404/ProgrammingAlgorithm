from python.HackerRank.compare_the_triplets.main import compareTriplets


def test_case_1():
    assert compareTriplets([5, 6, 7], [3, 6, 10]) == [1, 1]

def test_case_2():
    assert compareTriplets([17, 28, 30], [99, 16, 8]) == [2, 1]