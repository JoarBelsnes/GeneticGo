package main

import (
	"flag"
	"fmt"
	"github.com/handcraftsman/File"
	"strconv"
	"time"
)

const genericGeneSet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		println("Usage: go run samples/tsp.go ROUTEFILEPATH")
		return
	}
	var routeFileName = flag.Arg(0)
	if !File.Exists(routeFileName) {
		println("file " + routeFileName + " does not exist.")
		return
	}
	println("using route file: " + routeFileName)

	idToPointLookup := readPoints(routeFileName)
	println("read " + strconv.Itoa(len(idToPointLookup)) + " points...")

	calc := func(genes string) int {
		points := genesToPoints(genes, idToPointLookup)
		return getFitness(genes, points)
	}

	if File.Exists(routeFileName + ".opt.tour") {
		println("found optimal solution file: " + routeFileName + ".opt")
		optimalRoute := readOptimalRoute(routeFileName+".opt.tour", len(idToPointLookup))
		println("read " + strconv.Itoa(len(optimalRoute)) + " segments in the optimal route")
		points := getPointsInOptimalOrder(idToPointLookup, optimalRoute)
		genes := genericGeneSet[0:len(idToPointLookup)]
		idToPointLookup = make(map[string]Point, len(idToPointLookup))
		for i, v := range points {
			idToPointLookup[genericGeneSet[i:i+1]] = v
		}
		print("optimal route: " + genes)
		print("\t")
		println(getFitness(genes, points))
	}

	genes := genericGeneSet[0:len(idToPointLookup)]

	start := time.Now()

	disp := func(genes string) {
		points := genesToPoints(genes, idToPointLookup)
		print(genes)
		print("\t")
		print(getFitness(genes, points))
		print("\t")
		fmt.Println(time.Since(start))
	}

	var solver = new(genetic.Solver)
	solver.MaxSecondsToRunWithoutImprovement = 20
	solver.LowerFitnessesAreBetter = true

	var best = solver.GetBest(calc, disp, genes, len(idToPointLookup), 1)
	disp(best)
	print("Total time: ")
	fmt.Println(time.Since(start))
}
