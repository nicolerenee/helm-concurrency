#!/bin/bash

echo "Removing helm releases used for testing"
for i in {1..20}
do
  helm del --purge test$i > /dev/null 2>&1
done
