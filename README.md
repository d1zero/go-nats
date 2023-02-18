Simple NATS server & client implementation

To run app run
```shell
make nats
```
and in different terminals exec:
```shell
make server
```
and
```shell
make client
```

Client will start logging:
```
sended
```

and server will start logging smth like:
```
2023-02-18 20:46:55.124929 +0300 MSK m=+0.007460168
2023-02-18 20:46:56.126047 +0300 MSK m=+1.008563251
2023-02-18 20:46:57.127176 +0300 MSK m=+2.009677709
```