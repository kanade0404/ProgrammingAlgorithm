from python.HackerRank.plus_minus import plusMinus


def test_case_1(capfd):
    plusMinus([1, 1, 0, -1, -1])
    out, _ = capfd.readouterr()
    assert out == "0.400000\n0.400000\n0.200000\n"
    assert _ is ""


def test_case_2(capfd):
    plusMinus([-4, 3, -9, 0, 4, 1])
    out, _ = capfd.readouterr()
    assert out == "0.500000\n0.333333\n0.166667\n"
    assert _ is ""
