from python.book_programming_solution.chap02.code04 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["2", "1 3", "2 3"], "1", capsys, monkeypatch, main)
    assert_std(["1", "1 100000"], "99999", capsys, monkeypatch, main)
