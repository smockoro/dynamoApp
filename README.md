<!-- vim: set fileencoding=utf-8 : -->

# Message Application by DynamoDB
DynamoDBを使ったサンプルアプリです。
詳細は各ディレクトリの中のREADMEにも記載されています。
TODO: まだ書いてない…

## `app`ディレクトリ
DynamoDBへ`go-aws-sdk`を利用して下記をすることができるサンプルです。
1. テーブル作成
2. データ入力
3. テーブル削除

## `lambda`ディレクトリ
AWS lambdaで利用できる関数のサンプルです。
- recipient_select
　DynamoDBのメッセージ一覧から、Recipientを指定して取り出すことができます。
- put_message
 メッセージをDynamoDBへ入れます。

TODO: AWS API Gatewayへの対応ができていない。lambda上でサンプルがテスト動作することは確認済み
