<script setup>
import { ref, watch, computed } from 'vue'

const props = defineProps({
  src: { type: String, required: true },
  visible: { type: Boolean, default: false }
})

const emit = defineEmits(['update:visible'])

const scale = ref(1)
const cursorStyle = ref('default') // 初始指针样式

watch(() => props.visible, (val) => {
  if (!val) {
    scale.value = 1
    cursorStyle.value = 'default'
    isMouseOutside.value = false
  } else {
    isMouseOutside.value = true
  }
})

const handleClose = () => emit('update:visible', false)

const handleWheel = (e) => {
  e.preventDefault()
  if (e.deltaY < 0) {
    scale.value = Math.min(scale.value + 0.1, 5)
    cursorStyle.value = 'zoom-in'
  } else {
    scale.value = Math.max(scale.value - 0.1, 0.3)
    cursorStyle.value = 'zoom-out'
  }
}

const imgWidth = ref(0)
const imgHeight = ref(0)
const isMouseOutside = ref(false)

const handleMouseEnter = () => {
  cursorStyle.value = 'default'
  isMouseOutside.value = false
}
const handleMouseLeave = () => {
  cursorStyle.value = 'default'
  isMouseOutside.value = true
}

const imgStyle = computed(() => ({
  transform: `scale(${scale.value})`,
  transition: 'transform 0.12s ease',
  maxWidth: '90%',
  maxHeight: '90%',
  objectFit: 'contain',
  userSelect: 'none',
  WebkitUserDrag: 'none',
  cursor: cursorStyle.value
}))
</script>

<template>
  <div
    v-if="props.visible"
    class="preview-mask"
    @click.self="handleClose"
  >
    <img
      :src="props.src"
      :style="imgStyle"
      draggable="false"
      @wheel="handleWheel"
      @mouseenter="handleMouseEnter"
      @mouseleave="handleMouseLeave"
    />
    <div v-if="isMouseOutside" class="close-tip">点击任意位置关闭图片</div>
  </div>
</template>

<style scoped>
.preview-mask {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2000;
  cursor: pointer; /* 背景显示可点击关闭 */
}

.close-tip {
  position: absolute;
  bottom: 30px;
  width: 100%;
  text-align: center;
  color: rgba(255, 255, 255, 0.8);
  font-size: 14px;
  pointer-events: none; /* 不阻挡点击事件 */
}
</style>
