/**
 * AUTHOR: Samuel Beguiristain
 * Description: This code takes a CSV export file from SNIPE-IT and returns a summary of the inventory report required
 * Modify string to fit your needs
 * Version: 5
 **/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Prompt the user for the CSV file location
	fmt.Print("Enter the path to the CSV file: ")
	var filepath string
	fmt.Scanln(&filepath)

	// Parse the CSV data
	rows := parseCSVData(filepath)

	// Define the models to look for and the corresponding output sections
	models := map[string]string{
		"MacBook Pro (16-inch, 2019)":                           "Physical - MacOS - Platform Type 1",
		"Macbook Pro (16-inch, 2021)":                           "Physical - MacOS - Platform Type 1",
		"MacBook Pro (13-inch, 2020, Four Thunderbolt 3 ports)": "Physical - MacOS - Platform Type 2",
		"MacBook Pro (13-inch, 2020, Two Thunderbolt 3 ports)":  "Physical - MacOS - Platform Type 2",
		"MacBook Pro (13-inch, M1, 2020)":                       "Physical - MacOS - Platform Type 2",
		"MacBook Pro (14-inch, 2021)":                           "Physical - MacOS - Platform Type 2",
		"Latitude 5520":                                         "Physical - WinOS Platform Type 1",
		"Latitude 5530":                                         "Physical - WinOS Platform Type 1",
		"Latitude 3520":                                         "Physical - WinOS Platform Type 1",
	}

	// Count the number of instances of each status for each specified model
	output := countModelInstances(models, rows)

	// Define the desired order of the sections
	sectionOrder := []string{
		"Physical - Monitors",
		"Physical - MacOS - Platform Type 1",
		"Physical - MacOS - Platform Type 2",
		"Physical - WinOS Platform Type 1",
		"Physical - WinOS Platform Type 2",
	}

	// Print out the results in the desired order
	printResults(output, sectionOrder)
}

func parseCSVData(filepath string) [][]string {
	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the CSV data
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return rows
}

func countModelInstances(models map[string]string, rows [][]string) map[string]map[string]int {
	// Loop through the rows and count the number of instances of each status for each specified model
	output := make(map[string]map[string]int)
	for _, row := range rows {
		category := row[3] // Use the "Category" column
		if category == "Monitor" {
			if _, ok := output["Physical - Monitors"]; !ok {
				output["Physical - Monitors"] = make(map[string]int)
			}
			status := row[5]
			if status == "Checked-out (deployed)" || status == "Checked-out (deployed) (deployed)" {
				output["Physical - Monitors"]["Issued/Active"]++
			} else if status == "Available (deployable)" {
				output["Physical - Monitors"]["In-stock"]++
			} else if status == "Pending Return (deployed)" || status == "Pending Return (deployed) (deployed)" || status == "Pending Return (pending)" {
				output["Physical - Monitors"]["Pending Return"]++
			}
		} else if _, ok := models[row[1]]; ok {
			model := models[row[1]]
			if _, ok := output[model]; !ok {
				output[model] = make(map[string]int)
			}
			status := row[5]
			if status == "Checked-out (deployed)" || status == "Checked-out (deployed) (deployed)" {
				output[model]["Issued/Active"]++
			} else if status == "Available (deployable)" {
				output[model]["In-stock"]++
			} else if status == "Pending Return (deployed)" || status == "Pending Return (deployed) (deployed)" || status == "Pending Return (pending)" {
				output[model]["Pending Return"]++
			}
		} else if category == "PC" {
			model := "Physical - WinOS Platform Type 2"
			if _, ok := output[model]; !ok {
				output[model] = make(map[string]int)
			}
			status := row[5]
			if status == "Checked-out (deployed)" || status == "Checked-out (deployed) (deployed)" {
				output[model]["Issued/Active"]++
			} else if status == "Available (deployable)" {
				output[model]["In-stock"]++
			} else if status == "Pending Return (deployed)" || status == "Pending Return (deployed) (deployed)" || status == "Pending Return (pending)" {
				output[model]["Pending Return"]++
			}
		}
	}
	return output
}

func printResults(output map[string]map[string]int, sectionOrder []string) {
	// Print out the results in the desired order
	printedSections := make(map[string]bool) // Track the sections already printed
	for _, section := range sectionOrder {
		if _, ok := printedSections[section]; !ok {
			fmt.Printf("%s:\n", section)
			fmt.Printf("    Issued/Active: %d\n", output[section]["Issued/Active"])
			fmt.Printf("    In-stock: %d\n", output[section]["In-stock"])
			fmt.Printf("    Pending Return: %d\n", output[section]["Pending Return"])
			printedSections[section] = true // Mark the section as printed
		}
	}
}
