// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
gofmtはGoのソースコードをフォーマット（整形）するツールです。

パスを与えない場合、標準入力の内容を処理します。
ディレクトリを指定した場合は、すべての .go ファイルを再帰的に処理します
（ただし、"."ドットではじまるファイルは無視します）。
通常、gofmtは標準出力へフォーマットしたコードを表示します。

Usage:
        gofmt [flags] [path ...]

The flags are:
        -d
                フォーマットした内容を標準出力へ出力しません。
                もしファイルのフォーマットがgofmtに通したものと異なる場合は、
                diffを標準出力します。
        -e
                すべてのエラーを表示します。
        -l
                フォーマットした内容を標準出力へ出力しません。
                もしファイルのフォーマットがgofmtに通したものと異なる場合は、
                そのファイル名を標準出力します。
        -r rule
                再フォーマット前のソースコードへ置き換えルールを指定します。
                （Exampleを参照）
        -s
                置き換えルールを適用した後、もしあれば、コードの簡素化を試みます。
        -w
                フォーマットした内容を標準出力へ出力しません。
                もしファイルのフォーマットがgofmtに通したものと異なる場合は、
                gofmtのもので上書き保存します。

Formatting control flags:
        -comments=true
                コメント内容を含めます。falseを指定すると、すべてのコメントが省略されます。
        -tabs=true
                タブでインデントします。falseを指定すると、スペースが使われます。
        -tabwidth=8
                スペースでのタブ幅を指定します。


置き換えルールは -r フラグ以下の形式の文字列を次のように指定する必要があります:

        pattern -> replacement

patternとreplacementの両方は、Goの文法に従っている必要があります。
patternで、小文字の1字はsub-expressionsのワイルドカードとして活用できますので、
replacementで同じ文字へと置換できます。

gofmtが標準入力から読む場合、Goプログラムの全体か、プログラムの断片のどちらかで受け付けます。
プログラムの断片では、構文的に有効な宣言リスト、ステートメントリスト、式である必要があります。
そのような断片をフォーマットする場合、gofmtは先頭のインデントと末尾のスペースを保持します。
ですので、Goプログラムの個々ののセクションでgofmtを通してフォーマットすることができます。

Examples

余計な括弧のペアが付いているファイルを確認する:

        gofmt -r '(a) -> a' -l *.go

余計な括弧のペアを削除する:

        gofmt -r '(a) -> a' -w *.go

明示的なスライスサイズの指定から暗黙的なものへパッケージツリー全体を変換する:

        gofmt -r 'α[β:len(α)] -> α[β:]' -w $GOROOT/src/pkg

The simplify command

gofmt -sで起動すると、以下のような変換が可能な場合があります。

        配列、スライスやマップの複合したもの:
                []T{T{}, T{}}
        を簡素化すると:
                []T{{}, {}}

        スライス:
                s[a:len(s)]
        を簡素化すると:
                s[a:]

        range:
                for x, _ = range v {...}
        を簡素化すると:
                for x = range v {...}

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/gofmt/doc.go?r=6152955fc7819180f4fac15eee678407df87da0a
*/
package main

// BUG(rsc): The implementation of -r is a bit slow.
