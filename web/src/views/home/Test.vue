<template>
  <el-form ref="myForm" :model="formData" label-width="120px">
    <!-- 其他字段参数 -->
    <el-form-item label="姓名" prop="name">
      <el-input v-model="formData.name"></el-input>
    </el-form-item>

    <el-form-item label="描述" prop="description">
      <el-input v-model="formData.description"></el-input>
    </el-form-item>

    <el-form-item label="文件上传">
      <el-upload
        class="upload-demo"
        action="/your/upload/endpoint"
        :on-success="handleUploadSuccess"
        :before-upload="beforeUpload"
        multiple
        :limit="3"
        :file-list="fileList"
      >
        <el-button size="small" type="primary">点击上传</el-button>
      </el-upload>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="submitForm">提交</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { submitPaper } from "../../api";
export default {
  data() {
    return {
      fileList: [], // 存储上传的文件
      formData: {
        name: "",
        description: "",
        // 其他字段参数
      },
    };
  },
  methods: {
    handleUploadSuccess(response, file, fileList) {
      // 在文件上传成功后，将返回的数据（response）或其他信息存储在需要的变量中
      // response 包含了上传成功后后端返回的数据
      console.log(response);
      this.fileList = fileList;
    },
    beforeUpload(file) {
      // 可以在此处对文件进行校验或其他处理
    },
    submitForm() {
      this.$refs.myForm.validate((valid) => {
        if (valid) {
          // 构建包含文件和其他字段参数的 FormData 对象
          const formData = new FormData();
          formData.append("name", this.formData.name);
          formData.append("description", this.formData.description);
          this.fileList.forEach((file) => {
            formData.append("data", file.raw);
          });
          // 调用接口，实现文件上传
          submitPaper(formData).then((res) => {
            console.log(formData)
            console.log(res);
          });
        } else {
          return false;
        }
      });
    },
  },
};
</script>
