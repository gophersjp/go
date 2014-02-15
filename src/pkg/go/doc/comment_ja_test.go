package doc

import (
	"bytes"
	"testing"
)

// Test case that has valid custom language setting.
func TestToHTMLWithValidSetting(t *testing.T) {
	input := `{"language": "ja"}
これは導入部です。

これはヘッダです

これはヘッダではありません。

コレも、ヘッダではないです

ヘッダです

此れもヘッダでないです。

	インデントです。

此れはヘッダで須

これはヘッダでない！

Header

This is not a header.

This is also not a header!

This is also not a header！

Not, a header

not header`

	expected := `<p>
これは導入部です。
</p>
<h3 id="hdr-________">これはヘッダです</h3>
<p>
これはヘッダではありません。
</p>
<p>
コレも、ヘッダではないです
</p>
<h3 id="hdr-_____">ヘッダです</h3>
<p>
此れもヘッダでないです。
</p>
<pre>インデントです。
</pre>
<h3 id="hdr-________">此れはヘッダで須</h3>
<p>
これはヘッダでない！
</p>
<h3 id="hdr-Header">Header</h3>
<p>
This is not a header.
</p>
<p>
This is also not a header!
</p>
<p>
This is also not a header！
</p>
<p>
Not, a header
</p>
<p>
not header</p>
`

	executeTest(t, input, expected)
}

// Test case that has no custom language setting.
// Only alphabetical script check should be done.
func TestToHTMLWithNoSetting(t *testing.T) {
	input := `
これは導入部です。

これはヘッダです

これはヘッダではありません。

コレも、ヘッダではないです

ヘッダです

此れもヘッダでないです。

	インデントです。

此れはヘッダで須

これはヘッダでない！

Header

This is not a header.

This is also not a header!

This is also not a header！

Not, a header

not header`

	expected := `<p>
これは導入部です。
</p>
<p>
これはヘッダです
</p>
<p>
これはヘッダではありません。
</p>
<p>
コレも、ヘッダではないです
</p>
<p>
ヘッダです
</p>
<p>
此れもヘッダでないです。
</p>
<pre>インデントです。
</pre>
<p>
此れはヘッダで須
</p>
<p>
これはヘッダでない！
</p>
<h3 id="hdr-Header">Header</h3>
<p>
This is not a header.
</p>
<p>
This is also not a header!
</p>
<p>
This is also not a header！
</p>
<p>
Not, a header
</p>
<p>
not header</p>
`

	executeTest(t, input, expected)
}

// Test case that has custom language setting whose language is not supported.
// Only alphabetical script check should be done.
func TestToHTMLWithUnexistingLanguage(t *testing.T) {
	input := `{"language": "not_existing"}
これは導入部です。

これはヘッダです

これはヘッダではありません。

コレも、ヘッダではないです

ヘッダです

此れもヘッダでないです。

	インデントです。

此れはヘッダで須

これはヘッダでない！

Header

This is not a header.

This is also not a header!

This is also not a header！

Not, a header

not header`

	expected := `<p>
これは導入部です。
</p>
<p>
これはヘッダです
</p>
<p>
これはヘッダではありません。
</p>
<p>
コレも、ヘッダではないです
</p>
<p>
ヘッダです
</p>
<p>
此れもヘッダでないです。
</p>
<pre>インデントです。
</pre>
<p>
此れはヘッダで須
</p>
<p>
これはヘッダでない！
</p>
<h3 id="hdr-Header">Header</h3>
<p>
This is not a header.
</p>
<p>
This is also not a header!
</p>
<p>
This is also not a header！
</p>
<p>
Not, a header
</p>
<p>
not header</p>
`

	executeTest(t, input, expected)
}

// Test case that has invalid tag other than "language".
// Only alphabetical script check should be done.
func TestToHTMLWithInvalidTag(t *testing.T) {
	input := `{"tongue": "ja"}
これは導入部です。

これはヘッダです

これはヘッダではありません。

コレも、ヘッダではないです

ヘッダです

此れもヘッダでないです。

	インデントです。

此れはヘッダで須

これはヘッダでない！

Header

This is not a header.

This is also not a header!

This is also not a header！

Not, a header

not header`

	expected := `<p>
{&#34;tongue&#34;: &#34;ja&#34;}
これは導入部です。
</p>
<p>
これはヘッダです
</p>
<p>
これはヘッダではありません。
</p>
<p>
コレも、ヘッダではないです
</p>
<p>
ヘッダです
</p>
<p>
此れもヘッダでないです。
</p>
<pre>インデントです。
</pre>
<p>
此れはヘッダで須
</p>
<p>
これはヘッダでない！
</p>
<h3 id="hdr-Header">Header</h3>
<p>
This is not a header.
</p>
<p>
This is also not a header!
</p>
<p>
This is also not a header！
</p>
<p>
Not, a header
</p>
<p>
not header</p>
`

	executeTest(t, input, expected)
}

func executeTest(t *testing.T, input string, expected string) {
	var actual bytes.Buffer
	ToHTML(&actual, input, nil)
	if actual.String() != expected {
		t.Errorf("expected: %v\nbut was: %v", expected, actual.String())
	}
}
