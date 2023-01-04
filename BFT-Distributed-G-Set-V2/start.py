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

    os.system("ansible-playbook -i ./hosts start_scenario.yml -v")
    os.remove("hosts")
    os.system("cd /users/loukis/Thesis/BFT-Distributed-G-Set-V2/client /usr/local/go/bin/go run main.go")

def StartMutes(interactive=True, N=None):
    pass

def StartHalfAndHalf(interactive=True, N=None):
    pass
 

if __name__ == '__main__':
    # ask scenario
    # ask N (must be greater than hosts)
    # begin
    StartNormalInteractive()
