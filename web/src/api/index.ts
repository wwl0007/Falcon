import config from '@/config'
import axios from 'axios'

function get(path: string, ...rest: any[]) {
    return axios.get(config.API_HOST + path, ...rest);
}

function del(path: string, ...rest: any[]) {
    return axios.delete(config.API_HOST + path, ...rest);
}

export default {
    get,
    delete: del
}
