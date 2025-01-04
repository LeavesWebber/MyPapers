<template>
  <el-row>
    <el-col :span="24">
      <div class="box">
        <h1 style="text-align: center">Create Journal</h1>
        <el-form
          ref="form"
          :model="form"
          status-icon
          :rules="rules"
          label-width="145px"
        >
          <el-form-item
            style="width: 600px"
            label="Committee Name"
            prop="committee_name"
          >
            <el-select
              style="width: 100%"
              v-model="form.committee_name"
              filterable
              placeholder="select committee"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-form-item>

          <el-form-item style="width: 600px" label="Name" prop="name">
            <el-input v-model="form.name"></el-input>
          </el-form-item>
          <el-form-item label="Category" prop="categoryTags">
            <el-tag
              :key="tag"
              v-for="tag in form.categoryTags"
              closable
              :disable-transitions="false"
              @close="handleClose(tag)"
            >
              {{ tag }}
            </el-tag>
            <el-input
              style="width: 100px"
              class="input-new-tag"
              v-if="inputVisible"
              v-model="inputValue"
              ref="saveTagInput"
              size="small"
              @keyup.enter.native="handleInputConfirm"
              @blur="handleInputConfirm"
            >
            </el-input>
            <el-button
              v-else
              class="button-new-tag"
              size="small"
              @click="showInput"
              >+ Category</el-button
            >
          </el-form-item>
          <el-form-item
            style="width: 800px"
            label="Description"
            prop="description"
          >
            <div class="editor-container">
              <div ref="editor"></div>
              <div class="word-count">
                Words: {{ wordCount }} / 250 (minimum: 150)
              </div>
            </div>
          </el-form-item>

          <h3>President</h3>
          <el-form
            v-for="(form, index) in ruleForm"
            :model="form"
            :rules="rules2"
            size="mini"
            ref="ruleForm"
            label-width="100px"
            class="demo-ruleForm"
          >
            <el-row>
              <el-col :span="6">
                <el-form-item label="Name" prop="name">
                  <el-input v-model="form.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="Position" prop="position">
                  <el-input v-model="form.position"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="Start Time" prop="start_time">
                  <el-date-picker
                    v-model="form.start_time"
                    type="date"
                    placeholder="Start Time"
                    :picker-options="dateTimeStartFunc(index)"
                  >
                  </el-date-picker>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="End Time" prop="end_time">
                  <el-date-picker
                    v-model="form.end_time"
                    type="date"
                    placeholder="End Time"
                    :picker-options="dateTimeEndFunc(index)"
                  >
                  </el-date-picker>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <el-divider></el-divider>
          <div class="button-container">
            <el-form>
              <el-form-item size="mini" label-width="100px">
                <el-button type="primary" plain @click="add()">+</el-button>
                <el-button
                  type="danger"
                  plain
                  @click="reduce()"
                  :disabled="flag"
                  >-</el-button
                >
                <el-button type="warning" plain @click="resetForm()"
                  >Reset</el-button
                >
              </el-form-item>
            </el-form>
          </div>
          <h3>Vice President</h3>
          <el-form
            v-for="(form, index) in ruleForm2"
            :model="form"
            :rules="rules2"
            size="mini"
            ref="ruleForm2"
            label-width="100px"
            class="demo-ruleForm"
          >
            <el-row>
              <el-col :span="6">
                <el-form-item label="Name" prop="name">
                  <el-input v-model="form.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="Position" prop="position">
                  <el-input v-model="form.position"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="Start Time" prop="start_time">
                  <el-date-picker
                    v-model="form.start_time"
                    type="date"
                    placeholder="Start Time"
                    :picker-options="dateTimeStartFunc2(index)"
                  >
                  </el-date-picker>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="End Time" prop="end_time">
                  <el-date-picker
                    v-model="form.end_time"
                    type="date"
                    placeholder="End Time"
                    :picker-options="dateTimeEndFunc2(index)"
                  >
                  </el-date-picker>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <el-divider></el-divider>
          <div class="button-container">
            <el-form>
              <el-form-item size="mini" label-width="100px">
                <el-button type="primary" plain @click="add2()">+</el-button>
                <el-button
                  type="danger"
                  plain
                  @click="reduce2()"
                  :disabled="flag2"
                  >-</el-button
                >
                <el-button type="warning" plain @click="resetForm2()"
                  >Reset</el-button
                >
              </el-form-item>
            </el-form>
          </div>
          <h3>Members</h3>
          <el-form
            v-for="(form, index) in ruleForm3"
            :model="form"
            :rules="rules2"
            size="mini"
            ref="ruleForm3"
            label-width="100px"
            class="demo-ruleForm"
          >
            <el-row>
              <el-col :span="6">
                <el-form-item label="Name" prop="name">
                  <el-input v-model="form.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="Position" prop="position">
                  <el-input v-model="form.position"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="Start Time" prop="start_time">
                  <el-date-picker
                    v-model="form.start_time"
                    type="date"
                    placeholder="Start Time"
                    :picker-options="dateTimeStartFunc3(index)"
                  >
                  </el-date-picker>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="End Time" prop="end_time">
                  <el-date-picker
                    v-model="form.end_time"
                    type="date"
                    placeholder="End Time"
                    :picker-options="dateTimeEndFunc3(index)"
                  >
                  </el-date-picker>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <el-divider></el-divider>
          <div class="button-container">
            <el-form>
              <el-form-item size="mini" label-width="100px">
                <el-button type="primary" plain @click="add3()">+</el-button>
                <el-button
                  type="danger"
                  plain
                  @click="reduce3()"
                  :disabled="flag3"
                  >-</el-button
                >
                <el-button type="warning" plain @click="resetForm3()"
                  >Reset</el-button
                >
              </el-form-item>
            </el-form>
          </div>
          <el-form-item style="text-align: center">
            <el-button type="primary" @click="create('form')">Create</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import Quill from 'quill'
import 'quill/dist/quill.snow.css'
import { createJournal, getCommitteeList } from "../../api";

export default {
  data() {
    var validateWordCount = (rule, value, callback) => {
      const wordCount = value.trim().split(/\s+/).length;
      if (wordCount < 1 || wordCount > 20) {
        callback(new Error("Should contain 1 to 20 words"));
      } else {
        callback();
      }
    };
    var validateDescriptionWordCount = (rule, value, callback) => {
      const wordCount = value.trim().split(/\s+/).length;
      if (wordCount < 10 || wordCount > 10000) {
        callback(new Error("Should contain 10 to 10k words"));
      } else {
        callback();
      }
    };
    return {
      dateTimeStartFunc(index) {
        return {
          disabledDate: (time) => {
            if (this.ruleForm[index].end_time) {
              return (
                time.getTime() >
                new Date(this.ruleForm[index].end_time).getTime() - 8.64e7
              );
            }
          },
        };
      },
      dateTimeEndFunc(index) {
        return {
          disabledDate: (time) => {
            if (this.ruleForm[index].start_time) {
              return (
                time.getTime() <
                new Date(this.ruleForm[index].start_time).getTime() + 8.64e7
              );
            }
          },
        };
      },
      dateTimeStartFunc2(index) {
        return {
          disabledDate: (time) => {
            if (this.ruleForm2[index].end_time) {
              return (
                time.getTime() >
                new Date(this.ruleForm2[index].end_time).getTime() - 8.64e7
              );
            }
          },
        };
      },
      dateTimeEndFunc2(index) {
        return {
          disabledDate: (time) => {
            if (this.ruleForm2[index].start_time) {
              return (
                time.getTime() <
                new Date(this.ruleForm2[index].start_time).getTime() + 8.64e7
              );
            }
          },
        };
      },
      dateTimeStartFunc3(index) {
        return {
          disabledDate: (time) => {
            if (this.ruleForm3[index].end_time) {
              return (
                time.getTime() >
                new Date(this.ruleForm3[index].end_time).getTime() - 8.64e7
              );
            }
          },
        };
      },
      dateTimeEndFunc3(index) {
        return {
          disabledDate: (time) => {
            if (this.ruleForm3[index].start_time) {
              return (
                time.getTime() <
                new Date(this.ruleForm3[index].start_time).getTime() + 8.64e7
              );
            }
          },
        };
      },
      ruleForm: [
        {
          name: "",
          position: "",
          start_time: "",
          end_time: "",
          level: "president",
        },
      ],
      ruleForm2: [
        {
          name: "",
          position: "",
          start_time: "",
          end_time: "",
          level: "vice_president",
        },
      ],
      ruleForm3: [
        {
          name: "",
          position: "",
          start_time: "",
          end_time: "",
          level: "member",
        },
      ],
      form: {
        committee_id: 0,
        committee_name: "",
        name: "",
        category: "",
        categoryTags: [],
        description: "",
        presidents: [],
        vice_presidents: [],
        members: [],
      },
      rules: {
        committee_name: [
          {
            required: true,
            message: "Please select committee",
            trigger: "blur",
          },
        ],
        name: [
          { required: true, message: "Please input name", trigger: "blur" },
          { validator: validateWordCount, trigger: "blur" },
        ],
        categoryTags: [
          {
            required: true,
            message: "Please input categorys",
            trigger: "blur",
          },
        ],
        description: [
          {
            required: true,
            message: "Please input description",
            trigger: "blur",
          },
          { validator: validateDescriptionWordCount, trigger: "blur" },
        ],
      },
      rules2: {
        name: [
          { required: true, message: "Please input name", trigger: "blur" },
        ],
        position: [
          { required: true, message: "Please input position", trigger: "blur" },
        ],
        start_time: [
          {
            required: true,
            message: "Please input start time",
            trigger: "blur",
          },
        ],
        end_time: [
          { required: true, message: "Please input end time", trigger: "blur" },
        ],
      },
      flag: true,
      flag2: true,
      flag3: true,
      options: [],
      value: [],
      inputVisible: false,
      inputValue: "",
      editor: null,
      wordCount: 0,
      editorOption: {
        modules: {
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'],
            ['blockquote', 'code-block'],
            [{ 'header': 1 }, { 'header': 2 }],
            [{ 'list': 'ordered'}, { 'list': 'bullet' }],
            [{ 'color': [] }, { 'background': [] }],
            ['clean']
          ]
        },
        placeholder: 'Please input description...',
        theme: 'snow'
      },
    };
  },
  methods: {
    // 表单添加一行
    add() {
      var arr = {
        name: "",
        position: "",
        start_time: "",
        end_time: "",
        level: "president",
      };
      this.ruleForm.push(arr);
      this.flags();
    },
    // 表单减少一行
    reduce() {
      this.ruleForm.length = this.ruleForm.length - 1;
      this.flags();
    },
    // 判断数组长度
    flags() {
      if (this.ruleForm.length < 2) {
        this.flag = true;
      } else {
        //先赋值为true再赋为false, 不然会没反应
        this.flag = true;
        this.flag = false;
      }
    },
    // 重置方法
    resetForm() {
      this.ruleForm = [
        {
          name: "",
          position: "",
          start_time: "",
          end_time: "",
          level: "president",
        },
      ];
    },

    add2() {
      var arr = {
        name: "",
        position: "",
        start_time: "",
        end_time: "",
        level: "vice_president",
      };
      this.ruleForm2.push(arr);
      this.flags2();
    },
    reduce2() {
      this.ruleForm2.length = this.ruleForm2.length - 1;
      this.flags2();
    },
    flags2() {
      if (this.ruleForm2.length < 2) {
        this.flag2 = true;
      } else {
        //先赋值为true再赋为false, 不然会没反应
        this.flag2 = true;
        this.flag2 = false;
      }
    },
    resetForm2() {
      this.ruleForm2 = [
        {
          name: "",
          position: "",
          start_time: "",
          end_time: "",
          level: "vice_president",
        },
      ];
    },

    add3() {
      var arr = {
        name: "",
        position: "",
        start_time: "",
        end_time: "",
        level: "member",
      };
      this.ruleForm3.push(arr);
      this.flags3();
    },
    reduce3() {
      this.ruleForm3.length = this.ruleForm3.length - 1;
      this.flags3();
    },
    flags3() {
      if (this.ruleForm3.length < 2) {
        this.flag3 = true;
      } else {
        //先赋值为true再赋为false, 不然会没反应
        this.flag3 = true;
        this.flag3 = false;
      }
    },
    resetForm3() {
      this.ruleForm3 = [
        {
          name: "",
          position: "",
          start_time: "",
          end_time: "",
          level: "member",
        },
      ];
    },
    // 验证各个表单的逻辑
    validateForm() {
      let f1 = false;
      let f2 = false;
      let f3 = false;
      // 验证ruleForm
      for (let i = 0; i < this.ruleForm.length; i++) {
        this.$refs.ruleForm[i].validate((valid) => {
          if (valid) {
            f1 = true;
          }
        });
      }
      // 验证ruleForm2
      for (let i = 0; i < this.ruleForm2.length; i++) {
        this.$refs.ruleForm2[i].validate((valid) => {
          if (valid) {
            f2 = true;
          }
        });
      }
      // 验证ruleForm3
      for (let i = 0; i < this.ruleForm3.length; i++) {
        this.$refs.ruleForm3[i].validate((valid) => {
          if (valid) {
            f3 = true;
          }
        });
      }
      return f1 && f2 && f3;
    },
    create(form) {
      this.$refs[form].validate((valid) => {
        if (valid) {
          if (this.validateForm()) {
            // ruleForm有值才给form赋值
            // if (this.ruleForm[0].name) {
            this.form.presidents = this.ruleForm;
            // }
            // if (this.ruleForm2[0].name) {
            this.form.vice_presidents = this.ruleForm2;
            // }
            // if (this.ruleForm3[0].name) {
            this.form.members = this.ruleForm3;
            // }
            this.form.committee_id = this.form.committee_name;
            this.form.category = this.form.categoryTags.join(",");
            console.log(this.form, "form");
            console.log(this.ruleForm);
            console.log(this.ruleForm2);
            console.log(this.ruleForm3);
            createJournal(this.form).then((data) => {
              console.log(data);
              if (data.data.code === 1000) {
                this.$message({
                  message: "Create successfully",
                  type: "success",
                });
                this.$router.push("/center/selfJournal");
              } else {
                this.$message({
                  message: "Username may not exist",
                  type: "error",
                });
              }
            });
          } else {
            console.log("error create!!!");
          }
        } else {
          console.log("error create!!");
          return false;
        }
      });
    },
    handleClose(tag) {
      console.log(
        this.form.categoryTags.indexOf(tag),
        "this.form.categoryTags"
      );
      console.log(this.form.categoryTags, "this.form.categoryTags");

      this.form.categoryTags.splice(this.form.categoryTags.indexOf(tag), 1);
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.form.categoryTags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
    getCommitteeList() {
      getCommitteeList().then((res) => {
        console.log(res.data.data, "res.data.data");
        this.options = res.data.data.map((item) => {
          return { value: item.ID, label: item.name };
        });
        console.log(this.options, "this.options");
      });
    },
    initEditor() {
      this.editor = new Quill(this.$refs.editor, this.editorOption);
      
      // 监听内容变化
      this.editor.on('text-change', () => {
        // 获取纯文本内容
        let text = this.editor.getText();
        
        // 移除所有换行符和多余空格
        text = text.replace(/\n/g, ' ').trim();
        
        // 如果文本为空，设置字数为0
        if (!text || text === ' ') {
          this.wordCount = 0;
        } else {
          // 分别处理中文和英文
          const chineseChars = text.match(/[\u4e00-\u9fa5]/g) || [];
          
          // 移除中文字符，然后按空格分割英文单词
          const englishText = text.replace(/[\u4e00-\u9fa5]/g, '');
          const englishWords = englishText.split(/\s+/).filter(word => word.length > 0);
          
          // 总字数 = 中文字符数 + 英文单词数
          this.wordCount = chineseChars.length + englishWords.length;
        }
        
        // 更新 form.description，保留HTML格式
        this.form.description = this.editor.root.innerHTML;
      });

      // 如果有初始值，设置它
      if (this.form.description) {
        this.editor.root.innerHTML = this.form.description;
        
        // 初始化时也计算一次字数
        let text = this.editor.getText().replace(/\n/g, ' ').trim();
        
        if (!text || text === ' ') {
          this.wordCount = 0;
        } else {
          const chineseChars = text.match(/[\u4e00-\u9fa5]/g) || [];
          const englishText = text.replace(/[\u4e00-\u9fa5]/g, '');
          const englishWords = englishText.split(/\s+/).filter(word => word.length > 0);
          this.wordCount = chineseChars.length + englishWords.length;
        }
      }
    },
  },
  mounted() {
    this.getCommitteeList();
    this.initEditor();
  },
  beforeDestroy() {
    // 清理编辑器实例
    if (this.editor) {
      this.editor = null;
    }
  }
};
</script>
<style lang="less" scoped>
.box {
  // 左右下内边距
  padding: 0px 60px 60px 60px;
  // 取消父组件的居中
  text-align: left;
}
.el-date-editor.el-input {
  width: 155px;
}
.button-container {
  display: flex;
  justify-content: flex-end;
}
.editor-container {
  border-radius: 4px;
  
  .ql-toolbar {
    white-space: nowrap;
    overflow-x: auto;
    border-top-left-radius: 4px;
    border-top-right-radius: 4px;
  }

  .ql-container {
    height: 200px;
    border-bottom-left-radius: 4px;
    border-bottom-right-radius: 4px;
  }

  .ql-editor {
    height: 100%;
    font-size: 14px;
  }
}

.word-count {
  margin-top: 5px;
  text-align: right;
  color: #606266;
  font-size: 12px;
}

// 确保编辑器在较小屏幕上也能正常显示
@media screen and (max-width: 900px) {
  .editor-container {
    width: 100%;
    min-width: 300px;
  }
}
</style>