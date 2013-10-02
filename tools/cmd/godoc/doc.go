// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

godocはGoのソースコードからドキュメントを生成します。

2つのモデルがあります。

-httpフラグをつけずに起動すると、コマンドラインモードで実行し、標準出力へテキストの
ドキュメントを表示、終了します。
もしライブラリパッケージとコマンドで同じ名前になっている場合は、ライブラリパッケージ
よりも cmd/ ではじまるコマンドドキュメントを優先します。
-srcフラグが指定されると、godocはGoのソースコード形式でパッケージのエクスポートされた
インターフェースや、個別のエクスポートされたものの実装を表示します:

        godoc fmt                # fmtパッケージをドキュメンテーションします
        godoc fmt Printf         # fmt.Printfをドキュメンテーションします
        godoc cmd/go             # goコマンドをforce強制的にドキュメンテーションします
        godoc -src fmt           # Goのソースコード形式で、fmtパッケージを表示します
        godoc -src fmt Printf    # fmt.Printfの実装を表示します

-qフラグに検索クエリを指定することで、
ウェブサーバに対して検索結果を問い合わせることができます（コマンドラインモード時）。
サーバアドレスを-serverフラグで指定しなかった場合、godocは最初に localhost:6060
に接続を試み、次に http://golang.org へ接続します。

        godoc -q Reader
        godoc -q math.Sin
        godoc -server=:6060 -q sin

-httpフラグをつけて起動すると、ウェブサーバ(localhost:6060)が起動し、ウェブページでドキュメントを読むことができます。
（-httpフラグにはポート番号を以下のように指定します）

        godoc -http=:6060

Usage:
        godoc [flag] package [name ...]

The flags are:
        -v
                verbose mode
        -q
                引数に検索クエリを指定します。
                有効なクエリの例は、単なる名称（ToLowerといったもの）
                や、限定した名称（math.Sinといったもの）です。
        -src
                ソースコード（エクスポートされたもの）をコマンドラインモードで表示します。
        -tabwidth=4
                タブ幅をスペースの数で指定します。
        -timestamps=true
                show timestamps with directory listings
        -index
                enable identifier and full text search index
                (no search box is shown if -index is not set)
        -index_files=""
                glob pattern specifying index files; if not empty,
                the index is read from these files in sorted order
        -index_throttle=0.75
                index throttle value; a value of 0 means no time is allocated
                to the indexer (the indexer will never finish), a value of 1.0
                means that index creation is running at full throttle (other
                goroutines may get no time while the index is built)
        -links=true:
                link identifiers to their declarations
        -write_index=false
                write index to a file; the file name must be specified with
                -index_files
        -maxresults=10000
                maximum number of full text search results shown
                (no full text index is built if maxresults <= 0)
        -notes="BUG"
                regular expression matching note markers to show
                (e.g., "BUG|TODO", ".*")
        -html
                コマンドラインモードでHTMLを表示します。
        -goroot=$GOROOT
                Goのルートディレクトリを指定します。
        -http=addr
                HTTPサービスのアドレスを指定します (例: '127.0.0.1:6060' または ':6060')
        -server=addr
                コマンドラインサーチでのサーバアクセス先を指定します。
        -templates=""
                directory containing alternate template files; if set,
                the directory may provide alternative template files
                for the files in $GOROOT/lib/godoc
        -url=path
                print to standard output the data that would be served by
                an HTTP request for path
        -zip=""
                zip file providing the file system to serve; disabled if empty

godocは、環境変数 $GOROOT と $GOPATH （設定してあれば）を見てパッケージを検索します。
この動作は、 -gorootフラグで $GOROOT を変えることで変更することができます。

godocをウェブサーバとして実行する際に -index がセットされると、サーチインデックスを保有します。
インデックスは起動時に生成されます。

インデックスは、名称とフルテキストのサーチの情報（正規表現で検索できます）を含んでいます。
フルテキストサーチの結果を表示する限度は -maxresultsフラグで指定できます。
もし 0 を指定するとフルテキストサーチの結果は表示されません。
インデックスは名称のみとなり、フルテキストサーチのインデックスは生成されません。

godocが提供するウェブページのプレゼンテーションモードは、 URLパラメータに"m"でコントロールできます。
これは、コンマ区切りのリストを受け付けます。

        all     エクスポートされたものだけではなく、宣言されたすべてのドキュメントを表示します
        methods エクスポートしていない匿名フィールドだけでなく、すべての組み込みメソッドを表示します TODO:動作確認
                show all embedded methods, not just those of unexported anonymous fields
        src     ドキュメントではなく、その元のソースコードを表示します
        text    HTMLではなく、テキストフォーマット（コマンドライン用）で提供します
        flat    パッケージ表示を階層ではなく、フルパスを用いたフラット（インデントのない）なリストで提供します

例えば、 http://golang.org/pkg/math/big/?m=all を見ると、bigパッケージで宣言されたすべて（エクスポートされていないものも）のドキュメントを見ることができます。 
TODO: "godoc -src math/big .*" で意図したとおりにならないのはなんで？？
For instance, http://golang.org/pkg/math/big/?m=all,text shows the documentation
for all (not just the exported) declarations of package big, in textual form (as
it would appear when using godoc from the command line: "godoc -src math/big .*").

通常、godocは、基盤となっているOSのファイルシステムからファイルを提供します。
代わりに -zipフラグで .zip ファイルから提供することもできます。
ファイルパスは .zipファイルで保持され、パスセパレータとしてスラッシュ ('/') を使う必要があり、
それらはルートを ... ? TODO:動作確認 zipコマンドの使い方あってる？
By default, godoc serves files from the file system of the underlying OS.
Instead, a .zip file may be provided via the -zip flag, which contains
the file system to serve. The file paths stored in the .zip file must use
slash ('/') as path separator; and they must be unrooted. $GOROOT (or -goroot)
must be set to the .zip file directory path containing the Go root directory.
For instance, for a .zip file created by the command:

        zip go.zip $HOME/go

one may run godoc as follows:

        godoc -http=:6060 -zip=go.zip -goroot=$HOME/go

See "Godoc: documenting Go code" for how to write good comments for godoc:
http://golang.org/doc/articles/godoc_documenting_go_code.html

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools&r=3504d66cc4b60e42aa0e0c2ced58afb0f1a8cc82

*/
package main

import "code.google.com/p/go.tools/cmd/godoc"
