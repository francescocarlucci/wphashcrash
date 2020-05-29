package main

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "bytes"
  "strings"
  "crypto/md5"
)

func main() {

  hash := os.Args[1]

  rsalt := []rune(hash)
  salt := string(rsalt[4:12])

  itoa64 := "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

  hashloop := 1 << strings.Index(itoa64, string(hash[3]));

  dict, err := os.Open(os.Args[2])

  if err != nil {
    log.Fatal(err)
  }

  defer dict.Close()

  scanner := bufio.NewScanner(dict)

  for scanner.Scan() {

    pass := scanner.Text()

    //fmt.Println("Testing: " + pass + "...")

    hashedpass := md5.Sum([]byte(salt+pass))
    hashcount := 0

    for hashcount < hashloop {
      hashedpass = md5.Sum([]byte(string(hashedpass[:])+pass))
      hashcount = hashcount + 1
    }

    var c bytes.Buffer

    c.WriteString(string(itoa64[hashedpass[0] & 0x3f]))
    c.WriteString(string(itoa64[hashedpass[3] & 0x3f]))
    c.WriteString(string(itoa64[hashedpass[6] & 0x3f]))
    c.WriteString(string(itoa64[hashedpass[9] & 0x3f]))

    checkstring := c.String()

    if checkstring[0] == hash[12] && checkstring[1] == hash[16] && checkstring[2] == hash[20] && checkstring[3] == hash[24] {
      fmt.Println(pass + ":" + hash)
    }

  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

}
