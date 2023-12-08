package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"regexp"
	"sync"
	"math"
)	

type mapIslands [][3]uint64

type result struct {
	seed uint64
	location uint64
}

func constructMap(destmap mapIslands, line string) mapIslands {
	temp := strings.Split(line, " ")
	source,_ := strconv.Atoi(temp[1])
	destination := convertouInt(temp[0])
	span,_ := strconv.Atoi(temp[2])
	var tempmap [3]uint64
	tempmap[0] = destination
	tempmap[1] = uint64(source)
	tempmap[2] = uint64(span)
	destmap = append(destmap, tempmap)
 	return destmap
}

func findValue(destmap mapIslands, value uint64) uint64 {

	for _, m := range destmap {
		ss := m[1]
		se := ss + m[2]

		if ss <= value && value < se {
			ds := m[0]
			dist := value - ss
			return ds + dist
		}
	}
	return value
	
}

func convertouInt(input string) uint64 {
	value, _ := strconv.Atoi(input)
	
	return uint64(value)
}

func day5() {
	
	var seeds []uint64
	var INPUT string
	var soil  mapIslands
	var fertilizer mapIslands
	var water mapIslands
	var light mapIslands
	var temperature mapIslands
	var humidity mapIslands
	var location mapIslands
	var minLocation uint64
	INPUT = "input_5.txt"
	file, err := os.Open(INPUT)
	
	if err != nil {
		fmt.Println(err)
	}
	current_map := ""
	// source := ""
	destination_map := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match, _ := regexp.MatchString("^seeds: .*$", line)
		if match {
			temps := strings.Split(line," ")[1:]
			for _, Sseed := range temps {
				seed,_ := strconv.Atoi(Sseed)
				seeds = append(seeds, uint64(seed))
			}
		}
		match, _ = regexp.MatchString("^[a-z]+-to-[a-z]+ map:", line)
		if match {
			current_map = strings.Split(line," ")[0]
			// source = strings.Split(current_map, "-")[0]
			destination_map = strings.Split(current_map,"-")[2]
			
		}
		match, _ = regexp.MatchString("^[0-9 ]+$", line)
		if match {
			// here comes the hard part
			switch destination_map {
				case "soil":
					soil = constructMap(soil, line)
					break
				case "fertilizer":
					fertilizer = constructMap(fertilizer, line)
					break
				case "water":
					water = constructMap(water, line)
					break
				case "light":
					light = constructMap(light, line)
					
					break
				case "temperature":
					temperature = constructMap(temperature, line)
					break
				case "humidity":
					humidity = constructMap(humidity, line)
					break
				case "location":
					location = constructMap(location, line)
					break
			}
				
		}
	}
	// Now that I construct the maps, I exploit them Part1
	for i:=0; i < len(seeds); i++ {

		seed := seeds[i]
		soilS := findValue(soil, seed)
		fertilizerS := findValue(fertilizer, soilS)
		waterS := findValue(water, fertilizerS)
		lightS := findValue(light, waterS)
		temperatureS := findValue(temperature, lightS)
		humidityS := findValue(humidity, temperatureS)
		locationS := findValue(location, humidityS)
		
		if minLocation == 0 {
			minLocation = locationS 
		}
		if locationS < minLocation {
			minLocation = locationS
		}
	}
	fmt.Println("Part 1 minLocation: ", minLocation)
	// Ok now Part 2
	minLocation = uint64(math.MaxUint64)
	mutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Starting go subroutine %d, range %d\n", seeds[i], seeds[i+1])
			for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
				soilS := findValue(soil, seed)
				fertilizerS := findValue(fertilizer, soilS)
				waterS := findValue(water, fertilizerS)
				lightS := findValue(light, waterS)
				temperatureS := findValue(temperature, lightS)
				humidityS := findValue(humidity, temperatureS)
				locationS := findValue(location, humidityS)

				mutex.RLock()
				if locationS < minLocation {
					mutex.RUnlock()
					mutex.Lock()
					minLocation = locationS
					mutex.Unlock()
				} else {
					mutex.RUnlock()
				}
			}

			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Part 2 minLocation : %d\n", minLocation)
}

func main() {
	day5()	
}
