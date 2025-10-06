<script setup>
import { ref, computed, watch } from 'vue'
import GroupInfo from './components/group-info.vue'
import userInfo from './components/user-info.vue'

// 模拟联系人数据
const contacts = ref([
  { id: 1, name: 'Alice', phone: '111111111', email: 'alice@example.com', avatar: 'https://i.pravatar.cc/40?img=1', usercode: '123456789' },
  { id: 2, name: 'Bob', phone: '222222222', email: 'bob@example.com', avatar: 'https://i.pravatar.cc/40?img=2', usercode: '223456789' },
  { id: 3, name: 'Charlie', phone: '333333333', email: 'charlie@example.com', avatar: 'https://i.pravatar.cc/40?img=3', usercode: '323456789' },
  { id: 4, name: 'David', phone: '444444444', email: 'david@example.com', avatar: 'https://i.pravatar.cc/40?img=4', usercode: '423456789' },
  { id: 5, name: 'Eve', phone: '555555555', email: 'eve@example.com', avatar: 'https://i.pravatar.cc/40?img=5', usercode: '523456789' },
  { id: 6, name: 'Frank', phone: '666666666', email: 'frank@example.com', avatar: 'https://i.pravatar.cc/40?img=6', usercode: '623456789' },
  { id: 7, name: 'Grace', phone: '777777777', email: 'grace@example.com', avatar: 'https://i.pravatar.cc/40?img=7', usercode: '723456789' },
  { id: 8, name: 'Helen', phone: '888888888', email: 'helen@example.com', avatar: 'https://i.pravatar.cc/40?img=8', usercode: '823456789' },
  { id: 9, name: 'Ivan', phone: '999999999', email: 'ivan@example.com', avatar: 'https://i.pravatar.cc/40?img=9', usercode: '923456789' },
  { id: 10, name: 'Jack', phone: '101010101', email: 'jack@example.com', avatar: 'https://i.pravatar.cc/40?img=10', usercode: '103456789' },
  { id: 11, name: 'Kate', phone: '111111112', email: 'kate@example.com', avatar: 'https://i.pravatar.cc/40?img=11', usercode: '113456789' },
  { id: 12, name: 'Leo', phone: '121212121', email: 'leo@example.com', avatar: 'https://i.pravatar.cc/40?img=12', usercode: '123456780' },
  { id: 13, name: 'Mary', phone: '131313131', email: 'mary@example.com', avatar: 'https://i.pravatar.cc/40?img=13', usercode: '133456789' },
  { id: 14, name: 'Oscar', phone: '151515151', email: 'oscar@example.com', avatar: 'https://i.pravatar.cc/40?img=15', usercode: '153456789' },
  { id: 15, name: 'Paul', phone: '161616161', email: 'paul@example.com', avatar: 'https://i.pravatar.cc/40?img=16', usercode: '163456789' },
  { id: 16, name: 'Queen', phone: '171717171', email: 'queen@example.com', avatar: 'https://i.pravatar.cc/40?img=17', usercode: '173456789' },
  { id: 17, name: 'Robert', phone: '181818181', email: 'robert@example.com', avatar: 'https://i.pravatar.cc/40?img=18', usercode: '183456789' },
  { id: 18, name: 'Steve', phone: '191919191', email: 'steve@example.com', avatar: 'https://i.pravatar.cc/40?img=19', usercode: '193456789' },
  { id: 19, name: 'Tom', phone: '202020202', email: 'tom@example.com', avatar: 'https://i.pravatar.cc/40?img=20', usercode: '203456789' },
  { id: 20, name: 'Uma', phone: '212121212', email: 'uma@example.com', avatar: 'https://i.pravatar.cc/40?img=21', usercode: '213456789' },
  { id: 21, name: 'Victor', phone: '222222222', email: 'victor@example.com', avatar: 'https://i.pravatar.cc/40?img=22', usercode: '223456780' },
  { id: 22, name: 'Wendy', phone: '232323232', email: 'wendy@example.com', avatar: 'https://i.pravatar.cc/40?img=23', usercode: '233456789' },
  { id: 23, name: 'Xavier', phone: '242424242', email: 'xavier@example.com', avatar: 'https://i.pravatar.cc/40?img=24', usercode: '243456789' },
  { id: 24, name: 'Yvonne', phone: '252525252', email: 'yvonne@example.com', avatar: 'https://i.pravatar.cc/40?img=25', usercode: '253456789' },
  { id: 25, name: 'Zack', phone: '262626262', email: 'zack@example.com', avatar: 'https://i.pravatar.cc/40?img=26', usercode: '263456789' },
  { id: 26, name: 'Aaron', phone: '272727272', email: 'aaron@example.com', avatar: 'https://i.pravatar.cc/40?img=27', usercode: '273456789' },
  { id: 27, name: 'Bella', phone: '282828282', email: 'bella@example.com', avatar: 'https://i.pravatar.cc/40?img=28', usercode: '283456789' },
  { id: 28, name: 'Chris', phone: '292929292', email: 'chris@example.com', avatar: 'https://i.pravatar.cc/40?img=29', usercode: '293456789' },
  { id: 29, name: 'Diana', phone: '303030303', email: 'diana@example.com', avatar: 'https://i.pravatar.cc/40?img=30', usercode: '303456789' },
  { id: 30, name: 'Ethan', phone: '313131313', email: 'ethan@example.com', avatar: 'https://i.pravatar.cc/40?img=31', usercode: '313456789' },
  { id: 31, name: 'Fiona', phone: '323232323', email: 'fiona@example.com', avatar: 'https://i.pravatar.cc/40?img=32', usercode: '323456780' },
  { id: 32, name: 'George', phone: '333333334', email: 'george@example.com', avatar: 'https://i.pravatar.cc/40?img=33', usercode: '333456789' },
  { id: 33, name: 'Hannah', phone: '343434343', email: 'hannah@example.com', avatar: 'https://i.pravatar.cc/40?img=34', usercode: '343456789' },
  { id: 34, name: 'Ian', phone: '353535353', email: 'ian@example.com', avatar: 'https://i.pravatar.cc/40?img=35', usercode: '353456789' },
  { id: 35, name: 'Julia', phone: '363636363', email: 'julia@example.com', avatar: 'https://i.pravatar.cc/40?img=36', usercode: '363456789' },
  { id: 36, name: 'Kevin', phone: '373737373', email: 'kevin@example.com', avatar: 'https://i.pravatar.cc/40?img=37', usercode: '373456789' },
  { id: 37, name: 'Lily', phone: '383838383', email: 'lily@example.com', avatar: 'https://i.pravatar.cc/40?img=38', usercode: '383456789' },
  { id: 38, name: 'Mike', phone: '393939393', email: 'mike@example.com', avatar: 'https://i.pravatar.cc/40?img=39', usercode: '393456789' },
  { id: 39, name: 'Nina', phone: '404040404', email: 'nina@example.com', avatar: 'https://i.pravatar.cc/40?img=40', usercode: '403456789' },
  { id: 40, name: 'Oliver', phone: '414141414', email: 'oliver@example.com', avatar: 'https://i.pravatar.cc/40?img=41', usercode: '413456789' },
  { id: 41, name: 'Peter', phone: '424242424', email: 'peter@example.com', avatar: 'https://i.pravatar.cc/40?img=42', usercode: '423456780' },
  { id: 42, name: 'Quincy', phone: '434343434', email: 'quincy@example.com', avatar: 'https://i.pravatar.cc/40?img=43', usercode: '433456789' },
  { id: 43, name: 'Rachel', phone: '444444444', email: 'rachel@example.com', avatar: 'https://i.pravatar.cc/40?img=44', usercode: '443456789' },
  { id: 44, name: 'Sam', phone: '454545454', email: 'sam@example.com', avatar: 'https://i.pravatar.cc/40?img=45', usercode: '453456789' },
  { id: 45, name: 'Tina', phone: '464646464', email: 'tina@example.com', avatar: 'https://i.pravatar.cc/40?img=46', usercode: '463456789' },
  { id: 46, name: 'Ulysses', phone: '474747474', email: 'ulysses@example.com', avatar: 'https://i.pravatar.cc/40?img=47', usercode: '473456789' },
  { id: 47, name: 'Vera', phone: '484848484', email: 'vera@example.com', avatar: 'https://i.pravatar.cc/40?img=48', usercode: '483456789' },
  { id: 48, name: 'William', phone: '494949494', email: 'william@example.com', avatar: 'https://i.pravatar.cc/40?img=49', usercode: '493456789' },
  { id: 49, name: 'Xena', phone: '505050505', email: 'xena@example.com', avatar: 'https://i.pravatar.cc/40?img=50', usercode: '503456789' },
  { id: 50, name: 'Yosef', phone: '515151515', email: 'yosef@example.com', avatar: 'https://i.pravatar.cc/40?img=51', usercode: '513456789' },
])

const groups = ref([
  { id: 1, name: '前端学习群', members: 120, avatar: 'https://i.pravatar.cc/40?img=31' },
  { id: 2, name: 'Vue3 开发者', members: 85, avatar: 'https://i.pravatar.cc/40?img=32' },
  { id: 3, name: 'React 技术群', members: 200, avatar: 'https://i.pravatar.cc/40?img=33' },
  { id: 4, name: 'Go 语言交流群', members: 150, avatar: 'https://i.pravatar.cc/40?img=34' },
  { id: 5, name: 'Java 后端群', members: 98, avatar: 'https://i.pravatar.cc/40?img=35' },
  { id: 6, name: 'Python 爱好者', members: 240, avatar: 'https://i.pravatar.cc/40?img=36' },
  { id: 7, name: 'AI 技术分享', members: 300, avatar: 'https://i.pravatar.cc/40?img=37' },
  { id: 8, name: '大数据交流群', members: 176, avatar: 'https://i.pravatar.cc/40?img=38' },
  { id: 9, name: '算法研究群', members: 132, avatar: 'https://i.pravatar.cc/40?img=39' },
  { id: 10, name: '测试工程师群', members: 67, avatar: 'https://i.pravatar.cc/40?img=40' },
  { id: 11, name: 'UI/UX 设计群', members: 89, avatar: 'https://i.pravatar.cc/40?img=41' },
  { id: 12, name: '数据库交流群', members: 178, avatar: 'https://i.pravatar.cc/40?img=42' },
  { id: 13, name: 'Linux 运维群', members: 201, avatar: 'https://i.pravatar.cc/40?img=43' },
  { id: 14, name: '网络安全群', members: 142, avatar: 'https://i.pravatar.cc/40?img=44' },
  { id: 15, name: '全栈开发群', members: 110, avatar: 'https://i.pravatar.cc/40?img=45' },
  { id: 16, name: '创业交流', members: 90, avatar: 'https://i.pravatar.cc/40?img=46' },
  { id: 17, name: '机器学习群', members: 256, avatar: 'https://i.pravatar.cc/40?img=47' },
  { id: 18, name: '深度学习群', members: 220, avatar: 'https://i.pravatar.cc/40?img=48' },
  { id: 19, name: '区块链交流群', members: 134, avatar: 'https://i.pravatar.cc/40?img=49' },
  { id: 20, name: 'Rust 开发者', members: 88, avatar: 'https://i.pravatar.cc/40?img=50' },
  { id: 21, name: 'C++ 开发群', members: 177, avatar: 'https://i.pravatar.cc/40?img=51' },
  { id: 22, name: '算法比赛群', members: 199, avatar: 'https://i.pravatar.cc/40?img=52' },
  { id: 23, name: '开源项目群', members: 144, avatar: 'https://i.pravatar.cc/40?img=53' },
  { id: 24, name: '自由职业群', members: 76, avatar: 'https://i.pravatar.cc/40?img=54' },
  { id: 25, name: '远程工作群', members: 112, avatar: 'https://i.pravatar.cc/40?img=55' },
  { id: 26, name: '软件架构群', members: 168, avatar: 'https://i.pravatar.cc/40?img=56' },
  { id: 27, name: '程序人生', members: 300, avatar: 'https://i.pravatar.cc/40?img=57' },
  { id: 28, name: '技术面试群', members: 140, avatar: 'https://i.pravatar.cc/40?img=58' },
  { id: 29, name: '运维自动化', members: 121, avatar: 'https://i.pravatar.cc/40?img=59' },
  { id: 30, name: '职业发展群', members: 102, avatar: 'https://i.pravatar.cc/40?img=60' }
])


const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('')
const selectedFriend = ref(null)
const selectedGroup = ref(null)
const showContacts = ref(true)
const showGroups = ref(false)
const selectedLetter = ref(null)
const searchQuery = ref('')

const selectFriend = (friend) => {
  selectedFriend.value = friend
  selectedGroup.value = null
}

const selectGroup = (group) => {
  selectedGroup.value = group
  selectedFriend.value = null
}

const toggleContacts = () => {
  showContacts.value = !showContacts.value
}


const toggleGroups = () => {
  showGroups.value = !showGroups.value
}

const selectLetter = (letter) => {
  showContacts.value = true
  if (selectedLetter.value === letter) {
    selectedLetter.value = null
  } else {
    selectedLetter.value = letter
  }
}

const filteredContacts = computed(() => {
  let result = contacts.value

  if (selectedLetter.value) {
    result = result.filter(c => c.name[0].toUpperCase() === selectedLetter.value)
  }

  if (searchQuery.value.trim()) {
    const query = searchQuery.value.trim().toLowerCase()
    result = result.filter(c =>
      c.name.toLowerCase().includes(query) ||
      c.phone.toLowerCase().includes(query) ||
      c.usercode && c.usercode.toLowerCase().includes(query)
    )
  }
  return result
})

watch(searchQuery, (newVal) => {
  if (newVal.trim()) {
    showContacts.value = true
  }
})


</script>

<template>
  <a-layout class="relations-layout">
    <a-layout-sider width="260" class="relations-sider">
      <a-input
        v-model:value="searchQuery"
        placeholder="搜索联系人"
        allow-clear
        size="large"
        class="search"
      />

      <div class="relations-groups">
        <a-card size="small" :body-style="{ padding: 0 }">
            <div class="card-title" @click="toggleGroups">
            <span>群组</span>
            <span>{{ showGroups ? '▲' : '▼' }}</span>
          </div>

          <a-collapse>
            <div v-if="showGroups" class="contacts-list">
              <a-list :data-source="groups">
                <template #renderItem="{ item }">
                  <a-list-item
                    class="contact-item"
                    @click="selectGroup(item)"
                    :class="{ active: selectedGroup && selectedGroup.id === item.id }"
                  >
                      <a-avatar shape="square" size="large" :src="item.avatar" style="margin-right: 10px" />
                    {{ item.name }}
                  </a-list-item>
                </template>
              </a-list>
            </div>
          </a-collapse>
        </a-card>
      </div>

      <div class="relations-contacts">
        <a-card size="small" :body-style="{ padding: 0 }">
          <div class="card-title" @click="toggleContacts">
            <span>联系人</span>
            <span>{{ showContacts ? '▲' : '▼' }}</span>
          </div>

          <a-collapse>
            <div v-if="showContacts" class="contacts-list">
              <a-list :data-source="filteredContacts">
                <template #renderItem="{ item }">
                  <a-list-item
                    class="contact-item"
                    @click="selectFriend(item)"
                    :class="{ active: selectedFriend && selectedFriend.id === item.id }"
                  >
                      <a-avatar shape="square" size="large" :src="item.avatar" style="margin-right: 10px" />
                    {{ item.name }}
                  </a-list-item>
                </template>
              </a-list>
            </div>
          </a-collapse>
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
        <user-info :id="selectedFriend.id"></user-info>
      </a-card>

      <a-card v-else-if="selectedGroup" title="群组信息" bordered>
        <group-info :id="selectedGroup.id"></group-info>
      </a-card>

      <a-card v-else bordered>
        <div class="content">
          <div class="placeholder-wrapper">
            <img src="../../static/img/logo.png" alt="请选择联系人" class="placeholder-img" />
            <!-- <div class="placeholder-text">信息为空</div> -->
          </div>
        </div>
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
  overflow-y: auto;
  scrollbar-width: none;
  -ms-overflow-style: none;
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
  max-height: 73vh; 
  overflow-y: auto;
  scrollbar-width: none;
  -ms-overflow-style: none;
  background: #fff;
  padding: 0;
  flex: 1;
}

.relations-groups .contacts-list {
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
  scrollbar-width: none;
  -ms-overflow-style: none;
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

.search {
  margin: 10px 0 10px 0;
}

.content {
  width: 100%;
  height: 74vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

.placeholder-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.placeholder-img {
  max-width: 200px;
  max-height: 200px;
  object-fit: contain;
  filter: grayscale(100%);
  pointer-events: none;
  user-select: none;
  -webkit-user-drag: none;
}

.placeholder-text {
  margin-top: 10px;
  color: #999;
  font-size: 14px;
}
</style>