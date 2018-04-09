# Demo logger

```bash
 18:14:54 clab ./logger -s 10 -e="127.0.0.1:8081"
INFO[0000] Listening on127.0.0.1:8081                   
INFO[0004] Accept 127.0.0.1:50324 connection            
INFO[0027] Close 127.0.0.1:50324 connection, worked 23.335355923s 
INFO[0036] Accept 127.0.0.1:50328 connection            
INFO[0054] Close 127.0.0.1:50328 connection, worked 18.686647503s
INFO[0079] Accept 127.0.0.1:50332 connection            
INFO[0098] Close 127.0.0.1:50332 connection, worked 18.568321119s 
```

# Demo generator

```bash
 18:13:24 clab ./fibogen -s 8 -e "127.0.0.1:8081"
ERROR: generator: generate: got number of 129 bits, limit 129

 18:15:23 clab ./fibogen -s 10 -e "127.0.0.1:8081"
ERROR: generator: generate: got number of 129 bits, limit 129

 18:15:50 clab ./fibogen -s 100 -e "127.0.0.1:8081"
ERROR: generator: generate: got number of 129 bits, limit 129
```
