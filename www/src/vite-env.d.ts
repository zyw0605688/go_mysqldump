/// <reference types="vite/client" />

// 引入import，无法视为全局使用，单独写这个定义
// 全局请写入外部的env.d.ts
import type { IConfigApi } from "./service/visualConfig/index.ts";

declare global {
  interface Window {
    configApi: IConfigApi | any;
  }
}

export {};
// 使用export {} 将文件转为模块，在其他文件识别这个全局类型
