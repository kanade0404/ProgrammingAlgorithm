from python.HackerRank.extra_long_factories import extraLongFactorials

def test_case_1(capfd):
    extraLongFactorials(25)
    out, _ = capfd.readouterr()
    assert out == "15511210043330985984000000\n"
