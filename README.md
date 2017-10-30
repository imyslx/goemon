# Go-e-Mom

(WIP)
Goで作る監視アプリケーション Go-e-Mon です。  
命名に意味はないです。Go-Monitorにしようとして間にeを挟んでみたくなっただけです。

# Components

Server/Client型のアプリケーションです。  
バックエンドはMySQLとKafkaを利用する予定です。

Kafkaの後ろをどうするかは悩んでますがGridDB使ってみたいと思ってます。  
Pythonとかで別に切るかも。

* 管理用DB
  * MySQL
* 監視データの1次ストア
  * Apache Kafka
* 時系列監視データ (Repo内には置かない予定)
  * GridDB (KairosDB-connector)
* 発報用データ
  * 未定（毎度DB叩くのもバカっぽいので悩み中…。）
