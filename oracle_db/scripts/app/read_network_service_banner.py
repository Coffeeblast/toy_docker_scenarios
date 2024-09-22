# see also https://python-oracledb.readthedocs.io/en/latest/user_guide/connection_handling.html

import oracledb

# Database connection parameters
# Use the alias defined in tnsnames.ora
TNS_ALIAS = 'freepdb1_alias'  # Replace with your TNS alias

USERNAME = 'app_user'  # Replace with your username
PASSWORD = 'app_user_password'  # Replace with your password

# Connect to the Oracle database using the TNS alias
connection = oracledb.connect(user=USERNAME, password=PASSWORD, dsn=TNS_ALIAS, config_dir="/scripts/app")
cursor = connection.cursor()

# Query to fetch the first 5 rows from the random_data table
query = "SELECT network_service_banner FROM v$session_connect_info"

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