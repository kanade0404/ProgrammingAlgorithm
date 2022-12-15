from typing import Callable, List
from _pytest.capture import CaptureFixture
import io


def assert_std(i: List[str], expected: str, capsys: CaptureFixture[str], monkeypatch, fn: Callable):
    monkeypatch.setattr("sys.stdin", io.StringIO("\n".join(i)))
    fn()
    out, _ = capsys.readouterr()
    assert out == "%s\n" % expected
