#!/bin/bash

echo "++ With Expect Header ++"
for port in "10080" "20080"; do
  curl -X POST -H "Expect: 100-continue" http://localhost:$port
done

echo "++ Without Expect Header ++"
for port in "10080" "20080"; do
  curl -X POST http://localhost:$port
done

