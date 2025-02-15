const AUTH_BASE_URL = 'http://localhost:8888'; // etcd服务地址
const CONSUL_URL = 'http://localhost:8500'; // consul服务地址

const api = {
    // 从consul获取服务地址
    async getServiceAddress(serviceName) {
        try {
            const response = await fetch(`${CONSUL_URL}/v1/health/service/${serviceName}?passing=true`);
            const services = await response.json();
            
            if (!services || services.length === 0) {
                console.warn(`No healthy ${serviceName} service found`);
                throw new Error(`Service ${serviceName} not found`);
            }

            // 获取第一个健康的服务实例
            const service = services[0].Service;
            return `http://${service.Address}:${service.Port}`;
        } catch (error) {
            console.error('Failed to get service from Consul:', error);
            throw error;
        }
    },

    // 认证相关API - 使用etcd
    async register(username, password, confirmPassword) {
        try {
            const requestData = {
                email: username,
                password: password,
                confirm_password: confirmPassword
            };
            
            console.log('准备发送注册请求:', requestData);
            
            const response = await fetch(`${AUTH_BASE_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify(requestData),
            });
            
            const responseData = await response.json();
            console.log('收到注册响应:', responseData);
            
            if (!response.ok) {
                throw new Error(responseData.error || `注册失败: ${response.statusText}`);
            }
            
            return responseData;
        } catch (error) {
            console.error('注册请求失败:', error);
            throw error;
        }
    },

    async login(email, password) {
        try {
            const requestData = {
                email: email,
                password: password
            };
            
            console.log('准备发送登录请求:', requestData);
            
            const response = await fetch(`${AUTH_BASE_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify(requestData),
            });
            
            const responseData = await response.json();
            console.log('收到登录响应:', responseData);
            
            if (!response.ok) {
                throw new Error(responseData.error || `登录失败: ${response.statusText}`);
            }
            
            return responseData;
        } catch (error) {
            console.error('登录请求失败:', error);
            throw error;
        }
    },

    // 购物车相关API - 使用consul服务发现
    async addToCart(userId, productId, quantity) {
        try {
            const cartServiceUrl = await this.getServiceAddress('cart');
            const requestData = {
                user_id: userId,
                item: {
                    product_id: productId,
                    quantity: quantity
                }
            };

            const response = await fetch(`${cartServiceUrl}/cart/add`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify(requestData),
            });

            const responseData = await response.json();
            if (!response.ok) {
                throw new Error(responseData.error || '添加购物车失败');
            }
            return responseData;
        } catch (error) {
            console.error('添加购物车失败:', error);
            throw error;
        }
    },

    async getCart(userId) {
        try {
            const cartServiceUrl = await this.getServiceAddress('cart');
            const response = await fetch(`${cartServiceUrl}/cart/get?user_id=${userId}`, {
                credentials: 'include'
            });

            const responseData = await response.json();
            if (!response.ok) {
                throw new Error(responseData.error || '获取购物车失败');
            }
            return responseData;
        } catch (error) {
            console.error('获取购物车失败:', error);
            throw error;
        }
    },

    async emptyCart(userId) {
        try {
            const cartServiceUrl = await this.getServiceAddress('cart');
            const response = await fetch(`${cartServiceUrl}/cart/empty`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify({ user_id: userId }),
            });

            const responseData = await response.json();
            if (!response.ok) {
                throw new Error(responseData.error || '清空购物车失败');
            }
            return responseData;
        } catch (error) {
            console.error('清空购物车失败:', error);
            throw error;
        }
    }
}; 