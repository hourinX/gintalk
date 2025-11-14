<script setup>
import { defineProps, ref, defineEmits } from 'vue'
import { LeftOutlined, CloseOutlined } from '@ant-design/icons-vue'

const props = defineProps({
  show: Boolean
})

const emit = defineEmits(['update:show'])

const closeOverlay = () => {
  emit('update:show', false)
}

const albums = ref([
  {
    name: 'a子相册',
    photos: [
      'https://picsum.photos/100/100?random=1',
      'https://picsum.photos/100/100?random=2',
      'https://picsum.photos/100/100?random=3',
      'https://picsum.photos/100/100?random=4',
      'https://picsum.photos/100/100?random=5',
      'https://picsum.photos/100/100?random=6',
      'https://picsum.photos/100/100?random=7',
      'https://picsum.photos/100/100?random=8',
      'https://picsum.photos/100/100?random=9',
      'https://picsum.photos/100/100?random=10'
    ]
  },
  {
    name: 'b子相册',
    photos: [
      'https://picsum.photos/100/100?random=11',
      'https://picsum.photos/100/100?random=12',
      'https://picsum.photos/100/100?random=13'
    ]
  }
])
</script>

<template>
  <div v-if="show" class="group-overlay">
    <div class="overlay-inner">
      <span class="back-btn" @click="closeOverlay"><LeftOutlined /></span>
      <span class="title">群相册</span>
      <span class="close-btn" @click="closeOverlay"><CloseOutlined /></span>
    </div>

    <div class="content">
        <div
    class="sub-album"
    v-for="(album, index) in albums"
    :key="index"
  >
    <div class="sub-album-header">
      <span class="album-name">{{ album.name }}</span>
      <span class="album-count">{{ album.photos.length }} 张</span>
      <span class="arrow">&gt;</span>
    </div>

    <div class="photo-grid">
      <div v-for="(photo, i) in album.photos.slice(0, 9)" :key="i" class="photo-item">
        <img :src="photo" />
        <div v-if="i === 8 && album.photos.length > 9" class="overlay-plus">
          +{{ album.photos.length - 9 }}
        </div>
      </div>
    </div>
  </div>
    </div>

  </div>
</template>

<style scoped>
.group-overlay {
  position: absolute;
  top: 0;
  right: 0;
  width: 100%;
  height: auto;
  z-index: 999;
  border: 1px solid #f0f0f0;
  border-radius: 8px;
}

.overlay-inner {
  position: relative;
  padding: 24px;
  height: 50px;
  color: #000;
  background-color: aliceblue;
  font-size: 18px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  align-content: center;
}

.content {
  padding: 12px;
  display: flex;
  flex-direction: column;
}

.back-btn {
  position: absolute;
  top: 8px;
  left: 8px;
  cursor: pointer;
  font-weight: bold;
  font-size: 20px;
}

.close-btn {
  position: absolute;
  top: 8px;
  right: 8px;
  cursor: pointer;
  font-weight: bold;
  font-size: 20px;
}

.title {
  position: absolute;
  top: 8px;
  left: 44%;
  font-weight: bold;
  font-size: 20px;
  text-align: center;
}

.sub-album {
  margin-bottom: 24px;
}

.sub-album-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 4px 0;
  font-weight: bold;
  font-size: 16px;
  cursor: pointer;
}

.album-name {
  color: #333;
}

.album-count {
  color: #999;
  margin-left: 8px;
}

.arrow {
  color: #999;
  margin-left: auto;
}

.photo-grid {
  display: grid;
  grid-template-columns: repeat(3, 100px);
  grid-auto-rows: 100px;
  gap: 8px;
  grid-gap: 4px;
  margin-top: 8px;
}

.photo-item {
  position: relative;
  width: 100px;
  padding-top: 100px;
  overflow: hidden;
  border-radius: 4px;
}

.photo-item img {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.overlay-plus {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0,0.3);
  color: #fff;
  font-size: 24px;
  font-weight: bold;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 4px;
}
</style>
