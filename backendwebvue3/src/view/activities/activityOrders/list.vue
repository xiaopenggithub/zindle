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

      <el-table-column label="用户ID" prop="member_id" />
      <el-table-column label="活动ID" prop="activity_id" />
      <el-table-column label="状态" width="100">
        <template #default="scope">
          {{ formatStatus(scope.row.status)}}
        </template>
      </el-table-column>
      <el-table-column label="座位" prop="seat_number" width="70"/>

      <el-table-column label="创建时间" width="160">
        <template #default="scope">
          {{ formatTime(scope.row.created_at)}}
        </template>
      </el-table-column>

      <el-table-column label="操作" fixed="right" width="180" align="center">
        <template #default="scope">
          <el-button
            class="table-button"
            size="mini"
            type="primary"
            icon="el-icon-edit"
            @click="edit(scope.row)"
            >变更</el-button
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
        <el-form-item label="ID:" prop="id">
          <el-input v-model="formData.id" clearable placeholder="请输入ID" />
        </el-form-item>

        <el-form-item label="用户ID:" prop="member_id">
          <el-input
            v-model="formData.member_id"
            clearable
            placeholder="请输入用户ID"
          />
        </el-form-item>

        <el-form-item label="活动ID:" prop="activity_id">
          <el-input
            v-model="formData.activity_id"
            clearable
            placeholder="请输入活动ID"
          />
        </el-form-item>

        <el-form-item label="0报名1签到2取消3超时失约:" prop="status">
          <el-input
            v-model="formData.status"
            clearable
            placeholder="请输入0报名1签到2取消3超时失约"
          />
        </el-form-item>

        <el-form-item label="座位:" prop="seat_number">
          <el-input
            v-model="formData.seat_number"
            clearable
            placeholder="请输入座位"
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
  activityOrdersList,
  activityOrdersDelete,
  activityOrdersDeleteBatch,
  activityOrdersOne,
  activityOrdersAdd,
  activityOrdersUpdate,
} from "@/api/activityOrders"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import moment from 'moment';
import infoList from "@/mixins/infoList";
let defaultForm = {
  id: 0,
  member_id: 0,
  activity_id: 0,
  status: 0,
  seat_number: 1,
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
      listApi: activityOrdersList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: Object.assign({}, defaultForm),
      rules: {
        id: [{ required: true, message: "请输入ID", trigger: "blur" }],
        member_id: [
          { required: true, message: "请输入用户ID", trigger: "blur" },
        ],
        activity_id: [
          { required: true, message: "请输入活动ID", trigger: "blur" },
        ],
        status: [
          {
            required: true,
            message: "请输入0报名1签到2取消3超时失约",
            trigger: "blur",
          },
        ],
        seat_number: [
          { required: true, message: "请输入座位", trigger: "blur" },
        ],
      },
    };
  },
  async created() {
    await this.getTableData();
  },
  methods: {
    formatStatus(v){
      // 0报名1签到2取消3超时失约
      if(v==0){
        return '报名'
      }
      if(v==1){
        return '签到'
      }
      if(v==2){
        return '取消'
      }
      if(v==3){
        return '超时失约'
      }
      return '-'
    },
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

      const res = await activityOrdersDeleteBatch({ ids: ids.join(",") });
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
      const res = await activityOrdersOne({ id: row.id });
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
          const res = await activityOrdersDelete({ id: row.id });
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
          this.formData.member_id = parseInt(this.formData.member_id);
          this.formData.activity_id = parseInt(this.formData.activity_id);
          this.formData.status = parseInt(this.formData.status);
          this.formData.seat_number = parseInt(this.formData.seat_number);

          switch (this.type) {
            case "create":
              this.formData.id = 0;
              res = await activityOrdersAdd(this.formData);
              break;
            case "update":
              res = await activityOrdersUpdate(this.formData);
              break;
            default:
              res = await activityOrdersAdd(this.formData);
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
