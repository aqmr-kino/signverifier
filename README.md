# signverifier
OpenPGP形式でクリアテキスト署名されたメッセージの署名を検証し、正当な署名であることを確認できた場合のみファイル内容を出力するフィルタプログラムです。

メッセージは標準入力から読み取り、署名が正しい場合のみ、内容を標準出力へ出力します。  
何らかのエラー発生時は、標準エラー出力にエラーメッセージを出力し、標準出力には何も出力しません。

## ビルド方法
```sh
go build signverify.go
```

## 使い方
### Usage
```sh
./signverify publickey.pub
```

* publickey.pub : 署名検証に用いる公開鍵(PEM形式のGPG公開鍵)

### Sample
#### 署名済みテキストファイルの検証・内容出力
```
cat signedmessage.txt.asc | ./signverify publickey.pub
```

#### リモートサーバ上の署名済みスクリプトの実行
```
curl -s https://example.com/signedscript.sh.asc | ./signverify publickey.pub | bash
```
