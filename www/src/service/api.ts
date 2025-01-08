import http from "./http"

export const reload = async ()=>{
  return http.get("/other/reload")
}
