#!/bin/bash
/usr/sbin/sshd
su - yarn_nodemanager -c "$HADOOP_HOME/bin/yarn --daemon start nodemanager"
#${SPARK_HOME}/sbin/start-worker.sh
#tail -f /dev/null
tail -f $SPARK_HOME/logs/*.log $HADOOP_HOME/logs/*.log