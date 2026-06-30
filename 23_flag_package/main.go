package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Cek subcommand pertama
	if len(os.Args) > 1 && os.Args[1] == "myapp" {
		runMyApp()
		return
	}

	args := os.Args[1:]

	if len(args) == 0 {
		printBasicExample()
		return
	}

	// Parse flags dari args
	parsedFlags := parseArgs(args)

	fmt.Println("BASIC FLAG - String, Int, Bool")
	fmt.Println("---")

	name := getFlag(parsedFlags, "name", "Guest")
	age := getFlagInt(parsedFlags, "age", 25)
	active := getFlagBool(parsedFlags, "active", true)

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Active: %t\n", active)

	fmt.Println()
	fmt.Println("DEFAULT VALUES")
	fmt.Println("---")

	host := getFlag(parsedFlags, "host", "localhost")
	port := getFlagInt(parsedFlags, "port", 8080)
	verbose := getFlagBool(parsedFlags, "verbose", false)

	fmt.Printf("Server: %s:%d\n", host, port)
	fmt.Printf("Verbose: %t\n", verbose)

	fmt.Println()
	fmt.Println("MULTIPLE TAGS")
	fmt.Println("---")

	tags := getFlagSlice(parsedFlags, "tags")
	fmt.Printf("Tags: %v\n", tags)

	fmt.Println()
	fmt.Println("ENV FALLBACK")
	fmt.Println("---")

	dbHost := getEnvOrFlag("DB_HOST", "localhost")
	dbPort := getEnvOrFlag("DB_PORT", "5432")
	dbUser := getEnvOrFlag("DB_USER", "postgres")

	fmt.Printf("DB_HOST: %s\n", dbHost)
	fmt.Printf("DB_PORT: %s\n", dbPort)
	fmt.Printf("DB_USER: %s\n", dbUser)

	fmt.Println()
	fmt.Println("FLAG TYPES SUMMARY")
	fmt.Println("---")
	fmt.Println("- String   : --name=value")
	fmt.Println("- Int      : --port=8080")
	fmt.Println("- Bool     : --verbose atau --verbose=true")
	fmt.Println("- Slice    : --tags=go --tags=web")

	fmt.Println()
	fmt.Println("COMMANDS:")
	fmt.Println("  go run ./23_flag_package/ --name=Budi --age=30 --active=false")
	fmt.Println("  go run ./23_flag_package/ myapp --input=file.txt --format=csv")
}

func printBasicExample() {
	fmt.Println("BASIC FLAG EXAMPLE")
	fmt.Println("---")
	fmt.Println("Jalankan dengan flag:")
	fmt.Println("  go run ./23_flag_package/ --name=Budi --age=30 --active=false")
	fmt.Println()
	fmt.Println("Atau gunakan subcommand myapp:")
	fmt.Println("  go run ./23_flag_package/ myapp --input=file.txt --format=csv")
}

func runMyApp() {
	appFlags := flag.NewFlagSet("myapp", flag.ExitOnError)

	input := appFlags.String("input", "", "File input (wajib)")
	output := appFlags.String("output", "output.txt", "File output")
	format := appFlags.String("format", "json", "Format output (json/csv)")

	appFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: myapp [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		appFlags.PrintDefaults()
	}

	appFlags.Parse(os.Args[2:])

	if *input == "" {
		fmt.Println("Error: --input wajib diisi")
		appFlags.Usage()
		return
	}

	fmt.Println("MYAPP - Custom FlagSet")
	fmt.Println("---")
	fmt.Printf("Input: %s\n", *input)
	fmt.Printf("Output: %s\n", *output)
	fmt.Printf("Format: %s\n", *format)
}

func getFlag(flags map[string]string, key, fallback string) string {
	if val, ok := flags[key]; ok {
		return val
	}
	return fallback
}

func getFlagInt(flags map[string]string, key string, fallback int) int {
	if val, ok := flags[key]; ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}

func getFlagBool(flags map[string]string, key string, fallback bool) bool {
	if val, ok := flags[key]; ok {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return fallback
}

func getFlagSlice(flags map[string]string, key string) []string {
	if val, ok := flags[key]; ok {
		return strings.Split(val, ",")
	}
	return nil
}

func parseArgs(args []string) map[string]string {
	result := make(map[string]string)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--") {
			key := strings.TrimPrefix(arg, "--")

			if idx := strings.Index(key, "="); idx != -1 {
				result[key[:idx]] = key[idx+1:]
			} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				result[key] = args[i+1]
				i++
			} else {
				result[key] = "true"
			}
		} else if strings.HasPrefix(arg, "-") {
			key := strings.TrimPrefix(arg, "-")

			if idx := strings.Index(key, "="); idx != -1 {
				result[key[:idx]] = key[idx+1:]
			} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				result[key] = args[i+1]
				i++
			} else {
				result[key] = "true"
			}
		}
	}

	return result
}

func getEnvOrFlag(envKey, fallback string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return fallback
}

func init() {
	_ = time.Second
	_ = strconv.Itoa(0)
}
