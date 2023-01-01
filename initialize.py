import os

def check_ping(hostname):
    response = os.system("ping -c 1 " + hostname)
    # and then check the response...
    if response == 0:
        pingstatus = "Network Active"
    else:
        pingstatus = "Network Error"
    
    return pingstatus


if __name__ == '__main__':
    print(check_ping("node0"))