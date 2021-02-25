import * as parametrization_data from '../../../resources/services/dna.js'
import * as dna from './dna.js'
import { check, sleep } from 'k6';

export let options = parametrization_data.parametrization_test[__ENV.TYPE_TEST]

export default () => {
  const isMutant = Math.random() < 0.5;
  let res = dna.execute(isMutant)
  if (isMutant){
    check(res, { 'status was 200': r => r.status == 200 })
  }else {
    check(res, { 'status was 403': r => r.status == 403 })
  }

  sleep(1);
}

