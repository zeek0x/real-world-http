# 証明書の作成手順

## 雛形コピー

```console
$ sudo cp /etc/ssl/openssl.cnf .
```

## デフォルトセクション削除 （記載なし手順）

デフォルトセクションとは、`openssl.cnf` の先頭にある `[ section_name ]` 配下に属していないフィールド群のことである.

下記のフィールド群を削除する.

- `HOME`
- `oid_section`
- `openssl_conf`

`oid_section` で定義される `new_oids` セクションも削除する

## ルート認証局証明書作成

```console
# 秘密鍵を作成
$ openssl genrsa -out ca.key 2048

# 証明書要求
$ openssl req -new -sha256 -key ca.key -out ca.csr -config openssl.cnf

# 証明書を自分の鍵で署名して作成
$ openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile ./openssl.cnf
```

## ファイルの確認

```console
# 秘密鍵の確認
$ openssl rsa -in ca.key -text

# 証明書署名要求(CSR)の確認
$ openssl req -in ca.csr -text

# 証明書の確認
$ openssl x509 -in ca.crt -text
```

## サーバーの証明書作成

```console
# RSA 2049ビットの秘密鍵を作成
$ openssl genrsa -out server.key 2048

# 証明書署名要求(CSR)の作成
$ openssl req -new -nodes -sha256 -key server.key -out server.csr -config openssl.cnf
# Common Name で localhost を入力

# 証明書を自分の秘密鍵で署名して作成
$ openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Server
```
