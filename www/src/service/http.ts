import axios from "axios";


const instance = axios.create({
  baseURL: 'http://127.0.0.1:3028',
  timeout: 10000,
  headers: {'Content-Type': 'application/json;charset=UTF-8'}
});

instance.interceptors.response.use((res)=>{
  return res.data
})

export default instance
