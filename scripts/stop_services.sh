#!/bin/bash

# 查找并终止所有相关的服务进程
echo "Stopping all services..."
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

# 为每个服务停止进程
for service in "${services[@]}"
do
    # 使用更精确的匹配方式查找进程
    pid=$(pgrep -f "${service}")
    
    if [ -n "$pid" ]; then
        echo "Found ${service} service with PID: $pid"
        
        # 尝试优雅停止
        echo "Attempting graceful shutdown of ${service}..."
        kill -SIGINT $pid
        
        # 等待最多5秒让进程优雅退出
        for i in {5..1}
        do
            if kill -0 $pid 2>/dev/null; then
                echo "Waiting... $i seconds"
                sleep 1
            else
                echo "${service} stopped successfully"
                break
            fi
        done
        
        # 如果进程仍然存在，强制终止
        if kill -0 $pid 2>/dev/null; then
            echo "Force killing ${service} (PID: $pid)"
            kill -9 $pid
        fi
    else
        echo "No running process found for ${service}"
    fi
done

echo "All services shutdown completed"