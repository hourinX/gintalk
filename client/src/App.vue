<<<<<<< HEAD
<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

=======
>>>>>>> de58b6a0c79b5c88406fc6992919413a2249b927
<template>
  <div>
    <a href="https://vite.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>
<<<<<<< HEAD
  <HelloWorld msg="Vite + Vue" />
</template>

=======
  11111
  <HelloWorld msg="Vite + Vue" />
</template>

<script setup>
import HelloWorld from './components/HelloWorld.vue'
import {post} from "./utils/request.js"
import {onMounted} from "vue";

onMounted(async ()=>{
  console.log(111)
  await post("/user/register")
})
</script>

>>>>>>> de58b6a0c79b5c88406fc6992919413a2249b927
<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
