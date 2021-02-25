module.exports = {
    "url": "/stats",
    "headers":{
        "Content-Type": "application/json"
    },
    "parametrization_test": {
        "smoke_test": {
            vus: 1,
            duration: '1m',
            thresholds: {
              'http_req_duration': ['p(99)<1500'], // 99% of requests must complete below 1.5s
            },
            tags: { 
                stack: 'xmen',
                layer: 'detector',
                env: 'dev',
                service: 'stats',
                type_test: 'smoke_test' 
            },
        },
    }
}