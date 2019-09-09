package repl

import (
	"bufio"
	"first/lexer"
	"first/token"
	"fmt"
	"io"
)

const PROMPT = ">> "

// 改行が来るまで入力ソースから読み込み、読み込んだ行を取り出して、字句解析器のインスタンスに渡す
// 字句解析器が返す全てのトークンを表示する。EOFになるまで繰り返す
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

