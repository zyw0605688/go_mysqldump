import axios from "axios";

const instance = axios.create({
  timeout: 10000,
  headers: { "Content-Type": "application/json;charset=UTF-8" }
});

instance.interceptors.response.use((res) => {
  return res.data;
});

export default instance;
