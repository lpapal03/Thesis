#!/bin/bash

rm -rf results
mkdir results

./run_scenario.sh normal
./run_scenario.sh mute
./run_scenario.sh malicious
./run_scenario.sh half_and_half