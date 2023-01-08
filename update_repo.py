import os

if __name__ == '__main__':
    os.system("ansible-playbook -i ./hosts update_repo.yml -v")