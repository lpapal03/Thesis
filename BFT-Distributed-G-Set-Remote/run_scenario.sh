#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Please provide an input parameter."
    exit 1
fi

param=$1

for thread_num in {1..3}; do
    sed -i "s/NUM_THREADS=[0-9]*/NUM_THREADS=$thread_num/" config # Update the number of threads in the config file
    echo Starting with $thread_num threads
    ansible-playbook -i ./hosts end.yml
    ansible-playbook -i ./hosts start.yml
    while true; do
        echo Waiting for clients to finish...
        client_nodes=(node1 node2 node3 node4)
        done_count=0
        for node in "${client_nodes[@]}"; do
            ssh "$node" "pgrep BFT-Distributed > /dev/null"
            if [ $? -eq 0 ]; then
                echo "Experiment is still running on $node"
            else
                echo "Experiment is done on $node"
                let "done_count+=1"
            fi
        done

        # If every node is done, run the second script
        if [[ $done_count -eq ${#client_nodes[@]} ]]; then
            echo Start copying result files...
            rm -rf results/scenario-$param/threads-$thread_num
            mkdir results/scenario-$param
            mkdir results/scenario-$param/threads-$thread_num
            nodes=$(grep -v "^#" hosts | grep -v "^$" | grep -v "^node0$" | grep -v "^\[" | cut -d" " -f2)
            for node in $nodes; do
                echo Copying from $node...
                scp loukis@$node:/users/loukis/Thesis/BFT-Distributed-G-Set-Remote/client/scenario_results* /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/results/scenario-$param/threads-$thread_num/
                scp loukis@$node:/users/loukis/Thesis/BFT-Distributed-G-Set-Remote/server/scenario_results* /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/results/scenario-$param/threads-$thread_num/
            done
            break
        fi

        sleep 30 # wait for 30 seconds before running the first script again
    done
done
