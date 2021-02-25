module.exports = {
    "url": "/mutant",
    "headers":{
        "Content-Type": "application/json"
    },
    "parametrization_test": {
        "smoke_test": {
            vus: 1,  // 1 user looping for 1 minute
            duration: '1m',
            tags: { 
                stack: 'xmen',
                layer: 'detector',
                env: 'dev',
                service: 'dna',
                type_test: 'smoke_test' 
            },
        },
        "load_test": {
            stages: [
                { duration: "5m", target: 50 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
                { duration: "10m", target: 50 }, // stay at 100 users for 10 minutes
                { duration: "5m", target: 0 }, // ramp-down to 0 users
              ],
            tags: {
                stack: 'xmen',
                layer: 'detector',
                env: 'dev',
                service: 'dna',
                type_test: 'load_test'
            },
        },
        "stress_test":{
          stages: [
              { duration: '2m', target: 20 },
              { duration: '5m', target: 20 },
              { duration: '2m', target: 40 },
              { duration: '5m', target: 40 },
              { duration: '2m', target: 60 },
              { duration: '5m', target: 60 },
              { duration: '2m', target: 80 },
              { duration: '5m', target: 80 },
              { duration: '2m', target: 100 },
              { duration: '5m', target: 100 },
              { duration: '10m', target: 0 },
            ],
            tags: {
                stack: 'xmen',
                layer: 'detector',
                env: 'dev',
                service: 'dna',
                type_test: 'stress_test'
            },
      }
    }
}