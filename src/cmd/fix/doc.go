// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Fix finds Go programs that use old APIs and rewrites them to use
newer ones.  After you update to a new Go release, fix helps make
the necessary changes to your programs.
fixはGoのプログラムに古いAPIが使われていないかを探し、新しいAPIに書き換えます。
Goのリリースによって、あなたのプログラムに修正の必要が生じたときにfixは役立ちます。

Usage:
	go tool fix [-r name,...] [path ...]

Without an explicit path, fix reads standard input and writes the
result to standard output.
明確なパスの入力がなくても、fixは標準入力を読み込み、標準出力に結果を書き出します。

If the named path is a file, fix rewrites the named files in place.
If the named path is a directory, fix rewrites all .go files in that
directory tree.  When fix rewrites a file, it prints a line to standard
error giving the name of the file and the rewrite applied.
パスにファイルを指定した場合、fixはそのファイルを書き換えます。
パスにディレクトリを指定した場合、fixはそのディレクトリ・ツリー以下の全ての.goファイルを書き換えます。fixがファイルを書き換える際、標準エラー出力にそのファイル名と修正箇所が出力されます。

If the -diff flag is set, no files are rewritten. Instead fix prints
the differences a rewrite would introduce.
-diffフラグを指定すると、ファイルは書き換えられません。代わりに、fixは上書きされる箇所の差分を表示します。

The -r flag restricts the set of rewrites considered to those in the
named list.  By default fix considers all known rewrites.  Fix's
rewrites are idempotent, so that it is safe to apply fix to updated
or partially updated code even without using the -r flag.
-rフラグは指定されたリストの書き換えを禁止します。デフォルト設定では、fixは書き換えが必要な全てについて考慮します。
// FIXME
fixの書き換えは冪等(べきとう)です。そのため-rフラグさえ使わずに、更新されたコード、もしくは部分的に更新されたコードにfixを適用することは安全です。

Fix prints the full list of fixes it can apply in its help output;
to see them, run go tool fix -?.
// FIXME
fixはヘルプ出力に適用できる全ての修正リストを表示します。
それらを見るには、`go tool fix -?`を実行してください。

Fix does not make backup copies of the files that it edits.
Instead, use a version control system's ``diff'' functionality to inspect
the changes that fix makes before committing them.
fixは編集されたファイルのバックアップを作成しません。
コミットする前にfixによる変更を確認するには、バージョン管理システムの``diff''機能を使ってください。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/fix/doc.go?r=3633a89bb56d
*/
package main
