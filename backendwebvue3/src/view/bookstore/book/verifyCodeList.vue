<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="搜索关键词">
          <el-input
            v-model="searchInfo.keyword"
            placeholder="输入搜索账号"
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
        <!--
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
              <el-button
                size="mini"
                type="primary"
                @click="deleteBatch"
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
        -->
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

      <el-table-column label="号码(手机或邮箱)" prop="account" />
      <el-table-column label="验证码" prop="code" />

      <el-table-column label="类型" width="80" align="center">
        <template #default="scope">
          {{ scope.row.type | formatType }}
        </template>
      </el-table-column>

      <el-table-column label="状态" width="80" align="center">
        <template #default="scope">
          {{ scope.row.status | formatStatus }}
        </template>
      </el-table-column>

      <el-table-column label="创建时间" width="160">
        <template #default="scope">
          {{ scope.row.created_at | formatDate }}
        </template>
      </el-table-column>

      <el-table-column label="更新时间" width="160">
        <template #default="scope">
          {{ scope.row.updated_at | formatDate }}
        </template>
      </el-table-column>

      <!--
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
      -->
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
      v-model="dialogFormVisible"
      :before-close="closeDialog"      
      :title="type == 'create' ? '新增记录' : '编辑记录'"
    >
      <el-form
        :model="formData"
        label-position="right"
        label-width="100px"
        ref="form"
        :rules="rules"
      >
        <el-form-item label=":" prop="id">
          <el-input v-model="formData.id" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="号码(手机或邮箱):" prop="account">
          <el-input
            v-model="formData.account"
            clearable
            placeholder="请输入号码(手机或邮箱)"
          />
        </el-form-item>

        <el-form-item label="验证码:" prop="code">
          <el-input
            v-model="formData.code"
            clearable
            placeholder="请输入验证码"
          />
        </el-form-item>

        <el-form-item label="类型0手机1邮箱:" prop="type">
          <el-input
            v-model="formData.type"
            clearable
            placeholder="请输入类型0手机1邮箱"
          />
        </el-form-item>

        <el-form-item label="状态0未验证1已验证2验证错误:" prop="status">
          <el-input
            v-model="formData.status"
            clearable
            placeholder="请输入状态0未验证1已验证2验证错误"
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
  verifyCodeList,
  verifyCodeDelete,
  verifyCodeDeleteBatch,
  verifyCodeOne,
  verifyCodeAdd,
  verifyCodeUpdate,
} from "@/api/verifyCode"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
let defaultForm = {
  id: 0,
  account: "",
  code: "",
  type: 0,
  status: 0,
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
    formatType: function (v) {
      if (v == 0) {
        return "手机";
      }
      return "邮箱";
    },
    formatStatus: function (v) {
      if (v == 0) {
        return "未验证";
      } else if (v == 1) {
        return "已验证";
      } else {
        return "验证错误";
      }
    },
  },
  mixins: [infoList],
  data() {
    return {
      listApi: verifyCodeList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: Object.assign({}, defaultForm),
      rules: {
        id: [{ required: true, message: "请输入", trigger: "blur" }],
        account: [
          {
            required: true,
            message: "请输入号码(手机或邮箱)",
            trigger: "blur",
          },
        ],
        code: [{ required: true, message: "请输入验证码", trigger: "blur" }],
        type: [
          { required: true, message: "请输入类型0手机1邮箱", trigger: "blur" },
        ],
        status: [
          {
            required: true,
            message: "请输入状态0未验证1已验证2验证错误",
            trigger: "blur",
          },
        ],
      },
    };
  },
  async created() {
    await this.getTableData();
  },
  methods: {
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

      const res = await verifyCodeDeleteBatch({ ids: ids.join(",") });
      if (res.code == 200) {
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
      const res = await verifyCodeOne({ id: row.id });
      this.type = "update";
      if (res.code == 200) {
        this.formData = res.data.item;
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
          const res = await verifyCodeDelete({ id: row.id });
          if (res.code == 200) {
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
          this.formData.type = parseInt(this.formData.type);
          this.formData.status = parseInt(this.formData.status);

          switch (this.type) {
            case "create":
              this.formData.id = 0;
              res = await verifyCodeAdd(this.formData);
              break;
            case "update":
              res = await verifyCodeUpdate(this.formData);
              break;
            default:
              res = await verifyCodeAdd(this.formData);
              break;
          }
          if (res.code == 200) {
            this.$message({
              type: "success",
              // message: "创建/更改成功",
              message: res.message,
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
