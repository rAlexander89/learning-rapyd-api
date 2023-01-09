package rapyd

// "fmt"
// "log"
// "net/http"
// "net/url"
// "os"

// "github.com/joho/godotenv"

const (
	rapydSandbox = "https://sandboxapi.rapyd.net"
)

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("error loading env file")
// 	}

// 	sk := os.Getenv("SECRET_KEY")
// 	ak := os.Getenv("ACCESS_KEY")

// 	cli := http.DefaultClient
// 	url, _ := url.Parse(rapydSandbox)
// 	signer := NewSigner([]byte(ak), []byte(sk))
// 	client := NewClient(signer, cli, url)
// 	response, _ := client.GetCountries()

// 	fmt.Println("final response: ", response)
// }
