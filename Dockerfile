FROM ubuntu:18.04 as base_os
# 升级 apt-get
RUN sed -i "s@/archive.ubuntu.com/@/mirrors.tuna.tsinghua.edu.cn/@g" /etc/apt/sources.list \
    && sed -i "s@/security.ubuntu.com/@/mirrors.ustc.edu.cn/@g" /etc/apt/sources.list\
    && rm -Rf /var/lib/apt/lists/* \
    && apt-get update --fix-missing -o Acquire::http::No-Cache=True

# 是否要安装wget
${INSTALL_WGET}

# 加载go压缩包的方式
${LOAD_FILE_GO}

# 加载mongo压缩包的方式
${LOAD_FILE_MONGO}

# 安装 golang1.17.11
RUN set -x; tar -zxvf go1.17.11.linux-amd64.tar.gz -C /usr/local/\
    &&  echo "export GOROOT=/home/root/local/go" >> /etc/profile\
    &&  echo "export GOBIN=$GOROOT/bin" >> /etc/profile\
    &&  echo "export PATH=$GOROOT/bin:$PATH" >> /etc/profile\
    &&  /bin/bash -c "source /etc/profile"\
    &&  echo "export GOROOT=/home/root/local/go" >> ~/.bashrc\
    &&  echo "export GOBIN=$GOROOT/bin" >> ~/.bashrc\
    &&  echo "export PATH=$GOROOT/bin:$PATH" >> ~/.bashrc\
    &&  /bin/bash -c "source ~/.bashrc"

# 安装 mongo
COPY ["mongod.conf", "/usr/local/TEMPCTX/"]
RUN apt-get install libcurl4 openssl -y\
    && tar -zxvf mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz -C /usr/local/ \
    && mv /usr/local/mongodb-linux-x86_64-ubuntu1804-5.0.9 /usr/local/mongodb \
    && mv /usr/local/TEMPCTX/mongod.conf /usr/local/mongodb/mongod.conf\
    && echo "export PATH=$PATH:/usr/local/mongodb/bin" >> ~/.bashrc \
    && /bin/bash -c "source ~/.bashrc"\
    && mkdir -p /usr/local/mongodb/data/db \
    && mkdir -p /usr/local/mongodb/logs \
    && touch /usr/local/mongodb/logs/mongod.log

#mongodb的web端口
EXPOSE 28017
#连接端口
EXPOSE 27017

ENTRYPOINT ["/usr/local/mongodb/bin/mongod", "-f", "/usr/local/mongodb/mongod.conf"]

# 流程：构建->脱离模式运行容器->即时设置权限

# 构建镜像
# docker build -t ubuntu18-mongo5 --build-arg FILE_FROM=copy .
# 以脱离模式运行容器
# docker run --name ubuntu18-mongo5 -d -p 27017:27017 ubuntu18-mongo5:latest
# 以交互终端运行容器
# docker run --name ubuntu18-mongo5 -it -p 27017:27017 ubuntu18-mongo5:latest /bin/bash

# docker attach ubuntu18-mongo5
# docker exec -it ubuntu18-mongo5 /bin/bash

# 删除容器和镜像
# docker container rm ubuntu18-mongo5 && docker image rm ubuntu18-mongo5