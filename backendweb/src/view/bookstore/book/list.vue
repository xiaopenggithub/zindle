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
      <el-table-column label="封面" width="100" align="center">
        <template slot-scope="scope">
          <div v-if="scope.row.cover">
            <img :src="'/uploads1/' + scope.row.cover" width="80" height="80" />
          </div>
          <div v-else>-无-</div>
        </template>
      </el-table-column>

      <el-table-column label="标题" prop="title" />
      <el-table-column label="简介" prop="description" />
      <el-table-column label="数量" prop="quantity" />
      <el-table-column label="排序" prop="sort" />

      <el-table-column label="创建时间" width="160">
        <template slot-scope="scope">
          {{ scope.row.created_at | formatDate }}
        </template>
      </el-table-column>

      <el-table-column label="更新时间" width="160">
        <template slot-scope="scope">
          {{ scope.row.updated_at | formatDate }}
        </template>
      </el-table-column>

      <el-table-column label="操作" fixed="right" width="180" align="center">
        <template slot-scope="scope">
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
      :visible.sync="dialogFormVisible"
      :title="type == 'create' ? '新增记录' : '编辑记录'"
    >
      <el-form
        :model="formData"
        label-position="right"
        label-width="100px"
        ref="form"
        :rules="rules"
      >
        <el-form-item
          label="原封面:"
          v-if="formData.cover && formData.cover.indexOf('.') > 0"
        >
          <img :src="'uploads1/'+formData.cover" width="100" height="100" />
        </el-form-item>

        <el-form-item :minlength="6" label="上传封面">
          <el-upload
            ref="upload_coverUrl"
            accept=".jpg,.jpeg,.png,.gif,.bmp,.pdf,.JPG,.JPEG,.PBG,.GIF,.BMP,.mp4,.MP4,.apk,.APk"
            list-type="picture-card"
            :class="{ imgHide: imgHideUpload }"
            action=""
            :limit="1"
            :auto-upload="false"
            :file-list="fileList_coverUrl"
            :on-exceed="handleExceed_coverUrl"
            :before-upload="handleBeforeUpload_coverUrl"
            :on-preview="handlePictureCardPreview_coverUrl"
            :on-success="handleSuccess_coverUrl"
            :on-remove="handleRemove_coverUrl"
            :on-change="handleChange_coverUrl"
          >
            <i class="el-icon-plus"></i>
          </el-upload>
          <el-dialog :visible.sync="dialogVisible_coverUrl" append-to-body>
            <img width="100%" :src="dialogImageUrl_coverUrl" alt="" />
          </el-dialog>
        </el-form-item>

        <el-form-item label="标题:" prop="title">
          <el-input
            v-model="formData.title"
            clearable
            placeholder="请输入标题"
          />
        </el-form-item>

        <el-form-item label="简介:" prop="description">
          <el-input
            v-model="formData.description"
            clearable
            placeholder="请输入简介"
          />
        </el-form-item>

        <el-form-item label="数量:" prop="quantity">
          <el-input
            v-model="formData.quantity"
            clearable
            placeholder="请输入简介"
          />
        </el-form-item>

        <el-form-item label="排序:" prop="sort">
          <el-input
            v-model="formData.sort"
            clearable
            placeholder="请输入排序"
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
  bookList,
  bookDelete,
  bookDeleteBatch,
  bookOne,
  bookAdd,
  bookUpdate,
} from "@/api/book"; //  此处请自行替换地址
import axios from "axios";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { mapGetters, mapMutations } from "vuex";
import { store } from "@/store/index";
let defaultForm = {
  id: 0,
  title: "",
  description: "",
  quantity:0,
  cover: "",
  sort: 0,
  create_by: "",
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
      dialogVisible_coverUrl: false,
      fileList_coverUrl: [],
      dialogImageUrl_coverUrl: "",
      selectedfile_coverUrl: false,
      headers: {
        authorization: "",
        "Content-Type": "multipart/form-data",
      },
      imgHideUpload: false,
      coverFile:'',

      listApi: bookList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: Object.assign({}, defaultForm),
      rules: {
        id: [{ required: true, message: "请输入ID", trigger: "blur" }],
        title: [{ required: true, message: "请输入标题", trigger: "blur" }],
        description: [
          { required: true, message: "请输入简介", trigger: "blur" },
        ],
        cover: [{ required: true, message: "请输入封面", trigger: "blur" }],
        sort: [{ required: true, message: "请输入排序", trigger: "blur" }],
        quantity: [{ required: true, message: "请输入数量", trigger: "blur" }],
        create_by: [
          { required: true, message: "请输入创建者", trigger: "blur" },
        ],
        update_by: [
          { required: true, message: "请输入更新者", trigger: "blur" },
        ],
      },
    };
  },
  async created() {
    this.headers.authorization = this.token;
    await this.getTableData();
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"]),
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

      const res = await bookDeleteBatch({ ids: ids.join(",") });
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
      const res = await bookOne({ id: row.id });
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
          const res = await bookDelete({ id: row.id });
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
          this.formData.sort = parseInt(this.formData.sort);
          this.formData.quantity = parseInt(this.formData.quantity);
          let postData = new FormData();
          //postData = Object.assign({}, this.formData);
          
          console.log(this.fileList_coverUrl)
          postData.append('file',this.coverFile)

          postData.append('title',this.formData.title)
          postData.append('description',this.formData.description)
          postData.append('cover',this.formData.cover)
          postData.append('sort',this.formData.sort)
          postData.append('quantity',this.formData.quantity)
          postData.append('create_by',this.formData.create_by)
          postData.append('update_by',this.formData.update_by)


          const path = process.env.VUE_APP_BASE_API;
          let url = `${path}/book/add`;
          let res;
          switch (this.type) {
            case "create":
              //this.formData.id = 0;
              url = `${path}/book/add`;
              res = await axios.post(url, postData, { headers: this.headers });
              //res = await bookAdd(this.formData);
              break;
            case "update":
               postData.append('id',this.formData.id)
              //res = await bookUpdate(this.formData);
              url = `${path}/book/update`;
              res = await axios.put(url, postData, { headers: this.headers });
              break;
            default:
              url = `${path}/book/add`;
              res = await axios.post(url, postData, { headers: this.headers });
              // res = await bookAdd(this.formData);
              break;
          }
          console.log(res);
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
    // 上传 begin
    handleBeforeUpload_coverUrl(file) {
      console.log("before");
      if (
        !(
          file.type === "image/png" ||
          file.type === "image/gif" ||
          file.type === "image/jpg" ||
          file.type === "image/jpeg"
        )
      ) {
        this.$notify.warning({
          title: "警告",
          message:
            "请上传格式为image/png, image/gif, image/jpg, image/jpeg的图片",
        });
      }
      let size = file.size / 1024 / 1024 / 2;
      if (size > 2) {
        this.$notify.warning({
          title: "警告",
          message: "图片大小必须小于2M",
        });
      }
    },
    // 文件超出个数限制时的钩子
    handleExceed_coverUrl(files, fileList) {
      console.log(files, fileList);
      if (this.selectedfile_coverUrl || fileList.length >= 1) {
        this.$message.warning("上传的图片不能大于1张");
        return;
      }
    },
    // 点击文件列表中已上传的文件时的钩子
    handlePictureCardPreview_coverUrl(file) {
      this.dialogImageUrl_coverUrl = file.url;
      this.dialogVisible_coverUrl = true;
    },
    async handleSuccess_coverUrl(response, file, fileList) {
      if (response.code == 200) {
        console.log(response.data);
        this.showUpdateCover = false;
        // 更新store中头像
        await store.dispatch("user/changeCoverStore", response.data);
      }
    },
    handleChange_coverUrl(file, fileList) {
      this.coverFile=file.raw
      this.imgHideUpload = fileList.length >= 1;
      console.log(file);
      //this.$refs[this.Up].submit();
      this.selectedfile_coverUrl = true;
    },
    // 文件列表移除文件时的钩子
    handleRemove_coverUrl(file, fileList) {
      this.imgHideUpload = fileList.length >= 1;
      this.fileList_coverUrl = [];
      this.selectedfile_coverUrl = false;
    },
    // 点击上传
    uploadFile_coverUrl() {
      if (!this.selectedfile_coverUrl) {
        this.$message.warning("请选择一个文件");
        return;
      }
      this.$refs.upload_coverUrl.submit();
    },
    // 上传 end
  },
};
</script>

<style>
</style>
