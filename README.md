# PhotoN_Server

# 概要

go言語で作成したPhotonプロジェクトのサーバサイド
go-simple-archを踏襲して作成

# 実行方法

## Docker コンテナ起動

dockerフォルダに移動後下記コマンドで、
アプリケーション環境が作成され、実行される。  
下記にアクセスするとphpMyAdminに移動する。  
http://localhost:8080

```bash
# dockerフォルダに移動(カレントにいる場合)
cd docker

# docker初回起動
docker-compose up -d --build

# dockerコンテナを起動
docker-compose up -d

# dockerコンテナを停止
docker-compose stop

# docker関連のファイルをいじった場合、下記のコマンド
docker-compose down
docker-compose up -d --build
```

APIのテスト
``` bash
curl -i -X GET 'http://localhost:3000/api/v1/users'
```

## 実行ファイルの作成・実行

プロジェクトフォルダの直下(go.modと同じ階層)へ移動する。

```bash

# ビルド
go build

# 実行
./Photon_Server

```

# 使用パッケージ

パッケージは、公式以外は基本使用しない方針にする。  
初心者だとgo言語の仕様なのか、パッケージの機能なのか判断難しいからね。  
公式以外で採用したパッケージは以下の通り。

## github.com/gin-gonic/gin  

HTTPメソッドの判断とルーティング処理を公式のみで書くと、不必要に冗長になるので採用。  
似たようなパッケージでechoやgolliraもあるのだが、githubのスター数が一番多いのでginを採用した。

## github.com/gin-contrib/cors

CORS設定で必要。net/httpでやろうとすると、ginを使用する場合記述がめんどくさいので採用。  
CORS設定できないとローカルで動作確認ができない。
[CORSについて](https://developer.mozilla.org/ja/docs/Web/HTTP/CORS)

## github.com/go-sql-driver/mysql  

公式のデータベースパッケージを、mysqlで使用するために拡張してくれる。

## github.com/stretchr/testify

UnitTest用のパッケージ。

## gopkg.in/go-playground/assert.v1

UnitTestの結果判定を行うために採用。
