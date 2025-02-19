const API_BASE_URL = 'http://localhost:8888';

const api = {
    async register(username, password, confirmPassword) {
        try {
            const requestData = {
                email: username,
                password: password,
                confirm_password: confirmPassword
            };
            
            console.log('准备发送注册请求:', requestData);
            
            const response = await fetch(`${API_BASE_URL}/register`, {
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
            
            const response = await fetch(`${API_BASE_URL}/login`, {
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
}; 