import http from '../http.js'

/**
 * getDatabase 获取数据库
 * @param data
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getDatabase = () => {
    return http().request({
        url: 'api/get_database',
        method: 'GET'
    })
}

/**
 * getTableList 获取表列表
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getTableList = (data) => {
    return http().request({
        url: 'api/get_table_list',
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
        url: 'api/get_table',
        method: 'GET',
        params: data
    })
}