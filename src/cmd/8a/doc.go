// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

8aはPlan 9用のアセンブラです。元のドキュメントについては以下を参照してください。

	http://plan9.bell-labs.com/magic/man2html/1/8a

x86アーキテクチャをターゲットとしており、これらのツールには歴史的な理由から386として参照されます。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/8a/doc.go?r=3633a89bb56d
*/
package main
