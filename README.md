# drc

- Created: 2023/08/24
- Discord API v9,v10

## Description

- Discord Botに対してスラッシュコマンドを登録する
- スラッシュコマンドからAWS Lambdaを実行

## 手順

1. [Bot 作成](#bot-作成)
2. `make release`で作成したdrcを実行
3. [Lambda設定](https://github.com/jsbry/dal)
4. [INTERACTIONS ENDPOINT URL設定](#interactions-endpoint-url)

### Bot 作成

- <https://discord.com/developers/applications>
- New Application
- SETTINGS > Bot
    - PUBLIC BOTをOFF
    - REQUIRES OAUTH2 CODE GRANTをOFF
- SETTINGS > OAuth2 > URL Generator
    - SCOPESからbot, applications.commandsを選択
    - BOT PERMISSIONSからSend Messagesを選択
- GENERATED URLにアクセスしてBot追加

### INTERACTIONS ENDPOINT URL

- <https://discord.com/developers/applications>
- My ApplicationsからBotを選択
- SETTINGS > General Information
    - INTERACTIONS ENDPOINT URLにLambdaのURLを登録
