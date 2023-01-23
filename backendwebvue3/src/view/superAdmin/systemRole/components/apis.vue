<template>
  <div>
    <div class="clearfix sticky-button">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button class="fl-right" size="small" type="primary" @click="authApiEnter">{{ t('general.confirm') }}</el-button>
    </div>
    <div class="tree-content">
      <el-tree
        ref="apiTree"
        :data="apiTreeData"
        :default-checked-keys="apiTreeIds"
        :props="apiDefaultProps"
        default-expand-all
        highlight-current
        node-key="onlyId"
        show-checkbox
        :filter-node-method="filterNode"
        @check="nodeChange"
      />
    </div>
  </div>
</template>

<script setup>
import { systemApiList } from '@/api/systemApi'
import { systemRoleApiAdd, systemRoleApisByRoleId } from '@/api/systemRoleApi'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const apiDefaultProps = ref({
  children: 'children',
  label: 'description'
})
const filterText = ref('')
const apiTreeData = ref([])
const apiTreeIds = ref([])
const activeUserId = ref('')
const init = async() => {
  const res2 = await systemApiList({ page: 1, pageSize: 5000 })
  const apis = res2.data.list
  
  apiTreeData.value = buildApiTree(apis)  
  const res = await systemRoleApisByRoleId({
    id: props.row.id
  })
  
  activeUserId.value = props.row.id
  apiTreeIds.value = []
  res.data.list && res.data.list.forEach(item => {
    apiTreeIds.value.push('p:' + item.v1 + 'm:' + item.v2)
  })
}

init()

const needConfirm = ref(false)
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  authApiEnter()
}

// 创建api树方法
const buildApiTree = (apis) => {
  const apiObj = {}
  apis &&
        apis.forEach(item => {
          item.onlyId = 'p:' + item.path + 'm:' + item.method
          if (Object.prototype.hasOwnProperty.call(apiObj, item.api_group)) {
            apiObj[item.api_group].push(item)
          } else {
            Object.assign(apiObj, { [item.api_group]: [item] })
          }
        })
  const apiTree = []
  for (const key in apiObj) {
    const treeNode = {
      ID: key,
      description: key + t('api.group'),
      children: apiObj[key]
    }
    apiTree.push(treeNode)
  }
  return apiTree
}

// 关联关系确定
const apiTree = ref(null)
const authApiEnter = async() => {
  const checkArr = apiTree.value.getCheckedNodes(true)
  var casbinInfos = []
  checkArr && checkArr.forEach(item => {
    var casbinInfo = {
      p_type: 'p',
      v0: activeUserId.value + '',
      v1: item.path,
      v2: item.method,
      v3: '',
      v4: '',
      v5: ''
    }
    casbinInfos.push(casbinInfo)
  })
  const res = await systemRoleApiAdd({
    role_id: activeUserId.value,
    role_apis:casbinInfos
  })
  if (res.code === 200) {
    ElMessage({ type: 'success', message: t('api.setupSuccess') })
  }
}

defineExpose({
  needConfirm,
  enterAndNext
})

const filterNode = (value, data) => {
  if (!value) return true
  return data.description.indexOf(value) !== -1
}
watch(filterText, (val) => {
  apiTree.value.filter(val)
})

</script>

<style lang="scss" scoped>
@import "@/style/button.scss";
</style>
