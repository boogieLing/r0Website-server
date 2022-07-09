#!/bin/sh
file_from=""

apt_get_install_flag="false"
apt_get_install="RUN apt\-get install \-\-assume-yes apt\-utils \&\& apt\-get install wget \-y "

if [ ! -f "go1.17.11.linux-amd64.tar.gz" ];then
  echo "go1.17.11.linux-amd64.tar.gz not exist"
  if [ "$apt_get_install_flag" = "false" ];then
    sed -i "s/\${INSTALL_WGET}/${apt_get_install}/g" ./Dockerfile &&
    echo "${apt_get_install}"
  fi
  file_from="RUN wget https:\/\/golang.google.cn\/dl\/go1.17.11.linux\-amd64.tar.gz \-p \/tmp"
else
  echo "go1.17.11.linux-amd64.tar.gz exist !"
  file_from="COPY [\"go1.17.11.linux\-amd64.tar.gz\", \"\/tmp\/\"]"
fi

sed -i "s/\${LOAD_FILE_GO}/$file_from/g" ./Dockerfile &&
echo "${file_from}"

if [ ! -f "mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz" ];then
  echo "mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz not exist"
  if [ "$apt_get_install_flag" = "false" ];then
    sed -i "s/\${INSTALL_WGET}/$apt_get_install/g" ./Dockerfile &&
    echo "${apt_get_install}"
  fi
  file_from="RUN wget https:\/\/fastdl.mongodb.org\/linux\/mongodb\-linux\-x86_64\-ubuntu1804\-5.0.9.tgz \-p \/tmp"
else
  echo "mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz exist !"
  file_from="COPY [\"mongodb\-linux-x86_64\-ubuntu1804\-5.0.9.tgz\", \"\/tmp\/\"]"
fi

sed -i "s/\${LOAD_FILE_MONGO}/$file_from/g" ./Dockerfile &&
echo "${file_from}" &&
cat ./Dockerfile
# 构建编译应用程序的镜像
# docker build -t ubuntu18-mongo5 . &&

# 以脱离模式运行容器
# docker run --name ubuntu18-mongo5 -d -p 27017:27017 ubuntu18-mongo5:latest