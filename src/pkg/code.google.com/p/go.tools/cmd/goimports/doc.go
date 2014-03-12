/*

goimportsコマンドは、Goのimport節で不足しているimportを自動で追加したり、参照されていないimport文を削除するgofmt互換のツールです。
インストールは:

     $ go get code.google.com/p/go.tools/cmd/goimports

です。

goimportsは、自分のエディタのgofmt-on-saveフックの置き換えができます。
goimportsは、gofmtのコマンドラインインタフェースと同じであり、同じ方法で
コードのフォーマットができるようになっています。

emacsでは、(Go 1.2+)の go-mode.elを利用します:
   https://go.googlecode.com/hg/misc/emacs/go-mode.el
.emacsファイルに以下を追加しましょう:
   (setq gofmt-command "goimports")
   (add-to-list 'load-path "/home/you/goroot/misc/emacs/")
   (require 'go-mode-load)
   (add-hook 'before-save-hook 'gofmt-before-save)

For vim, set "gofmt_command" to "goimports":
    https://code.google.com/p/go/source/detail?r=39c724dd7f252
    https://code.google.com/p/go/source/browse#hg%2Fmisc%2Fvim
    etc

For GoSublime, follow the steps described here:
    http://michaelwhatcott.com/gosublime-goimports/

For other editors, you probably know what to do.

Happy hacking!

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/cmd/goimports/doc.go?repo=tools&r=8caf575e6beacc2237f7d78f8bd349c28f1b6b3e

*/
package main
