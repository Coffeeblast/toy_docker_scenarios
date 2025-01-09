FROM ubuntu:22.04

RUN apt-get update && apt-get upgrade

# installing ssh server
RUN apt-get install -y openssh-server

# creates a backup of unmodified sshd_config and sets it unwritable
RUN cp /etc/ssh/sshd_config /etc/ssh/sshd_config.original && chmod a-w /etc/ssh/sshd_config.original

RUN mkdir -p /run/sshd && \
    mkdir /scripts
COPY ./scripts/entrypoint.sh /scripts/entrypoint.sh
RUN chmod u+x /scripts/entrypoint.sh

# setting up a non-root user ... we will login into ssh_server via ssh as this user
ARG USR="somedude"
ARG PASS="somepass"

RUN useradd ${USR} -p $(openssl passwd -6 $PASS) --create-home -s /bin/bash && \
    usermod -aG sudo ${USR} && \
    mkdir /tmp_stuff

COPY ./ssh_keys /tmp_stuff

RUN mkdir /home/${USR}/.ssh && \
    cat /tmp_stuff/user_key/user_key.pub >> /home/${USR}/.ssh/authorized_keys && \
    chown -R ${USR}:${USR} /home/${USR}/.ssh && \
    chmod 600 /home/${USR}/.ssh/authorized_keys && \
    chmod 700 /home/${USR}/.ssh

# set up server signature keys
RUN rm /etc/ssh/ssh_host_* &&\
    cp /tmp_stuff/server_key/* /etc/ssh && \
    chmod 600 /etc/ssh/server_key && \
    chmod 644 /etc/ssh/server_key.pub && \
    echo "HostKey /etc/ssh/server_key" >> /etc/ssh/sshd_config && \
    rm -rf /tmp_stuff

#install python
RUN apt-get install -y software-properties-common && \
    add-apt-repository -y ppa:deadsnakes && \
    apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y python3.11 python3.11-dev python3.11-minimal python3.11-distutils python3.11-venv

# #install pip
# RUN mkdir tmp_files &&\
#     apt-get update &&\
#     apt-get install -y curl &&\
#     curl -o /tmp_files/get-pip.py https://bootstrap.pypa.io/get-pip.py &&\
#     python3.12 /tmp_files/get-pip.py &&\
#     rm -r tmp_files

# install java
# RUN apt-get update && apt-get install -y openjdk-8-jdk

# install python packages
# RUN python3.12 -m pip install py4j pyspark
# RUN python3.12 -m pip install jupyter

# install node.js
# RUN apt-get install -y npm # old version...

# RUN apt-get update && apt-get install python-software-properties
RUN apt-get update && apt-get install -y software-properties-common curl
# Here, we are installing node.js version 22, if you want to install version 23, you can replace setup_22.x with setup_23.x.
RUN curl -sSL https://deb.nodesource.com/setup_22.x | bash -
RUN apt-get -y install nodejs

# some utilities for analysing network, ping, nmap, nslookup ... 
RUN apt-get install -y nmap net-tools iputils-ping dnsutils nano

# install go
RUN mkdir tmp_files \
    && curl -SL -o tmp_files/go1.23.4.linux-amd64.tar.gz  https://go.dev/dl/go1.23.4.linux-amd64.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp_files/go1.23.4.linux-amd64.tar.gz \
    && rm -r tmp_files \
    && echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile

# install sqlite3 commandline app
RUN apt-get update && apt-get install sqlite3