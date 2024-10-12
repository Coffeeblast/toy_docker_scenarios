from pyspark.sql import SparkSession
# import os

# os.environ["JAVA_HOME"] = "/usr"
# os.environ["HADOOP_HOME"] = "/usr/local/lib/python3.10/dist-packages/pyspar"

spark = SparkSession.builder \
    .appName("OracleConnector") \
    .master("yarn")\
    .getOrCreate()

data = [
    [1, "whee"],
    [2, "glee"]
]

df = spark.createDataFrame(data, schema=["id", "val"])

df.show()