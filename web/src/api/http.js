import axios from "axios";
import {timeout, url} from "../config/index"


function http() {
    const headers = {}
    var http = axios.create({
        baseURL: url,
        timeout: timeout,
        headers: headers,
    });

    //添加请求拦截器
    http.interceptors.request.use(
        (cfg) => {
            return cfg;
        },
        (error) => {
            console.log(error);
            // @ts-ignore
            return Promise.reject(error);
        }
    );

    //添加响应拦截器
    http.interceptors.response.use(
        (response) => {
            return response;
        },
        (error) => {
            console.log(error.message);
            return Promise.reject(error);
        }
    );
    return http;
}

export default http;