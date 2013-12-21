// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルを変更しないで下さい。このファイルはmkdoc.shから生成されています。
// 他のファイルのドキュメントを編集し、mkdoc.shを再実行してこのドキュメントを生成して下さい。

/*
goはGo言語のソースコードを管理するためのツールです。

Usage:

	go command [arguments]

The commands are:

    build       パッケージのコンパイル（依存関係含む）を行います
    clean       オブジェクトファイルを削除します
    env         Goの環境変数を表示します
    fix         パッケージにgo tool fixを実行します
    fmt         パッケージのソースにgofmtを実行します
    get         パッケージのダウンロードとインストール（依存関係含む）を行います
    install     パッケージのコンパイルとインストール（依存関係含む）を行います
    list        パッケージのリストを表示します
    run         Goプログラムをコンパイルし実行します
    test        パッケージのテストを行ないます
    tool        指定したgo toolを実行します
    version     Goのバージョンを表示します
    vet         パッケージにgo tool vetを実行します

commandについてさらに詳しい情報を参照するには"go help [command]"を利用して下さい。

Additional help topics:

	c           Go・C間の呼び出し FIXME:calling between Go and C
    gopath      GOPATH環境変数
    importpath  インポートパス構文
    packages    パッケージリストの説明
    testflag    テストフラグの説明
    testfunc    テスト関数の説明

topicについてさらに詳しい情報を参照するには"go help [topic]"を利用して下さい。


パッケージのコンパイル（依存関係含む）

Usage:

	go build [-o output] [build flags] [packages]

buildはインポートパスに記述されたパッケージについて、依存関係も含めて
コンパイルを行います。ただし、コンパイルしたパッケージのインストールは行いません。

.goファイルのリストが引数として与えられた場合、単一のパッケージを指定したソースファイルのリストとして扱われます。

コマンドラインで単一のmainパッケージが指定された場合、
buildの結果として実行可能ファイルが出力されます。
それ以外の場合はパッケージをコンパイルしますが結果は破棄されます。
パッケージがbuildできることのチェックができるのみとなります。

-oフラグにより出力ファイル名を指定することができます。
指定されない場合には出力ファイル名は引数によりパッケージ名から
決定されます。例えば、パッケージp（pが'main'の場合は除きます）に
対してp.aとなります。パッケージがmainで複数のファイル名が指定されている場合、
最初に指定されたファイル名が出力ファイル名となります。例えば、
'go build f1.go f2.go'に対してはf1となります。ファイルが指定されない場合
('go build')、出力ファイル名はbuild対象を含むベースのディレクトリの名前となります。

buildフラグは、build、install、run、testコマンドで共通です：

	-a
		既に最新となっているパッケージを強制的にリビルドします。
	-n
		コマンドをプリントします。実行はしません。
	-p n
		並行で走るビルドの数を指定します。
		デフォルトはCPUの利用可能数です。
	-race
		データレースの検出を有効にします。
		linux/amd64、darwin/amd64、windows/amd64のみでサポートされます。
	-v
		コンパイルされるパッケージ名をプリントします。
	-work
		一時作業ディレクトリの名前をプリントします。
		ディレクトリが既に存在する場合には削除を行ないません。
	-x
		コマンドをプリントします。

	-ccflags 'arg list'
		5c、6c、8コンパイラ呼び出しに引数を渡します。
	-compiler name
		使用するコンパイラの名前をruntime.Compiler（gccgoもしくはgc）の通り記述します。
	-gccgoflags 'arg list'
		各gccgoコンパイラ／リンカ呼び出しに渡す引数を指定します。
	-gcflags 'arg list'
		各5g、6g、8gコンパイラ呼び出しに渡す引数を指定します。
	-installsuffix suffix
		出力結果をデフォルトビルドから分けて保持するために、
		パッケージインストールディレクトリに付加するサフィックスを指定します。
		-raceフラグを使用した場合は、インストールサフィックスは自動的にraceに
		セットされますが、明示的にセットした場合には_raceが付加されます。
	-ldflags 'flag list'
		各5l、6l、8lリンカ呼び出しに渡す引数を指定します。
	-tags 'tag list'
		ビルドの条件を満たすbuildタグのリストを指定します。
		buildタグの詳細な情報についてはgo/buildパッケージの
		ドキュメントをご覧ください。

リストフラグはスペースで区切った文字列のリストとして扱われます。リストの各要素の中にスペースを埋め込む場合は、シングルクオートまたはダブルクオートでくくります。

パッケージの詳細についてのさらに詳しい情報は 'go help packages'を参照してください。
パッケージやバイナリのインストール場所についてのさらに詳しい情報は、
'go help gopath' で確認できます。Go, C/C++間の呼び出しについての
さらに詳しい情報は 'go help c' で確認できます。

See also: go install, go get, go clean.


オブジェクトファイルの削除

Usage:

	go clean [-i] [-r] [-n] [-x] [packages]

cleanはパッケージのソースディレクトリからオブジェクトファイルを削除します。
go commandが大抵のオブジェクトを一時ディレクトリの中でビルドするので、
go cleanは主にその他のツールや手動でgo buildを呼び出したときに
残されたオブジェクトファイルのために用いられます。

具体的には、インポートパスに関連付けられた各ソースディレクトリの
以下のファイルについて削除します。

	_obj/            Makefilesにより残された、古いオブジェクトディレクトリ
	_test/           Makefilesにより残された、古いテストディレクトリ
	_testmain.go     Makefilesにより残された、古いgotestファイル
	test.out         Makefilesにより残された、古いテストログ
	build.out        Makefilesにより残された、古いテストログ
	*.[568ao]        Makefilesにより残された、古いオブジェクトファイル

	DIR(.exe)        go buildから生成
	DIR.test(.exe)   go test -cから生成
	MAINFILE(.exe)   go build MAINFILE.goから生成
	*.so             SWIGから生成
 
上記のリストについて、DIRはディレクトリの末端のパス要素を表し、
MAINFILEはパッケージをビルドしたときに含まれないディレクトリの
いずれかのGoソースファイルのベースの名前となります。

-iフラグを用いるとcleanはインストールされる関連アーカイブやバイナリ
（'go install'が生成するもの）を削除します。

-nフラグを用いるとcleanは実行する削除コマンドをプリントします。
ただし、実行はしません。

-rフラグを用いるとcleanはインポートパスに記述されたパッケージの
すべての依存関係に対して再帰的に実行されます。

-xフラグを用いるとcleanは実際に事項する削除コマンドをプリントします。

パッケージの詳細についてのさらに詳しい情報は 'go help packages'を参照してください。


Go環境情報のプリント

Usage:

	go env [var ...]

envはGo環境情報をプリントします。

デフォルトではenvはshellスクリプト（Windowsではバッチファイル）
として情報をプリントします。もし引数として1つ以上の変数名が
与えられた場合、envは順番に一行ずつその名前の変数の値をプリントします。


go tool fixをパッケージに実行

Usage:

	go fix [packages]

fixはインポートパスに名付けされたパッケージにGo fixコマンドを実行します。

fixについてのさらに詳しい情報は、'godoc fix'を参照してください。
パッケージの詳細についてのさらに詳しい情報は 'go help packages'を参照してください。

特定のオプションを指定してfixを実行するには、'go tool fix'を実行して下さい。

See also: go fmt, go vet.


gofmtをパッケージのソースに実行

Usage:

	go fmt [-n] [-x] [packages]

fmtはインポートパスに指定されたパッケージに'gofmt -l -w'コマンドを実行します。
修正するファイルの名前をプリントします。

gofmtについての詳しい情報は、'godoc gofmt'を参照してください。
パッケージの指定についての桑いい情報は、'go help packages'を参照して下さい。

-nフラグはこれから実行するコマンドをプリントします。
-xフラグはコマンドを実際に実行するときにプリントします。

gofmtに特定のオプションを付加して実行する場合は、gofmtとして実行してください。

See also: go fix, go vet.


パッケージおよび依存関係のダウンロードとインストール

Usage:

	go get [-d] [-fix] [-t] [-u] [build flags] [packages]

インポートパスに指定されたパッケージを依存関係と共にダウンロードし
インストールします。

-dフラグを付加するとパッケージのダウンロードのみを行い、インストールは行いません。

-fixフラグを付加するとパッケージをダウンロードする際、依存関係の解決や
コードのビルドを行う前にfix toolを実行します。

-tフラグを付加すると指定したパッケージのテストをビルドするのに
必要なパッケージも同時にダウンロードします。

-uフラグを付加すると指定した名前のパッケージとその依存関係を
ネットワークを通じてアップデートします。デフォルトでは、ネットワークを
通じて見つからないパッケージのチェックアウトはしますが、既存のパッケージに
ついてはアップデートのチェックをしません。

getでは'go build'および'go install'コマンドにある全てのフラグを使用して
インストールをコントロールできます。'go help build'を参照してください。

パッケージのチェックアウトやアップデートを行う際には、ローカルにインストールした
Go言語のバージョンにマッチしたブランチやタグを探します。一番重要なルールは、
もしローカルにインストールしたバージョン"go1"が動いていれば、"go1"という
名前のついたブランチやタグを探しに行きます。そのバージョンが存在しない場合には
そのパッケージの一番新しいバージョンが採用されます。

パッケージの指定についてのさらに詳しい情報は、'go help packages'をご参照ください。

'go get'がどのようにソースコードを探すかについての
さらに詳しい情報は、'go help importpath'をご参照ください。

See also: go build, go install, go clean.


パッケージおよび依存関係のコンパイルとインストール

Usage:

	go install [build flags] [packages]

installはインポートパスに指定された名前のパッケージを依存関係と共に
コンパイルしインストールします。

build flagsに関するさらに詳しい情報は、'go help build'をご参照ください。
パッケージの指定についてのさらに詳しい情報は、'go help packages'をご参照ください。

See also: go build, go get, go clean.


パッケージのリストアップ

Usage:

	go list [-e] [-race] [-f format] [-json] [-tags 'tag list'] [packages]

listはインポートパスに指定された名前のパッケージを1行ずつリストアップします。
デフォルトの出力としてパッケージのインポートパスが表示されます。

    code.google.com/p/google-api-go-client/books/v1
    code.google.com/p/goauth2/oauth
    code.google.com/p/sqlite

-fフラグを付加するとパッケージテンプレートの構文を適用した
フォーマットで表示されます。デフォルトの出力は-f '{{.ImportPath}}'と
同等となります。さらにstrings.Joinを呼び出す"join"を指定した
テンプレート機能を利用することもできます。テンプレートに渡される
structは以下のとおりです。

    type Package struct {
        Dir        string // パッケージのソースが含まれるディレクトリ
        ImportPath string // dirに含まれるパッケージのインポートパス
        Name       string // パッケージ名
        Doc        string // パッケージドキュメンテーションの文字列
        Target     string // インストールパス
        Goroot     bool   // このパッケージがGo rootにあるか？
        Standard   bool   // このパッケージがGoの標準ライブラリに属しているか？
        Stale      bool   // 'go install'がこのパッケージに対して何らかの作用をするか？
        Root       string // このパッケージを含むGo rootまたはGo path dir

        // ソースファイル
        GoFiles  []string       // .goソースファイル（CgoFiles、TestGoFiles、XTestGoFilesを除く）
        CgoFiles []string       // "C"をインポートする.goソースファイル
        IgnoredGoFiles []string // ビルド制約により無視される.goソースファイル
        CFiles   []string       // .cソースファイル
        CXXFiles []string       // .cc、.cxx、.cppソースファイル
        HFiles   []string       // .h、.hh、.hpp、.hxxソースファイル
        SFiles   []string       // .sソースファイル
        SwigFiles []string      // .swigファイル
        SwigCXXFiles []string   // .swigcxxファイル
        SysoFiles []string      // アーカイブに追加するための.sysoオブジェクト

        // Cgo命令
        CgoCFLAGS    []string // cgo: Cコンパイラ用のフラグ
        CgoCPPFLAGS  []string // cgo: Cプリプロセッサ用のフラグ
        CgoCXXFLAGS  []string // cgo: C++コンパイラ用のフラグ
        CgoLDFLAGS   []string // cgo: リンカ用のフラグ
        CgoPkgConfig []string // cgo: pkg-config名

        // 依存関係情報
        Imports []string // このパッケージで利用するインポートパス
        Deps    []string // （再帰的に）インポートされるすべての依存関係

        // エラー情報
        Incomplete bool            // このパッケージまたはいずれかの依存関係にエラーが発生
        Error      *PackageError   // パッケージ読み込みエラー
        DepsErrors []*PackageError // 複数のパッケージ読み込みエラー

        TestGoFiles  []string // パッケージ内の_test.goファイル
        TestImports  []string // TestGoFilesからのインポート
        XTestGoFiles []string // パッケージ外の_test.goファイル
        XTestImports []string // XTestGoFilesからのインポート
    }

-jsonフラグはテンプレートフォーマットの代わりにJSONフォーマットで
パッケージのデータをプリントします。

-eフラグは見つからないまたは異常によりエラーとなるパッケージの扱いを
変えます。デフォルトでは、listコマンドはエラーが含まれる各パッケージの
エラーを標準エラー出力にプリントし、通常のプリントではそのパッケージを
無視します。-eフラグを用いると、listコマンドはエラーを標準エラー出力に
プリントせず、代わりに通常のプリントと同様にエラーを含むパッケージを
処理します。エラーを含むパッケージは空でないImportPathやnilでない
Errorフィールドを含むことが多いですが、その他の情報が失われて
（ゼロリセットされて）いたりする場合もあります。

-tagsフラグで'go build'コマンドと同様にビルドタグのリストを指定します。

-raceフラグではraceディテクタによりパッケージデータが必要な依存関係を含むようになります。

パッケージの指定についてのさらに詳しい情報は、'go help packages'をご参照ください。


Goプログラムのコンパイルと実行

Usage:

	go run [build flags] gofiles... [arguments...]

runは指定された名前のGoソースファイルから構成されるmainパッケージを
コンパイルし、実行します。Goソースファイルは".go"の接尾辞で終わるファイルと
定義されているます。

build flagsについての更に詳しい情報は、'go help build'をご参照ください。

See also: go build.


パッケージのテスト

Usage:

	go test [-c] [-i] [build and test flags] [packages] [flags for test binary]

'go test'はインポートパスに指定された名前のパッケージのテストを
自動で行います。下記のフォーマットによりテスト結果の概要をプリントします。

	ok   archive/tar   0.011s
	FAIL archive/zip   0.022s
	ok   compress/gzip 0.033s
	...

テストを失敗した各パッケージに対して詳細出力が続きます。

'go test'は"*_test.go"ファイルパターンにマッチする名前のすべてのファイルと
一緒に各パッケージを再コンパイルします。
"_"（"_test.go"含む）や"."から始まる名前のファイルは無視されます。
これら追加ファイルにはテスト関数、ベンチマーク関数、見本関数を含ませる
ことができます。詳細は'go help testfunc'をご参照ください。
リストアップされたパッケージはそれぞれ個々のテストバイナリとして実行させれます。

パッケージ宣言をし、"_test"接尾辞を有するテストファイルは個々のパッケージとして
コンパイルされ、mainのテストバイナリと共にリンクされ実行されます。

デフォルトではgo testには引数は必要ありません。カレントディレクトリの
ソース（テスト含む）についてパッケージをコンパイル・テストします。

パッケージは一時ディレクトリでビルドされるため、テストを行わない
インストール作業に干渉することはありません。

build flagsに加えて、'go test'自身がハンドルするフラグには以下のものがあります。

	-c  テストバイナリをpkg.testにコンパイルしますが、実行はしません。
		（pkgの部分はパッケージのインポートパスの最後の要素となります。）

	-i
		テストの依存関係のパッケージをインストールします。
		テストの実行はしません。

テストバイナリはテスト実行を制御するフラグを受け付けます。これらのフラグは
'go test'でも利用できるものです。詳細は'go help testflag'をご参照ください。

テストバイナリに対してその他任意のフラグが必要な場合は、パッケージ名の後に
してする必要があります。go toolはマイナス記号から始まる最初の引数を
ひとつのフラグとして扱います。（マイナス記号自体は認識されません。）
その引数とその後に続くすべての引数がテストバイナリの引数として渡されます。

build flagsについてのさらに詳しい情報は、'go help build'をご参照ください。
パッケージの指定についてのさらに詳しい情報は、'go help packages'をご参照ください。

See also: go build, go vet.


特定のgo toolの実行

Usage:

	go tool [-n] command [args...]

toolは引数に指定したgo toolコマンドを実行します。
引数を指定しない場合には、既存のツールの一覧をプリントします。

-nフラグでは実行しようとしているコマンドをプリントしますが、実行はしません。

各toolコマンドについてのさらに詳しい情報は、'go tool command -h'をご参照ください。


Goバージョンのプリント

Usage:

	go version

versionはruntime.Versionで得られるのと同様に、Goのバージョンをプリントします。


パッケージへのgo tool vetの実行

Usage:

	go vet [-n] [-x] [packages]

vetはインポートパスに指定されている名前のパッケージに対し、Goのvetコマンドを実行します。

vetについてのさらに詳しい情報は、'godoc code.google.com/p/go.tools/cmd/vet'をご参照ください。
パッケージの指定についてのさらに詳しい情報は、'go help packages'をご参照ください。

特定のオプションを付加してvet toolを実行する場合は、'go tool vet'を実行します。

-nフラグは実行しようとしているコマンドをプリントします。
-xフラグは実行中にコマンドをプリントします。

See also: go fmt, go fix.


Go、C間の呼び出し

GoとC/C++コード感の呼び出しの方法は2つあります。

1つ目はGoのディストリビューションの一部であるcgo toolです。
使用方法については、cgoのドキュメント（godoc cmd/cgo）をご参照ください。

2つ目は言語間インタフェースの汎用ツールであるSWIGプログラムです。
SWIGについての情報はhttp://swig.org/をご参照ください。
go buildを実行する際、.swig拡張子を持つ任意のファイルがSWIGに渡されます。
-c++オプションにより、.swigcxx拡張子を持つ任意のファイルがSWIGに渡されます。

cgo、SWIGのいずれかを使用する場合、go buildは任意の.c、.s、.Sファイルを
Cコンパイラに渡し、任意の.cpp、.cxxファイルをC++コンパイラに渡します。
CCまたはCXX環境変数によりC、C++コンパイラそれぞれに対して使用するコンパイラを
指定することができます。


GOPATH環境変数

Go pathはインポート文を解決するために使用されます。
go/buildパッケージに実装、ドキュメントされています。

GOPATH環境変数はGoのコードを探す場所をリストアップします。
Unixでは、値はコロンで区切った文字列です。
Windowsでは、値はセミコロンで区切った文字列です。
Plan 9では、値はリストです。

GOPATHは、標準のGoツリー以外のパッケージをget、build、installするために
必ず設定する必要があります。

GOPATHにリストアップされている各ディレクトリは規定された構造を持つ必要があります。

src/ディレクトリはソースコードを保持します。'src'以下のディレクトリが
インポートパスや実行ファイル名決定します。

pkg/ディレクトリはインストール済みのパッケージオブジェクトを保持します。
Go treeにしたがって、各対象のオペレーティングシステムとアーキテクチャの
組み合わせでpkgのサブディレクトリが構成されます。
(pkg/GOOS_GOARCH).

DIRがGOPATHに含まれるディレクトリである場合は、DIR/src/foo/bar内に
ソースを有するパッケージが"foo/bar"としてインポートされ、
"DIR/pkg/GOOS_GOARCH/foo/bar.a"にコンパイルされた形でインストールされます。

bin/ディレクトリはコンパイル済みのコマンドを保持します。
各コマンドはそのソースディレクトリの名前となりますが、パス全体ではなく
最後の要素のみの名前となります。つまり、DIR/src/foo/quuxにソースがある
コマンドは、/foo/quuxではなく、DIR/bin/quuxにインストールされます。
DIR/binをPATHに追加することでインストールしたコマンドを利用できるように
foo/は取り除かれます。GOBIN環境変数がセットされている場合は、
DIR/binの代わりにそのディレクトリにコマンドがインストールされます。

以下はディレクトリレイアウトの一例です。

    GOPATH=/home/user/gocode

    /home/user/gocode/
        src/
            foo/
                bar/               （barパッケージ内のgoコード）
                    x.go
                quux/              （mainパッケージ内のgoコード）
                    y.go
        bin/
            quux                    （インストール済みのコマンド）
        pkg/
            linux_amd64/
                foo/
                    bar.a          （インストール済みのパッケージオブジェクト）

Goはソースコードを探す際、GOPATHにリストアップされている各ディレクトリを
検索しますが、新しいパッケージは常にリストの最初のディレクトリにダウンロードされます。


インポートパス構文

インポートパス（'go help packages'をご参照ください）はローカルファイルシステムに
格納されたパッケージを示します。一般的に、インポートパスは標準パッケージ
（例えば"unicode/utf8"）かいずれかのワークスペース（'go help gopath'を
ご参照ください）の中から見つかったパッケージのどちらか一方を示します。

相対インポートパス

./や../で始まるインポートパスは相対パスと呼ばれます。
toolchainは2つの方法により相対パスをショートカットとしてサポートします。

第一に、相対パスはコマンドライン上で簡略な表現として用いられます。
"unicode'としてインポートされたコードを含むディレクトリに対して
"unicode/utf8"のテストを実行したい場合には、フルパスで記述する
必要はなく、"go test ./utf8"と指定することができます。
逆の場合でも同様に、"go test .."では"unicode/utf8"ディレクトリから
"unicode"をテストできます。相対指定のパターンを組み合わせることも
可能で、たとえば"go test ./..."はすべてのサブディレクトリをテストします。
パターンの文法に関する詳しい情報は、'go help packages'をご参照ください。

第二に、Goプログラムをワークスペース外でコンパイルする場合、
同様にワークスペース外で近くにあるコードを参照するプログラムの
インポート文に相対パスを利用できます。
これにより通常のワークスペースの外での小規模な複数パッケージの
プログラムを用いた実験が容易になります。ただし、これらのプログラムは
"go install"でインストールすることはできません。
（インストール先のワークスペースがありません）
そのため、ビルドするごとにスクラッチから組み直す必要があります。
曖昧さ回避のため、Goプログラムでは単一ワークスペース内で相対パスを
用いることはできません。

リモートインポートパス

インポートパスによってはリビジョン管理システムを用いた
パッケージのソースコードの取得方法を同時に表しているものがあります。

いくつかのよく知られるコードホスティングサイトには特有の書き方があります。

	Bitbucket (Git, Mercurial)

		import "bitbucket.org/user/project"
		import "bitbucket.org/user/project/sub/directory"

	GitHub (Git)

		import "github.com/user/project"
		import "github.com/user/project/sub/directory"

	Google Code Project Hosting (Git, Mercurial, Subversion)

		import "code.google.com/p/project"
		import "code.google.com/p/project/sub/directory"

		import "code.google.com/p/project.subrepository"
		import "code.google.com/p/project.subrepository/sub/directory"

	Launchpad (Bazaar)

		import "launchpad.net/project"
		import "launchpad.net/project/series"
		import "launchpad.net/project/series/sub/directory"

		import "launchpad.net/~user/project/branch"
		import "launchpad.net/~user/project/branch/sub/directory"

その他のサーバがホストするコードについては、インポートパスにバージョン管理
のタイプの修飾子を指定するか、go toolにてhttps/httpでインポートパスを
動的に取得しHTMLの<meta>タグからコードがある場所を探します。

コードの場所を指定する際には、インポートパス形式

	repository.vcs/path

は、指定された名前のバージョン管理システムを用いて、.vcs接尾辞を含む場合・
含まない場合のいずれかについて与えられたレポジトリを明示し、次いで
レポジトリの内部のパスを指定します。サポートするバージョン管理システムは
以下のとおりです。

	Bazaar      .bzr
	Git         .git
	Mercurial   .hg
	Subversion  .svn

たとえば

	import "example.org/user/foo.hg"

は、Mercurialレポジトリのexample.org/user/fooまたはfoo.hgの
ルートディレクトリを表し、

	import "example.org/repo.git/foo/bar"

は、Gitレポジトリのexample.com/repoまたはrepo.gitの
foo/barディレクトリを表します。

バージョン管理システムが複数のプロトコルをサポートする場合、
ダウンロード時にそれぞれが順番に試行されます。例えば、
Gitのダウンロードはまずgit://を試行し、次にhttps://、
最後にhttp://を施行します。

インポートパスが良く知られたコードホスティングサイトではなく、
さらにバージョン管理の修飾子がない場合、go toolはhttps/httpにより
インポートを取ってくるように試行し、ドキュメントのHTML <head>内の
<meta>タグを探します。

metaタグは以下の形式です：

	<meta name="go-import" content="import-prefix vcs repo-root">

import-prefixはレポジトリルートに対応するインポートパスです。
"go get"で取ってくるパッケージのに完全一致するか接尾辞であるか必要があります。
完全一致でない場合は、<meta>タグの一致を検証するために接尾辞について
別のhttpリクエストが生成されます。

vcsは"git"、"hg"、"svn"などのうちの一つです。

repo-rootはスキームを含むが.vcs拡張子を含まないバージョン管理システムの
ルートです。

たとえば、

	import "example.org/pkg/foo"

の結果は以下のリクエストとなります：

	https://example.org/pkg/foo?go-get=1 (推奨)
	http://example.org/pkg/foo?go-get=1  (代替)

該当のページにmetaタグが含まれている場合

	<meta name="go-import" content="example.org git https://code.org/r/p/exproj">

go toolはhttps://example.org/?go-get=1が同じmetaタグを含むことを確認し、
GOPATH/src/example.orgにgit clone https://code.org/r/p/exprojします。

新たにダウンロードされたパッケージはGOPATH環境変数に挙げられている
最初のディレクトリに書き出されます。（'go help gopath'をご参照ください）

goコマンドは利用しているGoリリースの適切なパッケージのバージョンを
ダウンロードするようにします。
詳細については'go help install'を実行して下さい。


パッケージリストの説明

多くのコマンドが一連のパッケージに適用できます：

	go action [packages]

通常、[packages]はインポートパスのリストです。

ルートパスまたは'.'、'..'要素で始まるインポートパスはファイルシステムパスと
解釈され、そのディレクトリにあるパッケージであることを表します。

それ以外の場合、インポートパスPはGOPATH環境変数に挙げられている
あるDIRに対してDIR/src/Pディレクトリの中に存在するパッケージを
表します。（'go help gopath'をご参照ください。）

インポートパスが与えられていない場合、actionはカレントディレクトリの
パッケージに対して適用されます。

go toolでビルドされるパッケージに利用できないパスの予約名が3つあります：

- "main"はスタンドアロン実行ファイルのトップレベルパッケージを表します。

- "all"は全GOPATHツリーに存在するすべてのパッケージディレクトリに
展開します。たとえば、'go list all'はローカルシステムのすべてのパッケージを
リストアップします。

- "std"はallに似ていますが、標準Goライブラリのパッケージだけに
展開します。

インポートパスが1つ以上の"..."ワイルドカード（それぞれが空文字・スラッシュ
を伴う文字列を含むいずれの文字列にマッチする）を含む場合、
それらのパターンにマッチする名前を含むGOPATHツリーに存在するすべての
パッケージディレクトリに展開します。特別な場合としては、x/...は
xとxのサブディレクトリにマッチします。たとえば、net/...はnetとその
サブディレクトリのパッケージに展開します。

インポートパスにはリモートディレクトリからダウンロードされるパッケージの
名前をつけることができます。詳細については'go help importpath'をご参照ください。

プログラム内のいかなるパッケージでも必ず一意なインポートパスを
備えている必要があります。規則により、作成者に属する一意な接頭辞から
パスを始めるように取り決められています。たとえば、Googleで内部的に
使用されるパスはすべて'google'で始まり、リモートレポジトリを表すパスは
'code.google.com/p/project'のようにコードへのパスから始まります。

特別な場合として、パッケージリストが単一ディレクトリからの.goファイルの
一覧で構成されているときには、それらのファイルから作り上げられた単一の
合成パッケージにコマンドが適用されます。その際、それらファイルのすべての
ビルド制約、およびディレクトリ内のその他のファイルは無視されます。

"."や"_"から始まるファイル名はgo toolにより無視されます。

テストに関するフラグについての説明

'go test'コマンドは'go test'自体に適用されるフラグと生成されるテストバイナリ
に適用されるフラグの両方を扱います。

フラグのうちいくつかはプロファイルを制御し、"go tool pprof"に適合する
実行プロファイルを書き出します。更に詳しい情報については"go tool pprof help"
をご参照下さい。pprofの--alloc_space、--alloc_objects、--show_bytes
オプションは情報の表示のされ方を制御します。

以下のフラグは'go test'コマンドで利用でき、テスト実行の制御を行います：

	-bench regexp
		正規表現にマッチするベンチマークを実行します。
		デフォルトではベンチマークは実行されません。
		すべてのベンチマークを実行するには'-bench .'または'-bench=.'
		を利用します。

	-benchmem
		ベンチマークのためのメモリ割り当て統計情報をプリントします。

	-benchtime t
		time.Durationで規定された形式（たとえば、-benchtime 1h30s）で
		指定した時間tの間、各ベンチマークを繰り返し実行します。
		デフォルトは1秒（1s）です。

	-blockprofile block.out
		すべてのテストが終了する際に、指定したファイルにgoroutineの
		ブロックプロファイルを書き出します。

	-blockprofilerate n
		nに対しruntime.SetBlockProfileRateを呼び出すことにより、goroutine
		のブロックプロファイル内に与えられた詳細を制御します。
		'godoc runtime SetBlockProfileRate'をご参照ください。
		プロファイラはプログラムがブロックされているnナノ秒ごとに
		1つのブロックイベントとして平均的にサンプルするように努めます。
		デフォルトでは、-test.blockprofileをフラグなしで指定した場合には
		すべてのブロックイベントが記録されます。-test.blockprofilerate=1
		と同等となります。

	-cover
		カバレッジ解析を有効にします。

	-covermode set,count,atomic
		テストを実行するパッケージに対してカバレッジ解析モードを設定します。
		デフォルトは"set"です。
		値：
		set：bool：このステートメントが実行されるか？
		count：int：このステートメントが何回実行されるか？
		atomic：int：countと同等ですが、マルチスレッドのテストを正確に
			実行できます。ただし、著しくコストが掛かります。
		-coverを自動的に設定します。

	-coverpkg pkg1,pkg2,pkg3
		与えられたパッケージのリストに対する各テストにカバレッジ解析を
		適用します。デフォルトではテストを行うパッケージのみに対して
		各テストでの解析を行います。
		パッケージはインポートパスで指定されます。
		-coverを自動的に設定します。

	-coverprofile cover.out
		すべてのテストをパスした後に指定したファイルにカバレッジプロファイル
		を書き出します。
		-coverを自動的に設定します。

	-cpu 1,2,4
		テストもしくはベンチマークを実行するGOMAXPROCSの値のリストを
		指定します。デフォルトはGOMAXPROCSの現在の値となります。

	-cpuprofile cpu.out
		終了前に指定したファイルにCPUプロファイルを書き出します。

	-memprofile mem.out
		すべてのテストをパスした後に指定したファイルにメモリプロファイル
		を書き出します。

	-memprofilerate n
		runtime.MemProfileRateのセッティングにより、より正確な
		（ただしコストの掛かる）メモリプロファイルを有効にします。 
		'godoc runtime MemProfileRate'をご参照ください。
		すべてのメモリ割り当てをプロファイルするためには、
		-test.memprofilerate=1を利用し環境変数GOGC=offを設定して
		ガベッジコレクタを無効化してください。これにより、
		ガベッジコレクションを行うことなく利用可能なメモリをフルに
		使用してテストを実行できます。

	-outputdir directory
		指定したディレクトリにプロファイルからの出力ファイルを置きます。
		デフォルトは"go test"を実行しrているディレクトリです。

	-parallel n
		t.Parallelを呼び出すテスト関数の並行実行を許可します。
		このフラグの値は同時に実行するテストの最大数となります。
		デフォルトでは、GOMAXPROCSの値に設定されます。

	-run regexp
		正規表現にマッチするテストや例のみ実行します。

	-short
		長時間実行しているテストに対して実行時間を短縮するように要求します。
		デフォルトではオフになっていますが、all.bashを実行する際、
		Go treeによりサニティチェックが実行されますが、全数テストを
		実行するのに時間を費やさないようにオンに設定されます。

	-timeout t
		テストがtよりも長く実行している場合、panicを呼び出します。

	-v
		冗長な出力：実行するすべてのテストのログを出力します。また、
		テストが成功した場合にもLog、Logfの呼び出しからのすべての
		テキストをプリントします。

pkg.test（pkgはパッケージソースが含まれるディレクトリ名を表します）
と呼ばれるテストバイナリが、'go test -c'でビルドされた後直接呼び出されます。
テストバイナリディレクトリを呼び出す際、各標準フラグ名は
-test.run=TestMyFuncや-test.vのように、'test.'を接尾辞として付加する
必要があります。

'go test'を実行する際、上記の一覧に記述されていないフラグは変更なしに
渡されます。たとえば、コマンド

	go test -x -v -cpuprofile=prof.out -dir=testdata -update

はテストバイナリをコンパイルし、

	pkg.test -test.v -test.cpuprofile=prof.out -dir=testdata -update

として実行します。

カバレッジ以外のプロファイルを生成するテストフラグについても、プロファイルを
解析する際に利用できるようにpkg.testにテストバイナリを残します。

'go test'で認識されないフラグはすべてのパッケージ指定の後に位置している
必要があります。


テスト関数の説明

'go test'はtestの下にあるパッケージに対応する"*_test.go"ファイル内のテスト、ベンチマーク、見本関数を検索することを想定しています。

テスト関数はTestXXX（XXXは小文字から始まらない任意の英数文字列）という
名称で、下記のシグネチャを有する必要があります。

	func TestXXX(t *testing.T) { ... }

ベンチマーク関数はBenchmarkXXXという名称で、以下のシグネチャを有する
必要があります。

	func BenchmarkXXX(b *testing.B) { ... }

見本関数はテスト関数に似ていますが、成功・失敗を伝えるために*testing.T
を用いる代わりに、出力をos.Stdoutをプリントします。出力が、関数本体の
最後のコメントとして必ず記載される"Output:"コメントと比較されます。
（下記の例をご参照ください。）そのようなコメントない、もしくは"Output:"
の後が空白の見本はコンパイルされますが実行はされません。

godocは関数、定数、変数XXXの利用方法の説明のためにExampleXXXの本体を
表示します。T型や*T型をレシーバに持つメソッドMの見本はExampleT_Mと
名付けられます。_xxx（xxxは大文字以外で始まる接尾辞）を後ろに付けて
識別できるようにすることで、特定の関数、定数、変数について複数の
見本関数を定義することもできます。

以下は見本の例です：

	func ExamplePrintln() {
		Println("The output of\nthis example.")
		// Output: The output of
		// this example.
	}

テストファイルが単一の見本関数を有する場合（その他の関数や型や変数や
定数定義が少なくとも1つあり、テストやベンチマーク関数がない場合）は、
テストファイルの全体が見本として扱われます。

さらに詳しい情報についてはtestingパッケージのドキュメントをご参照ください。


本ドキュメントは以下のドキュメントを翻訳しています:https://code.google.com/p/go/source/browse/src/cmd/go/doc.go?r=c4d996668981f0bd23d3a6be0fff72d09334b380
*/
package main
