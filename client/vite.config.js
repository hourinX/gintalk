import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
<<<<<<< HEAD
=======
import path from "path"
>>>>>>> de58b6a0c79b5c88406fc6992919413a2249b927

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
<<<<<<< HEAD
=======
  resolve:{
    alias:{
      '@': path.resolve(__dirname, './src'),
    }
  }
>>>>>>> de58b6a0c79b5c88406fc6992919413a2249b927
})
