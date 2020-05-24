package environmentloader

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load .env files. 
// Existing .env files take precendence of .env files that are loaded later.
func Load() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "dev"
	}

	absolutePath := os.ExpandEnv("./")

	// If there's a .env.$env.local file, this one is loaded. Otherwise, it falls back to .env.$env
	godotenv.Load(absolutePath + ".env." + env + ".local")

	// If there's a .env.local file representing general local environment variables it's loaded now
	godotenv.Load(absolutePath + ".env." + env)

	// .env.local file is always ignored in test environment because tests should produce the same results for everyone
	if "test" != env {
		godotenv.Load(absolutePath + ".env.local")
	}

	// .env file
	godotenv.Load(absolutePath + ".env")

	checkMandatoryVariables()
}

// checkMandatoryVariables does all the necessary validations to make sure
// all mandatory variables are set up.
func checkMandatoryVariables() {
	mandatoryVariables := []string{
		"DATABASE_HOST",
		"DATABASE_PORT",
	}

	for _, mandatoryVariable := range mandatoryVariables {
		if _, exists := os.LookupEnv(mandatoryVariable); exists == false {
			log.Fatalf("\033[97;41m%s\033[0m\n", "### Environment variable not found: "+mandatoryVariable+" ###")
		}
	}
}
