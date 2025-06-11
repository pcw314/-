import { defineStore } from "pinia";
import { connect, io } from "socket.io-client";
import { useUserStore } from "@/store/module/user";
import { getChatList } from "@/api/center";

const userStore = useUserStore();
type sokcetType = {
  soc: any;
  unread: Number;
  timer?: any;
  chatList: any[];
};
export const useUserSo = defineStore("so", {
  state: (): sokcetType => {
    return {
      soc: null,
      unread: 0,
      timer: null,
      chatList: localStorage.getItem('chatList')? localStorage.getItem('chatList'):[],
    };
  },
  getters: {
    getSoc(): Object | null {
      return this.soc;
    },
    getUnread(): Object | null {
      return this.unread;
    },
  },
  actions: {
    connect() {
      let url =
        "http://127.0.0.1:8080/message/ws?token=" + userStore.token;
      let socket = (this.soc = new WebSocket(url));
      socket.onopen = function (event) {
        console.log("Connection open...");
        socket.send("Hello Server!");
      };

      // 收到消息时触发
      socket.onmessage = function (event) {
        window.dispatchEvent('changeChat');
        console.log("Message from server ", event.data);
      };

      // 连接关闭时触发
      socket.onclose = function (event) {
        console.log("Connection closed");
      };

      // 发生错误时触发
      socket.onerror = function (error) {
        console.log(error);
      };
    },
    reconnect() {
      setTimeout(() => {
        this.connect();
      }, 1000);
    },
    close() {
      this.soc.close();
    },
    getChatList() {
      getChatList().then((res) => {
        this.chatList = []
        if (res.code == 200) {
          if (res.data) {
            localStorage.setItem("chatList", JSON.stringify(res.data));
            this.chatList = res.data
          }
        }
      });
    },
    
  },
});
