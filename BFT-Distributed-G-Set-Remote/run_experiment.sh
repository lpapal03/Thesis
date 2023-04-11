#!/bin/bash

# Make sure there is no other hosts file in the directory

NUM_ITERATIONS=3

rm -rf results
mkdir results

echo Running normal
rm hosts
cp hosts_configurations/hosts_normal /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/hosts
for (( i=1; i<=NUM_ITERATIONS; i++ ))
do
  ./run_scenario.sh normal_iteration_"$i"
done
rm hosts

echo Running mute
rm hosts
cp hosts_configurations/hosts_mute /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/hosts
for (( i=1; i<=NUM_ITERATIONS; i++ ))
do
  ./run_scenario.sh mute_iteration_"$i"
done
rm hosts

echo Running malicious
rm hosts
cp hosts_configurations/hosts_malicious /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/hosts
for (( i=1; i<=NUM_ITERATIONS; i++ ))
do
  ./run_scenario.sh malicious_iteration_"$i"
done
rm hosts

echo Running half_and_half 
rm hosts
cp hosts_configurations/hosts_half_and_half /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/hosts
for (( i=1; i<=NUM_ITERATIONS; i++ ))
do
  ./run_scenario.sh half_and_half_iteration_"$i"
done
rm hosts