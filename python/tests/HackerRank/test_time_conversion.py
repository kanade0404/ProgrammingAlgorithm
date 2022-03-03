from python.HackerRank.time_conversion import timeConversion


def test_case_1():
    assert timeConversion("07:05:45PM") == "19:05:45"


def test_case_2():
    assert timeConversion("12:01:00PM") == "12:01:00"


def test_case_3():
    assert timeConversion("12:01:00AM") == "00:01:00"


def test_case_4():
    assert timeConversion("00:01:00AM") == "00:01:00"
