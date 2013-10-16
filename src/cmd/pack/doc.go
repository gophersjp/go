// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

packはPlan 9のar toolの変形です。オリジナルのドキュメントは以下にあります。

	http://plan9.bell-labs.com/magic/man2html/1/ar

packはアーカイブのファイルからすべてのGoの型情報を収集する__.PKGDEFという
Go言語特有の特別な領域を追加します。その領域はコンパイル時の
パッケージインポートの際、コンパイラが利用します。

Usage:
	go tool pack [uvnbailogS][mrxtdpq][P prefix] archive files ...

新たな'g'オプションは、ファイルがアーカイブに追加される際、packに__.PKGDEF領域
を保持させます。

新たな'S'オプションは、packに強制的にアーカイブが安全であると指定させます。

新たな'P'オプションは、既にアーカイブに格納・追加されているオブジェクトファイルの
行番号情報にあるファイル名から、プレフィックスをpackに除去させます。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/pack/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
*/
package main
