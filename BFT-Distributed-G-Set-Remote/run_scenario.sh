#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Usage: $0 <input_parameter>"
  exit 1
fi

input_parameter=$1

for i in {1..5}
do
  ./run_scenario_once.sh "$input_parameter"_iteration_"$i"
done

