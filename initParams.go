package main

import "flag"

type runParams struct {
	action   string
	reverse  bool
	fileName string
}

func initParams() runParams {
	var res runParams
	var action string
	var reverse bool
	var fileName string

	flag.StringVar(&action, "action", "count", "which action to perform")
	flag.StringVar(&action, "a", "count", "(shorthand for action)")
	flag.BoolVar(&reverse, "v", false, "reverse effect")
	flag.StringVar(&fileName, "f", "", "path to input file instead of stdin")

	flag.Parse()

	res.action = action
	res.reverse = reverse
	res.fileName = fileName

	return res
}
