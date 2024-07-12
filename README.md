# Drop

Transfer files securely while using power of ssh.

## Usage

### Run server

```bash
$ drop -h
usage: ./drop <PORT> <KEYFILE>
$ drop 2222 keys/id_ed25519
2024/07/12 03:01:32 Drop v0.0.1
2024/07/12 03:01:32 Listening Port: 2222
2024/07/12 03:01:32 Using Key File: keys/id_ed25519
```

### Send file

```bash
$ ssh -T server < file
New drop created: e1c2a171f5b90666
Waiting for receiver...
```

### Receive file

```bash
$ ssh -T server e1c2a171f5b90666 > file     # Save file
$ ssh -T server e1c2a171f5b90666 | tee file # Print and save file 
```

## Tips

- #### Sending live message 
    Just start ssh conn without giving file. And type messages after client connected.
    
- #### Anonim file sending & receiving
    ```bash
    $ torsocks ssh -T server < file
    ```

## TODO

- [ ] Add asciinema videos to README
- [ ] Create persistent mode to drops




