import os

def IsHostActive(hostname):
    response = os.system("ping -c 1 " + hostname)
    if response == 0:
        return True
    else:
        return False

if __name__ == '__main__':
    f = open("hosts", "w")
    for i in range(9999):
        active = IsHostActive("node" + str(i))
        if active:
            f.write("node" + str(i) + "\n")
        if not active:
            break
    f.close()
    os.system("ansible-playbook -i ./hosts update_repo.yml")
    os.remove("hosts")
    