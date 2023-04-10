#!/bin/bash

rm logfile
nohup ./run_experiment.sh > logfile 2>&1 &
