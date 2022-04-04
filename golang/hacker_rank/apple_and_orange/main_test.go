package apple_and_orange

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func Test_countApplesAndOranges(t *testing.T) {
	type args struct {
		s       int32
		t       int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Sample Test1",
			args{7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4}},
			`1
2`,
		},
		{
			"Sample Test1",
			args{7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6}},
			`1
1`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractStdout(t, countApplesAndOranges, tt.args.s, tt.args.t, tt.args.a, tt.args.b, tt.args.apples, tt.args.oranges); got != tt.want {
				t.Errorf("countApplesAndOranges() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Stdoutに書き込まれた文字列を抽出する関数
// (Stderrも同じ要領で出力先を変更できます)
func extractStdout(testing *testing.T, fnc func(int32, int32, int32, int32, []int32, []int32), s int32, t int32, a int32, b int32, apples []int32, oranges []int32) string {
	testing.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()
	// パイプの作成(r: Reader, w: Writer)
	r, w, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = w
	// テスト対象の関数を実行する
	fnc(s, t, a, b, apples, oranges)
	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	w.Close()
	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		testing.Fatalf("failed to read buf: %v", err)
	}
	// 文字列を取得する
	return strings.TrimRight(buf.String(), "\n")
}
