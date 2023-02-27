This Go program generates a summary of inventory report required from a CSV export file of inventory data from Snipe-IT.

## Requirements
Go 1.13 or later

## Usage
1. Enter the path to the CSV file when prompted by running the program.
2. The program will then parse the CSV data, count the number of instances of each status for each specified model, and print out the results in the desired order.

## Output Sections
The program generates output in the following sections:

* Physical - Monitors
* Physical - MacOS - Platform Type 1
* Physical - MacOS - Platform Type 2
* Physical - WinOS Platform Type 1
* Physical - WinOS Platform Type 2

## Example
![SCR-20230227-3gt](https://user-images.githubusercontent.com/125210256/221501072-a09dca5b-ab20-4ed3-a1f2-0d94c63dda74.png)
