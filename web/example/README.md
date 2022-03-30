# web example

This is a very simple example of how the `web` package can be used to build a web-based front end for GDB.

![Screenshot](http://i.imgur.com/Ql48hXT.png)

## Usage

1. move into this directory;

2. compile the dummy test program with:

    ```
    make -C test/ dummy
    ```

3. fetch the WebSocket library:

    ```
    go get golang.org/x/net/websocket
    ```

4. start the server with:

   ```
   go run main.go
   ```

   then load `http://localhost:8080/` in your browser

5. load the test program using a relative path `test/dummy`;

6. the rest should be straightforward...
