from python.book_programming_solution.chap04.code04 import main
from tests.asserts.assert_std import assert_std


def test_main(capsys, monkeypatch):
    assert_std(["51 15"], "3", capsys, monkeypatch, main)
