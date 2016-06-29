```
Usage of ./brute-force-calculator:
  -d  Enable debug mode
  -r value
      rate aka. Hash/sec (available suffix: G M K) (default 1G)
```

# Example

```
$ echo "rijyomhaknapnunyett2" | brute-force-calculator
Years: 197.79
$ echo "rijyomhaknapnunyett2" | brute-force-calculator -d
Hashes per Seconds: 1G
'rijyomhaknapnunyett2' with 'abcdefghijklmnopqrstuvwxyz0123456789'
Years: 197.79
$ brute-force-calculator -d rijyomhaknapnunyett2
Hashes per Seconds: 1G
'rijyomhaknapnunyett2' with 'abcdefghijklmnopqrstuvwxyz0123456789'
Years: 197.79
$ brute-force-calculator -d -r 100M rijyomhaknapnunyett2
Hashes per Seconds: 100M
'rijyomhaknapnunyett2' with 'abcdefghijklmnopqrstuvwxyz0123456789'
Years: 1977.92
```