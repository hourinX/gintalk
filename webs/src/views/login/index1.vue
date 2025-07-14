<template>
    <div class="login-page">
      <div class="container" :class="{ 'right-panel-active': isSignUpMode }" ref="container">
        <!-- 登录表单 -->
        <div class="form-container sign-in-container">
          <form class="form" @submit.prevent="handleSignIn">
            <div class="project-name">MyProject</div>
            <h1>登录</h1>
            <div class="social-container">
              <a href="#" class="social" @click.prevent="handleSocialLogin('google')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none">
                  <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                  <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                  <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                  <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                </svg>
              </a>
              <a href="#" class="social" @click.prevent="handleSocialLogin('wechat')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="#1AAD19">
                  <path d="M8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188L8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188zM8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188L8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188z"/>
                  <path d="M12.48 2.188c2.382 0 4.761 0.762 6.762 2.286 2.001 1.524 3.429 3.667 4.048 6.095 0.619 2.428 0.381 4.952-0.667 7.19-1.048 2.238-2.857 4.095-5.143 5.286-2.286 1.191-4.905 1.667-7.429 1.357-2.524-0.31-4.857-1.381-6.619-3.048C1.667 19.619 0.619 17.286 0.31 14.762c-0.31-2.524 0.167-5.143 1.357-7.429C2.857 4.905 4.714 3.095 6.952 2.048 9.19 1 11.714 0.762 14.143 1.381c2.428 0.619 4.571 2.048 6.095 4.048"/>
                </svg>
              </a>
              <a href="#" class="social" @click.prevent="handleSocialLogin('other')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="#C71D23">
                  <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                </svg>
              </a>
              <a href="#" class="social" @click.prevent="handleSocialLogin('github')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="#333">
                  <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
                </svg>
              </a>
            </div>
            <input 
              type="email" 
              placeholder="邮箱" 
              v-model="signInForm.email"
              required
            />
            <input 
              type="password" 
              placeholder="密码" 
              v-model="signInForm.password"
              required
            />
            <button type="submit" @click="createRipple">登录</button>
          </form>
        </div>
  
        <!-- 注册表单 -->
        <div class="form-container sign-up-container">
          <form class="form" @submit.prevent="handleSignUp">
            <div class="project-name">MyProject</div>
            <h1>注册</h1>
            <div class="social-container">
              <a href="#" class="social" @click.prevent="handleSocialLogin('google')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none">
                  <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                  <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                  <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                  <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                </svg>
              </a>
              <a href="#" class="social" @click.prevent="handleSocialLogin('wechat')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="#1AAD19">
                  <path d="M8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188L8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188zM8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188L8.691 2.188C8.691 2.188 8.691 2.188 8.691 2.188z"/>
                  <path d="M12.48 2.188c2.382 0 4.761 0.762 6.762 2.286 2.001 1.524 3.429 3.667 4.048 6.095 0.619 2.428 0.381 4.952-0.667 7.19-1.048 2.238-2.857 4.095-5.143 5.286-2.286 1.191-4.905 1.667-7.429 1.357-2.524-0.31-4.857-1.381-6.619-3.048C1.667 19.619 0.619 17.286 0.31 14.762c-0.31-2.524 0.167-5.143 1.357-7.429C2.857 4.905 4.714 3.095 6.952 2.048 9.19 1 11.714 0.762 14.143 1.381c2.428 0.619 4.571 2.048 6.095 4.048"/>
                </svg>
              </a>
              <a href="#" class="social" @click.prevent="handleSocialLogin('other')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="#C71D23">
                  <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                </svg>
              </a>
              <a href="#" class="social" @click.prevent="handleSocialLogin('github')">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="#333">
                  <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
                </svg>
              </a>
            </div>
            <input 
              type="text" 
              placeholder="姓名" 
              v-model="signUpForm.name"
              required
            />
            <input 
              type="email" 
              placeholder="邮箱" 
              v-model="signUpForm.email"
              required
            />
            <input 
              type="password" 
              placeholder="密码" 
              v-model="signUpForm.password"
              required
            />
            <button type="submit" @click="createRipple">注册</button>
          </form>
        </div>
  
        <!-- 覆盖层 -->
        <div class="overlay-container">
          <div class="overlay">
            <div class="overlay-panel overlay-left">
              <h1>欢迎回来！</h1>
              <p>要与我们保持联系，请使用您的个人信息登录</p>
              <button class="ghost" @click="switchToSignIn">登录</button>
            </div>
            <div class="overlay-panel overlay-right">
              <h1>还没账号？</h1>
              <p>输入您的个人详细信息，开始您的旅程</p>
              <button class="ghost" @click="switchToSignUp">注册</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, reactive, onMounted } from 'vue'
  
  // 响应式数据
  const isSignUpMode = ref(false)
  const container = ref(null)
  
  // 表单数据
  const signInForm = reactive({
    email: '',
    password: ''
  })
  
  const signUpForm = reactive({
    name: '',
    email: '',
    password: ''
  })
  
  // 方法
  const switchToSignUp = () => {
    isSignUpMode.value = true
  }
  
  const switchToSignIn = () => {
    isSignUpMode.value = false
  }
  
  const handleSignIn = () => {
    console.log('登录表单提交:', signInForm)
    // 这里可以添加登录逻辑
    alert(`登录信息: ${signInForm.email}`)
  }
  
  const handleSignUp = () => {
    console.log('注册表单提交:', signUpForm)
    // 这里可以添加注册逻辑
    alert(`注册信息: ${signUpForm.name} - ${signUpForm.email}`)
  }
  
  const handleSocialLogin = (provider) => {
    console.log(`使用 ${provider} 登录`)
    // 这里可以添加第三方登录逻辑
    alert(`使用 ${provider} 登录`)
  }
  
  // 创建涟漪效果
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
  
  // 组件挂载后的操作
  onMounted(() => {
    // 为社交登录按钮添加悬停效果
    const socialButtons = document.querySelectorAll('.social')
    socialButtons.forEach(social => {
      social.addEventListener('mouseenter', function() {
        this.style.transform = 'translateY(-3px) scale(1.1)'
      })
      
      social.addEventListener('mouseleave', function() {
        this.style.transform = 'translateY(0) scale(1)'
      })
    })
  })
  </script>
  
  <style scoped>
  .login-page {
    font-family: 'Arial', sans-serif;
    background-color: white;
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
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
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
  
  .form h1 {
    font-weight: bold;
    margin: 0;
    color: #333;
    font-size: 28px;
    margin-bottom: 20px;
  }
  
  .project-name {
    font-size: 16px;
    color: #666;
    margin-bottom: 10px;
    font-weight: 500;
  }
  
  .social-container {
    margin: 20px 0;
  }
  
  .social-container a {
    border: 1px solid #ddd;
    border-radius: 50%;
    display: inline-flex;
    justify-content: center;
    align-items: center;
    margin: 0 5px;
    height: 40px;
    width: 40px;
    text-decoration: none;
    transition: all 0.3s ease;
  }
  
  .social-container a:hover {
    background-color: #f5f5f5;
    transform: translateY(-2px);
  }
  
  .social-container img {
    width: 20px;
    height: 20px;
  }
  
  .form input {
    background-color: #f8f9fa;
    border: none;
    padding: 12px 15px;
    margin: 8px 0;
    width: 100%;
    border-radius: 8px;
    font-size: 14px;
    outline: none;
    transition: all 0.3s ease;
  }
  
  .form input:focus {
    background-color: #e9ecef;
    box-shadow: 0 0 0 2px rgba(255, 123, 0, 0.2);
  }
  
  button {
    border-radius: 20px;
    border: 1px solid #ff7b00;
    background-color: #ff7b00;
    color: white;
    font-size: 12px;
    font-weight: bold;
    padding: 12px 45px;
    letter-spacing: 1px;
    text-transform: uppercase;
    transition: all 0.3s ease;
    cursor: pointer;
    margin-top: 20px;
  }
  
  button:hover {
    background-color: #e66a00;
    transform: translateY(-2px);
    box-shadow: 0 5px 15px rgba(255, 123, 0, 0.3);
  }
  
  button.ghost {
    background-color: transparent;
    border-color: white;
    color: white;
  }
  
  button.ghost:hover {
    background-color: white;
    color: #ff7b00;
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
  
  