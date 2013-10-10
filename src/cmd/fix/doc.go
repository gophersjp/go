// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
fixはGoのプログラムに古いAPIが使われていないかを探し、新しいAPIに書き換えます。
Goのリリースによって、あなたのプログラムに修正の必要が生じたときにfixは役立ちます。

Usage:
	go tool fix [-r name,...] [path ...]

明確なパスの入力がなくても、fixは標準入力を読み込み、標準出力に結果を書き出します。

パスにファイルを指定した場合、fixはそのファイルを書き換えます。
パスにディレクトリを指定した場合、fixはそのディレクトリ・ツリー以下の全ての.goファイルを書き換えます。fixがファイルを書き換える際、標準エラー出力にそのファイル名と修正箇所が出力されます。

-diffフラグを指定すると、ファイルは書き換えられません。代わりに、fixは上書きされる箇所の差分を表示します。

-rフラグは指定されたリストの書き換えを禁止します。デフォルト設定では、fixは書き換えが必要な全てについて考慮します。
fixの書き換えは冪等(べきとう)です。そのため-rフラグを使用しなかったとしても、全部もしくは一部の更新されたコードへのfix適用の安全性は保証されます。

fixはヘルプ出力に適用できる修正の全リストを表示します。
それらを見るには、`go tool fix -?`を実行してください。

fixは編集するファイルのバックアップを作成しません。
コミットする前にfixによる変更を確認するには、バージョン管理システムの``diff''機能を使ってください。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/fix/doc.go?r=3633a89bb56d
*/
package main
