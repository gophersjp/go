// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// raceパッケージはデータ競合検知ロジックを実装しています。
// 公開インタフェースは提供されていません。
// データ競合検知についての詳細は http://golang.org/doc/articles/race_detector.html を
// ご覧ください。
//
// 本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/pkg/runtime/race/doc.go?r=30b1c2ff7d934e0dbbc9253293ba29caf3312507
package race
