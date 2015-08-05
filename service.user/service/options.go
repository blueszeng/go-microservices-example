package service

import "flag"

type options struct {
	Port string
}

func getOptions() *options {
	// Default options
	options := &options{}

	flag.StringVar(&options.Port, "port", "8081", "service port")

	flag.Parse()

	return options
}
