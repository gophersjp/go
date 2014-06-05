// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import "time"

// FindProcessはpidを用いて実行中のプロセスを探します。
// 返されるProcessは実際のOSのプロセスに関する情報を得るために使うことができます。
func FindProcess(pid int) (p *Process, err error) {
	return findProcess(pid)
}

// StartProcessはname、argv、attrにて示されるプログラム、引数、属性を用いて
// 新しいプロセスを開始します。
//
// StartProcessは低レベルのインターフェースです。
// os/execパッケージはより高レベルのインターフェースを提供します。
//
// もしエラーがある場合は、*PathError型のエラーになります。
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error) {
	return startProcess(name, argv, attr)
}

// ReleaseはProcess型の値pに関連付けられた全てのリソースを開放し、
// この先使えなくします。
// ReleaseはWaitが呼ばれていない場合のみ呼ぶ必要があります。
func (p *Process) Release() error {
	return p.release()
}

// KillはProcessをすぐに終了させます。
func (p *Process) Kill() error {
	return p.kill()
}

// WaitはProcessが終了するのを待ち、そのステータスを表すProcessStateと、
// あればerrorを返します。
// WaitはProcessに関連付けられた全てのリソースを開放します。
// ほとんどのOSでは、Processは現在のプロセスの子プロセスでなければならず、
// そうでなければerrorが返ります。
func (p *Process) Wait() (*ProcessState, error) {
	return p.wait()
}

// SignalはProcessにシグナルを送ります。
// Windowsにおける割り込みの送信は実装されていません。
func (p *Process) Signal(sig Signal) error {
	return p.signal(sig)
}

// UserTimeは終了したプロセスとその全ての子プロセスのユーザCPU時間を返します。
func (p *ProcessState) UserTime() time.Duration {
	return p.userTime()
}

// SystemTimeは終了したプロセスとその全ての子プロセスのシステムCPU時間を返します。
func (p *ProcessState) SystemTime() time.Duration {
	return p.systemTime()
}

// Exitedはプログラムが終了したかどうかを返します。
func (p *ProcessState) Exited() bool {
	return p.exited()
}

// Successはプログラムが成功(Unixならば終了ステータス0)で終了したかどうかを
// 返します。
func (p *ProcessState) Success() bool {
	return p.success()
}

// Sysはプロセスのシステム依存な終了情報を返します。
// 内容にアクセスするには適切な実際の型(Unixならばsyscall.WaitStatus)に
// 変換して下さい。
func (p *ProcessState) Sys() interface{} {
	return p.sys()
}

// SysUsageは終了したプロセスのシステム依存なリソース利用状況の情報を返します。
// 内容にアクセスするには適切な実際の型(Unixならば*syscall.Rusage)に変換して
// ください。
// (Unixでは、*syscall.Rusageがgetrusage(2)のmanページに定義されているrusage
// 構造体に相当します。)
func (p *ProcessState) SysUsage() interface{} {
	return p.sysUsage()
}

// Hostnameはカーネルが返すホスト名を返します。
func Hostname() (name string, err error) {
	return hostname()
}

// Readdirはfに関連付けられたディレクトリの内容を読み、Lstatが返すような
// FileInfoの値をn個まで持ったスライスを返します。
// 同じfに対して続けて呼ぶと、さらにFileInfoが得られます。
//
// nが0より大きい場合、Readdirは最大n個のFileInfo構造体を返します。
// この場合では、Readdirの返すスライスが空だと、その理由を示す非nilのerrorが
// 返ります。
// ディレクトリの最後では、errorはio.EOFになります。
//
// nが0以下の場合、Readdirはそのディレクトリの全てのFileInfoを1つのスライスで
// 返します。
// この場合では、Readdirが成功(ディレクトリの全てを読んだ)だと、スライスとnilの
// errorが返ります。
// ディレクトリを全て読む前にエラーが発生すると、Readdirはそこまで読んだ結果の
// FileInfoと非nilのerrorを返します。
func (f *File) Readdir(n int) (fi []FileInfo, err error) {
	if f == nil {
		return nil, ErrInvalid
	}
	return f.readdir(n)
}

// Readdirnamesはディレクトリfを読んで名前のスライスを返します。
//
// nが0より大きい場合、Readdirnamesは最大n個の名前を返します。
// この場合では、Readdirnamesの返すスライスが空だと、その理由を示す非nilのerror
// が返ります。
//
// nが0以下の場合、Readdirnamesはそのディレクトリの全ての名前を1つのスライスで
// 返します。
// この場合では、Readdirnamesが成功(ディレクトリの全てを読んだ)だと、スライスと
// nilのerrorが返ります。
// ディレクトリを全て読む前にエラーが発生すると、Readdirnamesはそこまで読んだ
// 結果の名前と非nilのerrorを返します。
func (f *File) Readdirnames(n int) (names []string, err error) {
	if f == nil {
		return nil, ErrInvalid
	}
	return f.readdirnames(n)
}
