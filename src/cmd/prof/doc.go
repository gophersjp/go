// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


/*

Profは基本的な機能を備えたリアルタイムプロファイラです。

実行コマンド、または既に実行されているコマンドのプロセスID(PID)を与えることで、
プログラムの状態を一定間隔で、振る舞いのレポートをサンプリングします。
オプションを付けずに実行した場合、サンプリング中に実行されているコードの
位置のヒストグラムが表示されます。

Profはリアルタイムプロファイラと呼ばれており、プログラムが休止中やI/O処理待ちのような
実行を停止しているときにプログラムの状態をサンプルする旧来のプロファイラとは違っています。
各スレッドは平等に統計処理にも貢献しています。


Usage:
	go tool prof -p pid [-t total_secs] [-d delta_msec] [6.out args ...]

The output modes (default -h) are:

	-P file.prof:
		プロファイルに関する設定をpprofのフォーマットでfile.profに記述します。
		現在のところ、このオプションはLinux amd64バイナリーと6l -eで作成された
		ELFデバッグinfoが必要です。
		詳細はhttp://code.google.com/p/google-perftoolsを参照してください。 
	-h: histograms
		それぞれのロケーションでサンプルを何回発生させるかを指定します。
	-f: dynamic functions
		それぞれのサンプル終了時に、実行中の関数の名前を表示させます。
	-l: dynamic file and line numbers
		それぞれのサンプル終了時に、実行した命令の行数やファイル名を表示させます。
	-r: dynamic registers
		それぞれのサンプル終了時に変数の値を表示させます。
	-s: dynamic function stack traces
		それぞれのサンプル終了時に、シンボリックスタックとレースを表示します。

-t フラグは実時間でサンプルする最大数を秒数で指定し、-d フラグはFlag -t sets the maximum real time to sample, in seconds, and -d
サンプリング間隔をミリセコンドで設定します。デフォルトではプログラムが完了するまで100ms周期でサンプリングされます。


profのインストールはアーキテクチャに依存しています。

*/
package main
