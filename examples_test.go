package xsdvalidate_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/terminalstatic/go-xsd-validate"
)

// An example on how to use the package.
// Init() is only required once before parsing and validating, and Cleanup() respectively when finished.
func Example() {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("examples/test1_split.xsd", xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xsdhandler.Free()

	xmlFile, err := os.Open("examples/test1_fail2.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	inXml, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	// Option 1:
	xmlhandler, err := xsdvalidate.NewXmlHandlerMem(inXml, xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xmlhandler.Free()

	err = xsdhandler.Validate(xmlhandler, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			fmt.Println(err)
			fmt.Printf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Message)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Path)
		default:
			fmt.Println(err)
		}
	}

	// Option 2:
	err = xsdhandler.ValidateMem(inXml, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			fmt.Println(err)
			fmt.Printf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Message)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Path)
		default:
			fmt.Println(err)
		}
	}
	// Output:
	// 3: Element 'shipto': This element is not expected. Expected is ( orderperson ).
	// Error in line: 3
	// Element 'shipto': This element is not expected. Expected is ( orderperson ).
	// shipto<shiporder
	// 3: Element 'shipto': This element is not expected. Expected is ( orderperson ).
	// Error in line: 3
	// Element 'shipto': This element is not expected. Expected is ( orderperson ).
	// shipto<shiporder
}
