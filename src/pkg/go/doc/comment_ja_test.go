package doc

import (
	"bytes"
	"testing"
)

func TestToHTML(t *testing.T) {
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

	var actual bytes.Buffer
	ToHTML(&actual, input, nil)
	if actual.String() != expected {
		t.Errorf("expected: %v\nbut was: %v", expected, actual.String())
	}
}
