package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	// go CloseHandler(conn)
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

// func CloseHandler(conn net.Conn) {
// 	c := make(chan os.Signal)
// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
// 	go func() {
// 		<-c
// 		fmt.Println("quit")
// 		// io.Copy(conn, os.Stdout) // NOTE: ignoring errors
// 		mustCopy(conn, os.Stdout)
// 		fmt.Println("TEST")
// 		os.Exit(1)
// 	}()
// }

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// io.Copy(dst, src)
