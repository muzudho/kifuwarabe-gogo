# kifuwarabe-gogo

コンピューター囲碁☆（＾～＾）  
[https://github.com/bleu48/GoGo](https://github.com/bleu48/GoGo) を元に練習☆（＾～＾）！  

大会版はこっち（＾～＾） 👉 [きふわらべUEC13](https://github.com/muzudho/kifuwarabe-uec13)  
# Overview

```plain
+----------------------------------------------------------------------+
| 📂presenter                                                          | 下層のライブラリに依存することがあります
+----------------------------------------------------------------------+

+--------------+--------------+----------+----------+    +-------------+
| 📂coding_obj | 📂config_obj | 📂input | 📂output |    | 📂entities | 上層または横に対しては 独立したパッケージ
+--------------+--------------+----------+----------+    +-------------+

+----------------------------------------------------------------------+
| fmt, os, math      ...                                               | Pythonの標準ライブラリ
|          math/rand                                                   |
+----------------------------------------------------------------------+
```

# Example - Lesson01

盤を表示して気分を盛り上げます  

```shell
# 実行ファイルを作成するために、以下のコマンドを打鍵してください。
# ソースコードを改造したあとだけで構いません
go build
# kifuwarabe-gogo.exe ファイルが作成されました。

# コンピューター囲碁エンジンを実行するために、以下の実行ファイルのステム（拡張子を省いた名前）を打鍵してください。
kifuwarabe-gogo Lesson01
```

# Example - Lesson02

石を取るデモです  

```shell
kifuwarabe-gogo Lesson02
```

# Example - Lesson03

ランダム打ちで１局終わらせます（９路盤を想定）  

```shell
kifuwarabe-gogo Lesson03
```

# Example - Lesson04

```shell
kifuwarabe-gogo Lesson04
```

# Example - Lesson05

```shell
kifuwarabe-gogo Lesson05
```

# Example - Lesson06

```shell
kifuwarabe-gogo Lesson06
```

# Example - Lesson07

```shell
kifuwarabe-gogo Lesson07
```

# Example - Lesson08

```shell
kifuwarabe-gogo Lesson08
```

# Example - Lesson09

## Run

```shell
go build

kifuwarabe-gogo Lesson09
# しばらく待つと自動で開始します
```

# Example - Lesson09a

## Run

```shell
go build

kifuwarabe-gogo Lesson09a
```

Input: (Example)  

```shell
genmove black
# 少し時間がかかります

genmove white
# 少し時間がかかります

quit
```

# Example - Test

```shell
kifuwarabe-gogo Test
```

## Memorandom

`go build -ldflags "-s -w"` でデバッグ情報を外せる？  
高速化した気はしないが……。  

📖 [installation](./doc/installation/install.md)  
📖 [References](./doc/references.md)  
📖 [Board](./doc/board.md)  
