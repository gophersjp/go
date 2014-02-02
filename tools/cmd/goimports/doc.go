/*

goimportsコマンドはGoのimport節にある不足しているimportを追加したり、参照されていないimport文を削除します。
インストールは:

     $ go get code.google.com/p/go.tools/cmd/goimports

このコマンドはgofmtコマンドのフォークですので、
自分のエディタのgofmt-on-saveフックの代用として利用できます。

emacsでは、(Go 1.2)の go-mode.elを利用します:
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

本ドキュメントは以下のドキュメントを翻訳しています: https://code.google.com/p/go/source/browse/cmd/goimports/doc.go?repo=tools&r=7dd9cfdeec4318ebfb8db987fe898645fa024f5d

*/
package main

import "code.google.com/p/go.tools/cmd/goimports"
