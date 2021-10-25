# yToolsBox

[README](README_en.md) | [中文文档](README.md)

yToolsBox 是一个 All-In-One 的 工具收纳与调度平台。目前支持收纳调度脚本(_.py _.sh)与 docker 容器。

## 为什么开发 yToolsBox ？

共享单车的出现，很大程度上解决了人们下班回家的“最后一公里”问题。

同样的，日常的工作中我们也会遇到很多与之类似的场景，比方说：我们会为了提升日常的工作效率，产出各种辅助工作的脚本或者 docker 镜像。但往往每个人产出的辅助工具都处于一种“各自为战”的割裂状态，大家互相都不了解对方的工具如何使用；亦或者有相关的使用说明文档，却鲜有人愿意阅读，最终还是以“还是你帮我弄一下吧”草草收场

如果我们把“解决各自工作中遇到的问题”当做“回家”，把为了解决这些问题而“产出的形形色色的工具”当做“回家的手段”，那么不难发现，这其实也是另一种形式的“最后一公里”问题。

因此，yToolsBox 的意义就在于，它是一个“交通工具”的收纳与调度平台，任何人都可以把自己产出的“交通工具”，按照指定的方式添加进这个平台，后续如果需要让其他人复用，那么直接进入工具盒选择配置并执行即可

## 部署

### 手动部署

#### 1. 依次执行如下指令

```shell
docker network create --driver bridge ytoolsbox_network

docker volume create ytoolsbox_db-data

docker run -itd -p 5432:5432 --name yToolsBox-db --network ytoolsbox_network -e POSTGRES_PASSWORD=test123456 -v ytoolsbox_db-data:/var/lib/postgresql/data postgres

docker run -itd -p 8081:8081 --name yToolsBox-api --network ytoolsbox_network -e HOST_SCRIPT_PATH=/home/yToolsBox/api/Script -v /home/yToolsBox/api/Script:/root/Script yanqiaoyu/ytoolsbox-api:v0.2.1  supervisord -c /etc/supervisord.conf

docker run -itd -p 80:80 --network ytoolsbox_network --name yToolsBox-dashboard yanqiaoyu/ytoolsbox-dashboard:v0.2.1
```

#### 2. 部署结果

![manu_deploy](/doc/pic/manu_deploy1.png)

#### 3. 验证结果

访问 http://yourIP 验证是否安装成功

### 用 docker-compose 部署

#### 1. docker-compose.yml

```yaml
version: '3'
services:
  yToolsBox-db:
    container_name: 'yToolsBox-db'
    image: postgres
    restart: always
    ports:
      - 5432:5432
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - network
    environment:
      POSTGRES_PASSWORD: test123456

  yToolsBox-api:
    container_name: 'yToolsBox-api'
    build:
      context: ./ytoolsbox-gin
      dockerfile: Dockerfile
    image: yanqiaoyu/ytoolsbox-api:v0.1.1
    depends_on:
      - yToolsBox-db
    networks:
      - network
    volumes:
      - /home/yToolsBox/api/Script:/root/Script
    # should be same as above path
    environment:
      HOST_SCRIPT_PATH: /home/yToolsBox/api/Script
    ports:
      - 8081:8081
    command:
      [
        'sh',
        'wait-for',
        'yToolsBox-db:5432',
        '--',
        './main',
        '-m',
        'production'
      ]

  yToolsBox-dashboard:
    container_name: 'yToolsBox-dashboard'
    build:
      context: ./ytoolsbox-vue
      dockerfile: Dockerfile
    image: yanqiaoyu/ytoolsbox-dashboard:v0.1.1
    networks:
      - network
    ports:
      - 80:80
    depends_on:
      - yToolsBox-db
      - yToolsBox-api

volumes:
  db-data:
networks:
  network:
    driver: bridge
```

#### 2.安装

执行

```shell
docker-compose up -d
```

#### 3.验证安装结果

同上

## 使用教程

### 1. 执行脚本类工具

#### 1.1 准备好一个脚本文件

假设我们现在有一个脚本 pwd.sh，具体内容如下



![usage_script_1](/doc/pic/usage_script_1.png)

#### 1.2 新增一个工具

进入工具盒界面



![usage_script_2](/doc/pic/usage_script_2.png)

#### 1.3 填写工具基础信息

![usage_script_3](/doc/pic/usage_script_3.png)

#### 1.4 填写工具配置信息

![usage_script_4](/doc/pic/usage_script_4.png)

![usage_script_5](/doc/pic/usage_script_5.png)

#### 1.5 确认添加

![usage_script_6](/doc/pic/usage_script_6.png)

#### 1.6 发起执行任务

![usage_script_7](/doc/pic/usage_script_7.png)

![usage_script_8](/doc/pic/usage_script_8.png)

![usage_script_9](/doc/pic/usage_script_9.png)

![usage_script_10](/doc/pic/usage_script_10.png)

#### 1.7 查看任务执行结果

![usage_script_11](/doc/pic/usage_script_11.png)

![usage_script_12](/doc/pic/usage_script_12.png)

### 2. 执行容器类工具



## 开发进度

截止至 2021 年 10 月 8 日 15:06:52
![developProgress](/doc/pic/developProgress.png)
