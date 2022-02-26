# 設計方針
* cmd/, pkg/に分割
* routerは設けず、api gwでdispatch
* 本当はdynamo以下に永続化処理を書いて色々やった方が良いんだろうけど、そこはサボる
    - まずhandlerの役割自体を大きく置いてる
    - CRUD含めて処理はpost, commentに一任
    
* third party, dbとデータの出入り口が二つある
    - dbに対してはpost_translator経由しよかな?
    
# 心配な部分
* lambda.Start(handlers.xxx)
    - `Handler()`関数省略
    - 不便ならラップすればいいよ
    
* handlers/xxx => entity/xxx
    - handlersは単なるreqestのバリデーションと関数呼び出し程度
- entity名以下のフォルダでcrud処理を担当させる
    - テストしやすくなったから多分あってる
 
* post.GetPosts()
    - 引数にoption構造体
    - Functional採用した方がいい？
    - afterも受付てないけど、必要なら設計をいじってみる
    
 * translatorとdynamodbで初期化方法が微妙に違う
    - コンストラクタ関数で生成するスコープが、translatorはsvcまで。dynamodbはconfigのみ
    - connectでsvcを生成。dbを利用する際はいちいち呼び出しする必要あり
      - svcとはいえあくまでdb。接続自体を単体テストしたい
      - 他dbライブラリとインターフェースを揃えたい
 
 * post, commentがtranslatorに依存してる
    - 外部ライブラリだけど、やってる事はシンプルだしまあいっか の精神

* 4そう構造にこだわるのやめた
    - translateとredditのエンドポイント分割して3走行層が一番スマート
    - データソースが増えるorネットワークパフォーマンスがボトルネックになる様ならpost丸ごとdbに入れる
        => その際usecase, domainレイヤを付与

* deeplの場合並列処理しないと遅すぎるかも
    - httpの並列処理について
# 制約
* post_translator.TranslateAllでbody > 1000は翻訳しない
    - 空文字で出力はされる
 
 # 疑問
 * sessionの有効期限。lambda自体の実行時間に依存してるのかな？
    - assume roleの制限が15分。オプションで変更可能らしいけど、lambdaの実行時間と一致している事から、設計上15分を想定していそう
  
 * もしerrorは出したくないけど、何かしらの形で知らせたいばあい
 
 * third party成のライブラリを触る時、どういう作りにすればええんや。
 
 * Translatorの作りわからん！
    - aws, deep　必要な引数が違う
    - type Translator interafce { client Client } の様に、interfce i ninterfaceにするとてスタビリティは高そう
    - けどそれって一つも実装をタッチしてないからテストする意味ってあるのかな
 * 命名規則
    - awsTranslator? aws_translatro?
    - NewClient_aws? NewAwsClient?
    
  * http clientのauth keyどこで？
 # 問題
 * guregu/dynamodb 構造体にマッピングできない！！！！
    - 応急処置でmap[string]string噛ませてる
    - 根本原因はあとで探す