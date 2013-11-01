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
                パッケージディレクトリを表示する際にタイムスタンプを表示します(-http時)。
        -index
                全文検索を有効にします
        -index_files=""
                インデックスファイル名のパターンを指定します。
                設定した場合、インデックスは整列順でファイルを読み込みます。
        -index_throttle=0.75
                インデックススロットルの値を指定します。
                値を0にすることは、インデクサーへの割り当て時間がない
                （インデクサは完了しない）という意味になります。
                値を1.0にすることは、インデックスの生成をフルスロットルで実行する
                （他のgoroutineはインデックスのビルドが完了するまでの間は動かない）という意味になります。
        -links=true:
                宣言されている箇所へのリンクを生成します。
        -write_index=false
                インデックスファイルを書き出します。
                ファイル名は -index_files で一緒に指定する必要があります。
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
                代替のテンプレートファイルを含むディレクトリを指定します。
                指定する場合、指定するディレクトリに static/ のファイルの代わりの
                テンプレートファイルを用意します。
        -url=path
                print to standard output the data that would be served by
                an HTTP request for path
        -zip=""
                zipファイルでコンテンツを提供します。未指定で無効です。

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

例えば、 http://golang.org/pkg/math/big/?m=all を開くと、 big パッケージで宣言されたすべて（エクスポートされていないものも）のドキュメントを見ることができます。
コマンドラインでは、 "godoc -src math/big .*" として閲覧することができます。

通常、godocは、基盤となっているOSのファイルシステムからファイルを提供します。
代わりに -zipフラグで .zip ファイルから提供することもできます。
ファイルパスは.zipファイルで保持し、パスセパレータはスラッシュ ('/') を使う必要があり、
それらはアンルート？？？(unroot)されている必要があります。
$GOROOT (または -goroot)は、Goのルートディレクトリを含む.zipファイルへのディレクトリパスを与える必要があります。
コマンドで.zipファイル生成するには、例えば:

        zip go.zip $HOME/go

のように（訳注1）してファイルを追加し、次のようにgodocを実行できます:

        godoc -http=:6060 -zip=go.zip -goroot=$HOME/go

godocのドキュメンテーションは go/docパッケージでHTMLやテキストへ変換されます。
フォーマットの詳細は http://golang.org/pkg/go/doc/#ToHTML を御覧ください。

コメントの良い書き方については、 "Godoc: documenting Go code"
http://golang.org/doc/articles/godoc_documenting_go_code.html を御覧ください。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/cmd/godoc/doc.go?repo=tools&r=0e399fef76b7c34144d51e7b64c6da5b5591ea51

訳注1:
zipコマンドの完全な例は:

        zip -r go.zip $HOME/go -i \*.go -i \*.html -i \*.css -i \*.js -i \*.txt -i \*.c -i \*.h -i \*.s -i \*.png -i \*.jpg -i \*.sh -i favicon.ico

        
*/
package main

import "code.google.com/p/go.tools/cmd/godoc"
