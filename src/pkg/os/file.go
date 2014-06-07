// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// osパッケージはOS機能へのプラットフォーム非依存なインターフェースを提供します。
// 設計はUnixライクですが、エラーの扱いはGoライクになっており、
// 関数呼び出しの失敗時にはエラー番号ではなく、error型の値を返します。
// errorからはより多くの情報が得られることがあります。
// 例えば、ファイル名を受け取る関数呼び出し(OpenやStatなど)が失敗したとき、
// errorをprintすると失敗したファイル名が得られ、またerrorを*PathError型の値
// として扱うことで、より詳細な情報が得られます。
//
// osパッケージは全てのOSに共通のインターフェースを提供することを目指しています。
// 共通ではない機能はシステム依存なパッケージであるsyscallにあります。
//
// 以下はファイルを開いてそれを読む簡単な例です。
//
//	file, err := os.Open("file.go") // For read access.
//	if err != nil {
//		log.Fatal(err)
//	}
//
// ファイルを開くのに失敗すると、errorは以下のようなエラーの内容を示す文字列
// になります。
//
//	open file.go: no such file or directory
//
// ファイルのデータはbyteのスライスへ読み込まれます。
// ReadとWriteは引数に渡されたスライスの長さをバイト数として返します。
//
//	data := make([]byte, 100)
//	count, err := file.Read(data)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("read %d bytes: %q\n", count, data[:count])
//
// 本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/pkg/os/file.go?r=5142686ded576fd930ba36b3aeac9b7c3a819d6b
package os
