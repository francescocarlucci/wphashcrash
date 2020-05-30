# WPHashCrash

Cracking WP passwords hashes using golang.

Blog post: [https://frenxi.com/cracking-wordpress-password-hash/](https://frenxi.com/cracking-wordpress-password-hash/)

This tool is related to my research about WordPress password hashes structure. By default, it accepts a single hash as input, but you can easly wrap it in a bash for loop to process multiple hashes.

At the moment, it processes 100K attempts in about 2 minutes on a small VPS.

No dictionary is included with the repo, you can find some good ones here: https://github.com/danielmiessler/SecLists/tree/master/Passwords/Common-Credentials

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
