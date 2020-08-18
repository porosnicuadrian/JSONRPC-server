package main

type ArgsSum struct {
	X, Y int
}

type ArgsWrite struct {
	X        int
	FilePath string
}

type ArgsRead struct {
	FilePath string
}
