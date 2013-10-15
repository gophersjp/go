// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// goplayはGo言語のコードで実験をするためのWebインタフェースです。
// Go Playgroundと同様のものです：http://golang.org/doc/play/
//
// goplayの利用方法：
//   $ cd $GOROOT/misc/goplay
//   $ go run goplay.go
// 続いて、Webブラウザで http://localhost:3999/ をロードします。
//
// Hello Worldプログラムが表示されるはずです。shift-enterキーにより
// コンパイル・実行ができます。チェックボックスをチェックして
// 有効にできる"compile-on-keypress"機能もあります。
//
// 警告！ WARNING! CUIDADO! ACHTUNG! ATTENZIONE!
// セキュリティに関する注意：goplayのWebインタフェースにアクセスすると
// 誰でもあなたのコンピュータ上で任意のコードを実行することができます。
// goplayはサンドボックスではなく、その他のセキュリティの仕組みもありません。
// goplayを信頼の置けない環境にデプロイしないでください。
// デフォルトでは、goplayはlocalhost上のみでリッスンしますが、
// -httpパラメータによりオーバーライドできます。ただし自己責任で行ってください。
//
//	本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/misc/goplay/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
package main

