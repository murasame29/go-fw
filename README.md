# go-fw
いろんなフレームワークでHelloWorld

以下のフラグを付けるとサーバを変更できます
```bash
# 起動方法
go run main.go -fw=[使いたいフレームワークのflag]
# 例: ginを使う
go run main.go -fw=gin
```

## フレームワーク一覧
| フレームワーク | flag | リポジトリリンク | 
| ---- | ---- | ---- |
| net/http | net| https://cs.opensource.google/go/go|
| gin | gin | https://github.com/gin-gonic/gin |
| echo | echo | https://github.com/labstack/echo |
| gorilla | gorilla | https://github.com/gorilla/mux |