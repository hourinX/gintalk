<script setup>
import { ref, reactive, watch, onMounted } from 'vue';
import { getLoginLogList } from '@/api/log';
import { message } from 'ant-design-vue';

const columns = [
    {
        title: '序号',
        key: 'id',
        customRender: ({ index }) => {
          return (pagination.page_no - 1) * pagination.page_size + index + 1;
        },
    },
    {
        title: '用户ID',
        dataIndex: 'user_id',
        key: 'user_id',
    },
    {
        title: '用户名',
        dataIndex: 'user_name',
        key: 'user_name',
    },
    {
        title: 'IP 地址',
        dataIndex: 'ip_address',
        key: 'ip_address',
    },
    {
        title: '登录时间',
        dataIndex: 'login_time',
        key: 'login_time',
    },
    {
        title: '设备信息',
        dataIndex: 'device_info',
        key: 'device_info',
    },
    {
        title: '登录状态',
        dataIndex: 'login_status',
        key: 'login_status',
    },
    {
        title: '登录方式',
        dataIndex: 'login_method',
        key: 'login_method',
    },
];

const pagination = reactive({
  page_no: 1,
  page_size: 10,
  total: 0,
  showTotal: (total) => `共 ${total} 条记录`,
});

const dataSource = ref([]);
const loading = ref(false);

const query = reactive({
    type: 0,
})

const getLoginLogData = async(type) => {
  let params = { 
    type: type,
    page_no: pagination.page_no,
    page_size: pagination.page_size,
  };
  loading.value = true;
  try{
    const res = await getLoginLogList(params);
    loading.value = false;
    if (res.code === 10000) {
      dataSource.value = res.data.list;
      pagination.total = res.data.count;
    }
  } catch(err){ 
    loading.value = false;
  } finally{
    loading.value = false;
  }
  
};

const handleTableChange = (pag) => {
  pagination.page_no = pag.current;
  pagination.page_size = pag.pageSize;
  getLoginLogData(query.type);
};

onMounted(() => {
  getLoginLogData(query.type);
});

watch(
  () => query.type,
  (newValue) => {
    pagination.page_no = 1;
    getLoginLogData(newValue);
    message.success('查询成功');
  }
);
</script>

<template>
  <div class="list-v">
    <a-form layout="inline" style="margin-bottom: 20px;">
        <a-row>
            <a-form-item label="时间范围">
                <a-radio-group v-model:value="query.type">
                    <a-radio-button :value="0">全部</a-radio-button>
                    <a-radio-button :value="1">本周</a-radio-button>
                    <a-radio-button :value="2">本月</a-radio-button>
                    <a-radio-button :value="3">本年度</a-radio-button>
                </a-radio-group>
            </a-form-item>
        </a-row>
    </a-form>

    <a-table :columns="columns" :data-source="dataSource" :row-key="record => record.id" :pagination="pagination" :loading="loading" @change="handleTableChange">
    </a-table>
  </div>
</template>

<style scoped>
.list-v {
  background-color: #fff;
  height: auto;
}
</style>