import type { IMaintenanceParams, IUnUpgradeItem } from "../interface/maintain";
import http, { type ResType } from "./http";

// 设备发现
export const GetDeviceList = async () => {
  return http.get("/discover");
};

// 上传文件
export const UploadFile = async (params: FormData, onUploadProgress: (progressEvent: ProgressEvent) => void) => {
  return http.upload("/upload", params, onUploadProgress);
};

/**
 * 防降级文件校验
 * @param params 
 * @returns 不可升级列表
 */
export const FileCheck = async (params: IMaintenanceParams): Promise<ResType<IUnUpgradeItem[]>> => {
  return http.post("/file_check", params);
};
// 开始维护
export const StartMaintain = async (params: IMaintenanceParams) => {
  return http.post("/maintain", params, 600000);
};
export const Reboot = async (data: { maintain_device_ip: string[] }) => {
  return http.post("/reboot", data);
};

export default {
  GetDeviceList,
  UploadFile,
  FileCheck,
  StartMaintain,
  Reboot,
};