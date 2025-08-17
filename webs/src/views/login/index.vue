<script setup>
import { ref, reactive, onMounted } from 'vue'
import { UserOutlined, LockOutlined, MailOutlined,PhoneOutlined } from '@ant-design/icons-vue'
import { login,register } from '@/api/login'
import { message,Modal } from 'ant-design-vue'
import { useRouter } from 'vue-router';

// 响应式数据
const isSignUpMode = ref(false)
const container = ref(null)
const loading = ref(false)
const remember = ref(false)
const router = useRouter();

const signInForm = reactive({
    user_code: '',
    password: '',
    captcha: '',
})

const handleSignIn = async ()=> {
    if (signInForm.user_code === "") {
        message.warning("请输入用户名")
        return
    }
    if (signInForm.password === "") {
        message.warning('请输入密码')
        return
    }
    // if (signInForm.captcha === "") {
    //     message.warning('请输入验证码')
    //     return
    // }
    loading.value = true
    
    const res = await login(signInForm)
    loading.value = false
    if (res.code === 10000 && res.data) {
      localStorage.setItem('access_token', res.data.access_token);
      localStorage.setItem('refresh_token', res.data.refresh_token);
      localStorage.setItem('expire_time', res.data.expire_time);
      if (remember.value) {
          localStorage.setItem('remember-account',JSON.stringify({
              user_code: signInForm.user_code,
              password: signInForm.password
          }))
          // TODO: navigate to homePage
          router.push({path: '/'})
      } else {
          const user = localStorage.getItem('remember-account')
          console.log('log2',user)
          if (user) {
              localStorage.removeItem('remember-account')
          }
      }
    }
}

const signUpForm = reactive({
  username: '',
  password: '',
  email: '',
  phone: '',
  gender: 1
})

// switch to register
const switchToSignUp = () => {
    Object.assign(signUpForm, {
        username: '',
        password: '',
        email: '',
        phone: '',
        gender: 1
    });
    isSignUpMode.value = true
}

// switch to login
const switchToSignIn = () => {
  isSignUpMode.value = false
}

// handle to register account
const handleSignUp = async () => {
  if (signUpForm.username === "") {
    message.warning("请输入用户名")
    return
  }
  if (signUpForm.password === "") {
    message.warning('请输入密码')
    return
  }
  if (signUpForm.phone === "") {
    message.warning('请输入手机号')
    return
  }
  if (signUpForm.email === "") {
    message.warning("请输入邮箱")
    return
  }

  loading.value = true
  const res = await register(signUpForm)
  loading.value = false
  if (res.code === 10000) {
    console.log('res',res)
    Modal.confirm({
        title: '注册成功',
        content: `您的账号已成功注册，登录账号为:${res.data},现在去登录吗？`,
        okText: '去登录',
        cancelText: '取消',
        onOk() {
          switchToSignIn(); 
          Object.assign(signInForm, {
            user_code: '',
            password: '',
            captcha: ''
          });

          Object.assign(signUpForm, {
            username: '',
            password: '',
            email: '',
            phone: '',
            gender: 1
          });
          message.info('已为您切换至登录页面');
        },
        onCancel() {
          message.info('您已取消，请自行选择登录或继续注册');
          Object.assign(signUpForm, {
            username: '',
            password: '',
            email: '',
            phone: '',
            gender: 1
          });
        },
    })
  }
}

// use ripple style
const createRipple = (event) => {
  const button = event.currentTarget
  const ripple = document.createElement('span')
  const rect = button.getBoundingClientRect()
  const size = Math.max(rect.width, rect.height)
  const x = event.clientX - rect.left - size / 2
  const y = event.clientY - rect.top - size / 2
  
  ripple.style.width = ripple.style.height = size + 'px'
  ripple.style.left = x + 'px'
  ripple.style.top = y + 'px'
  ripple.style.position = 'absolute'
  ripple.style.borderRadius = '50%'
  ripple.style.background = 'rgba(255, 255, 255, 0.6)'
  ripple.style.transform = 'scale(0)'
  ripple.style.animation = 'ripple 0.6s linear'
  ripple.style.pointerEvents = 'none'
  
  button.style.position = 'relative'
  button.style.overflow = 'hidden'
  button.appendChild(ripple)
  
  setTimeout(() => {
    ripple.remove()
  }, 600)
}

onMounted(() => {
  const socialButtons = document.querySelectorAll('.social')
  socialButtons.forEach(social => {
    social.addEventListener('mouseenter', function() {
      this.style.transform = 'translateY(-3px) scale(1.1)'
    })
    
    social.addEventListener('mouseleave', function() {
      this.style.transform = 'translateY(0) scale(1)'
    })
  });

  const account = localStorage.getItem('remember-account')
  if (account) {
    remember.value = true
    const user = JSON.parse(account)
    signInForm.user_code = user.user_code
    signInForm.password = user.password
  }
})
</script>

<template>
  <div class="login-page">
    <div class="container" :class="{ 'right-panel-active': isSignUpMode }" ref="container">
      
        <!-- 登录表单 -->
      <div class="form-container sign-in-container">
        <a-form :model="signInForm" class="form">
          <h1 style="font-size: 32px;font-family: 'Courier New', Courier, monospace;">GinTalk Login</h1>

          <div>
            <a-form-item>
                <a-input v-model:value="signInForm.user_code" placeholder="请输入用户名" size="large" class="input-v">
                    <template #prefix>
                        <UserOutlined style="margin-right: 0.5rem;margin-left: 0.5rem;" />
                    </template>
                </a-input>
            </a-form-item>

            <a-form-item>
                <a-input v-model:value="signInForm.password" placeholder="请输入密码" size="large" type="password" class="input-v">
                    <template #prefix>
                        <LockOutlined style="margin-right: 0.5rem;margin-left: 0.5rem;" />
                    </template>
                </a-input>
            </a-form-item>

            <a-form-item class="remember-v">
                <div class="remember-item-v">
                    <a-checkbox v-model:checked="remember">记住密码</a-checkbox>
                    <a class="font-v" href="">忘记密码？</a>
                </div>
            </a-form-item>

            <a-form-item>
                <a-button type="primary" :loading="loading" 
                    @click="handleSignIn" 
                    @click.native="createRipple" 
                    class="custom-button"
                >登录</a-button>
            </a-form-item>
          
          </div>
          
        </a-form>
      </div>

      <!-- 注册表单 -->
      <div class="form-container sign-up-container">
        <a-form :model="signUpForm" class="form">
          <h1 style="font-size: 30px;font-family: 'Courier New', Courier, monospace;">GinTalk Register</h1>
          
          <div>
            <a-form-item>
                <a-input v-model:value="signUpForm.username" placeholder="请输入用户名" size="large" class="input-v">
                    <template #prefix>
                        <UserOutlined style="margin-right: 0.5rem;margin-left: 0.5rem;" />
                    </template>
                </a-input>
            </a-form-item>

            <a-form-item>
                <a-input v-model:value="signUpForm.password" placeholder="请输入密码" type="password" size="large" class="input-v">
                    <template #prefix>
                        <LockOutlined style="margin-right: 0.5rem;margin-left: 0.5rem;" />
                    </template>
                </a-input>
            </a-form-item>

            <a-form-item>
                <a-input v-model:value="signUpForm.phone" placeholder="请输入手机号" size="large" class="input-v">
                    <template #prefix>
                        <PhoneOutlined style="margin-right: 0.5rem;margin-left: 0.5rem;" />
                    </template>
                </a-input>
            </a-form-item>

            <a-form-item>
                <a-input v-model:value="signUpForm.email" placeholder="请输入邮箱" size="large" class="input-v">
                    <template #prefix>
                        <MailOutlined style="margin-right: 0.5rem;margin-left: 0.5rem;" />
                    </template>
                </a-input>
            </a-form-item>

            <a-form-item label="性別">
                <a-radio-group v-model:value="signUpForm.gender" style="display: flex;justify-content: flex-start;margin-left: 20px;">
                    <a-radio :value="1">男</a-radio>
                    <a-radio :value="2">女</a-radio>
                </a-radio-group>
            </a-form-item>

            <a-form-item>
                <a-button type="primary" :loading="loading" 
                    @click="handleSignUp" 
                    @click.native="createRipple" 
                    class="custom-button"
                >注册</a-button>
            </a-form-item>
          </div>
        </a-form>
      </div>

      <!-- 覆盖层 -->
      <div class="overlay-container">
        <div class="overlay">
          <div class="overlay-panel overlay-left">
            <h1>欢迎回来！</h1>
            <p>要与我们保持联系，请使用您的个人信息登录</p>
            <a-button class="ghost" @click="switchToSignIn">登录</a-button>
          </div>
          <div class="overlay-panel overlay-right">
            <h1>还没账号？</h1>
            <p>输入您的个人详细信息，开始您的旅程</p>
            <a-button class="ghost" @click="switchToSignUp">注册</a-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  font-family: 'Arial', sans-serif;
  /* background-color: white; */
  background: linear-gradient(to right, #DCE3E8, #F5EEDF, #F5E2B8, #EBAA65, #dda77d);
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0;
  padding: 0;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.container {
  width: 800px;
  height: 500px;
  position: relative;
  background: white;
  border-radius: 20px;
  box-shadow:  0 20px 40px rgba(0, 0, 0, 0.1), /* Original bottom shadow */
  10px 0 20px rgba(0, 0, 0, 0.05), /* Right shadow */
  -10px 0 20px rgba(0, 0, 0, 0.05); /* Left shadow */
  overflow: hidden;
}

.form-container {
  position: absolute;
  top: 0;
  height: 100%;
  transition: all 0.6s ease-in-out;
}

.sign-in-container {
  left: 0;
  width: 50%;
  z-index: 2;
}

.sign-up-container {
  left: 0;
  width: 50%;
  opacity: 0;
  z-index: 1;
}

.container.right-panel-active .sign-in-container {
  transform: translateX(100%);
}

.container.right-panel-active .sign-up-container {
  transform: translateX(100%);
  opacity: 1;
  z-index: 5;
  animation: show 0.6s;
}

@keyframes show {
  0%, 49.99% {
    opacity: 0;
    z-index: 1;
  }
  50%, 100% {
    opacity: 1;
    z-index: 5;
  }
}

.overlay-container {
  position: absolute;
  top: 0;
  left: 50%;
  width: 50%;
  height: 100%;
  overflow: hidden;
  transition: transform 0.6s ease-in-out;
  z-index: 100;
}

.container.right-panel-active .overlay-container {
  transform: translateX(-100%);
}

.overlay {
  background: linear-gradient(135deg, #ff7b00, #ff9500);
  background-repeat: no-repeat;
  background-size: cover;
  background-position: 0 0;
  color: white;
  position: relative;
  left: -100%;
  height: 100%;
  width: 200%;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
}

.container.right-panel-active .overlay {
  transform: translateX(50%);
}

.overlay-panel {
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  padding: 0 40px;
  text-align: center;
  top: 0;
  height: 100%;
  width: 50%;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
}

.overlay-left {
  transform: translateX(-20%);
}

.container.right-panel-active .overlay-left {
  transform: translateX(0);
}

.overlay-right {
  right: 0;
  transform: translateX(0);
}

.container.right-panel-active .overlay-right {
  transform: translateX(20%);
}

.form {
  background-color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  padding: 0 50px;
  height: 100%;
  text-align: center;
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

.form h1 {
  font-weight: bold;
  margin: 0;
  color: #333;
  font-size: 28px;
  margin-bottom: 20px;
}

.input-v {
    width: 100%;
    min-height: 40px;
    border-radius: 4px;
    margin-bottom: 12px;
}

.project-name {
  font-size: 16px;
  color: #666;
  margin-bottom: 10px;
  font-weight: 500;
}

.font-title {
    text-align: center;
    font-size: 1.875rem;
    line-height: 2.25rem;
    font-weight: 700;
    margin-bottom: 2rem;
    color: #1f2937;
    font-family: 'Georgia', serif;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
}

/* 自定义 Ant Design 组件样式 */
.custom-input {
  margin: 8px 0;
  border-radius: 8px !important;
}

.custom-input .ant-input {
  background-color: #f8f9fa;
  border: none;
  padding: 12px 15px;
  font-size: 14px;
  border-radius: 8px;
}

.custom-input .ant-input:focus {
  background-color: #e9ecef;
  box-shadow: 0 0 0 2px rgba(255, 123, 0, 0.2);
}

.custom-input .ant-input-password {
  background-color: #f8f9fa;
  border: none;
  border-radius: 8px;
}

.custom-input .ant-input-password .ant-input {
  background-color: transparent;
}

.custom-button {
  border-radius: 20px !important;
  border: 1px solid #ff7b00 !important;
  background-color: #ff7b00 !important;
  color: white !important;
  font-size: 12px;
  font-weight: bold;
  padding: 12px 45px;
  letter-spacing: 1px;
  text-transform: uppercase;
  transition: all 0.3s ease;
  margin-top: 20px;
  height: auto !important;
}

.custom-button:hover {
  background-color: #e66a00 !important;
  border-color: #e66a00 !important;
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(255, 123, 0, 0.3) !important;
}

.ghost {
  background-color: transparent !important;
  border-color: white !important;
  color: white !important;
  border-radius: 20px !important;
  font-size: 12px;
  font-weight: bold;
  padding: 12px 45px;
  letter-spacing: 1px;
  text-transform: uppercase;
  transition: all 0.3s ease;
  height: auto !important;
}

.ghost:hover {
  background-color: white !important;
  color: #ff7b00 !important;
  border-color: white !important;
}

.overlay h1 {
  font-weight: bold;
  margin: 0;
  font-size: 24px;
  margin-bottom: 20px;
}

.overlay p {
  font-size: 14px;
  font-weight: 100;
  line-height: 20px;
  letter-spacing: 0.5px;
  margin: 20px 0 30px;
}

/* 隐藏 Ant Design 表单项的默认边距 */
.ant-form-item {
  margin-bottom: 0 !important;
}

@keyframes ripple {
  to {
    transform: scale(4);
    opacity: 0;
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .container {
    width: 90%;
    height: 600px;
    flex-direction: column;
  }
  
  .form-container,
  .overlay-container {
    width: 100%;
    height: 50%;
  }
  
  .sign-in-container,
  .sign-up-container {
    width: 100%;
  }
  
  .overlay-container {
    left: 0;
    top: 50%;
  }
  
  .container.right-panel-active .sign-in-container {
    transform: translateY(-100%);
  }
  
  .container.right-panel-active .sign-up-container {
    transform: translateY(-100%);
  }
  
  .container.right-panel-active .overlay-container {
    transform: translateY(-100%);
  }
}
</style>
