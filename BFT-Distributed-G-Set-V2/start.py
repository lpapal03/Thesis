import os

HOSTNAME = os.uname()[1].split('.')[0]

def GetHosts():
    hosts = []
    for i in range(9999):
        hostname = "node" + str(i)
        response = os.system("ping -c 1 " + hostname)
        if response == 0:
            hosts.append(hostname)
        else:
            return hosts

def StartNormalInteractive(N=None, c=None):
    servers = GetHosts()

    f = open("hosts", "w")
    f.write("[clients]\n")
    f.write(HOSTNAME + "\n\n")
    f.write("[servers]\n")
    for h in servers:
        if h != HOSTNAME:
            f.write(h+"\n")
    f.close()

    f = open("start_scenario.yml", "w")
    f.write("---\n")
    f.write("- name: Start servers\n")
    f.write("  hosts: servers\n")
    f.write("  become: true\n")
    f.write("  tasks:\n")

    f.write("    - name: Fetch file from" + HOSTNAME + "\n")
    f.write("      fetch:\n")
    f.write("        src: /users/loukis/Thesis/BFT-Distributed-G-Set-V2/hosts\n")
    f.write("        dest: /users/loukis/Thesis/BFT-Distributed-G-Set-V2/hosts\n")
    f.write("      when: ansible_hostname == '"+HOSTNAME+"'")
    
    f.write("    - name: Start servers\n")
    f.write("      command: /usr/local/go/bin/go run /users/loukis/Thesis/BFT-Distributed-G-Set-V2/server/main.go\n")
    f.close()

    os.system("ansible-playbook -i ./hosts start_scenario.yml -v")
    # os.remove("hosts")
    # os.remove("start_scenario.yml")
    # os.system("go run /users/loukis/Thesis/BFT-Distributed-G-Set-V2/client/main.go")

def StartMutes(interactive=True, N=None):
    pass

def StartHalfAndHalf(interactive=True, N=None):
    pass
 

if __name__ == '__main__':
    # ask scenario
    # ask N (must be greater than hosts)
    # begin
    StartNormalInteractive()
