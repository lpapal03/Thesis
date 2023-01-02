import os

def GetHosts():
    hosts = []
    for i in range(9999):
        hostname = "node" + str(i)
        response = os.system("ping -c 1 " + hostname)
        if response == 0:
            hosts.append(hostname)
        else:
            return hosts
    


if __name__ == '__main__':
    # for normal execution
    hosts = GetHosts()
    f = open("hosts", "w")
    for h in hosts:
        f.write(h+"\n")
    f.close

    f = open("start_servers.yml", "w")
    f.write("---\n")
    f.write("- name: My Playbook\n")
    f.write("  hosts: all\n")
    f.write("  become: true\n")
    f.write("  tasks:\n")
    f.write("    - name: Start servers\n")
    f.write("      script: /users/loukis/Thesis/BFT-Distributed-G-Set-V2/server/main.go\n")
    f.close()

    os.system("ansible-playbook -i ./hosts start_servers.yml")
    # os.remove("hosts")