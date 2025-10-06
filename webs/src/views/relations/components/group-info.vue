<script setup>
import { ref, computed, watch } from 'vue'
import { Progress, Avatar } from 'ant-design-vue'
import { size } from 'lodash'
import ImagePreview from '@/components/image-preview/index.vue'

const props = defineProps({
  id: {
    type: [String, Number],
    required: true
  }
})

watch(() => props.id, (newId) => {
  if (newId) {
    console.log('收到 id:', newId, typeof props.id)
  }
})

// 假数据（真实场景下从父组件传进来）
const group = ref({
  avatar: 'https://i.pravatar.cc/80?img=50',
  name: '前端开发交流群',
  code: 'G123456789',
  members: [
    { id: 1, name: 'Alice', gender: 'female', avatar: 'https://i.pravatar.cc/40?img=1' },
    { id: 2, name: 'Bob', gender: 'male', avatar: 'https://i.pravatar.cc/40?img=2' },
    { id: 3, name: 'Charlie', gender: 'male', avatar: 'https://i.pravatar.cc/40?img=3' },
    { id: 4, name: 'Diana', gender: 'female', avatar: 'https://i.pravatar.cc/40?img=4' },
    { id: 5, name: 'Eve', gender: 'female', avatar: 'https://i.pravatar.cc/40?img=5' },
    { id: 6, name: 'Frank', gender: 'male', avatar: 'https://i.pravatar.cc/40?img=6' }
  ],
  photos: [
    'https://picsum.photos/100/100?random=1',
    'https://picsum.photos/100/100?random=2',
    'https://picsum.photos/100/100?random=3'
  ]
})

// 成员数
const memberCount = computed(() => group.value.members.length)

// 计算男女比例
const genderStats = computed(() => {
  const males = group.value.members.filter(m => m.gender === 'male').length
  const females = group.value.members.filter(m => m.gender === 'female').length
  const total = males + females
  return {
    malePercent: total ? Math.round((males / total) * 100) : 0,
    femalePercent: total ? Math.round((females / total) * 100) : 0
  }
})

// 只显示前 5 个成员
const displayMembers = computed(() => {
  if (group.value.members.length <= 5) return group.value.members
  return group.value.members.slice(0, 5)
})

// 是否还有更多成员
const hasMoreMembers = computed(() => group.value.members.length > 5)


const previewVisible = ref(false)
const previewSrc = ref('')

// 点击头像时触发
const showPreview = (src) => {
  previewSrc.value = src
  previewVisible.value = true
}
</script>

<template>
  <!-- 群基本信息 -->
  <div class="group-header">
    <a-avatar :src="group.avatar" :size="80" shape="square" @click="showPreview(group.avatar)" />
      <div class="group-info">
        <h3>{{ group.name }}</h3>
        <p><b>群代码：</b>{{ group.code }}</p>
        <!-- <p><b>群人数：</b>{{ memberCount }}</p> -->
      </div>
  </div>

  <!-- 性别比例 -->
  <div class="gender-section">
    <p><b>性别比例</b></p>
     <div class="gender-progress">
      <span>男：{{ genderStats.malePercent }}%</span>
      <a-progress :percent="genderStats.malePercent" status="active" />
      <span>女：{{ genderStats.femalePercent }}%</span>
      <a-progress :percent="genderStats.femalePercent" status="active" stroke-color="#eb2f96" />
    </div>
  </div>

  <!-- 群相册 -->
  <div class="photo-section">
    <p><b>群相册</b></p>
    <div class="photo-list">
      <img v-for="(photo, index) in group.photos" :key="index" :src="photo" />
    </div>
  </div>

  <!-- 群成员 -->
  <div class="member-section">
    <p><b>群成员</b></p>
    <div class="member-list">
      <a-avatar
          v-for="m in displayMembers"
          :key="m.id"
          :src="m.avatar"
          :title="m.name"
          :size="50"
      />
      <span v-if="hasMoreMembers">...</span>
    </div>
  </div>

  <ImagePreview v-model:visible="previewVisible" :src="previewSrc" />
</template>

<style scoped>
.group-header {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
}
.group-info {
  margin-left: 12px;
}
.gender-section {
  margin-top: 16px;
}
.gender-progress {
  display: flex;
  flex-direction: column;
  gap: 4px;
}
.photo-section {
  margin-top: 16px;
}
.photo-list {
  display: flex;
  gap: 8px;
}
.photo-list img {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: 8px;
}
.member-section {
  margin-top: 16px;
}
.member-list {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 10px;
}
</style>