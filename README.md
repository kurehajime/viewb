# viewb

[![CircleCI](https://circleci.com/gh/kurehajime/viewb.svg?style=svg)](https://circleci.com/gh/kurehajime/viewb)

Convert the command to a web server

![viewb](https://cloud.githubusercontent.com/assets/4569916/9249386/589c2126-41ff-11e5-9e6f-b12daa6aadf0.png)

## Usage

```
$ viewb  <COMMAND> <ARGS>
```

Option:  
  -p=\<PORT\>: Port(default:8080)  
  -o: Open web browser  
  -user=\<USERNAME\>: Basic Authentication user name  
  -pass=\<PASSWORD\>: Basic Authentication password  
  -e=\<ENCODING\>: input encoding  


## How to install

[Download](https://github.com/kurehajime/viewb/releases)

or

Build yourself (Go lang).

```
go get -u github.com/kurehajime/viewb
```


## Example

##### Example 1 :Command to web server

```sh
$ viewb -p 8080 ls -la

http://localhost:8080
Stop: Ctrl+C
```

Open in browser http://localhost:8080

```
total 32  
drwxr-xr-x   6 user  staff   204  8  6 20:19 .  
drwx------+ 11 user  staff   374  8  6 20:17 ..  
-rw-r--r--@  1 user  staff  6148  8  6 20:19 .DS_Store  
-rw-r--r--   1 user  staff     5  8  6 20:18 Untitled-1.txt    
-rw-r--r--   1 user  staff  1557  8  6 20:19 Untitled-2.txt  
drwxr-xr-x   2 user  staff    68  8  6 20:19 test  
```

##### Example 2 :Script to web server

```sh
$ viewb ./HelloworldAndPingOne.sh

http://localhost:8080
Stop: Ctrl+C
```

Open in browser http://localhost:8080

```
Hello World!
PING 8.8.8.8 (8.8.8.8): 56 data bytes
64 bytes from 8.8.8.8: icmp_seq=0 ttl=54 time=60.380 ms

--- 8.8.8.8 ping statistics ---
1 packets transmitted, 1 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 60.380/60.380/60.380/0.000 ms
```

##### Example 3 :Basic Authentication

```sh
$ viewb -user laputa -pass balse echo booomb!

http://localhost:8080
Stop: Ctrl+C
```

Open in browser http://localhost:8080  
And login.

##### Example 5 :open browser 

```sh
$ viewb -o echo yah!

http://localhost:8080
Stop: Ctrl+C
```

Open automatically.




