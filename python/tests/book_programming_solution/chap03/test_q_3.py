from python.book_programming_solution.chap03.q_3 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["2", "1", "2"], "2", capsys, monkeypatch, main)
    assert_std(["3", "3", "2", "0"], "2", capsys, monkeypatch, main)
