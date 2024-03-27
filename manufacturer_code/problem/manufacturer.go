package vin

func Manufacturer(vin string) string {
	manufacturer := vin[:3] // get string from 0 to 2 = 0,1,2
	// if the last digit of the manufacture ID is a 9
	// the digits 12 to 14 are the second part of the ID
	if manufacturer[2] == '9' {
		manufacturer += vin[11:14]
	}

	return manufacturer
}

/*
So this function works correctly when given the right input, but it has some problems:

There is no guarantee that the input string is a VIN.
For strings shorter than three characters, the function causes a panic.
The optional second part of the ID is a feature of European VINs only.
The function would return wrong IDs for US cars having a 9 as the third digit of the manufacturer code.

*/
