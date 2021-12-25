# ep

A patterns expander

Examples:

eg1:
```bash
$ep 192.168.5.[1-3,8]
192.168.5.1
192.168.5.2
192.168.5.3
192.168.5.8
```

eg2:
```bash
$ ep foo-[a,b,c]-[1,2,3]
foo-a-1
foo-b-1
foo-c-1
foo-a-2
foo-b-2
foo-c-2
foo-a-3
foo-b-3
foo-c-3
```

eg3:
```bash
$ ep 192.168.5.[41,42,43] | while read -r ip; do
  ping -c 1 -w 1 $ip
done
```

eg4:
```bash
$ ep [80,443] | while read -r port; do
  nc -nv 192.168.5.41 $port
done
```
