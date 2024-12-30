import axios, { type InternalAxiosRequestConfig } from "axios";
import { ElMessage } from "element-plus";

axios.defaults.baseURL =
  import.meta.env.MODE === "development" ? "/zmos_ze05" : "/";
axios.defaults.timeout = 10000;
axios.defaults.headers.post["Content-Type"] = "application/json;charset=UTF-8";
axios.interceptors.request.use(
  (config): InternalAxiosRequestConfig<any> => {
    // 上传文件的接口
    // if (config.url === "/zmos_ze05/maintain") {
    //   axios.defaults.timeout = 600000;
    // }
    return config;
  },
  (error) => {
    return error;
  }
);
// 存储每个 URL 错误的时间戳
const errorCache = new Map();

// 响应拦截
axios.interceptors.response.use(
  (res) => {
    if (res && res.data && res.data.code !== 0 && res.data.code !== 200) {
      const errorKey = res.config.url;
      if (res.data.msg && res.data.msg !== "ok") {
        // 错误响应
        const currentTime = Date.now();
        // 如果缓存中没有这个 URL 或者已经超出冷却时间
        if (
          errorCache.get(errorKey) == undefined ||
          currentTime - errorCache.get(errorKey) > 2000
        ) {
          ElMessage.warning(res.data.msg); // 弹出错误消息
          // 更新缓存时间
          errorCache.set(errorKey, Date.now());
        }
      } else {
        errorCache.set(errorKey, undefined);
      }
    }
    return res;
  },
  (err) => { }
);

export interface ResType<T> {
  code: number;
  msg: string;
  data?: T | any;
  err?: string;
}
interface RequestParams {
  signal?: AbortSignal;
  [key: string]: any; // 允许其他参数
}

interface Http {
  get<T>(url: string, params?: RequestParams, timeout?: number): Promise<ResType<T>>;

  post<T>(url: string, params?: RequestParams, timeout?: number): Promise<ResType<T>>;

  delete<T>(url: string, params?: RequestParams): Promise<ResType<T>>;

  put<T>(url: string, params?: RequestParams): Promise<ResType<T>>;

  upload<T>(url: string, params: unknown, axiosOptions: any): Promise<ResType<T>>;
}

const http: Http = {
  get(url, params, timeout = 20000) {
    const { signal, ...restParams } = params || {};
    return new Promise((resolve, reject) => {
      axios
        .get(url, { params: restParams, signal, timeout: timeout }) // 将 signal 传递给 Axios
        .then((res) => {
          res?.data ? resolve(res.data) : reject("");
        })
        .catch((err) => {
          // 处理请求取消的情况
          if (err.name === "CanceledError") {
            console.log("Request canceled");
          } else {
            reject(err);
          }
        });
    });
  },
  post(url, params, timeout = 10000) {
    const { signal, ...restParams } = params || {};

    return new Promise((resolve, reject) => {
      axios
        .post(url, JSON.stringify(restParams), { signal, timeout })
        .then((res) => {
          resolve(res.data);
        })
        .catch((err) => {
          reject(err.data);
        });
    });
  },
  delete(url, params) {
    const { signal, ...restParams } = params || {};
    return new Promise((resolve, reject) => {
      axios
        .delete(url, { params: restParams, signal })
        .then((res) => {
          resolve(res.data);
        })
        .catch((err) => {
          reject(err.data);
        });
    });
  },
  put(url, params) {
    const { signal, ...restParams } = params || {};
    return new Promise((resolve, reject) => {
      axios
        .put(url, params, { signal })
        .then((res) => {
          resolve(res.data);
        })
        .catch((err) => {
          reject(err.data);
        });
    });
  },
  upload(url, formData, onUploadProgress) {
    return new Promise((resolve, reject) => {
      axios({
        method: "post",
        url,
        timeout: 600000,
        headers: {
          "Content-Type": "multipart/form-data"
        },
        data: formData,
        onUploadProgress: onUploadProgress
      })
        .then((res) => {
          resolve(res.data);
        })
        .catch((err) => {
          reject(err.data);
        });
    });
  }
};
export default http;
export const BaseAxios = axios;
