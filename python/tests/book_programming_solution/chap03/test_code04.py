from python.book_programming_solution.chap03.code04 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["3 10", "8", "5", "4", "4", "1", "9"], "12", capsys, monkeypatch, main)
