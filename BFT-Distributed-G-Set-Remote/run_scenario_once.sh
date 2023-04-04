#!/bin/bash

killall -9 2-Atomic-Adds

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
        nodes=($(awk '/^\[clients-automated\]/{flag=1;next}/^\[/{flag=0}flag{print $1}' hosts))
        # Check if every node is done with the process "2-Atomic-Adds"
        done_count=0
        for node in $nodes; do
            ssh $node "pgrep 2-Atomic-Adds > /dev/null && echo \"Node $node is done\" || echo \"Node $node is not done\""
        done | grep -v "Node.*is done" || ((done_count++))

        # If every node is done, run the second script
        if [[ $done_count -eq 0 ]]; then
            rm -rf results/scenario-$param/threads-$thread_num
            mkdir results/scenario-$param
            mkdir results/scenario-$param/threads-$thread_num
            nodes=$(grep -v "^#" hosts | grep -v "^$" | grep -v "^node0$" | grep -v "^\[" | cut -d" " -f2)
            for node in $nodes; do
                scp loukis@$node:/users/loukis/Thesis/BFT-Distributed-G-Set-Remote/client/scenario_results.txt /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/results/scenario-$param/threads-$thread_num/$node.txt
                scp loukis@$node:/users/loukis/Thesis/BFT-Distributed-G-Set-Remote/server/scenario_results.txt /users/loukis/Thesis/BFT-Distributed-G-Set-Remote/results/scenario-$param/threads-$thread_num/$node.txt
            done
            break
        fi

        sleep 5 # wait for 5 seconds before running the first script again
    done
done
