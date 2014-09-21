// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
testingパッケージはGoのパッケージの自動テストを支援します。
これは``go test''コマンドとともに使用することで、次のような形式の関数を自動実行します。
     func TestXxx(*testing.T)
Xxxはどんな英数字の文字列でもよく(ただし、最初の文字は[a-z]以外でなければなりません)、テストルーチンを識別するために用いられれます。

これらの関数内では、ErrorやFailといったメソッドを使用することで、失敗を報告することができます。

新しいテストスイートを書くには、名前が_test.goで終わり、ここで説明したTestXxx関数を含むファイルを作成してください。
ファイルは、テストされるものと同一のパッケージに配置してください。
このファイルは通常のパッケージビルドからは除外されますが、``go test''コマンドで実行されるときには含まれます。
詳細については、``go help test''や``go help testflag''を実行してみてください。

テストとベンチマークは、実行したくない状況では*Tや*BのSkipメソッドを呼び出すことでスキップできます:

    func TestTimeConsuming(t *testing.T) {
        if testing.Short() {
            t.Skip("skipping test in short mode.")
        }
     ...
    }

Benchmarks

次のような形式の関数
    func BenchmarkXxx(*testing.B)
は、ベンチマークとみなされ、"go test"コマンドに-benchフラグが与えられたときに実行されます。
ベンチマークは順次実行されます。

テストフラグの詳細については、こちらを参照してください:
http://golang-jp.org/pkg/cmd/go/#hdr-Go__________________

サンプルのベンチマーク関数は次のようになります:

    func BenchmarkHello(b *testing.B) {
        for i := 0; i < b.N; i++ {
            fmt.Sprintf("hello")
        }
    }

ベンチマーク関数は対象のコードをb.N回実行しなければなりません。
ベンチマークパッケージは、ベンチマーク関数が十分な時間継続するようになるまでb.Nを変化させます。
出力の
    BenchmarkHello    10000000    282 ns/op
は、ループが10000000回、一回あたり282ナノ秒で実行されたことを意味します。

もしベンチマークで処理がはじまる前に時間のかかる初期化が必要なら、タイマーをリセットすることもできます:

    func BenchmarkBigLen(b *testing.B) {
        big := NewBig()
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            big.Len()
        }
    }

もしベンチマークで並列処理の性能をテストする必要があれば、RunParallelヘルパー関数を使うこともできます;
このようなベンチマークはgo test -cpuフラグと一緒に使用できます:

    func BenchmarkTemplateParallel(b *testing.B) {
        templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
        b.RunParallel(func(pb *testing.PB) {
            var buf bytes.Buffer
            for pb.Next() {
                buf.Reset()
                templ.Execute(&buf, "World")
            }
        })
    }

Examples

testingパッケージは見本のコードも実行し、検証します。
見本関数は結果について"Output:"で始まる一行コメントを含むことができ、テストが実行された際の関数の標準出力と比較されます。
(比較では前後の空白は無視されます。) これらの見本関数の見本もあります:

    func ExampleHello() {
            fmt.Println("hello")
            // Output: hello
    }

    func ExampleSalutations() {
            fmt.Println("hello, and")
            fmt.Println("goodbye")
            // Output:
            // hello, and
            // goodbye
    }

見本関数でOutputのコメントがないものは、コンパイルはされますが実行はされません。

見本関数には命名規則があり、パッケージの見本として関数F、型T、型TのメソッドMについて書くには:

    func Example() { ... }
    func ExampleF() { ... }
    func ExampleT() { ... }
    func ExampleT_M() { ... }

複数の見本関数を一つのパッケージ/型/関数/メソッドについて書くには、個別のサフィックスを名前に追加してください。
サフィックスは、小文字で始まる必要があります。

    func Example_suffix() { ... }
    func ExampleF_suffix() { ... }
    func ExampleT_suffix() { ... }
    func ExampleT_M_suffix() { ... }

一つの見本関数だけを含む場合、ファイル全体が見本として提示されます。
これは他の関数や型、変数、定数の宣言を少なくとも一つ含み、テスト関数やベンチマーク関数を含まない必要があります。

Main

時には、テストの前後で追加の初期化や終了処理がテストプログラムで必要になることもあります。
また、メインスレッドで実行されるように制御する必要があるテストもあります。
これらをサポートするために、もしテストファイルが次の関数を含んでいた場合には:

     func TestMain(m *testing.M)

生成されたテストはテストを直接実行する代わりに、TestMain(m)を呼び出します。
TestMainはメインのゴルーチンで実行されるため、m.Runの呼び出しの前後で必要な初期化や終了処理を挟むことができます。
ここではm.Runの結果をもとにos.Exitを呼び出すべきです。

TestMainの最小実装は:

	func TestMain(m *testing.M) { os.Exit(m.Run()) }

実際のところ、この実装はTestMainが明示的に定義されなかった場合に使用されます。

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/testing/testing.go?r=804cc55d6b47e49e4bf46baa4ca9ec44b9684bed
*/
package testing
