# go-p2p
a peer to peer network implemented in Go.

- [中文](./README_CN.md)



## Implementation Principle

1. Create UDP server `S`; 
2. Establish `A` and `B` clients, and have sessions with `S` respectively, denoted as `SA` and `SB`; 
3. `S` acts as an intermediary to tell `B` the **IP+PORT** of `A` through `SB`, and similarly, the `A` is also told to `B`; 
4. `A` sends a **UDP** packet to the `B` public address, representing the handshake and opening the path of **A-&gt;B**; 
5. `B` sends a **UDP** packet to the `A` public address, and **A &lt;-&gt; B** session is established. 



## 2. Run step 

1. Run the server

   ```shell
   cd server
   go build server.go
   ./server
   ```

2. Run two clients

   ```shell
   cd client
   go build client.go
   ./client A 127.0.0.1 9527 1234
   ./client B 127.0.0.1 9527 4321
   ```

3. Use P2P networks for communication