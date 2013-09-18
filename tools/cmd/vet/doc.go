// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
vetコマンドはGoのソースコードをチェックし、Printfの呼び出しで引数が書式文字列(format string)で列んでないような欠陥をレポートします。
vetはヒューリスティック(heuristic)な問題を扱い、すべてのレポートが本当に問題であることを保証しません。しかし、コンパイラで見つけられないエラーを検出することができます。

コマンドの終了ステータス(exit code)は、コマンドの起動を間違うと"2"、問題がレポートされると"1"、"0"はそれ以外です。
コマンドは起こりうるすべての問題をチェックできず、信頼性の低いヒューリスティックに依存することを覚えておいてください。
ですから、ガイダンスとしてだけ使用すべきで、確固としたプログラムの正当性の指標として利用すべきではありません。

通常、すべてのチェックが実行されますが、フラグを付けることで範囲を指定することができます。

利用可能なオプション：

1. Printf群のチェック (-printf)

Printf群で関数を呼ぶものをチェックします。
Printf群:
        Print Printf Println
        Fprint Fprintf Fprintln
        Sprint Sprintf Sprintln
        Error Errorf
        Fatal Fatalf
        Panic Panicf Panicln
もし関数が'f'で終わるものであれば、その関数はfmt.Printfの作法での書式化文字をとることを想定します。もしそれ以外をとっていれば、vetは書式化文字列のような引数について一言言ってきます。

また、Printfの最初の引数のようなWriterを使う場合でもチェックします。

2. メソッドのチェック (-methods)

以下を含む、よく知られた名前のメソッドのための標準ではない識別子をチェックします:
        Format GobEncode GobDecode MarshalJSON MarshalXML
        Peek ReadByte ReadFrom ReadRune Scan Seek
        UnmarshalJSON UnreadByte UnreadRune WriteByte
        WriteTo

3. 構造体のタグのチェック (-structtags)

reflect.StructTag.Getで理解できる書式に従っていない構造体のタグをチェックします。

4. キーの無い複合的なリテラルのチェック (-composites)

フィードキーのシンタックスを使っていない複合的な構造体のリテラルをチェックします。


Usage:

        go tool vet [flag] [file.go ...]
        go tool vet [flag] [directory ...] # ディレクトリ以下の全ての .go ファイルを再帰的にスキャンします

The other flags are:
        -v
                Verbose mode
        -printfuncs
                A comma-separated list of print-like functions to supplement
                the standard list.  Each entry is in the form Name:N where N
                is the zero-based argument position of the first argument
                involved in the print: either the format or the first print
                argument for non-formatted prints.  For example,
                if you have Warn and Warnf functions that take an
                io.Writer as their first argument, like Fprintf,
                        -printfuncs=Warn:1,Warnf:1


本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/cmd/vet/doc.go?repo=tools&r=81e58ded571716d210ea5358d40cd9e10403b98b

*/
package main

import "code.google.com/p/go.tools/cmd/vet"
