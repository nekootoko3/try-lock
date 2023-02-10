# try-lock

実装方法によって思いがけずレコードが重複して作成してしまったり、外部 API が複数回呼び出してしまうことがあります。
そういった事象を PostgreSQL の行の排他ロックや UNIQUE 制約を利用して防止するための方法について試すためのリポジトリです。

## 仮定の状況

自社 DB で請求 ( bills ) を管理しており、実際の決済は外部 API を呼び出すことによって完了します。
1 つの請求に対して 1 度だけ決済が行われることを保証するための方法を検討します。

## 事前準備

- docker
  - それぞれ試す前に `docker compose up` を実行しておいてください
- go

## 試してみる

- [行の排他ロックを試してみる](./pessimistic/README.md)
- [UNIQUE 制約を試してみる](./uniq/README.md)


## 参照

- https://www.postgresql.jp/document/13/html/explicit-locking.html
