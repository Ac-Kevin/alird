# [ackevin/alird](https://github.com/Ac-Kevin/alird)

[![GitHub Stars](https://img.shields.io/github/stars/Ac-Kevin/alird.svg?color=94398d&labelColor=555555&logoColor=ffffff&style=for-the-badge&logo=github)](https://github.com/Ac-Kevin/alird)
[![GitHub Release](https://img.shields.io/github/release/Ac-Kevin/alird?color=94398d&labelColor=555555&logoColor=ffffff&style=for-the-badge&logo=github)](https://github.com/Ac-Kevin/alird/releases)
[![GitHub Package Repository](https://img.shields.io/static/v1.svg?color=94398d&labelColor=555555&logoColor=ffffff&style=for-the-badge&label=linuxserver.io&message=GitHub%20Package&logo=github)](https://github.com/Ac-Kevin/alird/packages)
[![Docker Pulls](https://img.shields.io/docker/pulls/ackevin/alird.svg?color=94398d&labelColor=555555&logoColor=ffffff&style=for-the-badge&label=pulls&logo=docker)](https://hub.docker.com/r/ackevin/alird)
[![Docker Stars](https://img.shields.io/docker/stars/ackevin/alird.svg?color=94398d&labelColor=555555&logoColor=ffffff&style=for-the-badge&label=stars&logo=docker)](https://hub.docker.com/r/ackevin/alird)

# 快速启动

```
cd docker
docker-compose up
```


# 说明(explain)

将本地的公网IP更新至阿里云域名解析

`update the local public network IP to the resolution operator`

<br />
<br />

# 配置文件(config.ini)

配置文件默认读取位置 : `./alird/config.ini`

可以使用 `CONFIG_FILE` 环境变量修改位置

```
[domain]
name=域名
host=需要解析的关键字，如 www 和 api。多个host用;分割。
```
如
```
[domain]
;域名
name=china.cn
; 使用;分割RR 如： www;www2
host=www;
```

# 日志(logs)

默认输出日志位置： `./alird/logs`

> 保留 <em>7</em> 天内日志


# 环境变量(environment)

| Name            | Value                                                                                       |
| --------------- | ------------------------------------------------------------------------------------------- |
| ALI_ACCESSKEYID  | 阿里云 AccessKey ID ，需要登录阿里云后台获取。 [点击这里获取](https://ram.console.aliyun.com/overview)，注意配置域名权限 |
| ALI_ACCESSKEY_SECRET | 阿里云api 密钥,同上方获取方式，创建 AccessKey ID 时出现 <em>一次</em> ，请务必保存。                                                      |
|CONFIG_FILE     | 配置文件路径                                                                                |
|LOG_FILEDIR|日志存储位置|
|INTERVAL_TIME|检测间隔时间 (s) 默认30s|
# 映射文件和文件夹

| 类型   | 本机地址           | 容器地址          |
| ------ | :----------------- | :---------------- |
| 文件   | ./alird/config.ini | /alird/config.ini |
| 文件夹 | ./alird/logs       | /alird/logs       |





