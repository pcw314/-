<template>
  <div class="contain" :class="isSignUp ? 'sign-up-mode' : ''">
    <div class="forms-contain">
      <div class="signin-signup">
        <form action="#" class="sign-in-form">
          <div class="flex flex-row tab">
            <div @click="loginRole(1)" :class="activeTab === 1 ? 'active' : ''">
              我要接单
            </div>
            <div @click="loginRole(2)" :class="activeTab === 2 ? 'active' : ''">
              发布任务
            </div>
            <div @click="loginRole(3)" :class="activeTab === 3 ? 'active' : ''">
              管理后台
            </div>
          </div>
          <div class="input-field">
            <el-icon :size="20" class="fas">
              <User />
            </el-icon>
            <input type="text" placeholder="输入您的账号" v-model="username" />
          </div>
          <div class="input-field">
            <el-icon :size="20" class="fas">
              <Lock />
            </el-icon>
            <input type="password" placeholder="输入您的密码" v-model="password" />
          </div>
          <input type="submit" value="登 录" class="btn solid" @click.prevent="login" />
        </form>
        <form action="#" class="sign-up-form">
          <div class="flex flex-row tab">
            <div @click="loginRole(1)" :class="activeTab === 1 ? 'active' : ''">
              我要接单
            </div>
            <div @click="loginRole(2)" :class="activeTab === 2 ? 'active' : ''">
              发布任务
            </div>
            <!-- <div @click="loginRole(3)" :class="activeTab === 3 ? 'active' : ''">
              管理后台
            </div> -->
          </div>
          <div class="input-field">
            <el-icon :size="20" class="fas">
              <User />
            </el-icon>
            <input type="text" placeholder="输入您的账号" v-model="username" />
          </div>
          <div class="input-field">
            <el-icon :size="20" class="fas">
              <Lock />
            </el-icon>
            <input type="password" placeholder="输入您的密码" v-model="password" />
          </div>
          <input type="submit" class="btn" value="注 册" @click.prevent="register" />
        </form>
      </div>
    </div>

    <div class="panels-contain">
      <div class="panel left-panel">
        <div class="content">
          <h3>新用户?</h3>
          <p>
            好兄弟,你来了,我们的网站就差你的加入了,点击下方注册按钮加入我们吧!!
          </p>
          <button class="btn transparent" id="sign-up-btn" @click="signUpBtn">
            注册
          </button>
        </div>
        <img src="@/static/img/log.svg" class="image" alt="" />
      </div>
      <div class="panel right-panel">
        <div class="content">
          <h3>已经是我们自己人了吗?</h3>
          <p>那好兄弟,你直接点击登录按钮,登录到我们这优秀的系统里!!</p>
          <button class="btn transparent" id="sign-in-btn" @click="signInBtn">
            登 录
          </button>
        </div>
        <img src="@/static/img/register.svg" class="image" alt="" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { loginAdminApi, menuApi, userRegister } from "@/api/login";
import { getCompanyInfo,getUserInfoById } from "@/api/common"
import { useUserStore } from "@/store/module/user";
import { ElMessage } from 'element-plus'
import { useRouter } from "vue-router";
const route = useRouter()
const userStore = useUserStore();
// const formObj = reactive({ username: "", password: "",role:1, });
const activeTab = ref<number>(1);
const username = ref<string>("");
const password = ref<string>("");
const loginRole = (index: number) => {
  activeTab.value = index;
};
const login = () => {
  if (!username.value) {
    return ElMessage.error('账号不能为空')
  } if (!username.value) {
    return ElMessage.error('密码不能为空')
  }

  loginAdminApi({ username: username.value, password: password.value, role: activeTab.value })
    .then(async (res) => {
      await userStore.setToken(res.data);
     if(activeTab.value!=3){
      ElMessage.success(res.Message ? res.Message : "数据异常")
      let funcName = activeTab.value==1?getUserInfoById:getCompanyInfo
      console.log(funcName,'funcName')
      funcName(0).then(async res => {
        if (res.code == 200) {
          userStore.setUserInfo(res.data)
          if(!res.data.school_id){
          route.push({ path: "/userInit" })
          }else{
          route.push({ path: "/center/index" })

          }

        }
      }).catch(res=>{
        userStore.loginOut()
      })
     }else{
     await getMenue();

     }
    })
    .catch(() => {
    });
};
const getMenue = () => {
  menuApi().then((res) => {
    userStore.setRoleRouters(res.data);
  });
};
const isSignUp = ref<boolean>(false);

const signUpBtn = () => {
  isSignUp.value = true;
};
const signInBtn = () => {
  isSignUp.value = false;
};
const register = () => {
  userRegister({ username: username.value, password: password.value, role: activeTab.value }).then(res => {
    if (res.code === 200) {
      isSignUp.value = false;

    }
    ElMessage.success(res.Message ? res.Message : "数据异常")
  })
}
</script>

<style scoped lang="scss">
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700;800&display=swap");

body,
input {
  font-family: "Poppins", sans-serif;
}

.contain {
  position: relative;
  min-height: 100vh;

  overflow: hidden;
}

.forms-contain {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
}

.tab {
  margin-bottom: 20px;

  div {
    position: relative;
    cursor: pointer;
    width: 100px;
    padding: 10px;
    transition: all 3s ease;
  }

  div::after {
    content: "";
    position: absolute;
    bottom: 0;
    left: 0;
    /* 初始位置在中间 */
    width: 0;
    /* 初始宽度为0 */
    height: 1px;
    /* 边框的厚度 */
    background-color: #4481eb;
    /* 边框颜色 */
    transition: width 0.5s ease-out;
    /* 动画效果 */
  }

  .active::after {
    bottom: 0;

    width: 100%;
    /* 鼠标悬停时宽度变为100% */
  }
}

.signin-signup {
  position: absolute;
  top: 50%;
  transform: translate(-50%, -50%);
  left: 75%;
  width: 50%;
  transition: 1s 0.7s ease-in-out;
  display: grid;
  grid-template-columns: 1fr;
  z-index: 5;
}

form {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  padding: 0rem 5rem;
  transition: all 0.2s 0.7s;
  overflow: hidden;
  grid-column: 1 / 2;
  grid-row: 1 / 2;
}

form.sign-up-form {
  opacity: 0;
  z-index: 1;
}

form.sign-in-form {
  z-index: 2;
}

.title {
  font-size: 2.2rem;
  color: #444;
  margin-bottom: 10px;
}

.input-field {
  max-width: 380px;
  width: 100%;
  background-color: #f0f0f0;
  margin: 10px 0;
  height: 55px;
  border-radius: 55px;
  display: grid;
  grid-template-columns: 15% 85%;
  padding: 0 0.4rem;
  position: relative;
}

.input-field .fas {
  justify-self: center;
  align-self: center;
  line-height: 55px;
  color: #acacac;
  transition: 0.5s;
  font-size: 1.1rem;
}

.input-field input {
  background: none;
  outline: none;
  border: none;
  line-height: 1;
  font-weight: 600;
  font-size: 1.1rem;
  color: #333;
}

.input-field input::placeholder {
  color: #aaa;
  font-weight: 500;
}

.social-icon {
  height: 46px;
  width: 46px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0 0.45rem;
  color: #333;
  border-radius: 50%;
  border: 1px solid #333;
  text-decoration: none;
  font-size: 1.1rem;
  transition: 0.3s;
}

.social-icon:hover {
  color: #4481eb;
  border-color: #4481eb;
}

.btn {
  width: 150px;
  background-color: #5995fd;
  border: none;
  outline: none;
  height: 49px;
  border-radius: 49px;
  color: #fff;
  text-transform: uppercase;
  font-weight: 600;
  margin: 10px 0;
  cursor: pointer;
  transition: 0.5s;
}

.btn:hover {
  background-color: #4d84e2;
}

.panels-contain {
  position: absolute;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  display: grid;
  grid-template-columns: repeat(2, 1fr);
}

.contain:before {
  content: "";
  position: absolute;
  height: 2000px;
  width: 2000px;
  top: -10%;
  right: 48%;
  transform: translateY(-50%);
  background-image: linear-gradient(-45deg, #4481eb 0%, #04befe 100%);
  transition: 1.8s ease-in-out;
  border-radius: 50%;
  z-index: 6;
}

.image {
  width: 100%;
  transition: transform 1.1s ease-in-out;
  transition-delay: 0.4s;
}

.panel {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  justify-content: space-around;
  text-align: center;
  z-index: 6;
}

.left-panel {
  pointer-events: all;
  padding: 3rem 17% 2rem 12%;
}

.right-panel {
  pointer-events: none;
  padding: 3rem 12% 2rem 17%;
}

.panel .content {
  color: #fff;
  transition: transform 0.9s ease-in-out;
  transition-delay: 0.6s;
}

.panel h3 {
  font-weight: 600;
  line-height: 1;
  font-size: 1.5rem;
}

.panel p {
  font-size: 0.95rem;
  padding: 0.7rem 0;
}

.btn.transparent {
  margin: 0;
  background: none;
  border: 2px solid #fff;
  width: 130px;
  height: 41px;
  font-weight: 600;
  font-size: 0.8rem;
}

.right-panel .image,
.right-panel .content {
  transform: translateX(800px);
}

/* ANIMATION */

.contain.sign-up-mode:before {
  transform: translate(100%, -50%);
  right: 52%;
}

.contain.sign-up-mode .left-panel .image,
.contain.sign-up-mode .left-panel .content {
  transform: translateX(-800px);
}

.contain.sign-up-mode .signin-signup {
  left: 25%;
}

.contain.sign-up-mode form.sign-up-form {
  opacity: 1;
  z-index: 2;
}

.contain.sign-up-mode form.sign-in-form {
  opacity: 0;
  z-index: 1;
}

.contain.sign-up-mode .right-panel .image,
.contain.sign-up-mode .right-panel .content {
  transform: translateX(0%);
}

.contain.sign-up-mode .left-panel {
  pointer-events: none;
}

.contain.sign-up-mode .right-panel {
  pointer-events: all;
}

@media (max-width: 870px) {
  .contain {
    min-height: 800px;
    height: 100vh;
  }

  .signin-signup {
    width: 100%;
    top: 95%;
    transform: translate(-50%, -100%);
    transition: 1s 0.8s ease-in-out;
  }

  .signin-signup,
  .contain.sign-up-mode .signin-signup {
    left: 50%;
  }

  .panels-contain {
    grid-template-columns: 1fr;
    grid-template-rows: 1fr 2fr 1fr;
  }

  .panel {
    flex-direction: row;
    justify-content: space-around;
    align-items: center;
    padding: 2.5rem 8%;
    grid-column: 1 / 2;
  }

  .right-panel {
    grid-row: 3 / 4;
  }

  .left-panel {
    grid-row: 1 / 2;
  }

  .image {
    width: 200px;
    transition: transform 0.9s ease-in-out;
    transition-delay: 0.6s;
  }

  .panel .content {
    padding-right: 15%;
    transition: transform 0.9s ease-in-out;
    transition-delay: 0.8s;
  }

  .panel h3 {
    font-size: 1.2rem;
  }

  .panel p {
    font-size: 0.7rem;
    padding: 0.5rem 0;
  }

  .btn.transparent {
    width: 110px;
    height: 35px;
    font-size: 0.7rem;
  }

  .contain:before {
    width: 1500px;
    height: 1500px;
    transform: translateX(-50%);
    left: 30%;
    bottom: 68%;
    right: initial;
    top: initial;
    transition: 2s ease-in-out;
  }

  .contain.sign-up-mode:before {
    transform: translate(-50%, 100%);
    bottom: 32%;
    right: initial;
  }

  .contain.sign-up-mode .left-panel .image,
  .contain.sign-up-mode .left-panel .content {
    transform: translateY(-300px);
  }

  .contain.sign-up-mode .right-panel .image,
  .contain.sign-up-mode .right-panel .content {
    transform: translateY(0px);
  }

  .right-panel .image,
  .right-panel .content {
    transform: translateY(300px);
  }

  .contain.sign-up-mode .signin-signup {
    top: 5%;
    transform: translate(-50%, 0);
  }
}

@media (min-width: 900px) {}

@media (max-width: 570px) {
  form {
    padding: 0 1.5rem;
  }

  .image {
    display: none;
  }

  .panel .content {
    padding: 0.5rem 1rem;
  }

  .contain {
    padding: 1.5rem;
  }

  .contain:before {
    bottom: 72%;
    left: 50%;
  }

  .contain.sign-up-mode:before {
    bottom: 28%;
    left: 50%;
  }
}

.input-field input {
  background-color: #f0f0f0 !important;
  border-radius: 14px;
}
</style>
