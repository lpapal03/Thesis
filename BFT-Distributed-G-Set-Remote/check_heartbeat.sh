#!/bin/bash

# Loop through nodes
done_counter=0
for num in {1..29}; do
  ssh node$num "pgrep BFT-Distributed > /dev/null"
  if [ $? -eq 0 ]; then
    echo "Experiment is running on node$num"
  else
    echo "Experiment is not running on node$num"
    let "done_counter+=1"
  fi
done

echo "Machines done:" $done_counter

