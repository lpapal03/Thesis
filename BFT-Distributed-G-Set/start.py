import os
import sys

def StartNormal():
    os.system("ansible-playbook -i ./hosts ansible/start_normal.yml -v")
    os.system("cd /users/loukis/Thesis/BFT-Distributed-G-Set/client; /usr/local/go/bin/go run main.go")
 

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("No arguments provided")
        exit(1)
    if sys.argv[1] == "n":
        StartNormal()
    else:
        print("Invalid selection")
