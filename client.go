package main

import "flag"

//Клиент.
func main() {
  port := flag.String("port", "", "Порт в режиме 'Клиент'")
  key := flag.String("key", "", "Начальный ключ")
  flag.Parse()
  clientWork(*port, *key)
}
