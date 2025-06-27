

## 	elasticsearch8.18.3重置密码
重置elasticsearch密码出现错误：
ERROR: Failed to determine the health of the cluster., with exit code 69

解决办法：
1、修改es配置文件：/etc/elasticsearch/elasticsearch.yml

将xpack.security.http.ssl的enable改为false
```shell
xpack.security.http.ssl:
  enabled: false
```

2、cd到es的根目录/bin，执行重置密码命令，并记下执行结果（密码）：

```shell
./elasticsearch-reset-password -u elastic
```
```shell
Password for the [elastic] user successfully reset.
New value: EgMi8g=91RMvcWVlmteZ
```
3、将第一步中的enable重新设置为true，否则登录不了

4、重启es:
```shell
sudo systemctl restart elasticsearch
```
5、使用elastic和生成的密码重新登录http://localhost:5601







​	

