from python.book_programming_solution.chap04.code02 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["5"], "15", capsys, monkeypatch, main)
    assert_std(["1"], "1", capsys, monkeypatch, main)

