<script setup>
import { ref, computed } from 'vue'

// 模拟联系人数据
const contacts = ref([
  { id: 1, name: 'Alice', phone: '111111111', email: 'alice@example.com', avatar: 'https://i.pravatar.cc/40?img=1' },
  { id: 2, name: 'Bob', phone: '222222222', email: 'bob@example.com', avatar: 'https://i.pravatar.cc/40?img=2' },
  { id: 3, name: 'Charlie', phone: '333333333', email: 'charlie@example.com', avatar: 'https://i.pravatar.cc/40?img=3' },
  { id: 4, name: 'David', phone: '444444444', email: 'david@example.com', avatar: 'https://i.pravatar.cc/40?img=4' },
  { id: 5, name: 'Eve', phone: '555555555', email: 'eve@example.com', avatar: 'https://i.pravatar.cc/40?img=5' },
  { id: 6, name: 'Frank', phone: '666666666', email: 'frank@example.com', avatar: 'https://i.pravatar.cc/40?img=6' },
  { id: 7, name: 'Grace', phone: '777777777', email: 'grace@example.com', avatar: 'https://i.pravatar.cc/40?img=7' },
  { id: 8, name: 'Helen', phone: '888888888', email: 'helen@example.com', avatar: 'https://i.pravatar.cc/40?img=8' },
  { id: 9, name: 'Ivan', phone: '999999999', email: 'ivan@example.com', avatar: 'https://i.pravatar.cc/40?img=9' },
  { id: 10, name: 'Jack', phone: '101010101', email: 'jack@example.com', avatar: 'https://i.pravatar.cc/40?img=10' },
  { id: 11, name: 'Kate', phone: '111111112', email: 'kate@example.com', avatar: 'https://i.pravatar.cc/40?img=11' },
  { id: 12, name: 'Leo', phone: '121212121', email: 'leo@example.com', avatar: 'https://i.pravatar.cc/40?img=12' },
  { id: 13, name: 'Mary', phone: '131313131', email: 'mary@example.com', avatar: 'https://i.pravatar.cc/40?img=13' },
  { id: 14, name: 'Nancy', phone: '141414141', email: 'nancy@example.com', avatar: 'https://i.pravatar.cc/40?img=14' },
  { id: 15, name: 'Oscar', phone: '151515151', email: 'oscar@example.com', avatar: 'https://i.pravatar.cc/40?img=15' },
  { id: 16, name: 'Paul', phone: '161616161', email: 'paul@example.com', avatar: 'https://i.pravatar.cc/40?img=16' },
  { id: 17, name: 'Queen', phone: '171717171', email: 'queen@example.com', avatar: 'https://i.pravatar.cc/40?img=17' },
  { id: 18, name: 'Robert', phone: '181818181', email: 'robert@example.com', avatar: 'https://i.pravatar.cc/40?img=18' },
  { id: 19, name: 'Steve', phone: '191919191', email: 'steve@example.com', avatar: 'https://i.pravatar.cc/40?img=19' },
  { id: 20, name: 'Tom', phone: '202020202', email: 'tom@example.com', avatar: 'https://i.pravatar.cc/40?img=20' },
  { id: 21, name: 'Uma', phone: '212121212', email: 'uma@example.com', avatar: 'https://i.pravatar.cc/40?img=21' },
  { id: 22, name: 'Victor', phone: '222222222', email: 'victor@example.com', avatar: 'https://i.pravatar.cc/40?img=22' },
  { id: 23, name: 'Wendy', phone: '232323232', email: 'wendy@example.com', avatar: 'https://i.pravatar.cc/40?img=23' },
  { id: 24, name: 'Xavier', phone: '242424242', email: 'xavier@example.com', avatar: 'https://i.pravatar.cc/40?img=24' },
  { id: 25, name: 'Yvonne', phone: '252525252', email: 'yvonne@example.com', avatar: 'https://i.pravatar.cc/40?img=25' },
  { id: 26, name: 'Zack', phone: '262626262', email: 'zack@example.com', avatar: 'https://i.pravatar.cc/40?img=26' },
  { id: 27, name: 'Aaron', phone: '272727272', email: 'aaron@example.com', avatar: 'https://i.pravatar.cc/40?img=27' },
  { id: 28, name: 'Bella', phone: '282828282', email: 'bella@example.com', avatar: 'https://i.pravatar.cc/40?img=28' },
  { id: 29, name: 'Chris', phone: '292929292', email: 'chris@example.com', avatar: 'https://i.pravatar.cc/40?img=29' },
  { id: 30, name: 'Diana', phone: '303030303', email: 'diana@example.com', avatar: 'https://i.pravatar.cc/40?img=30' }
])


// 字母索引
const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('')

const selectedFriend = ref(null)

const showContacts = ref(true)

const selectFriend = (friend) => {
  selectedFriend.value = friend
}

const toggleContacts = () => {
  showContacts.value = !showContacts.value
}

const showGroups = ref(false)

const toggleGroups = () => {
  showGroups.value = !showGroups.value
}

const selectedLetter = ref(null)

const selectLetter = (letter) => {
  if (selectedLetter.value === letter) {
    selectedLetter.value = null
  } else {
    selectedLetter.value = letter
  }
}

const filteredContacts = computed(() => {
  if (!selectedLetter.value) return contacts.value
  return contacts.value.filter(c => c.name[0].toUpperCase() === selectedLetter.value)
})
</script>

<template>
  <a-layout class="relations-layout">
    <a-layout-sider width="260" class="relations-sider">
      <div class="relations-groups">
        <a-card size="small" :body-style="{ padding: 0 }">
            <div class="card-title" @click="toggleGroups">
            <span>群组</span>
            <span>{{ showGroups ? '▲' : '▼' }}</span>
          </div>
          <div v-if="showGroups" class="groups-content">
            <p>这里以后放群组列表</p>
          </div>
        </a-card>
      </div>

      <div class="relations-contacts">
        <a-card size="small" :body-style="{ padding: 0 }">
          <div class="card-title" @click="toggleContacts">
            <span>联系人</span>
            <span>{{ showContacts ? '▲' : '▼' }}</span>
          </div>

          <div v-if="showContacts" class="contacts-list">
            <a-list :data-source="filteredContacts">
              <template #renderItem="{ item }">
                <a-list-item
                  class="contact-item"
                  @click="selectFriend(item)"
                  :class="{ active: selectedFriend && selectedFriend.id === item.id }"
                >
                    <a-avatar :src="item.avatar" style="margin-right: 10px" />
                  {{ item.name }}
                </a-list-item>
              </template>
            </a-list>
          </div>
        </a-card>
      </div>
    </a-layout-sider>

    <a-layout-content class="relations-content">
      <div class="alphabet-bar">
        <a-tag
          v-for="letter in alphabet"
          :key="letter"
          :color="selectedLetter === letter ? 'orange' : 'default'"
          class="letter"
          @click="selectLetter(letter)"
        >
          {{ letter }}
        </a-tag>
      </div>

      <a-card v-if="selectedFriend" title="好友信息" bordered>
        <p><b>姓名：</b>{{ selectedFriend.name }}</p>
        <p><b>电话：</b>{{ selectedFriend.phone }}</p>
        <p><b>Email：</b>{{ selectedFriend.email }}</p>
      </a-card>

      <a-card v-else bordered>
        <p>请选择一个联系人</p>
      </a-card>
    </a-layout-content>
  </a-layout>
</template>

<style scoped>
.relations-layout {
  background: #fff;
  height: 100%;
}

.relations-sider {
  background: #f7f7f7;
  padding: 10px;
  display: flex;
  flex-direction: column;
  height: 88vh;
  box-sizing: border-box;
}

.relations-groups,
.relations-contacts {
  margin-bottom: 10px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex: 1;
}

.relations-contacts .contacts-list {
  max-height: 77vh; 
  overflow-y: auto;
  background: #fff;
  padding: 0;
  flex: 1;
}

.relations-sider {
  /* background: #f7f7f7; */
  background: #fff;
  padding: 10px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.groups-section {
  margin-bottom: 20px;
}

.card-title {
  padding: 10px 15px;
  cursor: pointer;
  font-weight: 600;
  display: flex;
  justify-content: space-between;
  background: #fafafa;
  border-bottom: 1px solid #f0f0f0;
}

.contacts-list {
  max-height: 400px;
  overflow-y: auto;
}

.contact-item {
  cursor: pointer;
  padding: 10px 15px;
  transition: background 0.3s;
}

.contact-item:hover {
  background: #f0f0f0;
}

.contact-item.active {
  background: #fff0ec;
  color: #e15536;
  font-weight: 600;
}

.relations-content {
  padding: 20px;
  position: relative;
}

.alphabet-bar {
  display: flex;
  gap: 8px;
  margin-bottom: 15px;
  flex-wrap: wrap;
}

.letter {
  cursor: pointer;
  padding: 4px 6px;
  border-radius: 4px;
  transition: background 0.3s;
  font-size: 12px;
  color: #666;
}

.letter:hover {
  background: #f0f0f0;
  color: #000;
}

.friend-info {
  padding: 15px;
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  background: #fafafa;
}

.friend-placeholder {
  color: #999;
  padding: 20px;
}
</style>