#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Please provide an input parameter."
    exit 1
fi

param=$1

for thread_num in {1..5}; do
    sed -i "s/NUM_THREADS=[0-9]*/NUM_THREADS=$thread_num/" config # Update the number of threads in the config file
    echo Starting with $thread_num threads
    ansible-playbook -i ./hosts end.yml
    ansible-playbook -i ./hosts start.yml
    while true; do
        echo Waiting for clients to finish...
        # Get the list of nodes under the [clients-automated] tag from a remote machine
        client_nodes=(node1 node2 node3 node4)
        # Check if every node is done with the process "BFT-Distributed-G-Set-Remote"
        done_count=0
        for node in $client_nodes; do
            ssh "$node" "pgrep BFT-Distributed > /dev/null"
            if [ $? -eq 0 ]; then
                echo "Experiment is running on $node"
            else
                echo "Experiment is not running on $node"
                let "done_count+=1"
            fi

        # If every node is done, run the second script
        if [[ $done_count -eq ${#client_nodes[@]} ]]; then
            rm -rf results/scenario-$param/threads-$thread_num
            mkdir results/scenario-$param
            mkdir results/scenario-$param/threads-$thread_num
            nodes=$(grep -v "^#" hosts | grep -v "^$" | grep -v "^node0$" | grep -v "^\[" | cut -d" " -f2)
            for node in $nodes; do
                scp loukis@$node:/users/loukis/Thesis/BFT-Distributed-G-Set-Remote/client/scenario_results* /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/results/scenario-$param/threads-$thread_num/
                scp loukis@$node:/users/loukis/Thesis/BFT-Distributed-G-Set-Remote/server/scenario_results* /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/results/scenario-$param/threads-$thread_num/
            done
            break
        fi

        sleep 10 # wait for 5 seconds before running the first script again
    done
done
