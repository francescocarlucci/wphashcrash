# WPHashCrash

**Under development**

Cracking WP passwords hashes using golang.

This tool is related to my research about WordPress password hashes structure. By default, it accepts a single hash as input, but you can easly wrap it in a bash for loop to process multiple hashes.

At the moment, it processes 100K attempts in about 2 seconds.

No dictionary is included with the rempo, you can find some good ones here: https://github.com/danielmiessler/SecLists/tree/master/Passwords/Common-Credentials

Usage:

```
go get github.com/francescocarlucci/wphashcrash
wphashcrash 'hash' path/to/dictionary.txt
```

### ToDo

- multi-thread
- verbosity
- improve inputs (multiple hashes, email:hash)

Contributions are welcome!

By [frenxi](https://frenxi.com)
