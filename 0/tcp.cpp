#include <cstring>
#include <iostream>
#include <netdb.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <thread>

void handle_connection(int sock_fd) {

  std::cout << "new handle connection on " << sock_fd << std::endl;
  int buffer[1024];

  int recived_buffer = 0;
  while (true) {
    recived_buffer = recv(sock_fd, &buffer, sizeof buffer, 0);

    std::cout << "received on " << sock_fd << std::endl;
    if (recived_buffer <= 0) {
      std::cout << "recv error is finnished: " << recived_buffer << std::endl;
      break;
    }
    std::cout << "sending " << sock_fd << std::endl;
    send(sock_fd, &buffer, recived_buffer, 0);
    std::cout << "sent " << sock_fd << std::endl;
  }
}

int main() {
  constexpr char *MY_PORT = (char *)"7";
  constexpr int BACKLOG = 5;
  struct sockaddr_storage their_addr[5];
  struct addrinfo hints, *res;
  socklen_t addr_size[5];
  int sockfd, new_fd;

  std::memset(&hints, 0, sizeof hints);
  hints.ai_family = AF_UNSPEC;
  hints.ai_socktype = SOCK_STREAM;
  hints.ai_flags = AI_PASSIVE;

  int status = getaddrinfo(NULL, (char *)MY_PORT, &hints, &res);
  if (status != 0) {
    std::cout << "getaddrinfo error: " << gai_strerror(status) << std::endl;
    return 1;
  }

  // make socket
  sockfd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
  if (sockfd == -1) {
    std::cout << "socket error " << std::endl;
    return 2;
  }

  if (bind(sockfd, res->ai_addr, res->ai_addrlen) == -1) {
    std::cout << "bind error " << std::endl;
    return 3;
  }

  if (listen(sockfd, BACKLOG) == -1) {
    std::cout << "listen error " << std::endl;
    return 4;
  }

  int i = 0;
  std::thread *threads[5]{};
  while (true) {
    addr_size[i] = sizeof their_addr[i];

    std::cout << "new connection begining" << std::endl;
    int new_sock_fd =
        accept(sockfd, (struct sockaddr *)&their_addr[i], &addr_size[i]);
    std::cout << "new connection" << std::endl;

    threads[i] = new std::thread(handle_connection, new_sock_fd);

    std::cout << "new trhead" << std::endl;

    i++;

    std::cout << "HELLO AFTER I++" << std::endl;
    if (new_sock_fd < 0) {
      std::cout << "accept error" << std::endl;
      return 5;
    }
    if (i >= 5) {
      std::cout << "we breaking loop" << std::endl;
      break;
    }

    std::cout << "HELLO xxx IF" << std::endl;
  }

  std::cout << "CLOSED WHILE" << std::endl;

  for (auto t : threads) {
    if (t != nullptr) {
      t->join();
    }
  }

  std::cout << "all good";
  return 0;
}
