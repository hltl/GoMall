document.addEventListener('DOMContentLoaded', () => {
    // 获取表单元素
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');

    // 登录表单处理
    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            // 获取表单数据
            const loginEmail = document.getElementById('loginEmail');
            const loginPassword = document.getElementById('loginPassword');
            
            if (!loginEmail || !loginPassword) {
                console.error('登录表单元素未找到');
                return;
            }

            const formData = {
                email: loginEmail.value.trim(),
                password: loginPassword.value
            };

            // 表单验证
            if (!formData.email) {
                alert('请输入邮箱地址');
                return;
            }
            if (!formData.password) {
                alert('请输入密码');
                return;
            }

            try {
                const response = await api.login(formData.email, formData.password);
                console.log('登录成功:', response);
                // 存储token
                if (response.token) {
                    localStorage.setItem('token', response.token);
                    alert('登录成功！');
                    // 这里可以添加登录成功后的跳转逻辑
                } else {
                    alert('登录成功，但未收到token');
                }
            } catch (error) {
                console.error('登录失败:', error);
                alert(error.message);
            }
        });
    } else {
        console.error('登录表单未找到');
    }

    // 注册表单处理
    if (registerForm) {
        registerForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            // 获取表单元素
            const registerEmail = document.getElementById('registerEmail');
            const registerPassword = document.getElementById('registerPassword');
            const registerConfirmPassword = document.getElementById('registerConfirmPassword');
            
            if (!registerEmail || !registerPassword || !registerConfirmPassword) {
                console.error('注册表单元素未找到');
                return;
            }

            // 获取表单数据
            const formData = {
                email: registerEmail.value.trim(),
                password: registerPassword.value,
                confirmPassword: registerConfirmPassword.value
            };

            // 表单验证
            if (!formData.email) {
                alert('请输入邮箱地址');
                return;
            }
            if (!formData.password) {
                alert('请输入密码');
                return;
            }
            if (!formData.confirmPassword) {
                alert('请输入确认密码');
                return;
            }
            if (formData.password !== formData.confirmPassword) {
                alert('两次输入的密码不一致！');
                return;
            }

            try {
                console.log('发送注册请求:', {
                    email: formData.email,
                    password: formData.password,
                    confirm_password: formData.confirmPassword
                });
                const response = await api.register(
                    formData.email,
                    formData.password,
                    formData.confirmPassword
                );
                console.log('注册成功:', response);
                if (response.token) {
                    localStorage.setItem('token', response.token);
                }
                alert('注册成功！请登录');
                switchForm('login');
            } catch (error) {
                console.error('注册失败:', error);
                alert(error.message);
            }
        });
    } else {
        console.error('注册表单未找到');
    }
});

function switchForm(formType) {
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');
    const tabs = document.querySelectorAll('.tab-btn');

    if (formType === 'login') {
        loginForm.classList.add('active');
        registerForm.classList.remove('active');
        tabs[0].classList.add('active');
        tabs[1].classList.remove('active');
    } else {
        loginForm.classList.remove('active');
        registerForm.classList.add('active');
        tabs[0].classList.remove('active');
        tabs[1].classList.add('active');
    }
} 