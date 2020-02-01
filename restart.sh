#!/bin/bash
cd /home/hajnxg/go/src/nCoV-API
# 停止进程
ps -ef | grep "nCoV-API" | grep -v grep | awk '{print $2}' | xargs kill -9 
# 启动http服务
nohup ./nCoV-API -config "./config/app_prod_http.conf" >> access_http.log &
 
nohup ./nCoV-API -config "./config/app_prod.conf" >> access_https.log &
