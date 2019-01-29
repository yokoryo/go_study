# 準備
こんな感じでフォルダを作って宿題を提出する

![](/images/2018-11-19-20-29-53.png)


# 第1回目
- ex01：任意のLinux(Windows)コマンドを実行する
- ex02：引数を受け取って、出力する os.Argsを使いなさい
- ex03：引数を受け取って、出力する flagを使いなさい
- ex04：引数をflagで受け取って、任意のLinux(Windows)にコマンドに渡しなさい
例えばlsコマンドに-aや-lを渡しなさい

# 第2回目
- ex01：１行のtxt形式のファイルwを読み取り、内容すべて表示しなさい。   
※os.Openを使いなさい
また、os.Open関数はfileとerrを2つを返します。
errが返されたときわかるように、　errハンドリングしなさい。   
例えばファイルがないとき、エラーになるはずです。   
[sampleテキスト](/sample/sample002/kon.txt)   

- ex02：複数行のtxt形式のファイルを読み取り、１行ずつ１秒間隔で表示しなさい。   
入力をバッファリングし、表示することで出来ます。   
[sampleテキスト](/sample/sample002/all.txt)   

- ex03:１行のJSON形式のファイルを読み取り、内容すべて表示しなさい。   
jsonパッケージを利用します。また構造体(stract)の理解が必要です。   
[sample JSON](/sample/sample002/kon.json)   

- ex04：１行のJSON形式のファイルを読み取り、任意のキーとバリューを表示しなさい。   
- ex05:複数行のJSON形式のファイルを読み取り、任意のキーとバリューを表示しなさい。   
[sample JSON](/sample/sample002/go_study_member.json)   

# 第3回目
- ex01：sampleのプログラムを実行し、結果を確認しなさい   
[sampleプログラム](/sample/sample001)
- ex02：ex01のプログラムに追加して、/helloのときは「hello web server!」と表示するようにしなさい   
- ex03：URLを引数に指定し、その中身を表示しなさい   
httpパッケージでGETした値を、変数に格納し、Bodyをioutilパッケージで取得しましょう。   
- ex04：ex03で取得したWEBページのhttpレスポンスコードを表示しなさい。   


# 第4回目 Webクライアント（スクレイピング）
- ex01
  - 以下のsampleサーバプログラムを実行して動作することを確認して下さい。
  - [sampleプログラム](/sample/sample004)
- ex02
  - ex01にクライアントプログラムを追加して下さい。
  - サーバへリクエストを行い、結果を標準出力に表示して下さい。   
- ex03
  - ex02で作ったプログラムを修正します。
  - goqueryモジュールを利用して、aタグの個数をカウントして標準出力に表示して下さい。
- ex04
  - [GitHubトレンド](https://github.com/trending)のTop10を教えて下さい。
  - 不要な情報は見たくないので、ex03で作ってプログラムを修正して「リポジトリ名」「スター数」だけを出力して下さい。

# 第5回目 DB接続と操作
`注意：自習問題 002 (docker tutorialとSQL操作)を完了していることが前提です`  
- ex01
  - DBに接続してtasksテーブルのselect結果を標準出力しなさい  
  ※DBはMySQLを利用しています。ブランクインポートを利用してください。
- ex02
  - ex01を改修して、WEB上でselect結果を表示しなさい。
- ex03 
  - ex02を改修して、URLパラメータにNoをしているすると、そのレコードだけが表示されるようにしなさい。  
  （例） http://localhost:port/task?no=1
  - また、Noが表示しないときは、それがわかるようにしなさい。
- ex04
  - ex03を改修して、POSTメソッドでex03のURLにアクセスした時、tasksテーブルの特定のレコードのend_flagにendに更新されるようにしなさい。

# 自習問題 001 (アルゴリズムとbenchmark関数)
本問題は『楽しく学ぶ アルゴリズムとプログラミングの図鑑』を参考にしてください。
- ex01 (P75)
  - 合計値を求めるアルゴリズムを書きなさい
- ex02 (P88)
  - 平均値を求めるアルゴリズムを書きなさい
- ex03 (P99)
  - 最大値と最小値を探すアルゴリズムを書きなさい
- ex04 (P110)
  - ２つのデータを交換するアルゴリズムを書きなさい
- ex05 (P123)
  - リニアサーチを実装しなさい
- ex06 (P136)
  - バイナリサーチを実装しなさい
- ex07 (P161)
  - バブルソートを実装しなさい
- ex08 (P179)
  - 選択ソートを実装しなさい
- ex09 (P199)
  - 挿入ソートを実装しなさい
- ex10 (P220)
  - シェルソートを実装しなさい
- ex11 (P244)
  - クイックソートを実装しなさい
- ex12 (P272)
  - オブジェクトのソートをGoで書き直しなさい
- ex13 (P274)
  - シャッフルをGoで書き直しなさい
- ex14 (オリジナル)
  - ex05,ex06のプログラムをGo言語のbenchmark関数を使って計測しなさい  
  ※できるだけソートする値の前提条件を揃えて実施してください。
- ex15 (オリジナル)
  - ex07〜ex11のプログラムをGo言語のbenchmark関数を使って計測しなさい  
　※できるだけソートする値の前提条件を揃えて実施してください。

# 自習問題 002 (docker tutorialとSQL操作)
- ex01
  - 以下のsampleを実施しなさい
  - [docker sample001](/docker/sample001)
- ex02
  - 以下のsampleを実施しなさい
  - [docker sample002](/docker/sample002)
- ex03
  - 以下のsampleを実施しなさい
  - [docker sample003](/docker/sample003)




