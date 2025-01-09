import http from "./http";

const getBaseUrl = ()=>{
  // @ts-ignore
  return import.meta.env.MODE === "development" ? "http://127.0.0.1:3028" : window.location.origin
}

export const reload = async () => {
  return http.get( `${getBaseUrl()}/other/reload`);
};

export const DbsByDsn = async (data) => {
  return http.post( `${getBaseUrl()}/other/getDbsByDsn`,data);
};

export const s3list = async () => {
  return http.get(`${getBaseUrl()}/s3/list`);
};

export const s3delete = async (id) => {
  return http.delete(`${getBaseUrl()}/s3/delete?ID=${id}`);
};

export const s3update = async (data) => {
  return http.post(`${getBaseUrl()}/s3/update`,data);
};

export const dbList = async () => {
  return http.get(`${getBaseUrl()}/db/list`);
};

export const dbDelete = async (id) => {
  return http.delete(`${getBaseUrl()}/db/delete?ID=${id}`);
};

export const dbUpdate = async (data) => {
  return http.post(`${getBaseUrl()}/db/update`,data);
};

