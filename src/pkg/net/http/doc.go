// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
httpパッケージはHTTPクライアントとサーバの実装を提供します。

Get、Head、Post、PostFormはHTTP(またはHTTPS)リクエストを作成します。

	resp, err := http.Get("http://example.com/")
	...
	resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	...
	resp, err := http.PostForm("http://example.com/form",
		url.Values{"key": {"Value"}, "id": {"123"}})

クライアントは終了時にレスポンスボディをクローズしなければなりません。

	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// ...

HTTPクライアントヘッダ、リダイレクトポリシーなどの設定を制御するには
Clientを作成してください。

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Get("http://example.com")
	// ...

	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	// ...

プロキシ、TLS設定、キープアライブ、圧縮などの設定を制御するには
Transportを作成してください。

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")

ClientとTransportは複数のゴルーチンからの並行アクセスに対して安全であり、
効率のために一度だけ作成されて再利用されるべきです。

ListenAndServeは与えられたアドレスとハンドラを使ってHTTPサーバを開始します。
ハンドラはたいていはnilであり、その場合はDefaultServeMuxが使われます。
HandleとHandleFuncはDefaultServeMuxにハンドラを追加します。

	http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

独自にServerを作成することで、サーバの振る舞いをさらに制御することが可能です。

	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/src/pkg/net/http/doc.go?r=09a1cf0d94b2b6cf672f8dcd04c6962e0916bc4e
*/
package http
