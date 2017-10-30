# Go-e-Mom Client

**(WIP)**
Goで作る監視アプリケーション Go-e-Mon のClientアプリケーションです。

goroutineとTickerを利用して一定間隔毎に情報収集を行い、バックエンドのDBに記録します。

# Components

Server/Client型のアプリケーションです。
バックエンドはKafkaを利用する予定です。

* 監視データの1次ストア
  * Apache Kafka