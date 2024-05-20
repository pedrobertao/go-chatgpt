package auth

type Config struct {
	ApiKey    string
	SecretKey string
	health    func()
}
