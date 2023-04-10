#!/bin/bash

rm -rf results
mkdir results

./run_scenario normal
./run_scenario mute
./run_scenario malicious
./run_scenario half_and_half