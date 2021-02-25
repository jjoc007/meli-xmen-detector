#!/bin/bash

# Example stats execution smoke test
k6 run -e TYPE_TEST=smoke_test --summary-export results/stats/summary.json src/services/stats/index.js