package main

import (
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// write me a function that receive array of int and return a array of int that only prime number
func primeNumber(arr []int) []int {
	var prime []int
	for _, v := range arr {
		if isPrime(v) {
			prime = append(prime, v)
		}
	}
	return prime
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
s
func main() {
	println(X(450000, 3, 4))
}

// func addStringNumber(a, b string) int {
// 	x1, _ := strconv.Atoi(a)
// 	x2, _ := strconv.Atoi(b)

// 	return x1 + x2
// }

func X(fa int, in int, year int) int {
	installment := year * 12
	return ((fa * (in / 100) * year) + fa) / installment
}

// func main() {
// 	slog.SetLogLoggerLevel(slog.LevelDebug)

// 	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
// 	defer stop()

// 	db, err := sql.Open("sqlite3", "./languages.sqlite")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// mux := http.NewServeMux()
// 	mux := gin.Default()

// 	helloHandler := func(w http.ResponseWriter, req *http.Request) {
// 		_, err := io.WriteString(w, "Hello, world!\n")
// 		if err != nil {
// 			slog.Error(err.Error())
// 		}
// 	}
// 	// mux.HandleFunc("/hello", helloHandler)
// 	mux.("/hello", helloHandler)

// 	programmingHandler := programming.ProgrammingHandler{
// 		Db: db,
// 	}
// 	// programmingHandler := programming.NewProgrammingHandler(db)

// 	mux.HandleFunc("/languages", programmingHandler.Languages)

// 	idleConnsClosed := make(chan struct{})
// 	fmt.Println("idle con", idleConnsClosed)
// 	port := os.Getenv("PORT")
// 	srv := http.Server{
// 		Addr:              ":" + port,
// 		Handler:           mux,
// 		ReadHeaderTimeout: 5 * time.Second,
// 	}

// 	go func() {
// 		<-ctx.Done()
// 		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer cancel()

// 		fmt.Println("shutting down...")
// 		if err := srv.Shutdown(ctx); err != nil {
// 			log.Printf("HTTP server Shutdown: %v", err)
// 		}
// 		// close(idleConnsClosed)
// 		idleConnsClosed <- struct{}{}
// 	}()

// 	slog.Debug("listening and serving on HTTP port :" + port)
// 	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
// 		log.Panicf("HTTP server ListenAndServe: %v", err)
// 	}

// 	x := <-idleConnsClosed
// 	fmt.Println(x)
// 	fmt.Println("bye")

// }

// func AddString(nums string) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	if len(nums) == 1 {
// 		r, _ := strconv.Atoi(nums)
// 		return r
// 	}

// 	list := strings.Split(nums, ",")

// 	result := 0
// 	for _, v := range list {
// 		num, _ := strconv.Atoi(v)
// 		result += num
// 	}
// 	return result
// }

// func couple(s string) []string {
// 	result := []string{}
// 	for i, j := 0, 1; i < len(s); i = i + 2 {
// 		cha := string(s[i])
// 		result = append(result, char)
// 	}
// }

// func reverse(arr [4]int) [4]int {
// 	newList := [4]int{}

// 	for i, j := len(arr)-1, 0; i >= 0; i-- {
// 		newList[j] = arr[i]
// 		j++
// 	}

// 	return newList
// }

// func add(a, b int) int {
// 	return a + b
// }
