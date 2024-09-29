rm -r spark/downloads hadoop/downloads
mkdir spark/downloads hadoop/downloads
curl -o spark/downloads/spark.tgz https://dlcdn.apache.org/spark/spark-3.5.3/spark-3.5.3-bin-hadoop3.tgz
curl -o hadoop/downloads/hadoop.tar.gz https://dlcdn.apache.org/hadoop/common/stable/hadoop-3.4.0.tar.gz