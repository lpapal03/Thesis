import os

if __name__ == '__main__':
    os.system("ansible-playbook -i ./hosts ansible/start.yml -v")
    # os.system("cd /users/loukis/Thesis/BFT-Distributed-G-Set/client; /usr/local/go/bin/go run main.go")