<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="搜索关键词">
          <el-input
            v-model="searchInfo.keyword"
            placeholder="输入搜索关键词"
            size="mini"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            size="mini"
            icon="el-icon-search"
            @click="search"
            >查询</el-button
          >
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            size="mini"
            icon="el-icon-plus"
            @click="openDialog"
            >新增</el-button
          >
        </el-form-item>
        <el-form-item>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false"
                >取消</el-button
              >
              <el-button size="mini" type="primary" @click="deleteBatch"
                >确定</el-button
              >
            </div>
            <el-button
              slot="reference"
              icon="el-icon-delete"
              size="mini"
              type="danger"
              >批量删除</el-button
            >
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      :data="tableData"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" fixed="left" />
      <el-table-column label="ID" prop="id" width="70" />

      <el-table-column label="活动" prop="title" />
      <el-table-column label="人数" prop="quantity" width="80" />
      <el-table-column label="开始时间" width="160">
        <template #default="scope">
          {{ formatTime(scope.row.time_start) }}
        </template>
      </el-table-column>
      <el-table-column label="结束时间" width="160">
        <template #default="scope">
          {{ formatTime(scope.row.time_end) }}
        </template>
      </el-table-column>

      <el-table-column label="创建时间" width="160">
        <template #default="scope">
          {{ formatTime(scope.row.created_at) }}
        </template>
      </el-table-column>

      <el-table-column label="操作" fixed="right" width="220" align="center">
        <template #default="scope">
          <el-button
            class="table-button"
            size="mini"
            type="primary"
            icon="el-icon-edit"
            @click="edit(scope.row)"
            >修改</el-button
          >
          <el-button
            @click="remove(scope.row)"
            size="mini"
            type="danger"
            icon="el-icon-delete"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      layout="total, sizes, prev, pager, next, jumper"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />

    <el-dialog
      :before-close="closeDialog"
      v-model="dialogFormVisible"
      :title="type == 'create' ? '新增记录' : '编辑记录'"
    >
      <el-form
        :model="formData"
        label-position="right"
        label-width="100px"
        ref="form"
        :rules="rules"
      >
        <!--
        <el-form-item label="活动地址:" prop="address">
          <el-input
            v-model="formData.address"
            clearable
            placeholder="请输入活动地址"
          />
        </el-form-item>

        
        <el-form-item label="封面:" prop="cover">
          <el-input
            v-model="formData.cover"
            clearable
            placeholder="请输入封面"
          />
        </el-form-item>
        

        <el-form-item label="创建者:" prop="create_by">
          <el-input
            v-model="formData.create_by"
            clearable
            placeholder="请输入创建者"
          />
        </el-form-item>

        <el-form-item label="简介:" prop="description">
          <el-input
            v-model="formData.description"
            clearable
            placeholder="请输入简介"
          />
        </el-form-item>

        
        <el-form-item label="ID:" prop="id">
          <el-input v-model="formData.id" clearable placeholder="请输入ID" />
        </el-form-item>

        <el-form-item label="排序:" prop="sort">
          <el-input
            v-model="formData.sort"
            clearable
            placeholder="请输入排序"
          />
        </el-form-item>
        <el-form-item label="更新者:" prop="update_by">
          <el-input
            v-model="formData.update_by"
            clearable
            placeholder="请输入更新者"
          />
        </el-form-item>
        -->
        <el-form-item label="标题:" prop="title">
          <el-input
            v-model="formData.title"
            clearable
            placeholder="请输入标题"
          />
        </el-form-item>
        <el-form-item label="数量:" prop="quantity">
          <el-input
            v-model="formData.quantity"
            clearable
            placeholder="请输入数量"
          />
        </el-form-item>

        <el-form-item label="开始时间:" prop="time_start">
          <!--
          <el-input
            v-model="formData.time_start"
            clearable
            placeholder="请输入开始时间"
          />
          -->
          <el-date-picker
            v-model="formData.time_start"
            type="datetime"
            placeholder="请输入开始时间"
          />
        </el-form-item>

        <el-form-item label="结束时间:" prop="time_end">
          <!--
          <el-input
            v-model="formData.time_end"
            clearable
            placeholder="请输入结束时间"
          />
          -->
          <el-date-picker
            v-model="formData.time_end"
            type="datetime"
            placeholder="请输入结束时间"
          />
        </el-form-item>

        
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  activityList,
  activityDelete,
  activityDeleteBatch,
  activityOne,
  activityAdd,
  activityUpdate,
} from "@/api/activity"; //  此处请自行替换地址
import moment from "moment"; //导入模块
moment.locale("zh-cn"); //设置语言 或 moment.lang('zh-cn');
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
let defaultForm = {
  address: "",
  cover: "",
  create_by: "",
  description: "",
  id: 0,
  quantity: 10,
  sort: 0,
  time_end: "",
  time_start: "",
  title: "龙南佳苑自习室开放",
  update_by: "",
};
export default {
  name: "SystemUser",
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
  },
  mixins: [infoList],
  data() {
    return {
      listApi: activityList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: Object.assign({}, defaultForm),
      rules: {
        address: [
          { required: true, message: "请输入活动地址", trigger: "blur" },
        ],
        cover: [{ required: true, message: "请输入封面", trigger: "blur" }],
        create_by: [
          { required: true, message: "请输入创建者", trigger: "blur" },
        ],
        description: [
          { required: true, message: "请输入简介", trigger: "blur" },
        ],
        id: [{ required: true, message: "请输入ID", trigger: "blur" }],
        quantity: [{ required: true, message: "请输入数量", trigger: "blur" }],
        sort: [{ required: true, message: "请输入排序", trigger: "blur" }],
        time_end: [
          { required: true, message: "请输入结束时间", trigger: "blur" },
        ],
        time_start: [
          { required: true, message: "请输入开始时间", trigger: "blur" },
        ],
        title: [{ required: true, message: "请输入标题", trigger: "blur" }],
        update_by: [
          { required: true, message: "请输入更新者", trigger: "blur" },
        ],
      },
    };
  },
  async created() {
    await this.getTableData();
  },
  methods: {
    formatTime(v) {
      return moment(v).format("YYYY-MM-DD HH:mm:ss");
    },
    // 条件搜索前端看此方法
    search() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    async deleteBatch() {
      const ids = [];
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: "warning",
          message: "请选择要删除的数据",
        });
        return;
      }
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.id);
        });

      const res = await activityDeleteBatch({ ids: ids.join(",") });
      if (res.data.code == 200) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        //if (this.tableData.length == ids.length) {
        //  this.page--;
        //}
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    async edit(row) {
      const res = await activityOne({ id: row.id });
      this.type = "update";
      if (res.data.code == 200) {
        this.formData = res.data.data.item;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.$refs.form.resetFields();
      this.formData = Object.assign({}, defaultForm);
      this.dialogFormVisible = false;
    },
    async remove(row) {
      this.$confirm("此操作将永久删除所有角色下该api, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          const res = await activityDelete({ id: row.id });
          if (res.data.code == 200) {
            this.$message({
              type: "success",
              message: "删除成功!",
            });
            //if (this.tableData.length == 1) {
            //  this.page--;
            //}
            this.getTableData();
          }
        })
        .catch(() => {
          console.log("已取消删除");
          /*
          this.$message({
            type: "info",
            message: "已取消删除"
          });
          */
        });
    },
    async enterDialog() {
      this.$refs.form.validate(async (valid) => {
        if (valid) {
          let res;
          this.formData.quantity = parseInt(this.formData.quantity);
          this.formData.sort = parseInt(this.formData.sort);

          this.formData.time_end = this.formatTime(this.formData.time_end);
          this.formData.time_start = this.formatTime(this.formData.time_start);
          
            
          switch (this.type) {
            case "create":
              this.formData.id = 0;
              res = await activityAdd(this.formData);
              break;
            case "update":
              res = await activityUpdate(this.formData);
              break;
            default:
              res = await activityAdd(this.formData);
              break;
          }
          if (res.data.code == 200) {
            this.$message({
              type: "success",
              // message: "创建/更改成功",
              message: res.data.message,
            });
            this.closeDialog();
            this.getTableData();
          }
        }
      });
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
  },
};
</script>

<style>
</style>
