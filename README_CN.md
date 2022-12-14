# go-p2p
一个用 Go 实现的 p2p 网络。

- [English](./README.md)


## 1. 实现原理

1. 建立 UDP 服务器 `S`;
2. 建立 `A`、`B` 客户端，分别与 `S` 会话，记为 `SA`、`SB`；
3. `S` 作为中介，将 `A`的 **IP+PORT** 通过 `SB` 告诉 `B`，同理，将 `A` 的也告诉 `B`;
4. `A` 向 `B` 公网地址发送一个 **UDP** 包，代表握手，打通 A->B 的路径；
5. `B` 向 `A` 公网地址发送也一个 **UDP** 包，A <-> B 的会话就建立完毕了。


## 2. 运行步骤

1. 运行服务端
    ```shell
    cd server
    go build server.go
    ./server
    ```
2. 运行两个客户端
    ```shell
    cd client
    go build client.go
    ./client A 127.0.0.1 9527 1234
    ./client B 127.0.0.1 9527 4321
    ```
3. 使用 p2p 网络进行通讯