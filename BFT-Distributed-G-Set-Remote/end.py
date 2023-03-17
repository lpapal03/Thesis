import os

if __name__ == '__main__':
    os.system("ansible-playbook -i ./hosts end.yml -v")