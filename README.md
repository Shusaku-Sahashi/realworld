## TODO

https://github.com/gothinkster/realworld/tree/master/spec

### Backend

https://github.com/gothinkster/realworld/tree/master/api

- [ ] Authenticate users via JWT (login/signup pages + logout button on settings page)
- [ ] CRU\* users (sign up & settings page - no deleting required)
- [ ] CRUD Articles
- [ ] CR\*D Comments on articles (no updating required)
- [ ] GET and display paginated lists of articles
- [ ] Favorite articles
- [ ] Follow other users

## gin での ambiguous エラーに関して

gin 内部で使用している`ugorji/go`のモジュール構成が変更された影響で下記の様なエラーが発生する。
2 つの依存県警を見ていることとなり、import が一位に定まらないのが原因らしい。

```
build github.com/app/realworld: cannot load github.com/ugorji/go/codec: ambiguous import: found github.com/ugorji/go/codec in multiple modules:
        github.com/ugorji/go v1.1.4 (/Users/shusaku/go/pkg/mod/github.com/ugorji/go@v1.1.4/codec)
        github.com/ugorji/go/codec v0.0.0-20190320090025-2dc34c0b8780 (/Users/shusaku/go/pkg/mod/github.com/ugorji/go/codec@v0.0.0-20190320090025-2dc34c0b8780)
```

これを回避するために、`go get modulename@none`で回避可能。
https://github.com/gin-gonic/gin/issues/1673
