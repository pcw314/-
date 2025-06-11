import { watch, ref } from "vue";
import { isString } from "@/utils/index";
import { WEB_TITLE } from "@/constants/index";

export const useTitle = (newTitle?: string) => {
  const title = ref(newTitle ? newTitle : WEB_TITLE);
  watch(
    title,
    (n, o) => {
      if (isString(n) && n !== o && document) {
        document.title = n;
      }
    },
    { immediate: true }
  );

  return title;
};
