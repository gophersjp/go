// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

5aはPlan 9用のアセンブラです。元のドキュメントについては以下を参照してください。

        http://plan9.bell-labs.com/magic/man2html/1/8a

ARMアーキテクチャをターゲットとしており、これらのツールにはarmとして参照されます。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/5a/doc.go?r=3633a89bb56d
*/
package main
