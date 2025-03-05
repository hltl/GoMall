#!/bin/bash

# 定义服务目录
SERVICES_DIR="app"
BASE_DIR=$(pwd)

# 定义要启动的服务列表
services=(
    "auth"
    "user" 
    "frontend"
    "cart"
    "order"
    "payment"
    "checkout"
    "product"
)

# 创建日志目录
mkdir -p logs

# 循环启动每个服务
for service in "${services[@]}"
do
    echo "Starting $service service..."
    cd "$BASE_DIR/$SERVICES_DIR/$service"
    go run . > "$BASE_DIR/logs/$service.log" 2>&1 &
    echo "$service service started with PID $!"
done

echo "All services started successfully!"