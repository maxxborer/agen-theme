// go example:

struct config {
		port int
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "port to listen on")
	flag.Parse()

	fmt.Printf("listening on port %d\n", cfg.port)
}
