package main

import (
	"net/http"

	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/coolguru/euler-prime"
	"github.com/labstack/echo"
)

const invalidInputError string = "Invalid input"

/*
AboutHandler for / endpoint
*/
func AboutHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "find the Xth Y digit prime number in the expansion of Euler's number. Add more info there.")
}

func calcEulerPrime(x, y string) (int64, error) {
	xNum, errX := strconv.Atoi(x)

	if errX != nil {
		return 0, errors.New(invalidInputError)
	}

	yNum, errY := strconv.Atoi(y)

	if errY != nil {
		return 0, errors.New(invalidInputError)
	}

	return eulerprime.EulerPrime(xNum, yNum)
}

func processCSVFile(filename string) string {

	// open a file
	if file, err := os.Open(filename); err == nil {

		// make sure it gets closed
		defer file.Close()

		result := []string{}

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			inputs := strings.Split(scanner.Text(), ",")
			if len(inputs) == 2 {
				ePrime, errE := calcEulerPrime(inputs[0], inputs[1])
				if errE != nil {
					result = append(result, errE.Error())
				} else {
					result = append(result, strconv.FormatInt(ePrime, 10))
				}
			}
		}

		// check for errors
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return strings.Join(result, "<br/>")
	}

	return "Error processing file"
}

// CalculateEulerPrimeHandler to calculate Euler prime for input x and y
func CalculateEulerPrimeHandler(c echo.Context) error {

	x := c.QueryParam("x")
	y := c.QueryParam("y")

	ePrime, err := calcEulerPrime(x, y)

	if err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, (strconv.FormatInt(ePrime, 10)))
}

//UploadAndCalculateEulerPrimeHandler for uploading csv file, 1) Uploads a csv file 2) Processes a csv file. There can optimization done to directly consume the input of csv file.
func UploadAndCalculateEulerPrimeHandler(c echo.Context) error {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, processCSVFile(file.Filename))
}
