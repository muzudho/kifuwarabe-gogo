# kifuwarabe-gogo

コンピューター囲碁☆（＾～＾）  
[https://github.com/bleu48/GoGo](https://github.com/bleu48/GoGo) を元に始めるが、大改造する予定☆（＾～＾）  

# Example - V09

## Run

```shell
# 実行ファイルを作成するために、以下のコマンドを打鍵してください。
go build
# kifuwarabe-gogo.exe ファイルが作成されました。

# コンピューター囲碁エンジンを実行するために、以下の実行ファイルのステム（拡張子を省いた名前）を打鍵してください。
kifuwarabe-gogo V09
# しばらく待つと自動で開始します
```

# Example - V09a

## Run

```shell
# 実行ファイルを作成するために、以下のコマンドを打鍵してください。
go build
# kifuwarabe-gogo.exe ファイルが作成されました。

# コンピューター囲碁エンジンを実行するために、以下の実行ファイルのステム（拡張子を省いた名前）を打鍵してください。
kifuwarabe-gogo V09a
```

Input: (Example)  

```shell
genmove b
# 少し時間がかかります

genmove w
# 少し時間がかかります

quit
```

## Memorandom

`go build -ldflags "-s -w"` でデバッグ情報を外せる？  
高速化した気はしないが……。  
