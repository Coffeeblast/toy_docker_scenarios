KEY_NAMES=(user_key server_key)

for key_name in "${KEY_NAMES[@]}"; 
do
     rm -r $key_name
    mkdir $key_name
    ssh-keygen -t rsa -b 4096 -C "login as user from host to docker container" -f $key_name/$key_name -N ''
done
