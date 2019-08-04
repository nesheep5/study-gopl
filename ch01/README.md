# 第一章 チュートリアル

## 1.1 ハロー、ワールド

## 1.2 コマンドライン引数
* https://godoc.org/os#pkg-variables

### 練習問題
#### EX1.1 
echoプログラムを修正して、そのプログラムを起動したコマンド名である`os.Args[0]` も表示するようにしなさい。
#### EX1.2 
echoプログラムを修正して、個々の引数のインデックスと値の組を1行ごとに表示しなさい。
#### EX1.3 
非効率な可能性のあるバージョンと `strings.Join` を使ったバージョンとで、実行時間の差を計測しなさい。

## 1.3 重複した行を見つける
* https://godoc.org/bufio#Scanner

### 練習問題
#### EX1.4
重複した行のそれぞれが含まれていたすべてのファイルの名前を表示するように `dup2`を修正しなさい。

## 1.4 GIFアニメーション
* https://godoc.org/image
* https://godoc.org/image/gif
* https://godoc.org/image/color

![lissajous](https://github.com/nesheep5/study-gopl/blob/master/ch01/lissajous/out.gif)
![ex05](https://github.com/nesheep5/study-gopl/blob/master/ch01/ex05/out.gif)
![ex05](https://github.com/nesheep5/study-gopl/blob/master/ch01/ex06/out.gif)
### 練習問題
#### EX1.5
もっともらしくするために、リサージュプログラムのカラーパレットを背景を黒として緑の線になるように修正しなさい。  
ウェブカラー `#RRGGBB` を作成するために `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`を使いなさい。  
ここで16進表記のそれぞれの値は、画素の赤、緑、青の成分の強度を表しています。
#### EX1.6
リサージュプログラムを修正して`palette`にもっと値を追加し、何らかの興味深い方法で`SetColorIndex`の第3引数を変更して複数の色で画像を生成するようにしなさい。

## 1.5 URLからの取得
* https://godoc.org/net/http
* https://godoc.org/io/ioutil#ReadAll
### 練習問題
#### EX1.7
関数呼び出し `io.Copy(dst, src)`は、`src`から読み込み`dst`へ書き込みます。  
ストリーム全体を保持するのに十分な大きさのバッファを要求することなくレスポンスの内容を `os.Stdout`へコピーするために、`ioutil.ReadAll`の代わりにその関数を使いなさい。  
なお、`io.Copy`のエラー結果は必ず検査するようにしなさい。
#### EX1.8
`fetch`を修正して、個々の引数のURLに接頭辞 `http://`がなければ追加するようにしなさい。`strings.HasPrefix`を使いたくなるかも知れません。
#### EX1.9
`fetch`を修正して、`resp.Status`に設定されているHTTPステータスコードも表示するようにしなさい。

## 1.6 URLからの平行な取得
### 練習問題
#### EX1.10
大量のデータを生成するウェブサイトを見つけなさい。  
報告される時間が大きく変化するかを調べるために `fetchall`を2回続けて実行して、キャッシュされているかどうかを調査しなさい。  
毎回同じ内容を得ているでしょうか。 `fetchall`を修正して、その出力をファイルへ保存するようにして調べられるようにしなさい。
#### EX1.11
`alexa.com` にある上位１００万件のウェブサイトのように、より長い引数リストで`fetchall`を試しなさい。  
あるウェブサイトが応答しない場合には、プログラムはどのように振る舞うべきでしょうか。(※8.9節参照)

## 1.7 ウェブサーバ
### 練習問題
#### EX1.12
リサージュ図形のサーバを修正して、URLからパラメータ値を読み取るようにしなさい。  
例えば、`http://localhost:8000/?cycles=20`のようなURLでは、周回の回数をデフォルトの5ではなく20に設定するようにしなさい。  
文字列パラメータを変数に変換するために `strconv.Atoi`関数を使いなさい。  
その関数のドキュメントは `go doc strconv.Atoi` で見ることができます。

## 1.8 残りの項目
