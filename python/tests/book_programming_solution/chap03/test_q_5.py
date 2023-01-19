from python.book_programming_solution.chap03.q_5 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["3", "2", "4", "6"], "1", capsys, monkeypatch, main)
    assert_std(["3", "8", "12", "40"], "2", capsys, monkeypatch, main)
