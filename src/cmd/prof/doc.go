// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


/*

Profは基本的な機能を備えたリアルタイムプロファイラです。

実行コマンド、または既に実行されているコマンドのプロセスID(PID)を与えることで、
プログラムの状態を一定間隔で、振る舞いのレポートをサンプリングします。
オプションを付けずに実行した場合、実行中にサンプリングされたコードの
位置のヒストグラムが表示されます。

Profはリアルタイムなプロファイラであるため、従来のプロファイラとは異なり、たとえスリープ状態やI/Oの待機中といった実行中ではないときでもプログラムの状態をサンプリングします。
各スレッドは平等に統計処理を行ないます。

Usage:
	go tool prof -p pid [-t total_secs] [-d delta_msec] [6.out args ...]

The output modes (default -h) are:

	-P file.prof:
		プロファイルに関する設定をpprofのフォーマットでfile.profに書き込みます。
		現在のところ、このオプションはLinux amd64バイナリ上のみで動作し、
		またELFデバッグ情報を生成するにはバイナリが6l -eを使用して書き出されている必要があります。
		詳細はhttp://code.google.com/p/google-perftoolsを参照してください。 
	-h: histograms
		それぞれのロケーションでサンプルを何回発生させるかを指定します。
	-f: dynamic functions
		それぞれのサンプル時に、実行中の関数の名前を表示させます。
	-l: dynamic file and line numbers
		それぞれのサンプル時に、実行した命令の行数やファイル名を表示させます。
	-r: dynamic registers
		それぞれのサンプル時に変数の値を表示させます。
	-s: dynamic function stack traces
		それぞれのサンプル時に、シンボリックスタックトレースを表示します。

-t フラグはサンプリングを行う最大実時間を秒数で設定し、
-d フラグはサンプリング間隔をミリ秒で設定します。
デフォルトではプログラムが完了するまで100ms周期でサンプリングされます。

profはgo tool profコマンドでインストールされます。また、アーキテクチャ依存です。

本ドキュメントは以下のドキュメントを翻訳しています:  https://code.google.com/p/go/source/browse/src/cmd/prof/doc.go?r=3633a89bb56d9276a9fe55435b849f931bfa6393
*/
package main
