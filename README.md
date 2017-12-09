# Go-e-Mom

(WIP)  
Goで作る監視アプリケーション Go-e-Mon です。  
命名に意味はないです。Go-Monitorにしようとして間にeを挟んでみたくなっただけです。

# Consept

Basicな機能を備えつつ分散に特化した構成です。
正直Zabbixなどに勝てる程に機能を充実させる気は無いです。

# Components

Server/Client型のアプリケーションです。  
バックエンドはMySQLとKafkaを利用する予定です。

Kafkaの後ろをどうするかは悩んでますがGridDB使ってみたいと思ってます。  
ConnectorはPythonとかで別に切るかも。

* 管理用DB
  * MySQL
* 監視データの1次ストア
  * Apache Kafka
* 時系列監視データ (Repo内には置かない予定)
  * GridDB (KairosDB-connector)
* 発報用データ
  * 未定（毎度DB叩くのもバカっぽいので悩み中…。）
