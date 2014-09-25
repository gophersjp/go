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
    env         Go環境変数をプリントします
    fix         パッケージにgo tool fixを実行します
    fmt         パッケージのソースにgofmtを実行します
    generate    ソースを処理してGoのファイルを生成します
    get         パッケージのダウンロードとインストール（依存関係含む）を行います
    install     パッケージのコンパイルとインストール（依存関係含む）を行います
    list        パッケージのリストを表示します
    run         Goプログラムをコンパイルし実行します
    test        パッケージのテストを行ないます
    tool        指定したgo toolを実行します
    version     Goのバージョンをプリントします
    vet         パッケージにgo tool vetを実行します

commandについてさらに詳しい情報を参照するには"go help [command]"を利用して下さい。

Additional help topics:

    c           Go・C間の呼び出し
    filetype    ファイルタイプ
    gopath      GOPATH環境変数
    importpath  インポートパス構文
    packages    パッケージリストの説明
    testflag    テストフラグの説明
    testfunc    テスト関数の説明

topicについてさらに詳しい情報を参照するには"go help [topic]"を利用して下さい。


Goパッケージと依存関係のコンパイル

Usage:

	go build [-o output] [-i] [build flags] [packages]

buildはインポートパスに記述されたパッケージについて、依存関係も含めてコンパイルを行います。ただし、コンパイルしたパッケージのインストールは行いません。

.goファイルのリストが引数として与えられた場合、単一のパッケージを指定したソースファイルのリストとして扱われます。

コマンドラインで単一のmainパッケージが指定された場合、buildの結果として実行可能ファイルが出力されます。それ以外の場合はパッケージをコンパイルしますが結果は破棄されます。パッケージがbuildできることのチェックができるのみとなります。

-oフラグにより出力ファイル名を指定することができます。指定されない場合には出力ファイル名は引数によりパッケージ名から決定されます。例えば、パッケージp（pが'main'の場合は除きます）に対してp.aとなります。パッケージがmainで複数のファイル名が指定されている場合、最初に指定されたファイル名が出力ファイル名となります。例えば、'go build f1.go f2.go'に対してはf1となります。ファイルが指定されない場合('go build')、出力ファイル名はbuild対象を含むベースのディレクトリの名前となります。

-iフラグはターゲットが依存するパッケージをインストールします。

buildフラグは、build、install、run、testコマンドで共通です。

	-a
		既に最新となっているパッケージを強制的にリビルドします。
	-n
		コマンドをプリントします。実行はしません。
	-p n
		並行で走るビルドの数を指定します。
		デフォルトはCPUの利用可能数です。
	-race
		データレースの検出を有効にします。
		linux/amd64、freebsd/amd64、darwin/amd64、windows/amd64のみでサポートされます。
	-v
		コンパイルされるパッケージ名をプリントします。
	-work
		一時作業ディレクトリの名前をプリントします。
		該当ディレクトリが既に存在する場合でも削除されることはありません。
	-x
		コマンドをプリントします。

	-ccflags 'arg list'
		5c、6c、8コンパイラ呼び出しに引数を渡します。
	-compiler name
		使用するコンパイラ名をruntime.Compiler（gccgoもしくはgc）の通り記述します。
	-gccgoflags 'arg list'
		各gccgoコンパイラ／リンカ呼び出しに渡す引数を指定します。
	-gcflags 'arg list'
		各5g、6g、8gコンパイラ呼び出しに渡す引数を指定します。
	-installsuffix suffix
		出力結果をデフォルトビルドから分けて保持するために、
		パッケージインストールディレクトリに付与するサフィックスを指定します。
		-raceフラグを使用した場合は、インストールサフィックスは自動的にraceに
		セットされますが、明示的にセットした場合には_raceが付加されます。
	-ldflags 'flag list'
		各5l、6l、8lリンカ呼び出しに渡す引数を指定します。
	-tags 'tag list'
		ビルドの条件を満たすbuildタグのリストを指定します。
		buildタグの詳細な情報についてはgo/buildパッケージの
		ビルド制約(Build Constraints) についてのドキュメントをご覧ください。

リストフラグはスペースで区切った文字列のリストとして扱われます。リストの各要素の中にスペースを埋め込む場合は、シングルクオートまたはダブルクオートでくくります。

パッケージの詳細についてのさらに詳しい情報は 'go help packages'を参照してください。パッケージやバイナリのインストール場所についてのさらに詳しい情報は、'go help gopath' で確認できます。Go, C/C++間の呼び出しについてのさらに詳しい情報は 'go help c' で確認できます。

See also: go install, go get, go clean.


Goオブジェクトファイルの削除

Usage:

	go clean [-i] [-r] [-n] [-x] [build flags] [packages]

cleanはパッケージのソースディレクトリからオブジェクトファイルを削除します。go commandが大抵のオブジェクトを一時ディレクトリの中でビルドするので、go cleanは主にその他のツールや手動でgo buildを呼び出したときに残されたオブジェクトファイルのために用いられます。

具体的には、インポートパスに関連付けられた各ソースディレクトリの以下のファイルについて削除します。

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

上記のリストについて、DIRはディレクトリの末端のパス要素を表し、MAINFILEはパッケージをビルドしたときに含まれないディレクトリ内のいずれかのGoソースファイルのベースの名前となります。

-iフラグを用いるとcleanはインストールされる関連アーカイブやバイナリ（'go install'が生成するもの）を削除します。

-nフラグを用いるとcleanは実行する削除コマンドをプリントします。ただし、実行はしません。

-rフラグを用いるとcleanはインポートパスに記述されたパッケージのすべての依存関係に対して再帰的に実行されます。

-xフラグを用いるとcleanは実行中に随時削除コマンドをプリントします。

build flagsに関するさらに詳しい情報は、'go help build'を参照してください。

パッケージの詳細についてのさらに詳しい情報は 'go help packages'を参照してください。


Go環境情報のプリント

Usage:

	go env [var ...]

envはGo環境情報をプリントします。

デフォルトではenvはshellスクリプト（Windowsではバッチファイル）として情報をプリントします。引数として1つ以上の変数名が与えられた場合、envは順番に一行ずつその名前の変数の値をプリントします。


Goパッケージにgo tool fixを実行

Usage:

	go fix [packages]

fixはインポートパスに指定された名前のパッケージにGo fixコマンドを実行します。

fixについてのさらに詳しい情報は、'godoc fix'を参照してください。パッケージの詳細についてのさらに詳しい情報は 'go help packages'を参照してください。

特定のオプションを指定してfixを実行するには、'go tool fix'を実行して下さい。

See also: go fmt, go vet.


Goパッケージのソースにgofmtを実行

Usage:

	go fmt [-n] [-x] [packages]

fmtはインポートパスに指定されたパッケージに'gofmt -l -w'コマンドを実行します。修正するファイルの名前をプリントします。

gofmtについての詳しい情報は、'godoc gofmt'を参照してください。パッケージの指定についての詳しい情報は、'go help packages'を参照して下さい。

-nフラグはこれから実行するコマンドをプリントします。-xフラグは実行中に随時コマンドをプリントします。

gofmtに特定のオプションを付加して実行する場合は、gofmtとして実行してください。

See also: go fix, go vet.


Goのファイルをソースを処理して生成

Usage:

	go generate [-run regexp] [file.go... | packages]

generateはファイルのディレクティブ中で記述されたコマンドを実行します。
これらのコマンドではどのような処理をすることもできますが、ここでの意図はyaccの実行のようなGoのソースファイルの生成もしくは更新です。

go generateはgo buildやgo get、go testなどによって自動で実行されることはありません。
必ず、明示的に実行する必要があります。

ディレクティブは、次の形式の一行コメントとして記述され、

	//go:generate command argument...

コマンド(command)は実行されるべきジェネレータで、ローカルで実行可能な実行ファイルに対応します(注: "//go"にスペースはありません)。
これはshellのパス中にある(gofmt)か、完全に修飾されたパス(/usr/you/bin/mytool)か、以下に述べるコマンドの別名である必要があります。

引数(argument...)はスペースで区切られたトークンか、ダブルクオートされた文字列で、ジェネレータへ個々の引数として実行時に与えられます。

クオートされた文字列はGoの文法を使い、実行前に評価されます; クオートされた文字列は1個の引数としてジェネレータに渡されます。

go generateはジェネレータの実行時にいくつかの変数を設定します:

	$GOARCH
		実行時のアーキテクチャ (arm, amd64など)
	$GOOS
		実行時のオペレーティングシステム (linux, windowsなど)
	$GOFILE
		ファイルの名前
	$GOPACKAGE
		ディレクティブを含むファイルのパッケージ名

変数の展開やクオートされた文字列の評価以外には、コマンドライン上でのglob処理のような特別な処理は行われません。

コマンドが実行される直前には、英数字から成る環境変数への参照、例えば $GOFILE や $HOME はコマンドライン上ですべて展開されます。
変数展開のための文法はオペレーティングシステムによらず $NAME です。
評価順序の関係で、クオートされた文字列の内部であっても変数は展開されます。
もし変数名 NAME がセットされていなければ、 $NAME は空文字列に展開されます。

次の形式のディレクティブ、

	//go:generate -command xxx args...

は、引数argsとして実行されるコマンドxxxを定義します。
ソースファイルのこれ以降の部分のみで有効です。
これは、コマンドの別名をつくったり、複数語からなるジェネレータを扱うために使うことができます。
例えば、

	//go:generate -command yacc go tool yacc

は、ジェネレータ"go tool yacc"を表すコマンド"yacc"を定義します。

generateは、複数のパッケージをコマンドラインで与えられた順序で一つずつ処理します。
一つのパッケージ内では、generateはソースファイルを名前順に一つずつ処理します。
一つのソースファイル内では、generateはファイル中で現れた順に一つずつジェネレータを起動します。

もしジェネレータが一つでもエラーの終了ステータスを返せば、generateはそのパッケージの残りの処理をすべてスキップします。

ジェネレータはパッケージのソースディレクトリ上で実行されます。

go generateは一つの指定フラグを受け付けます:

	-run=""
		もし空でなければ、正規表現として解釈され、マッチするコマンドをもつディレクティブのみが実行されます。

また、-v、-nや-xなど、標準的なビルドフラグも受け付けます。
フラグ-vは処理されるパッケージとファイルの名前を出力します。
フラグ-nは実行されるであろうコマンドを出力します。
フラグ-xは実行されたコマンドを出力します。

パッケージの指定についてのさらに詳しい情報は、'go help packages'を参照してください。


Goパッケージおよび依存関係のダウンロードとインストール

Usage:

	go get [-d] [-fix] [-t] [-u] [build flags] [packages]

インポートパスに指定されたパッケージを依存関係と共にダウンロードしインストールします。

-dフラグを付加するとパッケージのダウンロードのみを行い、インストールは行いません。

-fixフラグを付加するとパッケージをダウンロードする際、依存関係の解決やコードのビルドを行う前にfix toolを実行します。

-tフラグを付加すると指定したパッケージのテストをビルドするのに必要なパッケージも同時にダウンロードします。

-uフラグを付加すると指定した名前のパッケージとその依存関係をネットワークを通じてアップデートします。デフォルトでは、ネットワークを通じて足りないパッケージのチェックアウトはしますが、既存のパッケージについてはアップデートのチェックをしません。

getではインストールを制御するためのbuild flagsも受け付けます。'go help build'を参照してください。


パッケージのチェックアウトやアップデートを行う際には、ローカルにインストールされたGo言語のバージョンにマッチしたブランチやタグを探します。一番重要なルールとして、ローカルにインストールしたバージョン"go1"で動作していれば、"go1"という名前のついたブランチやタグを探しに行きます。そのバージョンが存在しない場合にはそのパッケージの一番新しいバージョンが採用されます。

パッケージの指定についてのさらに詳しい情報は、'go help packages'を参照してください。

'go get'がどのようにソースコードを探すかについてのさらに詳しい情報は、'go help importpath'を参照してください。

See also: go build, go install, go clean.


Goパッケージおよび依存関係のコンパイルとインストール

Usage:

	go install [build flags] [packages]

installはインポートパスに指定された名前のパッケージを依存関係と共にコンパイルしインストールします。

build flagsに関するさらに詳しい情報は、'go help build'を参照してください。パッケージの指定についてのさらに詳しい情報は、'go help packages'を参照してください。

See also: go build, go get, go clean.


Goパッケージの一覧

Usage:

	go list [-e] [-f format] [-json] [build flags] [packages]

listはインポートパスに指定された名前のパッケージを1行ずつリストアップします。デフォルトの出力としてパッケージのインポートパスが表示されます。

    code.google.com/p/google-api-go-client/books/v1
    code.google.com/p/goauth2/oauth
    code.google.com/p/sqlite

-fフラグを付加するとパッケージテンプレートの構文を適用したフォーマットで表示されます。デフォルトの出力は-f '{{.ImportPath}}'と同等となります。テンプレートに渡されるstructは以下のとおりです。

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


テンプレート関数"join"はstrings.Joinを呼び出します。

テンプレート関数"context"は次のビルドコンテキストを返します:

	type Context struct {
		GOARCH        string   // ターゲットアーキテクチャ
		GOOS          string   // ターゲットオペレーティングシステム
		GOROOT        string   // Go root
		GOPATH        string   // Go path
		CgoEnabled    bool     // cgoが使用できるかどうか
		UseAllFiles   bool     // ファイルを+build行やファイル名によらず使用するかどうか
		Compiler      string   // ターゲットのパスを計算するときに想定するコンパイラ
		BuildTags     []string // +build行にマッチするビルド制約
		ReleaseTags   []string // 互換性のあるリリースの一覧
		InstallSuffix string   // インストールディレクトリの名前に使用するサフィックス
	}

これらのフィールドの詳細な意味はgo/buildパッケージのContext型のドキュメントを参照してください。

-jsonフラグはテンプレートフォーマットの代わりにJSONフォーマットでパッケージのデータをプリントします。

-eフラグは見つからないまたは異常によりエラーとなるパッケージの扱いを変えます。デフォルトでは、listコマンドはエラーが含まれる各パッケージのエラーを標準エラー出力にプリントし、通常のプリントではそのパッケージを無視します。-eフラグを用いると、listコマンドはエラーを標準エラー出力にプリントせず、代わりに通常のプリントと同様にエラーを含むパッケージを処理します。エラーを含むパッケージは空でないImportPathやnilでないErrorフィールドを含むことが多いですが、その他の情報が失われて（ゼロリセットされて）いたりする場合もあります。

build flagsについての更に詳しい情報は、'go help build'を参照してください。

パッケージの指定についてのさらに詳しい情報は、'go help packages'を参照してください。


Goプログラムのコンパイルと実行

Usage:

	go run [build flags] [-exec xprog] gofiles... [arguments...]

runは指定された名前のGoソースファイルから構成されるmainパッケージをコンパイルし、実行します。Goソースファイルは".go"のサフィックスで終わるファイルと定義されています。

デフォルトでは、'go run'はコンパイルしたバイナリを直接実行します: 'a.out arguments...'。
もし-execフラグが与えられれば、'go run'はxprogを使ってバイナリを実行します: 'xprog a.out arguments...'。
もし-execフラグが与えられず、GOOSやGOARCHはシステムの初期値によって異なりますが、go_$GOOS_$GOARCH_execというプログラムが現在の検索パスにあれば、'go run'はそのプログラムを使ってバイナリを実行します。
例えば、'go_nacl_386_exec a.out arguments...'のように。
これによってシミュレータやその他の実行方法がある場合にはクロスコンパイルしたバイナリを実行できるようになります。

build flagsについての更に詳しい情報は、'go help build'を参照してください。

See also: go build.


Goパッケージのテスト

Usage:

	go test [-c] [-i] [build and test flags] [packages] [flags for test binary]

'go test'はインポートパスに指定された名前のパッケージのテストを自動で行います。下記のフォーマットでテスト結果の概要をプリントし:

	ok   archive/tar   0.011s
	FAIL archive/zip   0.022s
	ok   compress/gzip 0.033s
	...

テストに失敗した各パッケージの詳細出力が続きます。

'go test'は"*_test.go"ファイルパターンにマッチする名前をもつすべてのファイルと一緒に、各パッケージを再コンパイルします。"_"（"_test.go"含む）や"."から始まる名前のファイルは無視されます。これら追加ファイルはテスト関数、ベンチマーク関数、Example関数を含むことができます。詳細は'go help testfunc'を参照してください。リストアップされたパッケージはそれぞれ個々のテストバイナリとして実行されます。

パッケージ宣言をし、"_test"サフィックスを有するテストファイルは個々のパッケージとしてコンパイルされ、mainのテストバイナリと共にリンクされ実行されます。

デフォルトではgo testには引数は必要ありません。カレントディレクトリのソース（テスト含む）についてパッケージをコンパイル・テストします。

パッケージは一時ディレクトリでビルドされるため、テストを行わないインストール作業に干渉することはありません。

build flagsに加えて、'go test'自身がハンドルするフラグには以下のものがあります。

	-c  テストバイナリをpkg.testにコンパイルしますが、実行はしません。
		（pkgの部分はパッケージのインポートパスの最後の要素となります。）

	-i
		テストの依存関係のパッケージをインストールします。
		テストの実行はしません。

	-exec xprog
		テストバイナリをxprogを使って実行します。
		この振る舞いは'go run'と同じです。詳細は'go help run'を参照してください。

テストバイナリはテスト実行を制御するフラグを受け付けます。これらのフラグは'go test'でも利用できるものです。詳細は'go help testflag'を参照してください。

テストバイナリに対してその他任意のフラグが必要な場合は、パッケージ名の後に指定する必要があります。go toolはマイナス記号から始まる最初の引数をひとつのフラグとして扱います。（マイナス記号自体は認識されません。）その引数とその後に続くすべての引数がテストバイナリの引数として渡されます。

build flagsについてのさらに詳しい情報は、'go help build'を参照してください。パッケージの指定についてのさらに詳しい情報は、'go help packages'を参照してください。

See also: go build, go vet.


Goコマンドで特定のgo toolを実行

Usage:

	go tool [-n] command [args...]

toolは引数に指定したgo toolコマンドを実行します。引数を指定しない場合には、既存のツールの一覧をプリントします。

-nフラグでは実行しようとしているコマンドをプリントしますが、実行はしません。

各toolコマンドについてのさらに詳しい情報は、'go tool command -h'を参照してください。


Goのバージョンのプリント

Usage:

	go version

versionはruntime.Versionで得られるのと同様に、Goのバージョンをプリントします。


Goパッケージへのgo tool vetの実行

Usage:

	go vet [-n] [-x] [packages]

vetはインポートパスに指定されている名前のパッケージに対し、Goのvetコマンドを実行します。

vetについてのさらに詳しい情報は、'godoc code.google.com/p/go.tools/cmd/vet'を参照してください。パッケージの指定についてのさらに詳しい情報は、'go help packages'を参照してください。

特定のオプションを付加してvet toolを実行する場合は、'go tool vet'を実行します。

-nフラグは実行しようとしているコマンドをプリントします。-xフラグは実行中に随時コマンドをプリントします。

See also: go fmt, go fix.


Go、C間の呼び出し

GoとC/C++コード間の呼び出しの方法は2つあります。

1つ目はGoのディストリビューションの一部であるcgo toolです。使用方法については、cgoのドキュメント（godoc cmd/cgo）を参照してください。

2つ目は言語間インタフェースの汎用ツールであるSWIGプログラムです。SWIGについての情報はhttp://swig.org/を参照してください。go buildを実行する際、.swig拡張子を持つ任意のファイルがSWIGに渡されます。-c++オプションにより、.swigcxx拡張子を持つ任意のファイルがSWIGに渡されます。

cgo、SWIGのいずれかを使用する場合、go buildは任意の.c、.m、.s、.SファイルをCコンパイラに渡し、任意の.cpp、.cxxファイルをC++コンパイラに渡します。CCまたはCXX環境変数によりC、C++コンパイラそれぞれに対して使用するコンパイラを指定することができます。

Goコマンドとファイルタイプ

goコマンドはそれぞれのディレクトリにある一部分のファイルの中身を対象とします。
どのファイルが対象になるかは、ファイルの拡張子で決まります。
それらの拡張子は:

	.go
		Goのソースファイル
	.c, .h
		Cのソースファイル。
		もしパッケージがcgoを使用していれば、OS固有のコンパイラ
		(一般にはgcc)でコンパイルされます。 そうでなければ、
		Go固有のコンパイラ、5c、6c、8cなどでコンパイルされます。
	.cc, .cpp, .cxx, .hh, .hpp, .hxx
		C++のソースファイル。cgoもしくはSWIGとのみ使用でき、
		常にOS固有のコンパイルでコンパイルされます。
	.m
		Objective-Cのソースファイル。cgoでのみ使用でき、
		常にOS固有のコンパイルでコンパイルされます。
	.s, .S
		アセンブラソースファイル。
		もしパッケージがcgoを使用していれば、OS固有のアセンブラ
		(一般にはgcc (sic))でコンパイルされます。 そうでなければ、
		Go固有のアセンブラ、5a、6a、8aなどでアセンブルされます。
	.swig, .swigcxx
		SWIGの定義ファイル。
	.syso
		システムオブジェクトファイル。

これらのファイルタイプは.sysoを除きビルド制約を含めることができます。
ただし、goコマンドはビルド制約の検出を、空行でも//-スタイルの行コメントでもない
最初の項目をみつけた時点で停止します。


GOPATH環境変数

Go pathはインポート文を解決するために使用されます。go/buildパッケージに実装、ドキュメントされています。

GOPATH環境変数はGoのコードを探しにいく場所をリストアップします。Unixでは、値はコロンで区切った文字列です。Windowsでは、値はセミコロンで区切った文字列です。Plan 9では、値はリストです。

GOPATHは、標準のGoツリー以外のパッケージをget、build、installするために必ず設定する必要があります。

GOPATHにリストアップされている各ディレクトリは規定された構造を持つ必要があります。

src/ディレクトリはソースコードを保持します。'src'以下のディレクトリがインポートパスや実行ファイル名を決定します。

pkg/ディレクトリはインストール済みのパッケージオブジェクトを保持します。Go treeにしたがって、それぞれの対象オペレーティングシステムとアーキテクチャの組み合わせでpkgのサブディレクトリが構成されます。(pkg/GOOS_GOARCH).

DIRがGOPATHに含まれるディレクトリである場合は、DIR/src/foo/bar内にソースを有するパッケージが"foo/bar"としてインポートされ、"DIR/pkg/GOOS_GOARCH/foo/bar.a"にコンパイルされた形でインストールされます。

bin/ディレクトリはコンパイル済みのコマンドを保持します。各コマンドはそのソースディレクトリの名前となりますが、パス全体ではなく最後の要素のみの名前となります。つまり、DIR/src/foo/quuxにソースがあるコマンドは、/foo/quuxではなく、DIR/bin/quuxにインストールされます。DIR/binをPATHに追加することでインストールしたコマンドを利用できるようにfoo/は取り除かれます。GOBIN環境変数がセットされている場合は、DIR/binの代わりにそのディレクトリにコマンドがインストールされます。

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

Goはソースコードを探す際、GOPATHにリストアップされている各ディレクトリを検索しますが、新しいパッケージは常にリストの最初のディレクトリにダウンロードされます。


Goのインポートパス構文

インポートパス（'go help packages'を参照してください）はローカルファイルシステムに格納されたパッケージを示します。一般的に、インポートパスは標準パッケージ（例えば"unicode/utf8"）かいずれかのワークスペース（'go help gopath'を参照してください）の中から見つかったパッケージのどちらか一方を示します。

Goの相対インポートパス

./や../で始まるインポートパスは相対パスと呼ばれます。toolchainは2つの方法により相対パスをショートカットとしてサポートします。

1つ目に、相対パスはコマンドライン上で簡易表現として用いられます。"unicode"としてインポートされたコードを含むディレクトリに対して"unicode/utf8"のテストを実行したい場合には、フルパスで記述する必要はなく、"go test ./utf8"と指定することができます。逆の場合でも同様に、"go test .."では"unicode/utf8"ディレクトリから"unicode"をテストできます。相対指定のパターンを組み合わせることも可能で、たとえば"go test ./..."はすべてのサブディレクトリをテストします。パターンの文法に関する詳しい情報は、'go help packages'を参照してください。

2つ目に、Goプログラムをワークスペース外でコンパイルする場合、同様にワークスペース外で近くにあるコードを参照するプログラムのインポート文に相対パスを利用できます。これにより通常のワークスペースの外での小規模な複数パッケージのプログラムを用いた実験が容易になります。ただし、これらのプログラムは"go install"でインストールすることはできません。（インストール先のワークスペースがありません）そのため、ビルドするごとにスクラッチから組み直す必要があります。曖昧さ回避のため、Goプログラムでは単一ワークスペース内で複数の相対パスを用いることはできません。

Goのリモートインポートパス

インポートパスによってはリビジョン管理システムを用いたパッケージのソースコードの取得方法を同時に表しているものがあります。

よく知られるコードホスティングサイトには特有の書き方があります。

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

	IBM DevOps Services (Git)

		 import "hub.jazz.net/git/user/project"
		 import "hub.jazz.net/git/user/project/sub/directory"

その他のサーバがホストするコードについては、インポートパスにバージョン管理のタイプの修飾子を指定するか、go toolにてhttps/httpでインポートパスを動的に取得しHTMLの<meta>タグからコードがある場所を探します。

コードの場所を指定する際、以下のインポートパス形式

	repository.vcs/path

は、特定のレポジトリを明示し、指定された名前のバージョン管理システムを利用します。（.vcsサフィックスを含んでも含まなくても問題ありません）次いでレポジトリの内部のパスを指定します。サポートするバージョン管理システムは以下のとおりです。

	Bazaar      .bzr
	Git         .git
	Mercurial   .hg
	Subversion  .svn

たとえば

	import "example.org/user/foo.hg"

は、Mercurialレポジトリのexample.org/user/fooまたはfoo.hgのルートディレクトリを表し、

	import "example.org/repo.git/foo/bar"

は、Gitレポジトリのexample.org/repoまたはrepo.gitのfoo/barディレクトリを表します。

バージョン管理システムが複数のプロトコルをサポートする場合、ダウンロード時にそれぞれが順番に試行されます。例えば、Gitのダウンロードはまずgit://を試行し、次にhttps://、最後にhttp://を試行します。

インポートパスが既知のコードホスティングサイトではなく、さらにバージョン管理の修飾子がない場合、go toolはhttps/httpによりインポートを取得するように試行し、ドキュメントのHTML <head>内の<meta>タグを探します。

metaタグは以下の形式です。

	<meta name="go-import" content="import-prefix vcs repo-root">

import-prefixはレポジトリルートに対応するインポートパスです。"go get"で取得するパッケージのプリフィックスとなっているか完全一致する必要があります。完全一致でない場合は、<meta>タグの一致検証のため、プリフィックスに対してまた別のhttpリクエストが生成されます。

vcsは"git"、"hg"、"svn"などのうちの一つです。

repo-rootはスキームを含み.vcs拡張子を含まないバージョン管理システムのルートです。

たとえば、

	import "example.org/pkg/foo"

の結果は以下のリクエストとなります。

	https://example.org/pkg/foo?go-get=1 (推奨)
	http://example.org/pkg/foo?go-get=1  (代替)

該当のページに以下のmetaタグが含まれている場合、

	<meta name="go-import" content="example.org git https://code.org/r/p/exproj">

go toolはhttps://example.org/?go-get=1が同じmetaタグを含むことを確認し、GOPATH/src/example.orgにgit clone https://code.org/r/p/exprojします。

新規にダウンロードしたパッケージはGOPATH環境変数に挙げられている最初のディレクトリに書き出されます。（'go help gopath'を参照してください）

goコマンドは、利用しているGoリリースの適切なパッケージのバージョンをダウンロードするようにします。詳細については'go help install'を実行して下さい。


Goコマンドでのパッケージリストについて

多くのコマンドが以下のように一連のパッケージに適用できます。

	go action [packages]

通常、[packages]はインポートパスのリストです。

ルートパスまたは'.'、'..'要素で始まるインポートパスはファイルシステムパスと解釈され、そのディレクトリにあるパッケージであることを表します。

それ以外の場合、インポートパス"P"はGOPATH環境変数に挙げられているあるDIRに対してDIR/src/Pディレクトリの中に存在するパッケージを表します。（'go help gopath'を参照してください。）

インポートパスが与えられていない場合、actionはカレントディレクトリのパッケージに対して適用されます。

go toolでビルドされるパッケージに利用できないパスの予約名が以下のとおり3つあります。

- "main"はスタンドアロン実行ファイルのトップレベルパッケージを表します。

- "all"は全GOPATHツリーに存在するすべてのパッケージディレクトリへ展開されます。たとえば、'go list all'はローカルシステムのすべてのパッケージをリストアップします。

- "std"はallに似ていますが、標準Goライブラリのパッケージだけへ展開されます。

インポートパスが1つ以上の"..."ワイルドカード（それぞれが空文字・スラッシュを伴う文字列を含む任意の文字列にマッチする）を含む場合、それらのパターンにマッチする名前を含むGOPATHツリーに存在するすべてのパッケージディレクトリに展開されます。特別な場合としては、x/...はxとxのサブディレクトリにマッチします。たとえば、net/...はnetとそのサブディレクトリのパッケージに展開されます。

インポートパスにはリモートディレクトリからダウンロードするパッケージを指定することもできます。詳細については'go help importpath'を参照してください。

プログラム内のいかなるパッケージも必ず一意なインポートパスをもつ必要があります。規則により、作成者に属する一意なプリフィックスからパスを始めるように取り決められています。たとえば、Googleで内部的に使用されるパスはすべて'google'で始まり、リモートレポジトリを表すパスは'code.google.com/p/project'のようにコードへのパスから始まります。

特別な場合として、パッケージリストが単一のディレクトリにある.goファイルのリストであるときには、それらのファイルのみから構成される擬似的なパッケージにコマンドが適用されます。その際、それらのファイルのビルド制約と、同じディレクトリの他のファイルは無視されます。

"."や"_"から始まるファイル名はgo toolにより無視されます。


Goコマンドによるテストのフラグについて

'go test'コマンドは'go test'自体に適用されるフラグと生成されるテストバイナリに適用されるフラグの両方を扱います。

フラグのうちいくつかはプロファイルを制御し、"go tool pprof"に適合する実行プロファイルを書き出します。更に詳しい情報については"go tool pprof help"を参照してください。pprofの--alloc_space、--alloc_objects、--show_bytesオプションは情報の表示のされ方を制御します。

以下のフラグは'go test'コマンドで利用でき、テスト実行の制御を行います。

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
		runtime.SetBlockProfileRateをnを指定して呼び出すことにより、goroutine
		のブロックプロファイルに供給する詳細度を制御します。
		'godoc runtime SetBlockProfileRate'を参照してください。
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
		ただし、-raceが有効になっている場合は"atomic"です。
		値：
		set：bool：このステートメントが実行されるか？
		count：int：このステートメントが何回実行されるか？
		atomic：int：countと同等ですが、マルチスレッドのテストを正確に
			実行できます。ただし、著しくコストが掛かります。
		-coverを自動的に設定します。

	-coverpkg pkg1,pkg2,pkg3
		与えられたパッケージリストに対し、各テストでカバレッジ解析を
		適用します。デフォルトではテストを行うパッケージのみに対して
		各テストでの解析を行います。
		パッケージはインポートパスとして指定されているものです。
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
		'godoc runtime MemProfileRate'を参照してください。
		すべてのメモリ割り当てをプロファイルするためには、
		-test.memprofilerate=1を設定し、pprofツールに
		--alloc_spaceフラグを与えてください。

	-outputdir directory
		指定したディレクトリにプロファイルからの出力ファイルを置きます。
		デフォルトは"go test"を実行しているディレクトリです。

	-parallel n
		t.Parallelを呼び出すテスト関数の並行実行を許可します。
		このフラグの値は同時に実行するテストの最大数となります。
		デフォルトでは、GOMAXPROCSの値に設定されます。

	-run regexp
		正規表現にマッチするテストやExample関数のみ実行します。

	-short
		長時間実行しているテストに対して実行時間を短縮するように要求します。
		デフォルトではオフになっていますが、all.bashを実行する際には
		Go treeによりサニティチェックが実行されるため、全数テストを
		実行するのに時間を費やさないようにオンに設定されます。

	-timeout t
		テストがtよりも長く実行している場合、panicを呼び出します。

	-v
		冗長な出力：実行するすべてのテストのログを出力します。また、
		テストが成功した場合にもLog、Logfの呼び出しからのすべての
		テキストをプリントします。

pkg.test（pkgはパッケージソースが含まれるディレクトリ名を表します）という名前のテストバイナリが、'go test -c'でビルドされた後直接呼び出されます。テストバイナリディレクトリを呼び出す際、各標準フラグ名は-test.run=TestMyFuncや-test.vのように、'test.'をプリフィックスとして付加する必要があります。

'go test'を実行する際、上記の一覧に記述されていないフラグは変更なしとして扱われます。たとえば、コマンド

	go test -x -v -cpuprofile=prof.out -dir=testdata -update

はテストバイナリをコンパイルし、

	pkg.test -test.v -test.cpuprofile=prof.out -dir=testdata -update

として実行します。

カバレッジ以外のプロファイルを生成するテストフラグについても、プロファイルを解析する際に利用できるようにpkg.testにテストバイナリを残します。

'go test'で認識されないフラグはすべてのパッケージ指定の後に位置している必要があります。


Goコマンドによるテストのテスト関数について

'go test'はテスト中のパッケージに付随する"*_test.go"ファイルから、テスト、ベンチマーク、Example関数を検出しようとします。

テスト関数はTestXXX（XXXは小文字以外から始まる任意の英数文字列）という名称で、下記のシグネチャを有する必要があります。

	func TestXXX(t *testing.T) { ... }

ベンチマーク関数はBenchmarkXXXという名称で、以下のシグネチャを有する必要があります。

	func BenchmarkXXX(b *testing.B) { ... }

Example関数はテスト関数に似ていますが、成功・失敗を伝えるために*testing.Tを用いる代わりに、出力をos.Stdoutにプリントします。その出力は"Output:"コメントと比較されます。これは関数本体の最後のコメントである必要があります（下記の例を参照してください。）そのようなコメントがない、もしくは"Output:"の後が空白のExample関数はコンパイルされますが実行はされません。

godocは関数、定数、変数XXXの利用方法の説明のためにExampleXXXの内容を表示します。T型や*T型をレシーバに持つメソッドMの用例はExampleT_Mと名付けられます。_xxx（xxxは大文字以外で始まるサフィックス）を後ろに付けて識別できるようにすることで、特定の関数、定数、変数について複数のExample関数を定義することもできます。

以下はExample関数の例です：

	func ExamplePrintln() {
		Println("The output of\nthis example.")
		// Output: The output of
		// this example.
	}

テストファイルが単一のExample関数を有する場合（その他の関数や型や変数や定数定義が少なくとも1つあり、テストやベンチマーク関数がない場合）は、テストファイルの全体が用例として扱われます。

さらに詳しい情報についてはtestingパッケージのドキュメントを参照してください。


本ドキュメントは以下のドキュメントを翻訳しています:https://code.google.com/p/go/source/browse/src/cmd/go/doc.go?r=528863ac401ff24391b01bbaf88080e5b85c0534
*/
package main
