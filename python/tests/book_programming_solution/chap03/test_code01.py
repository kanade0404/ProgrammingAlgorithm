from python.book_programming_solution.chap03.code01 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["7 5", "4", "3", "12", "7", "11"], "True", capsys, monkeypatch, main)
