## 個人開発練習用リポジトリ
１．概　要  
CtoCサイトのひな形を作成  
RESTfulを意識
CSS-in-JS非採用  
直交性と疎結合、静的型付けとDockerを重視  

２．技術スタック  
フロントエンド：TypeScript、React.js、Sass、Html  
バックエンド：Go(Gin,sqlc)、PostgreSQL  
開発環境：Docker、VSCode、Git&Github、A5:SQL Mk-2、Windows10  

３．起動手順
```
# ディレクトリk_904set_devにあるdocker-compose.ymlでコンテナを起動

# APIコンテナへ入系
docker exec -it k_904set_dev-214api_dev-1 bash

# 同じ階層にmain.goが存在することを確認する

# 以下のコマンド群をまとめてシェルで流す
go mod init github.com/tatsuhidehirakawa/myportfolio_k && go get github.com/kyleconroy/sqlc/cmd/sqlc && go install github.com/kyleconroy/sqlc/cmd/sqlc && sqlc version && sqlc generate --file sqlc/sqlc.yaml && cd sqlc/build_sqlc && go mod init sqlc_pkg && cd ../.. && go get sqlc_pkg | echo 'replace sqlc_pkg => ./sqlc/build_sqlc' >> go.mod && go mod tidy && go run main.go
```
４．API定義  
☞[エンドポイント](/k_214api_dev/README.md)、[データ形式](/k_214api_dev/README.md)  
（リンク先参照）  

５．サーバ構成図(全体)  
```
REST(ULCODC$SS)_Architecture
+-----------+  +------------------------------------------------+
|           |  |           +-----------+          +-----------+ |
|           |  |           |   *214    |          |   *314    | |
|           |  |           | APIServer |          |  DBServer | |
|           |  |           |  Golang   |          | PostgreSQL| |
|           |  |           |(Gin/sqlc) |          |           | |
|           |  | +--JSON-->|#443  #5432|<---SQL-->|#5432      | |
|  CLIENT   |  | |         +-----------+          +-----------+ |
| (Browser) |  | |         +-----------+                        |
|        #80|<---+         |   *114    |                        |
|           |  | |         | WebServer |                        |
|           |  | |         |React,Sass |                        |
|           |  | |  HTML   |           |                        |
|           |  | +-CSS,JS--|#8080      | act: Ubuntu 22.04 LTS  |
|           |  |           +-----------+      Amazon Linux 2022 |
+-----------+  +------------------------------------------------+
```
６．ディレクトリ構成図(全体／抜粋)  
```
myportfolio_k
  ├─k_214api_dev                   # 214番APIServer
  ├─k_314dbs_dev  [データ形式](/k_214api_dev/README.md)                 # 314番DBServer
  └─k_904set_dev                   # 管理系ファイル群
     └─docker-compose.yml          # 起動用ymlファイル
```

<!--
ディレクトリ命名規則  

|桁番号|     意 味      | 記 号 |     例     | レベル感 | 備 考 |
| :---: |:---      | :---: | :---      | :---    | :---  |
|桁1|プロジェクト記号  |   k   |個人開発練習|         |       |
|  桁2  |  サーバ番号     | _101  |101サーバ   |         |       |
|&nbsp;&nbsp;&nbsp;&nbsp;〃&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;〃&nbsp;&nbsp;&nbsp;&nbsp;| _201  |201サーバ   |         |       |
|  〃   |&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;〃&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;| _202  |202サーバ   |         |       |
|  〃   |      〃        | _301  |301サーバ   |&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;||
|  桁3  |主たる技術スタック| _re  |React.js    |作り込み中|※CodeSandbox<br>(https://codesandbox.io/s/weathered-violet-86tgok)|
|  〃   |      〃        | _go  |Golang      |将来検討予定||
|  〃   |      〃        | _re  |Rails       |錬成中||
|  〃   |      〃        | _pg  |PostgreSQL  |とりあえずサーバは立てた||
|  桁4  | サーバ種別      | _wbs  |Webサーバ   |||
|  〃   |      〃        | _aps  |APサーバ    |||
|  〃   |      〃        | _dbs  |dbサーバ    |||
|  桁5  |生成種別        |   _g   |「git clone」から生成|||
|  〃   |      〃        |   _d  |「docker compose run」から生成|||
|  〃   |      〃        |   _c  |CodeSandboxからインポート|||
-->



