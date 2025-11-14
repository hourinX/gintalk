<script setup>
import { ref, computed, watch } from 'vue'
import { Progress, Avatar } from 'ant-design-vue'
import { size } from 'lodash'
import ImagePreview from '@/components/image-preview/index.vue'
import GroupAlbum from './group-album.vue'

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
    'https://picsum.photos/100/100?random=3',
    'https://picsum.photos/100/100?random=4',
    'https://picsum.photos/100/100?random=5'
  ]
})

const showOverlay = ref(false)
const toggleOverlay = () => {
  showOverlay.value = !showOverlay.value
}

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

const clickPhoto = () => {
  console.log('点击了群相册图片')
  showOverlay.value = true
}

const genderCircleColor = computed(() => {
  const male = genderStats.value.malePercent
  const female = 100 - male

  return {
    '0%': '#1890ff',
    [male + '%']: '#1890ff',

    [(male + 0.01) + '%']: '#eb2f96',
    '100%': '#eb2f96'
  }
})

</script>

<template>
  <div v-if="showOverlay" class="group-overlay" @click="toggleOverlay">
    <div class="overlay-inner">
      <GroupAlbum v-model:show="showOverlay" @click.stop></GroupAlbum>
    </div>
  </div>
  
  <!-- 群基本信息 -->
  <div class="group-header">
    <a-avatar :src="group.avatar" :size="80" shape="square" @click="showPreview(group.avatar)" />
      <div class="group-info">
        <h3>{{ group.name }}</h3>
        <p><b>群代码：</b>{{ group.code }}</p>
        <!-- <p><b>群人数：</b>{{ memberCount }}</p> -->
      </div>
  </div>

  <div class="progress-section">
    <!-- 性别比例 -->
    <div class="gender-section">
      <div class="gender-circle-wrapper">
        <a-progress
          type="circle"
          :percent="100"
          :width="120"
          :stroke-width="10"
          :stroke-color="genderCircleColor"
        >
          <template #format="性别比例">
            <span style="color: #c4b7d6;font-size: 16px;font-weight: bold;">{{ "性别比例" }}</span>
          </template>
        </a-progress>
      </div>

      <div class="gender-legend">
        <span class="male">男 {{ genderStats.malePercent }}%</span>
        <span class="female">女 {{ 100 - genderStats.malePercent }}%</span>
      </div>
    </div>

    <!-- 性别比例 -->
    <div class="gender-section" style="margin-left: 60px;">
      <div class="gender-circle-wrapper">
        <a-progress
          type="circle"
          :percent="100"
          :width="120"
          :stroke-width="10"
          :stroke-color="genderCircleColor"
        >
          <template #format="性别比例">
            <span style="color: #c4b7d6;font-size: 16px;font-weight: bold;">{{ "性别比例" }}</span>
          </template>
        </a-progress>
      </div>

      <div class="gender-legend">
        <span class="male">男 {{ genderStats.malePercent }}%</span>
        <span class="female">女 {{ 100 - genderStats.malePercent }}%</span>
      </div>
    </div>

    <!-- 性别比例 -->
    <div class="gender-section" style="margin-left: 60px;">
      <div class="gender-circle-wrapper">
        <a-progress
          type="circle"
          :percent="100"
          :width="120"
          :stroke-width="10"
          :stroke-color="genderCircleColor"
        >
          <template #format="性别比例">
            <span style="color: #c4b7d6;font-size: 16px;font-weight: bold;">{{ "性别比例" }}</span>
          </template>
        </a-progress>
      </div>

      <div class="gender-legend">
        <span class="male">男 {{ genderStats.malePercent }}%</span>
        <span class="female">女 {{ 100 - genderStats.malePercent }}%</span>
      </div>
    </div>
  </div>
  

  <!-- 群相册 -->
  <div class="photo-section">
    <p><b>群相册</b></p>
    <div class="photo-list">
      <img v-for="(photo, index) in group.photos" :key="index" :src="photo" @click="clickPhoto" />
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

  <div class="group-intro">
    <p><b>群简介</b></p>
    <div class="text-intro">
      <text>{{ group.intro || '暂无群简介...' }}</text>
    </div>
  </div>

  <ImagePreview v-model:visible="previewVisible" :src="previewSrc" />
</template>

<style scoped>
.group-overlay {
  position: absolute;
  top: 27px;
  right: 0;
  width: 48%;
  height: auto;
  z-index: 999;
  border-radius: 8px;
}

.overlay-inner {
  padding: 16px;
  color: #fff;
  font-size: 18px;
}

.group-header {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
}

.group-info {
  margin-left: 12px;
}

.progress-section {
  margin-top: 16px;
  display: flex;
  flex-direction: row;
}

.gender-section {
  text-align: center;
}

.gender-circle-wrapper {
  display: flex;
  justify-content: center;
  margin: 12px 0;
}

.gender-legend {
  margin-top: 10px;
  display: flex;
  justify-content: center;
  gap: 10px;
  font-weight: bold;
}

.gender-legend .male {
  color: #1890ff;
}

.gender-legend .female {
  color: #eb2f96;
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
  margin-top: 8px;
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

.group-intro {
  margin-top: 16px;
}

.text-intro {
  margin-top: 8px;
}
</style>