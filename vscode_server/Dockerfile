FROM ubuntu:22.04

RUN apt-get update && apt-get upgrade
RUN apt-get install -y wget gpg
# RUN wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg &&\
#     install -D -o root -g root -m 644 packages.microsoft.gpg /etc/apt/keyrings/packages.microsoft.gpg &&\
#     echo "deb [arch=amd64,arm64,armhf signed-by=/etc/apt/keyrings/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main" | tee /etc/apt/sources.list.d/vscode.list > /dev/null && \
#     rm -f packages.microsoft.gpg 
# RUN apt-get install -y apt-transport-https && \
#     apt-get update && \ 
#     apt-get install -y code

# installing ssh server
RUN apt-get install -y openssh-server

# creates a backup of unmodified sshd_config and sets it unwritable
RUN cp /etc/ssh/sshd_config /etc/ssh/sshd_config.original && chmod a-w /etc/ssh/sshd_config.original

RUN mkdir -p /run/sshd
RUN mkdir /scripts
COPY ./scripts/entrypoint.sh /scripts/entrypoint.sh
RUN chmod u+x /scripts/entrypoint.sh

# setting up a non-root user ... we will login into ssh_server via ssh as this user
ARG USR="somedude"
ARG PASS="somepass"

RUN useradd ${USR} -p $(openssl passwd -6 $PASS) --create-home -s /bin/bash

# append USR into sudo group
RUN usermod -aG sudo ${USR}
#USER ${USR}

# set up user ssh keys
RUN mkdir /tmp_stuff
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

# this environment variable needs to be set to anything, so that code does not yap about wsl
# ENV DONT_PROMPT_WSL_INSTALL=whee
# RUN mkdir /code_user_data
# RUN code --install-extension ms-vscode-remote.remote-ssh --no-sandbox --user-data-dir /code_user_data

