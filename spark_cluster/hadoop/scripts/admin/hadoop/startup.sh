#!/bin/bash
/usr/sbin/sshd
$HADOOP_HOME/sbin/start-dfs.sh
$HADOOP_HOME/sbin/start-yarn.sh
#tail -f /dev/null
tail -f $HADOOP_HOME/logs/*.log