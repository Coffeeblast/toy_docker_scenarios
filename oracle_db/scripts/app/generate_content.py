import oracledb
import random
import string
from datetime import datetime, timedelta

# Database connection parameters
HOST = 'oracle_db'
PORT = '1521'  # Default Oracle port
SERVICE_NAME = 'FREEPDB1'
USERNAME = 'app_user'
PASSWORD = 'app_user_password'

# Connect to the Oracle database
dsn = f"{HOST}:{PORT}/{SERVICE_NAME}"
connection = oracledb.connect(user=USERNAME, password=PASSWORD, dsn=dsn)
cursor = connection.cursor()

# Create a table
table_creation_query = """
CREATE TABLE random_data (
    id NUMBER GENERATED BY DEFAULT AS IDENTITY,
    int_field NUMBER,
    double_field FLOAT,
    string_field VARCHAR2(50),
    datetime_field TIMESTAMP
)
"""

# Execute table creation
try:
    cursor.execute(table_creation_query)
    print("Table created successfully.")
except oracledb.DatabaseError as e:
    print(f"Error creating table: {e}")

# Function to generate random data
def random_string(length=10):
    letters = string.ascii_letters
    return ''.join(random.choice(letters) for i in range(length))

# Insert randomly generated rows
insert_query = """
INSERT INTO random_data (int_field, double_field, string_field, datetime_field)
VALUES (:1, :2, :3, :4)
"""

for _ in range(100):
    int_field = random.randint(1, 100)
    double_field = random.uniform(1.0, 100.0)
    string_field = random_string(random.randint(5, 10))
    datetime_field = datetime.now() + timedelta(days=random.randint(0, 30))

    # Execute insert
    cursor.execute(insert_query, (int_field, double_field, string_field, datetime_field))

# Commit the transaction
connection.commit()
print("100 random rows inserted successfully.")

# Close the cursor and connection
cursor.close()
connection.close()