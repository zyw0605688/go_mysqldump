/// <reference types="vite/client" />

// 没有引入import，将视为全局使用
declare module "*.vue" {
  import type { DefineComponent } from "vue";
  const component: DefineComponent<{}, {}, Component>;
  export default component;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

declare module "element-plus/dist/locale/zh-cn.mjs";
