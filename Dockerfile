FROM ubuntu:18.04

RUN sed -i "s@/archive.ubuntu.com/@/mirrors.tuna.tsinghua.edu.cn/@g" /etc/apt/sources.list \
    && sed -i "s@/security.ubuntu.com/@/mirrors.ustc.edu.cn/@g" /etc/apt/sources.list\
    && rm -Rf /var/lib/apt/lists/* \
    && apt-get update --fix-missing -o Acquire::http::No-Cache=True\
    && apt-get install --assume-yes apt-utils\
    && apt-get install wget -y
# 安装 golang1.17.11
RUN set -x; wget https://golang.google.cn/dl/go1.17.11.linux-amd64.tar.gz\
    &&  tar -zxvf go1.17.11.linux-amd64.tar.gz -C /usr/local/\
    &&  echo "export GOROOT=/home/root/local/go" >> /etc/profile\
    &&  echo "export GOBIN=$GOROOT/bin" >> /etc/profile\
    &&  echo "export PATH=$GOROOT/bin:$PATH" >> /etc/profile\
    &&  /bin/bash -c "source /etc/profile"\
    &&  echo "export GOROOT=/home/root/local/go" >> ~/.bashrc\
    &&  echo "export GOBIN=$GOROOT/bin" >> ~/.bashrc\
    &&  echo "export PATH=$GOROOT/bin:$PATH" >> ~/.bashrc\
    &&  /bin/bash -c "source ~/.bashrc"

COPY ./mongod.conf /usr/local/TEMPCTX/
# 安装 mongo
RUN apt-get install libcurl4 openssl -y\
    && wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz\
    && tar -zxvf mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz -C /usr/local/ \
    # && mkdir -p /usr/local/mongodb\
    && mv /usr/local/mongodb-linux-x86_64-ubuntu1804-5.0.9 /usr/local/mongodb \
    && mv /usr/local/TEMPCTX/mongod.conf /usr/local/mongodb/mongod.conf\
    # && chown -R root /usr/local/mongodb \
    && echo "export PATH=$PATH:/usr/local/mongodb/bin" >> ~/.bashrc \
    && /bin/bash -c "source ~/.bashrc"\
    && mkdir -p /usr/local/mongodb/data/db \
    && mkdir -p /usr/local/mongodb/logs \
    && touch /usr/local/mongodb/logs/mongod.log
    # && /usr/local/mongodb/bin/mongod -f /usr/local/mongodb/mongod.conf \

#mongodb的web端口
EXPOSE 28017
#连接端口
EXPOSE 27017

COPY ./mongo-add-user.js /usr/local/mongodb/
COPY ./mongo-add-admin.js /usr/local/mongodb/
COPY ./mongo-ini.sh /usr/local/mongodb/
RUN ["chmod", "+x", "usr/local/mongodb/mongo-ini.sh"]
# RUN /bin/bash -c "/usr/local/mongodb/bin/mongod -f /usr/local/mongodb/mongod.conf"\
#     && sleep 5

# CMD ["/usr/local/mongodb/bin/mongod", "-f", "/usr/local/mongodb/mongod.conf"]
# CMD ["/usr/local/mongodb/bin/mongo", "admin", "/usr/local/mongodb/mongo-add-admin.js"]

# RUN /bin/bash -c "/usr/local/mongodb/bin/mongo admin /usr/local/mongodb/mongo-add-admin.js"

# ENTRYPOINT ["/usr/local/mongodb/bin/mongod", "-f", "/usr/local/mongodb/mongod.conf"]
ENTRYPOINT ["/usr/local/mongodb/mongo-ini.sh"]
# RUN /bin/bash -c "/usr/local/mongodb/bin/mongod -f /usr/local/mongodb/mongod.conf"\
#     && sleep 5

# RUN /bin/bash -c "/usr/local/mongodb/bin/mongo r0Website /usr/local/mongodb/mongo-add-user.js"

# 构建镜像
# docker build -t ubuntu18-mongo5 .
# 以脱离模式运行容器
# docker run --name ubuntu18-mongo5 -d -p 27017:27017 ubuntu18-mongo5:latest
# 以交互终端运行容器
# docker run --name ubuntu18-mongo5 -it -p 27017:27017 ubuntu18-mongo5:latest /bin/bash

# docker attach ubuntu18-mongo5
# docker exec -it ubuntu18-mongo5 /bin/bash

# 删除容器和镜像
# docker container rm ubuntu18-mongo5 && docker image rm ubuntu18-mongo5