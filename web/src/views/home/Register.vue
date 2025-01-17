<template>
  <div class="background">
    <!-- <div>
      <img class="image" src="../../images/1.jpg" />
    </div> -->
    <div class="box">
      <el-form
        :model="ruleForm"
        :inline="true"
        status-icon
        :rules="rules"
        ref="ruleForm"
        label-width="180px"
        class="demo-ruleForm"
        size="mini"
      >
        <el-form-item label="UserName" prop="username">
          <el-input v-model="ruleForm.username"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input
            type="password"
            v-model="ruleForm.password"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="FirstName" prop="first_name">
          <el-input v-model="ruleForm.first_name"></el-input>
        </el-form-item>
        <el-form-item label="LastName" prop="last_name">
          <el-input v-model="ruleForm.last_name"></el-input>
        </el-form-item>
        <!-- <el-form-item label="Sex" prop="data">
          <el-input v-model="ruleForm.sex"></el-input>
        </el-form-item> -->
        <el-form-item label="Email" prop="email">
          <el-input v-model="ruleForm.email"></el-input>
        </el-form-item>
        <el-form-item label="Department" prop="department">
          <el-input v-model="ruleForm.department"></el-input>
        </el-form-item>
        <el-form-item label="Phone" prop="phone">
          <el-input v-model="ruleForm.phone"></el-input>
        </el-form-item>
        <el-form-item label="Address" prop="address">
          <el-input v-model="ruleForm.address"></el-input>
        </el-form-item>
        <el-form-item label="Education" prop="education">
          <el-input v-model="ruleForm.education"></el-input>
        </el-form-item>
        <el-form-item label="Title" prop="title">
          <el-input v-model="ruleForm.title"></el-input>
        </el-form-item>
        <el-form-item label="Research" prop="research">
          <el-input v-model="ruleForm.research"></el-input>
        </el-form-item>
        <el-form-item label="Blockchain Addr" prop="block_chain_address">
          <el-input v-model="ruleForm.block_chain_address"></el-input>
        </el-form-item>
        <el-form-item label="Affiliations" prop="affiliation">
          <el-input v-model="ruleForm.affiliation"></el-input>
        </el-form-item>
        <el-form-item label="AffiliationType" prop="affiliation_type">
          <el-input v-model="ruleForm.affiliation_type"></el-input>
        </el-form-item>
        <div class="button">
          <el-button type="primary" round @click="submitForm('ruleForm')"
            >Register</el-button
          >
          <el-button type="primary" plain round @click="resetForm('ruleForm')"
            >Reset</el-button
          >
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
import { register,SendMail } from "../../api";
import { MPScontractInstance } from "@/constant";
const contractInstance = MPScontractInstance;
export default {
  data() {
    var checkData = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("data is required"));
      } else {
        callback();
      }
    };
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("Password is required"));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        id: 0,
        username: "",
        password: "",
        first_name: "",
        last_name: "",
        // sex: "",
        email: "",
        department: "",
        phone: "",
        address: "",
        education: "",
        title: "",
        research: "",
        block_chain_address: "",
        affiliation: "",
        affiliation_type: "",
      },
      SendMails:{
        MailReceiver:"",
        Verification:"",
      } ,
      verificationCode:"",
      rules: {
        username: [
          { required: true, trigger: "blur", message: "please input username" },
        ],
        password: [
          { required: true, trigger: "blur", message: "please input password" },
        ],
        first_name: [
          {
            required: true,
            trigger: "blur",
            message: "please input first name",
          },
        ],
        last_name: [
          {
            required: true,
            trigger: "blur",
            message: "please input last name",
          },
        ],
        email: [
          { required: true, trigger: "blur", message: "please input email",type:'email'},
        ],
        block_chain_address: [
          {
            required: true,
            trigger: "blur",
            message: "please input block chain address",
          },
        ],
      },
    };
  },
  methods: {
    async registe_gift(block_chain_address) {
      const functionArgs = [
          block_chain_address
        ];
      const functionName="registerUser"
      const result = await contractInstance.methods[functionName](
          ...functionArgs
        ).send({
          from: window.ethereum.selectedAddress,
          gasPrice: "0",
        });
      
      },
    generateCode() {
      const charset = 'abcdefghijklmnopqrstuvwxyz0123456789';
      let code = '';
      for (let i = 0; i < 6; i++) {
        const randomIndex = Math.floor(Math.random() * charset.length);
        code += charset[randomIndex];
      }
      this.verificationCode = code;
      return this.verificationCode;
    },
 
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.SendMails.MailReceiver=this.ruleForm.email;
            this.SendMails.Verification=this.generateCode();
            SendMail(this.SendMails)
          console.log(this.ruleForm);
          register(this.ruleForm).then(({ data }) => {
            console.log(data.data);
            if (data.code === 1000) {
              // localStorage.setItem("token", data.data.token); // 用localStorage缓存token值
              this.$alert("Register success", {
                confirmButtonText: "ok",
              });
              this.registe_gift(this.ruleForm.block_chain_address)
              this.$router.push("/home");
            }

            if (data.code === 1003) {
              this.$alert("User Existed!", {
                confirmButtonText: "ok",
              });
            }
          });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },
};
</script>
<style lang="less" scoped>
.image {
  width: 100%;
  height: 950px;
}
// .box {
//   border-radius: 4px;
//   box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
//   position: absolute;
//   left: 50%;
//   top: 50%;
//   transform: translate(-50%, -50%);
//   background-color: #f2f6fc;
//   height: 550px;
//   width: 70%;
//     .demo-ruleForm {
//       // display: inline-block;
//       padding-right: 60px;
//       margin-top: 80px;
//     }
//   .button {
//     margin-left: 330px;
//   }
// }
.background {
  // 盒子沾满整个屏幕
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-image: url("../../images/login.jpg");
  background-size: cover;
}
.box {
  // 盒子放在页面中间
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  // 盒子的宽度
  width: 850px;
  // 高度被内容撑开
  display: inline-block;

  // 盒子透明
  background-color: rgba(255, 255, 255, 0.5);
  // 盒子的圆角
  border-radius: 8px;
  padding-top: 25px;
  .button {
    margin-left: 330px;
  }
  padding-bottom: 20px;
}
</style>
