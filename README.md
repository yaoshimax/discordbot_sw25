## SW2.5.go

### 概要

[discordgo](https://github.com/bwmarrin/discordgo)をつかった
SW2.5のキャラ作成ダイスを行ってくれるDiscord用 bot

### 使い方
#### 準備

- 前提として、discordの開発者用ページ (https://discordapp.com/developers/applications) からアプリケーションを作成し、botを追加することで、bot用の認証トークンを取得してください。
  - Botを一般公開しない場合、設定画面のAuthorization Flowで Public Bot のフラグはOFFにすると良いです。

- 認証トークンは以下の`Click to Reveal Token` をクリックすると表示されます。

![tokenl](./img/token.png)


- sw35.goの`{{REPLACE_AUTHENTICATION_TOKEN}}` を認証トークンの文字列に置換してください

#### ビルド
 
- ビルドは以下により実行できます

```console
$ go build sw25.go
```

#### botをサーバに登録する方法
- discordの開発者用シーンからアプリケーション管理画面の `Oauth2` を選択し、URIをコピーします
![get_uri](./img/add_application.png)
- コピーしたURIをブラウザでひらくと、「サーバに追加」以下のセレクトボックスからサーバを選択できるので、botを追加したいサーバを選択します
![select_server](./img/select_server.png)

以上で指定のサーバにbotのユーザを追加できます

#### botの起動と利用の仕方

ビルドされたバイナリファイルを実行するだけでbotが起動します。以下はWindowsの例です。

```console
$ sw25.exe
```
この場合、botの終了はCtrl+Cで行えます。


Discordからこのbotを利用するときは、botユーザがサーバでオンライン表記になっている事を確認の上で
`/人間` 等のように、SW2.5の種族名をチャットに打ち込んでください。

現在はルールブック3にまで掲載されている種族名が利用可能のはずです。

![bot_example](./img/bot_result.png)
