# CloudRunの動作確認用リポジトリ
[![CircleCI](https://circleci.com/gh/yusuke0701/cloud-run-test-project.svg?style=svg)](https://circleci.com/gh/yusuke0701/cloud-run-test-project)

### 概要
gRCP-gatewayをCloudRunで動かすサンプル。
CircleCIでCICDしている。
CloudRunで動かすにあたり、ContainerRegistryにdockerイメージを置いている。

### 主な構成
* サーバー
    * CloudRun/Go(1.13)
* ライブラリ
    * grpc-gateway: https://github.com/grpc-ecosystem/grpc-gateway

### 開発環境の構築
1. CloudSDKのインストール
    1. 公式URL: https://cloud.google.com/sdk/
1. golangのインストール
    1. 参考URL: https://golang.org/doc/install
1. protobufのインストール
    1. 公式URL: https://developers.google.com/protocol-buffers/

### 動作確認用URL一覧
* report_API: [POST] https://gateway-5dnafyaz7q-ue.a.run.app/v1/report
    * bodyは空で良い
* bake_API: [POST] https://gateway-5dnafyaz7q-ue.a.run.app/v1/bake
    * TODO: bodyのサンプル

### サーバーのURL一覧

* gRPC_Server: https://server-5dnafyaz7q-ue.a.run.app
* gRPC_Gateway: https://gateway-5dnafyaz7q-ue.a.run.app

### 注意点
* 新規サービスを立ち上げると、認証の状態をデフォルトになる。\
    CRunのサービスのコンソール画面で、権限にallUsersを追加することで、一般公開可能。
* 普通に呼び出しても動かないので、動作確認用のエンドポイントを使用すること

### 参考にしたURL一覧
1. Ubuntu に最新版のprotobufをインストール: https://qiita.com/Sylba2050/items/8ee3228fae50d35ce38d
1. gateway側の認証: https://github.com/petomalina/cloudrun-grpc/blob/master/cmd/client/main.go