print("oracledb_test")

import oracledb

ip_address = "oracle_db"
with oracledb.connect(user="app_user", password="app_user_password", dsn=f"{ip_address}/FREEPDB1") as connection:
    with connection.cursor() as cursor:
        sql = """select sysdate from dual"""
        for r in cursor.execute(sql):
            print(r)