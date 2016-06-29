```
Usage of ./brute-force-calculator:
  -d  Enable debug mode
  -r value
      rate aka. Hash/sec (available suffix: G M K)
      or use 'auto' to calculate on the fly (default 1G)
```

# Example

```
$ echo "rijyomhaknapnunyett2" | brute-force-calculator
Years: 197.79
```

```
$ echo "rijyomhaknapnunyett2" | brute-force-calculator -d
Hashes per Seconds: 1G
'rijyomhaknapnunyett2' with 'abcdefghijklmnopqrstuvwxyz0123456789'
Years: 197.79
```

```
$ brute-force-calculator -d rijyomhaknapnunyett2
Hashes per Seconds: 1G
'rijyomhaknapnunyett2' with 'abcdefghijklmnopqrstuvwxyz0123456789'
Years: 197.79
```

```
$ brute-force-calculator -d -r 100M rijyomhaknapnunyett2
Hashes per Seconds: 100M
'rijyomhaknapnunyett2' with 'abcdefghijklmnopqrstuvwxyz0123456789'
Years: 1977.92
```

```
$ brute-force-calculator -d -r auto rijyomhaknapnunyett2
Calculate HasPerSec value based on your computation power with MD5...
...................................................................................................
100000000 hashes over 20 seconds => 5000000
Hashes per Seconds: 5M
'rijyomhaknapnunyett2' with 'abcdefghijklmnopqrstuvwxyz0123456789'
Years: 39558.31
```