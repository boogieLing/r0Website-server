FROM ubuntu:18.04

RUN sed -i "s@/archive.ubuntu.com/@/mirrors.tuna.tsinghua.edu.cn/@g" /etc/apt/sources.list \
    && sed -i "s@/security.ubuntu.com/@/mirrors.ustc.edu.cn/@g" /etc/apt/sources.list\
    && rm -Rf /var/lib/apt/lists/* \
    && apt-get update --fix-missing -o Acquire::http::No-Cache=True\
    && apt-get install --assume-yes apt-utils\
    && apt-get install wget -y
# 安装 golang1.17.11
RUN set -x; wget https://golang.google.cn/dl/go1.17.11.linux-amd64.tar.gz\
    &&  tar -zxvf go1.15.15.linux-amd64.tar.gz -C /usr/local/\
    &&  echo "export GOROOT=/home/root/local/go" >> /etc/profile\
    &&  echo "export GOBIN=$GOROOT/bin" >> /etc/profile\
    &&  echo "export PATH=$GOROOT/bin:$PATH" >> /etc/profile\
    &&  source /etc/profile\
    &&  echo "export GOROOT=/home/root/local/go" >> ~/.bashrc\
    &&  echo "export GOBIN=$GOROOT/bin" >> ~/.bashrc\
    &&  echo "export PATH=$GOROOT/bin:$PATH" >> ~/.bashrc\
    &&  source ~/.bashrc\
    &&  go version
COPY ./mongod.conf /usr/local/mongodb
# 安装 mongo
RUN apt-get install libcurl4 openssl\
    && wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz\
    && tar -zxvf mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz -C /usr/local/ \
    && mv /usr/local/mongodb-linux-x86_64-ubuntu1804-5.0.9/ /usr/local/mongodb \
    && chown -R root /usr/local/mongodb \
    && echo "export PATH=$PATH:/usr/local/mongodb/bin" >> ~/.bashrc \
    && source ~/.bashrc \
    && mkdir -p /usr/local/mongodb/data/db \
    && mkdir -p /usr/local/mongodb/logs \
    && touch /usr/local/mongodb/logs/mongod.log\
    && /usr/local/mongodb/bin/mongod -f /usr/local/mongodb/mongod.conf