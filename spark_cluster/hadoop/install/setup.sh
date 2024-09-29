mkdir /run/sshd 

# import environment variables
source $HADOOP_HOME/etc/hadoop/hadoop-env.sh

user_list=($HDFS_NAMENODE_USER $HDFS_SECONDARYNAMENODE_USER $HDFS_DATANODE_USER)
for tmp_usr in "${user_list[@]}"; 
do
     echo "Setting up user $tmp_usr"
     useradd $tmp_usr --create-home -s /bin/bash
     usermod -aG $HDFS_GROUP $tmp_usr
    
     # setup ssh keys
     mkdir /home/$tmp_usr/.ssh
     ssh-keygen -t rsa -P '' -C "$tmp_usr@localhost" -f "/home/$tmp_usr/.ssh/id_rsa"
     cat /home/$tmp_usr/.ssh/id_rsa.pub >> /home/$tmp_usr/.ssh/authorized_keys  
     chown -R $tmp_usr "/home/$tmp_usr/.ssh"
     chmod 0600 /home/$tmp_usr/.ssh/authorized_keys

done

# format hdfs
#su - hdfs_namenode -c "$HADOOP_HOME/bin/hdfs namenode -format"
$HADOOP_HOME/bin/hdfs namenode -format