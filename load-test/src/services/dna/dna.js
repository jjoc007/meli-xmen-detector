import http from 'k6/http';
import * as general_data from '../../../resources/services/general.js'
import * as dna_data from '../../../resources/services/dna.js'

const HEADERS = {
    headers: dna_data.headers
}

export let execute = (isMutant) => {
    const BODY_IS_MUTANT = {
        "dna": [
            "CCCCGA", "CAGTCC", "GTTAGT", "CGAAGG", "TGTCTA", "TCCCCC"
        ]
    }

    const BODY_IS_NOT_MUTANT = {
        "dna": [
            "CCCCGA", "CAGTCC", "GTTAGT", "CGAAGG", "TGTCTA", "TCCATC"
        ]
    }

    return http.post(`${general_data.base_url}${dna_data.url}`,
        JSON.stringify(isMutant ? BODY_IS_MUTANT : BODY_IS_NOT_MUTANT),
        HEADERS)
}
