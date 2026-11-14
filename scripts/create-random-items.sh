#!/bin/bash
#numbers=(2024 2025 2026)
#picked=$(shuf -n 1 -e "${numbers[@]}")
#echo "$picked"
for e in $(shuf -i 1-6 -n 3 -r); do
    echo $e
done

