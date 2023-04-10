#!/bin/bash

# Make sure there is no other hosts file in the directory

rm -rf results
mkdir results

echo Running normal
mv hosts_normal hosts
for i in {1..5}
do
  ./run_scenario.sh normal_iteration_"$i"
done
mv hosts hosts_normal

echo Running mute
mv hosts_mute hosts
for i in {1..5}
do
  ./run_scenario.sh mute_iteration_"$i"
done
mv hosts hosts_mute

echo Running malicious
mv hosts_malicious hosts
for i in {1..5}
do
  ./run_scenario.sh malicious_iteration_"$i"
done
mv hosts hosts_malicious

echo Running half_and_half 
mv hosts_half_and_half hosts
for i in {1..5}
do
  ./run_scenario.sh half_and_half_iteration_"$i"
done
mv hosts hosts_half_and_half
