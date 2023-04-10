#!/bin/bash

# Make sure there is no other hosts file in the directory

rm -rf results
rm logfile
mkdir results

echo Running normal
mv hosts_normal hosts
nohup ./run_scenario.sh normal > logfile 2>&1
mv hosts hosts_normal

echo Running mute
mv hosts_mute hosts
nohup ./run_scenario.sh mute > logfile 2>&1
mv hosts hosts_mute

echo Running malicious
mv hosts_malicious hosts
nohup ./run_scenario.sh malicious > logfile 2>&1
mv hosts hosts_malicious

echo Running half_and_half
mv hosts_half_and_half hosts
nohup ./run_scenario.sh half_and_half > logfile 2>&1
mv hosts hosts_half_and_half