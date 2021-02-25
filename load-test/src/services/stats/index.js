import * as stats from './stats.js'
import * as parametrization_data from '../../../resources/services/stats.js'
import { check, sleep } from 'k6';

export let options = parametrization_data.parametrization_test[__ENV.TYPE_TEST]

export default () => {
  let res = stats.execute()
  check(res, { 'status was 200': r => r.status == 200 })
  sleep(1)
}