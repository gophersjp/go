// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

gcは、Plan 9(修正済のもの)ツールチェーン(tool chain)の一部として機能するGoコンパイラファミリのための一般的な名称です。
Cコンパイラドキュメントはこちらを参照してください:

        http://plan9.bell-labs.com/sys/doc/comp.pdf     (ツールの概要)
        http://plan9.bell-labs.com/sys/doc/compiler.pdf (Cコンパイラアーキテクチャ)

上記ドキュメントにはツールチェーンの全体の設計が書かれています。
オプティマイザのようないくつかの適用部分を置いておけば、Goコンパイラはまったく新しいプログラムです。

このコンパイラは、".go"の拡張子のGoファイル群を読み込みます。
これらはすべて、ひとつのパッケージの一部でなければなりません。
出力は、リンカ(6lなど)へ入力する前段階、コンパイル済みパッケージの"binary assembly"を表す単一の中間ファイルです。

生成されたファイルには、パッケージ内でエクスポートしたシンボルについての型情報や、他のパッケージからパッケージ内でインポートしたシンボルで使う型に関する情報が含まれています。
それゆえ、パッケージPの依存のファイルを読み込むために、PのクライアントCをコンパイルしているときには必要ありません。Pのコンパイル済みの出力だけです。
It is therefore not necessary when compiling client C of package P to read the files of P's dependencies, only the compiled output of P.
TODO: 意味確認。

Command Line

Usage:
        go tool 6g [flags] file...
指定するファイルはGoのソースコードファイルであり、かつ、すべて同じパッケージの一部である必要があります。
環境に合わせて適切に6gを8gや5gに変えてください。

Flags:
        -o file
                出力ファイルを指定します。デフォルトでは6gならfile.6です
        -e
                全てのエラーを表示します。デフォルトでは10個のエラーを出力するとコンパイラを終了します
        -p path
                指定したパスがソースコードの最後のインポートパスだとみなし、
                それに依存するパッケージをインポートしようと試みます。 TODO: 意味不明
                assume that path is the eventual import path for this code,
                and diagnose any attempt to import a package that depends on it.
        -D path
                パスへの相対として、相対インポートを扱います  TODO: 要確認
                treat a relative import as relative to path
        -L
                エラーでファイル名を表示する際に、ファイルパス全体を表示します
        -I dir1 -I dir2
                dir1とdir2をインポートしたパッケージをチェックするためにパスのリストへ追加します TODO: 要確認
                add dir1 and dir2 to the list of paths to check for imported packages
        -N
                最適化を無効にします
        -S
                アセンブリ言語のテキストを標準出力します（コードのみ）
        -S -S
                アセンブリ言語のテキストを標準出力します（コードとデータ）
        -u
                safeとされていないパッケージのインポートを拒否します TODO: 要確認
        -V
                コンパイラのバージョン情報を表示します
        -race
                レースコンディションの検出器を有効にしてコンパイルします

他にも、デバッグ用のフラグがいくつかあります。
このコマンドを引数なしで走らせ、全てのUsageを見てください。

Compiler Directives

コンパイラは、行の最初に // のコメントの形式で２つのコンパイラディレクティブを受け入れます。
ディレクティブは、ディレクティブではないコメントと区別するために、スラッシュとディレクティブの名前との間にスペースを必要としません。
コメントですので、ディレクティブを認識しないツールは、
他のコメントのようにディレクティブをスキップすることができます。

    //line path/to/file:linenumber

//lineディレクティブは、ソースコードの行を
指定したファイルパスと行番号から来たものとして記録します。
連続している行は、指定した行番号を次のディレクティブまでインクリメントして記録します。
コンパイラやデバッガが生成器へ元々の入力で行を表示するので、このディレクティブは、一般的にマシンで生成されたコード中で登場します。

    //go:noescape

//go:noescapeディレクティブは、
bodyを持たない関数（Goで書かれていない実装があるということ）
でなければならないファイルの次の宣言が
ヒープ、または、関数からの戻り値にエスケープ(escape)する引数として渡るポインタを
なにひとつ許可しないということを記録します。
この情報は、関数を呼び出すGoコードのコンパイラのエスケープ解析中に利用できます。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/gc/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
*/
package main
