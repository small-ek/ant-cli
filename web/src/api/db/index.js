import http from '../http.js'

/**
 * getDatabase 获取数据库
 * @param data
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getDatabase = () => {
    return http().request({
        url: 'api/database',
        method: 'GET'
    })
}

/**
 * getTableList 获取表列表
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getTableList = (data) => {
    return http().request({
        url: 'api/table-list',
        method: 'GET',
        params: data
    })
}

/**
 * getTable 获取表
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getTable = (data) => {
    return http().request({
        url: 'api/table-field',
        method: 'GET',
        params: data
    })
}

/**
 * previewCode 预览代码
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const previewCode = (data) => {
    return http().request({
        url: 'api/code',
        method: 'POST',
        data: data
    })
}

/**
 * GenerateCode 生成代码
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const generateCode = (data) => {
    return http().request({
        url: 'api/generate_code',
        method: 'POST',
        data: data
    })
}