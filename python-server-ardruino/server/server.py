import socket

def start_server(host="localhost",port=8080):
    sock = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    sock.bind((host,port))
    sock.listen(1)
    print(f"Server on host: {host} and port: {port} ")
    while True:
        cli_sock,cli_addr = sock.accept()
        print(f"Socket accepted from {cli_addr}")
        data = cli_sock.recv(4096)
        if data: 
            print(f"Message: {data.decode()}")
            cli_sock.sendall("Success".encode())
        else:
            print("No data received")
        cli_sock.close()
    

start_server()
