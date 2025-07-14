<script setup>
import { onMounted, reactive,ref } from 'vue'
import { UserOutlined, LockOutlined,SwapOutlined } from '@ant-design/icons-vue';
import { login } from '@/api/login'
import { message } from 'ant-design-vue';

const formState = reactive({
  user_code: '',
  password: '',
  captcha: '',
});

const loading = ref(false)
const remember = ref(false)

const handleLogin = async ()=> {
    if (formState.user_code === "") {
        message.warning("请输入用户名")
        return
    }
    if (formState.password === "") {
        message.warning('请输入密码')
        return
    }
    loading.value = true
    
    const res = await login(formState)
    if (res.code === 10000 && res.data) {
        console.log('log1',remember.value)
        if (remember.value) {
            localStorage.setItem('remember-account',JSON.stringify({
                user_code: formState.user_code,
                password: formState.password
            }))
        } else {
            const user = localStorage.getItem('remember-account')
            console.log('log2',user)
            if (user) {
                localStorage.removeItem('remember-account')
            }
        }
    }
    loading.value = false
}

onMounted(() => {
    const account = localStorage.getItem('remember-account')
    if (account) {
        remember.value = true
        const user = JSON.parse(account)
        formState.user_code = user.user_code
        formState.password = user.password
    }
})
</script>

<template>
    <div class="login-v">
        <div class="background-v">
            
        </div>

        <div class="login-box">
            <h2 class="text-center text-3xl font-bold mb-8 text-gray-800">GinTalk Login</h2>
            <a-form :model="formState">
                <a-form-item>
                    <a-input v-model:value="formState.user_code" placeholder="请输入用户名" size="large">
                        <template #prefix>
                            <UserOutlined style="margin-right: 0.5rem;" />
                        </template>
                    </a-input>
                </a-form-item>

                <a-form-item>
                    <a-input-password v-model:value="formState.password" placeholder="请输入密码" size="large">
                        <template #prefix>
                            <LockOutlined style="margin-right: 0.5rem;" />
                        </template>
                    </a-input-password>
                </a-form-item>

                <a-form-item class="remember-v">
                    <div class="remember-item-v">
                        <a-checkbox v-model:checked="remember">记住密码</a-checkbox>
                        <a class="font-v" href="">忘记密码？</a>
                    </div>
                    
                </a-form-item>

                <a-form-item>
                    <a-button class="login-btn" type="primary" :loading="loading" @click="handleLogin">登录</a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<style scoped>
.login-v {
    width: 100%;
    height: 100vh;
    position: relative;
}

.particles-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}

.login-box {
  width: 400px;
  padding: 30px;
  background-color: white;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

h2 {
    text-align: center;
    font-size: 1.875rem;
    line-height: 2.25rem;
    font-weight: 700;
    margin-bottom: 2rem;
    color: #1f2937;
    font-family: 'Georgia', serif;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
}

.login-btn {
    width: 100%;
}

.remember-v {
    margin-bottom: 20px;
}

.remember-item-v {
    display: flex;
    justify-content: space-between;
}

.font-v {
    color: #807c7c;
}
</style>