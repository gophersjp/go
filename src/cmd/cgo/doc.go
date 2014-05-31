// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

CgoはCのコードを呼び出すGoパッケージの作成を可能にします。

Cgoをgoコマンドで使う

Cgoを使うには、普通のGoのコードで擬似パッケージ"C"をインポートします。
これにより、GoのコードがC.size_tのような型や、 C.stdoutのような値、C.putcharのような関数を参照できるようになります。

もし"C"のインポートがコメントの直後にあれば、そのコメントはプリアンプルとみなされ、パッケージのC部分をコンパイルする際のヘッダとなります。
例えば:

	// #include <stdio.h>
	// #include <errno.h>
	import "C"

参考までに、$GOROOT/misc/cgo/stdioと$GOROOT/misc/cgo/gmpをみてください。
また、cgoを使う入門としては、"C? Go? Cgo!"をみてください: http://golang.org/doc/articles/c_go_cgo.html 。

CFLAGSとCPPFLAGS、CXXFLAGS、LDFLAGSはこのコメント中に擬似ディレクティブ#cgoで定義でき、CやC++コンパイラの振る舞いを調整することが出来ます。
複数のディレクティブで定義された値は連結され、一緒になります。
ディレクティブはビルド制約のリストを含めることができ、その制約を満たすシステムでのみ使用されるようにもできます。
(制約の書き方の詳細については、 http://golang.org/pkg/go/build/#hdr-Build_Constraints をみてください)。
例えば:

	// #cgo CFLAGS: -DPNG_DEBUG=1
	// #cgo amd64 386 CFLAGS: -DX86=1
	// #cgo LDFLAGS: -lpng
	// #include <png.h>
	import "C"

かわりに、CPPFLAGSとLDFLAGSはpkg-configツールから得ることもできます。
'#cgo pkg-config:'というディレクティブの後にパッケージ名を書きます。
例えば:

	// #cgo pkg-config: png cairo
	// #include <png.h>
	import "C"

ビルドの際には、CGO_CFLAGSとCGO_CPPFLAGS、CGO_CXXFLAGS、CGO_LDFLAGS環境変数がさきほどのディレクティブから得られたフラグに追加されます。
パッケージ固有のフラグは環境変数ではなくディレクティブで設定し、手を加えていない環境でもビルドできるべきです。

あるパッケージ内のすべてのcgo CPPFLAGSとCFLAGSディレクティブは連結され、そのパッケージ内のCのファイルをコンパイルするのに使用されます。
あるパッケージ内のすべてのCPPFLAGSとCXXFLAGSディレクティブは連結され、そのパッケージ内のC++のファイルをコンパイルするのに使用されます。
プログラムに含まれるすべてのパッケージのLDFLAGSディレクティブは連結され、リンク時に使用されます。
すべてのpkg-configディレクティブは連結されてpkg-configに同時に送られて、適切なコマンドラインフラグとして追加されます。

Goツールが一つ以上のGoのファイルが特殊なインポート"C"をみつけたとき、ほかのGoではないファイルをディレクトリ内からみつけてGoのパッケージの一部としてコンパイルします。
.cや.s、.SファイルはCコンパイラでコンパイルされます。
.ccや.cpp、.cxxファイルはC++コンパイラでコンパイルされます。
.hや.hh、.hpp、.hxxファイルは個別にはコンパイルされませんが、これらのヘッダファイルが変更された際にはCとC++のファイルは再コンパイルされます。
標準のCとC++コンパイラはCCとCXX環境変数で変更でき、これらの環境変数にはコマンドラインオプションを含めることも出来ます。

クロスコンパイルビルド中にcgoを有効にするには、CGO_ENABLED環境変数を1にセットしてGoツールをmake.bashでビルドしてください。
また、CC_FOR_TARGETにターゲット用のCクロスコンパイラを設定してください。
CCはホストのためのコンパイルに使用されます。

Goツールがビルドされた後は、goコマンドを起動する際のCC_FOR_TARGETは無視されます。
CC_FOR_TARGETの値はmake.bashを動かすときのデフォルトコンパイラです。
しかし、CC_FOR_TARGETではなく環境変数CCを設定することで、goツールを動かすときのコンパイラを設定することができます。

CXX_FOR_TARGETはC++のコードで同じように働きます。

GoからCを参照する

Goのファイルの中からは、Cの構造体のフィールド名としてGoのキーワードを使っている場合は、アンダースコアを先頭につけることでアクセスできます:
もしxが"type"という名前のフィールドをもつCの構造体を指しているなら、x._typeでアクセスできます。
Cの構造体のフィールドでGoで表現できないもの、例えばビットフィールドや間違ったアライメントをもつものは、
Goの構造体からは省略され、適切なパディングによって置き換えられ、その次のフィールドか構造体の最後に達します。

Cの標準の数値型は次の名前で使用できます:
C.char, C.schar (signed char), C.uchar (unsigned char),
C.short, C.ushort (unsigned short), C.int, C.uint (unsigned int),
C.long, C.ulong (unsigned long), C.longlong (long long),
C.ulonglong (unsigned long long), C.float, C.double。
Cのvoid*型はGoのunsafe.Pointerで表現されます。

構造体型や共用体型、列挙型に直接アクセスするためには、C.struct_statのように先頭にstruct_やunion_、enum_をつけてください。

GoはCの共用体型を一般的な形ではサポートしていないため、
Cの共用体型は同じ長さをもつGoのバイト配列として表現されます。

Goの構造体はCの型をフィールドととして埋め込めない。

CgoはCの型を、等価な非公開のGoの型に置き換えます。
この変換は非公開であるため、GoのパッケージはCの型をその公開APIとして出すべきではありません:
あるGoのパッケージで使われているCの型は、他のパッケージで使われている同じCの型とは異なります。

すべてのCの関数(void関数も含む)は多値代入コンテキストで呼ぶことができ、返り値(あれば)とCのerrno値をerrorとして取得できます
(関数がvoidを返す場合には、_を使って返り値をスキップしてください)。
例えば:

	n, err := C.sqrt(-1)
	_, err := C.voidFunc()

関数ポインタの呼び出しは現在はサポートされていませんが、Cの関数ポインタをもつGoの変数は宣言でき、それを渡すことでGoとCでやりとりすることはできます。
CのコードはGoから受け取った関数ポインタを呼び出すこともできます。
例えば:

	package main

	// typedef int (*intFunc) ();
	//
	// int
	// bridge_int_func(intFunc f)
	// {
	//		return f();
	// }
	//
	// int fortytwo()
	// {
	//	    return 42;
	// }
	import "C"
	import "fmt"

	func main() {
		f := C.intFunc(C.fortytwo)
		fmt.Println(int(C.bridge_int_func(f)))
		// Output: 42
	}

Cでは、固定長配列の関数引数は配列の先頭要素へのポインタを必要とします。
Cコンパイラはこれを検出して呼び出しを調整しますが、Goではできません。
Goでは、先頭要素へのポインタを明示的に渡す必要があります: C.f(&x[0])。

いくつかの特殊な関数によってGoとCの型をコピーして変換することが出来ます。
Goでの擬似的な定義は:

	// Go string to C string
	// Cの文字列はmallocを使ってCのヒープに割り当てられます。
	// C.freeを使って解放するのは呼び出し元の責任です(C.freeが必要なときはstdlib.hをインクルードしてください)。
	func C.CString(string) *C.char

	// C string to Go string
	func C.GoString(*C.char) string

	// C string, length to Go string
	func C.GoStringN(*C.char, C.int) string

	// C pointer, length to Go []byte
	func C.GoBytes(unsafe.Pointer, C.int) []byte

CからGoを参照する

Goの関数は次の方法でCから使えるように出来ます:

	//export MyFunction
	func MyFunction(arg1, arg2 int, arg3 string) int64 {...}

	//export MyFunction2
	func MyFunction2(arg1, arg2 int, arg3 string) (int64, *C.char) {...}

これらは、Cのコードからは次のように見えます:

	extern int64 MyFunction(int arg1, int arg2, GoString arg3);
	extern struct MyFunction2_return MyFunction2(int arg1, int arg2, GoString arg3);

これは生成されるヘッダである _cgo_export.h に含まれ、cgoの入力ファイルからプリアンプルがコピーされた後に作られます。
複数の返り値を持つ関数は、構造体を返す関数として対応付けられます。
すべてのGoの型がCの型として使いやすく対応付けられるわけではありません。

ファイルのプリアンプルに //export と書く際には制約があります:
これは二つの異なるCの出力ファイルにコピーされるため、定義を含んでいてはならず、宣言だけである必要があります。
定義は他のファイルのプリアンプルか、Cのソースファイルに置く必要があります。

Cgoを直接使う

Usage:
	go tool cgo [cgo options] [-- compiler options] file.go

Cgoは入力のfile.goを四つの出力ファイルに変換します:
二つのGoソースファイル、6c (もしくは8cか5c)のための一つのCファイル、gccのための一つのCファイル。

コンパイラへのオプションは解釈されずにパッケージのC部分をコンパイルする際にCコンパイラに渡されます。

次のオプションがcgoを直接使う場合に使用可能です。

	-dynimport file
		fileからインポートされるシンボルの一覧を出力します。
		標準出力か、 -dynout 引数で指定されたところへ出力されます。
		cgoのパッケージをビルドする際に、go buildから使われます。
	-dynout file
		ファイルに -dynimport の出力を書き出します。
	-dynlinker
		ダイナミックリンカを -dynimport の出力の一部として含めます。
	-godefs
		入力ファイルをGoの文法で出力します。
		Cのパッケージ名は実際の値で置き換えられます。
		これはsyscallパッケージを新しいターゲットに構築する際に使われます。
	-cdefs
		オプション -godefs と似ていますが、Cの文法でファイルに出力します。
		これはruntimeパッケージを新しいターゲットに構築する際に使われます。
	-objdir directory
		生成されたファイルをすべてdirectoryに出力します。
	-gccgo
		gcコンパイラではなく、gccgoコンパイラ向けに出力を生成します。
	-gccgoprefix prefix
		オプション -fgo-prefix としてgccgoで使われます。
	-gccgopkgpath path
		オプション -fgo-pkgpath としてgccgoで使われます。
	-import_runtime_cgo
		設定されていれば(デフォルト値) runtime/cgo を出力結果でインポートします。
	-import_syscall
		設定されていれば(デフォルト値) syscall を出力結果でインポートします。
	-debug-define
		デバッグ用のオプション。#defineをプリントする。
	-debug-gcc
		デバッグ用のオプション。Cコンパイラの実行と出力をトレースする。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/cmd/cgo/doc.go?r=78cda70d487e1dd239c7a385e9b97d3df3340be6
*/
package main

/*
Implementation details.

Cgo provides a way for Go programs to call C code linked into the same
address space. This comment explains the operation of cgo.

Cgo reads a set of Go source files and looks for statements saying
import "C". If the import has a doc comment, that comment is
taken as literal C code to be used as a preamble to any C code
generated by cgo. A typical preamble #includes necessary definitions:

	// #include <stdio.h>
	import "C"

For more details about the usage of cgo, see the documentation
comment at the top of this file.

Understanding C

Cgo scans the Go source files that import "C" for uses of that
package, such as C.puts. It collects all such identifiers. The next
step is to determine each kind of name. In C.xxx the xxx might refer
to a type, a function, a constant, or a global variable. Cgo must
decide which.

The obvious thing for cgo to do is to process the preamble, expanding
#includes and processing the corresponding C code. That would require
a full C parser and type checker that was also aware of any extensions
known to the system compiler (for example, all the GNU C extensions) as
well as the system-specific header locations and system-specific
pre-#defined macros. This is certainly possible to do, but it is an
enormous amount of work.

Cgo takes a different approach. It determines the meaning of C
identifiers not by parsing C code but by feeding carefully constructed
programs into the system C compiler and interpreting the generated
error messages, debug information, and object files. In practice,
parsing these is significantly less work and more robust than parsing
C source.

Cgo first invokes gcc -E -dM on the preamble, in order to find out
about simple #defines for constants and the like. These are recorded
for later use.

Next, cgo needs to identify the kinds for each identifier. For the
identifiers C.foo and C.bar, cgo generates this C program:

	<preamble>
	#line 1 "not-declared"
	void __cgo_f_xxx_1(void) { __typeof__(foo) *__cgo_undefined__; }
	#line 1 "not-type"
	void __cgo_f_xxx_2(void) { foo *__cgo_undefined__; }
	#line 1 "not-const"
	void __cgo_f_xxx_3(void) { enum { __cgo_undefined__ = (foo)*1 }; }
	#line 2 "not-declared"
	void __cgo_f_xxx_1(void) { __typeof__(bar) *__cgo_undefined__; }
	#line 2 "not-type"
	void __cgo_f_xxx_2(void) { bar *__cgo_undefined__; }
	#line 2 "not-const"
	void __cgo_f_xxx_3(void) { enum { __cgo_undefined__ = (bar)*1 }; }

This program will not compile, but cgo can use the presence or absence
of an error message on a given line to deduce the information it
needs. The program is syntactically valid regardless of whether each
name is a type or an ordinary identifier, so there will be no syntax
errors that might stop parsing early.

An error on not-declared:1 indicates that foo is undeclared.
An error on not-type:1 indicates that foo is not a type (if declared at all, it is an identifier).
An error on not-const:1 indicates that foo is not an integer constant.

The line number specifies the name involved. In the example, 1 is foo and 2 is bar.

Next, cgo must learn the details of each type, variable, function, or
constant. It can do this by reading object files. If cgo has decided
that t1 is a type, v2 and v3 are variables or functions, and c4, c5,
and c6 are constants, it generates:

	<preamble>
	__typeof__(t1) *__cgo__1;
	__typeof__(v2) *__cgo__2;
	__typeof__(v3) *__cgo__3;
	__typeof__(c4) *__cgo__4;
	enum { __cgo_enum__4 = c4 };
	__typeof__(c5) *__cgo__5;
	enum { __cgo_enum__5 = c5 };
	__typeof__(c6) *__cgo__6;
	enum { __cgo_enum__6 = c6 };

	long long __cgo_debug_data[] = {
		0, // t1
		0, // v2
		0, // v3
		c4,
		c5,
		c6,
		1
	};

and again invokes the system C compiler, to produce an object file
containing debug information. Cgo parses the DWARF debug information
for __cgo__N to learn the type of each identifier. (The types also
distinguish functions from global variables.) If using a standard gcc,
cgo can parse the DWARF debug information for the __cgo_enum__N to
learn the identifier's value. The LLVM-based gcc on OS X emits
incomplete DWARF information for enums; in that case cgo reads the
constant values from the __cgo_debug_data from the object file's data
segment.

At this point cgo knows the meaning of each C.xxx well enough to start
the translation process.

Translating Go

[The rest of this comment refers to 6g and 6c, the Go and C compilers
that are part of the amd64 port of the gc Go toolchain. Everything here
applies to another architecture's compilers as well.]

Given the input Go files x.go and y.go, cgo generates these source
files:

	x.cgo1.go       # for 6g
	y.cgo1.go       # for 6g
	_cgo_gotypes.go # for 6g
	_cgo_defun.c    # for 6c
	x.cgo2.c        # for gcc
	y.cgo2.c        # for gcc
	_cgo_export.c   # for gcc
	_cgo_main.c     # for gcc

The file x.cgo1.go is a copy of x.go with the import "C" removed and
references to C.xxx replaced with names like _Cfunc_xxx or _Ctype_xxx.
The definitions of those identifiers, written as Go functions, types,
or variables, are provided in _cgo_gotypes.go.

Here is a _cgo_gotypes.go containing definitions for C.flush (provided
in the preamble) and C.puts (from stdio):

	type _Ctype_char int8
	type _Ctype_int int32
	type _Ctype_void [0]byte

	func _Cfunc_CString(string) *_Ctype_char
	func _Cfunc_flush() _Ctype_void
	func _Cfunc_puts(*_Ctype_char) _Ctype_int

For functions, cgo only writes an external declaration in the Go
output. The implementation is in a combination of C for 6c (meaning
any gc-toolchain compiler) and C for gcc.

The 6c file contains the definitions of the functions. They all have
similar bodies that invoke runtime·cgocall to make a switch from the
Go runtime world to the system C (GCC-based) world.

For example, here is the definition of _Cfunc_puts:

	void _cgo_be59f0f25121_Cfunc_puts(void*);

	void
	·_Cfunc_puts(struct{uint8 x[1];}p)
	{
		runtime·cgocall(_cgo_be59f0f25121_Cfunc_puts, &p);
	}

The hexadecimal number is a hash of cgo's input, chosen to be
deterministic yet unlikely to collide with other uses. The actual
function _cgo_be59f0f25121_Cfunc_puts is implemented in a C source
file compiled by gcc, the file x.cgo2.c:

	void
	_cgo_be59f0f25121_Cfunc_puts(void *v)
	{
		struct {
			char* p0;
			int r;
			char __pad12[4];
		} __attribute__((__packed__, __gcc_struct__)) *a = v;
		a->r = puts((void*)a->p0);
	}

It extracts the arguments from the pointer to _Cfunc_puts's argument
frame, invokes the system C function (in this case, puts), stores the
result in the frame, and returns.

Linking

Once the _cgo_export.c and *.cgo2.c files have been compiled with gcc,
they need to be linked into the final binary, along with the libraries
they might depend on (in the case of puts, stdio). 6l has been
extended to understand basic ELF files, but it does not understand ELF
in the full complexity that modern C libraries embrace, so it cannot
in general generate direct references to the system libraries.

Instead, the build process generates an object file using dynamic
linkage to the desired libraries. The main function is provided by
_cgo_main.c:

	int main() { return 0; }
	void crosscall2(void(*fn)(void*, int), void *a, int c) { }
	void _cgo_allocate(void *a, int c) { }
	void _cgo_panic(void *a, int c) { }

The extra functions here are stubs to satisfy the references in the C
code generated for gcc. The build process links this stub, along with
_cgo_export.c and *.cgo2.c, into a dynamic executable and then lets
cgo examine the executable. Cgo records the list of shared library
references and resolved names and writes them into a new file
_cgo_import.c, which looks like:

	#pragma cgo_dynamic_linker "/lib64/ld-linux-x86-64.so.2"
	#pragma cgo_import_dynamic puts puts#GLIBC_2.2.5 "libc.so.6"
	#pragma cgo_import_dynamic __libc_start_main __libc_start_main#GLIBC_2.2.5 "libc.so.6"
	#pragma cgo_import_dynamic stdout stdout#GLIBC_2.2.5 "libc.so.6"
	#pragma cgo_import_dynamic fflush fflush#GLIBC_2.2.5 "libc.so.6"
	#pragma cgo_import_dynamic _ _ "libpthread.so.0"
	#pragma cgo_import_dynamic _ _ "libc.so.6"

In the end, the compiled Go package, which will eventually be
presented to 6l as part of a larger program, contains:

	_go_.6        # 6g-compiled object for _cgo_gotypes.go *.cgo1.go
	_cgo_defun.6  # 6c-compiled object for _cgo_defun.c
	_all.o        # gcc-compiled object for _cgo_export.c, *.cgo2.c
	_cgo_import.6 # 6c-compiled object for _cgo_import.c

The final program will be a dynamic executable, so that 6l can avoid
needing to process arbitrary .o files. It only needs to process the .o
files generated from C files that cgo writes, and those are much more
limited in the ELF or other features that they use.

In essence, the _cgo_import.6 file includes the extra linking
directives that 6l is not sophisticated enough to derive from _all.o
on its own. Similarly, the _all.o uses dynamic references to real
system object code because 6l is not sophisticated enough to process
the real code.

The main benefits of this system are that 6l remains relatively simple
(it does not need to implement a complete ELF and Mach-O linker) and
that gcc is not needed after the package is compiled. For example,
package net uses cgo for access to name resolution functions provided
by libc. Although gcc is needed to compile package net, gcc is not
needed to link programs that import package net.

Runtime

When using cgo, Go must not assume that it owns all details of the
process. In particular it needs to coordinate with C in the use of
threads and thread-local storage. The runtime package, in its own
(6c-compiled) C code, declares a few uninitialized (default bss)
variables:

	bool	runtime·iscgo;
	void	(*libcgo_thread_start)(void*);
	void	(*initcgo)(G*);

Any package using cgo imports "runtime/cgo", which provides
initializations for these variables. It sets iscgo to 1, initcgo to a
gcc-compiled function that can be called early during program startup,
and libcgo_thread_start to a gcc-compiled function that can be used to
create a new thread, in place of the runtime's usual direct system
calls.

Internal and External Linking

The text above describes "internal" linking, in which 6l parses and
links host object files (ELF, Mach-O, PE, and so on) into the final
executable itself. Keeping 6l simple means we cannot possibly
implement the full semantics of the host linker, so the kinds of
objects that can be linked directly into the binary is limited (other
code can only be used as a dynamic library). On the other hand, when
using internal linking, 6l can generate Go binaries by itself.

In order to allow linking arbitrary object files without requiring
dynamic libraries, cgo will soon support an "external" linking mode
too. In external linking mode, 6l does not process any host object
files. Instead, it collects all the Go code and writes a single go.o
object file containing it. Then it invokes the host linker (usually
gcc) to combine the go.o object file and any supporting non-Go code
into a final executable. External linking avoids the dynamic library
requirement but introduces a requirement that the host linker be
present to create such a binary.

Most builds both compile source code and invoke the linker to create a
binary. When cgo is involved, the compile step already requires gcc, so
it is not problematic for the link step to require gcc too.

An important exception is builds using a pre-compiled copy of the
standard library. In particular, package net uses cgo on most systems,
and we want to preserve the ability to compile pure Go code that
imports net without requiring gcc to be present at link time. (In this
case, the dynamic library requirement is less significant, because the
only library involved is libc.so, which can usually be assumed
present.)

This conflict between functionality and the gcc requirement means we
must support both internal and external linking, depending on the
circumstances: if net is the only cgo-using package, then internal
linking is probably fine, but if other packages are involved, so that there
are dependencies on libraries beyond libc, external linking is likely
to work better. The compilation of a package records the relevant
information to support both linking modes, leaving the decision
to be made when linking the final binary.

Linking Directives

In either linking mode, package-specific directives must be passed
through to 6l. These are communicated by writing #pragma directives
in a C source file compiled by 6c. The directives are copied into the .6 object file
and then processed by the linker.

The directives are:

#pragma cgo_import_dynamic <local> [<remote> ["<library>"]]

	In internal linking mode, allow an unresolved reference to
	<local>, assuming it will be resolved by a dynamic library
	symbol. The optional <remote> specifies the symbol's name and
	possibly version in the dynamic library, and the optional "<library>"
	names the specific library where the symbol should be found.

	In the <remote>, # or @ can be used to introduce a symbol version.

	Examples:
	#pragma cgo_import_dynamic puts
	#pragma cgo_import_dynamic puts puts#GLIBC_2.2.5
	#pragma cgo_import_dynamic puts puts#GLIBC_2.2.5 "libc.so.6"

	A side effect of the cgo_import_dynamic directive with a
	library is to make the final binary depend on that dynamic
	library. To get the dependency without importing any specific
	symbols, use _ for local and remote.

	Example:
	#pragma cgo_import_dynamic _ _ "libc.so.6"

	For compatibility with current versions of SWIG,
	#pragma dynimport is an alias for #pragma cgo_import_dynamic.

#pragma cgo_dynamic_linker "<path>"

	In internal linking mode, use "<path>" as the dynamic linker
	in the final binary. This directive is only needed from one
	package when constructing a binary; by convention it is
	supplied by runtime/cgo.

	Example:
	#pragma cgo_dynamic_linker "/lib/ld-linux.so.2"

#pragma cgo_export_dynamic <local> <remote>

	In internal linking mode, put the Go symbol
	named <local> into the program's exported symbol table as
	<remote>, so that C code can refer to it by that name. This
	mechanism makes it possible for C code to call back into Go or
	to share Go's data.

	For compatibility with current versions of SWIG,
	#pragma dynexport is an alias for #pragma cgo_export_dynamic.

#pragma cgo_import_static <local>

	In external linking mode, allow unresolved references to
	<local> in the go.o object file prepared for the host linker,
	under the assumption that <local> will be supplied by the
	other object files that will be linked with go.o.

	Example:
	#pragma cgo_import_static puts_wrapper

#pragma cgo_export_static <local> <remote>

	In external linking mode, put the Go symbol
	named <local> into the program's exported symbol table as
	<remote>, so that C code can refer to it by that name. This
	mechanism makes it possible for C code to call back into Go or
	to share Go's data.

#pragma cgo_ldflag "<arg>"

	In external linking mode, invoke the host linker (usually gcc)
	with "<arg>" as a command-line argument following the .o files.
	Note that the arguments are for "gcc", not "ld".

	Example:
	#pragma cgo_ldflag "-lpthread"
	#pragma cgo_ldflag "-L/usr/local/sqlite3/lib"

A package compiled with cgo will include directives for both
internal and external linking; the linker will select the appropriate
subset for the chosen linking mode.

Example

As a simple example, consider a package that uses cgo to call C.sin.
The following code will be generated by cgo:

	// compiled by 6g

	type _Ctype_double float64
	func _Cfunc_sin(_Ctype_double) _Ctype_double

	// compiled by 6c

	#pragma cgo_import_dynamic sin sin#GLIBC_2.2.5 "libm.so.6"

	#pragma cgo_import_static _cgo_gcc_Cfunc_sin
	#pragma cgo_ldflag "-lm"

	void _cgo_gcc_Cfunc_sin(void*);

	void
	·_Cfunc_sin(struct{uint8 x[16];}p)
	{
		runtime·cgocall(_cgo_gcc_Cfunc_sin, &p);
	}

	// compiled by gcc, into foo.cgo2.o

	void
	_cgo_gcc_Cfunc_sin(void *v)
	{
		struct {
			double p0;
			double r;
		} __attribute__((__packed__)) *a = v;
		a->r = sin(a->p0);
	}

What happens at link time depends on whether the final binary is linked
using the internal or external mode. If other packages are compiled in
"external only" mode, then the final link will be an external one.
Otherwise the link will be an internal one.

The directives in the 6c-compiled file are used according to the kind
of final link used.

In internal mode, 6l itself processes all the host object files, in
particular foo.cgo2.o. To do so, it uses the cgo_import_dynamic and
cgo_dynamic_linker directives to learn that the otherwise undefined
reference to sin in foo.cgo2.o should be rewritten to refer to the
symbol sin with version GLIBC_2.2.5 from the dynamic library
"libm.so.6", and the binary should request "/lib/ld-linux.so.2" as its
runtime dynamic linker.

In external mode, 6l does not process any host object files, in
particular foo.cgo2.o. It links together the 6g- and 6c-generated
object files, along with any other Go code, into a go.o file. While
doing that, 6l will discover that there is no definition for
_cgo_gcc_Cfunc_sin, referred to by the 6c-compiled source file. This
is okay, because 6l also processes the cgo_import_static directive and
knows that _cgo_gcc_Cfunc_sin is expected to be supplied by a host
object file, so 6l does not treat the missing symbol as an error when
creating go.o. Indeed, the definition for _cgo_gcc_Cfunc_sin will be
provided to the host linker by foo2.cgo.o, which in turn will need the
symbol 'sin'. 6l also processes the cgo_ldflag directives, so that it
knows that the eventual host link command must include the -lm
argument, so that the host linker will be able to find 'sin' in the
math library.

6l Command Line Interface

The go command and any other Go-aware build systems invoke 6l
to link a collection of packages into a single binary. By default, 6l will
present the same interface it does today:

	6l main.a

produces a file named 6.out, even if 6l does so by invoking the host
linker in external linking mode.

By default, 6l will decide the linking mode as follows: if the only
packages using cgo are those on a whitelist of standard library
packages (net, os/user, runtime/cgo), 6l will use internal linking
mode. Otherwise, there are non-standard cgo packages involved, and 6l
will use external linking mode. The first rule means that a build of
the godoc binary, which uses net but no other cgo, can run without
needing gcc available. The second rule means that a build of a
cgo-wrapped library like sqlite3 can generate a standalone executable
instead of needing to refer to a dynamic library. The specific choice
can be overridden using a command line flag: 6l -linkmode=internal or
6l -linkmode=external.

In an external link, 6l will create a temporary directory, write any
host object files found in package archives to that directory (renamed
to avoid conflicts), write the go.o file to that directory, and invoke
the host linker. The default value for the host linker is $CC, split
into fields, or else "gcc". The specific host linker command line can
be overridden using command line flags: 6l -extld=clang
-extldflags='-ggdb -O3'.  If any package in a build includes a .cc or
other file compiled by the C++ compiler, the go tool will use the
-extld option to set the host linker to the C++ compiler.

These defaults mean that Go-aware build systems can ignore the linking
changes and keep running plain '6l' and get reasonable results, but
they can also control the linking details if desired.
*/
