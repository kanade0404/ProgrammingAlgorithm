from python.HackerRank.mini_max_sum import miniMaxSum


def test_case_1(capfd):
    miniMaxSum([1, 3, 5, 7, 9])
    out, _ = capfd.readouterr()
    assert out == "16 24\n"


def test_case_2(capfd):
    miniMaxSum([1, 2, 3, 4, 5])
    out, _ = capfd.readouterr()
    assert out == "10 14\n"
