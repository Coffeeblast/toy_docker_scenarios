import oracledb

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

# Query to fetch the first 5 rows from the random_data table
query = "SELECT * FROM random_data FETCH FIRST 5 ROWS ONLY"

try:
    # Execute the query
    cursor.execute(query)
    
    # Fetch and print the rows
    rows = cursor.fetchall()
    for row in rows:
        print(row)
except oracledb.DatabaseError as e:
    print(f"Error fetching data: {e}")

# Close the cursor and connection
cursor.close()
connection.close()