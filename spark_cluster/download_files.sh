rm -r spark/downloads
mkdir spark/downloads
curl -o spark/downloads/spark.tgz https://dlcdn.apache.org/spark/spark-3.5.3/spark-3.5.3-bin-without-hadoop.tgz

rm -r hadoop/downloads
mkdir hadoop/downloads
curl -o hadoop/downloads/hadoop.tar.gz https://dlcdn.apache.org/hadoop/common/stable/hadoop-3.4.0.tar.gz