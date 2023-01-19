from python.book_programming_solution.chap03.q_4 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["1 2 3 4"], "3", capsys, monkeypatch, main)
    assert_std(["1 0"], "1", capsys, monkeypatch, main)
    assert_std(["0"], "0", capsys, monkeypatch, main)
