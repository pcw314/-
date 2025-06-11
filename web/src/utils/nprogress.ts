// 页面进度条组件封装
import NProgress from "nprogress";
import "nprogress/nprogress.css";
NProgress.configure({ showSpinner: false }); // 显示右上角螺旋加载提示
// 打开进度条
export const NProgressStart = (color: string = "#5180ee") => {
  document
    .getElementsByTagName("body")[0]
    .style.setProperty("--nprogress-background-color", color);
  NProgress.start();
};

// 关闭进度条
export const NProgressDone = () => {
  NProgress.done();
};
