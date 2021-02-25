import http from 'k6/http';
import * as general_data from '../../../resources/services/general.js'
import * as stats_data from '../../../resources/services/stats.js'

const HEADERS = {
    headers: stats_data.headers
}

export let execute = () => {
    return http.get(`${general_data.base_url}${stats_data.url}`, HEADERS)
}