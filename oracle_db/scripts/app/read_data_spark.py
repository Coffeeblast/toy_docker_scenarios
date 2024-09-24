from pyspark.sql import SparkSession
import os

os.environ["JAVA_HOME"] = "/usr"
os.environ["HADOOP_HOME"] = "/usr/local/lib/python3.10/dist-packages/pyspar"

spark = SparkSession.builder \
    .appName("OracleConnector") \
    .config("spark.jars", "/jarfiles/ojdbc8.jar") \
    .getOrCreate()

# oracle_connection_properties = {
#     "user": "app_user",
#     "password": "app_user_password",
#     "driver": "oracle.jdbc.OracleDriver"
# }

jdbc_url = "jdbc:oracle:thin:@//oracle_db:1521/FREEPDB1"

# df = spark.read.jdbc(
#     url=jdbc_url,
#     table="random_data",  # You can also use a SQL query in place of a table name
#     properties=oracle_connection_properties
# )

df = spark.read.format("jdbc")\
    .option("url", jdbc_url)\
    .option("user", "app_user") \
    .option("password", "app_user_password") \
    .option("query", "select * from random_data fetch first 5 rows only")\
    .option("driver", "oracle.jdbc.OracleDriver")\
    .load()

df.show()