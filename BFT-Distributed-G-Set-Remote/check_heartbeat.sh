#!/bin/bash

# Loop through nodes
for num in {0..29}; do
  ssh node$num "pgrep BFT-Distributed > /dev/null"
  if [ $? -eq 0 ]; then
    echo "TEST is running on node$num"
  else
    echo "TEST is not running on node$num"
  fi
done

# Check if all nodes are done
if ! pgrep TEST > /dev/null; then
  echo "Done"
fi
