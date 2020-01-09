package main

import (
    "fmt"
    "net"
)

// Обработка подключения.
func handleConnection(conn net.Conn, key string) {
    defer conn.Close()
    serverProtector := NewProtector(key)

    for {
        // считываем полученные в запросе данные
        input := make([]byte, (1024 * 4))
        n, err := conn.Read(input)
        if n == 0 || err != nil {
            fmt.Println("Read error:", err)
            break
        }
        source := string(input[0:n])

        // Ответ сервера.
        target := serverProtector.nextSessionKey(key)

        // выводим на консоль сервера диагностическую информацию
        fmt.Println("Начальный ключ: " + source)
        fmt.Println("Ключ, сгенерированный сервером: " + serverProtector.nextSessionKey(key))

        //проверка равенства ключей.
        //Если не равны - соединения обрывается.
        if source != serverProtector.nextSessionKey(key) {
          fmt.Println("Ключи не совпадают, соединение прервано")
          conn.Write([]byte("Ключи не совпадают, соединение прервано"))
          break
        }

        fmt.Println("Ключи совпадают, соединение продолжается")
        fmt.Println()

        // отправка данных клиенту
        conn.Write([]byte(target))
        key = serverProtector.nextSessionKey(key)
    }
}

//Работа Сервера.
func serverWork(port string, key string)  {
  listener, err := net.Listen("tcp", "127.0.0.1:" + port)
  if err != nil {
      fmt.Println(err)
      return
  }

  defer listener.Close()
  fmt.Println("Сервер работает \n Прослушивается порт: " + port)
  fmt.Println("Начальный ключ: " + key)
  for {
      conn, err := listener.Accept()
      if err != nil {
          fmt.Println(err)
          conn.Close()
          continue
      }
      //обработка подключения
      go handleConnection(conn, key)
  }
}

//Работа Клиента.
func clientWork(port string, key string)  {
  conn, err := net.Dial("tcp", "127.0.0.1:" + port)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer conn.Close()

  clientProtector := NewProtector(key)

  for{
      var source string
      fmt.Print("Отправить сообщение? [y/n] ")
      _, err := fmt.Scanln(&source)
      if err != nil {
          fmt.Println("Некорректный ввод", err)
          continue
      }
      if source != "y" {
          continue
      }
      fmt.Println("Хэш-ключ: ", clientProtector.nextSessionKey(key))
      source = clientProtector.nextSessionKey(key)
      // отправка сообщения серверу
      if n, err := conn.Write([]byte(source));
      n == 0 || err != nil {
          fmt.Println(err)
          return
      }
      // получение ответа
      buff := make([]byte, 1024)
      n, err := conn.Read(buff)
      if err !=nil{ break}

      fmt.Println("Ответ от сервера: " + string(buff[0:n]))
      //проверка равенства ключей.
      //Если не равны - соединения обрывается.
      if clientProtector.nextSessionKey(key) != string(buff[0:n]) {
        fmt.Println("Ключи не совпадают, соединение прервано")
        break
      }
      fmt.Println("Ключи совпадают, соединение продолжается")
      fmt.Println()

      key = clientProtector.nextSessionKey(key)
  }
}
