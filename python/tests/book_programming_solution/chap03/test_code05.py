from python.book_programming_solution.chap03.code05 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["2 3", "1", "2"], "Yes", capsys, monkeypatch, main)
    assert_std(["2 1", "1", "2"], "Yes", capsys, monkeypatch, main)
    assert_std(["3 10", "1", "2", "8"], "Yes", capsys, monkeypatch, main)
