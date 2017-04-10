package main

import (
	"fmt"
	"github.com/dgruber/go-cfclient"
	"os"
)

func listIsolationSegments(c *cfclient.Client) {
	is, err := c.ListIsolationSegments()
	if err != nil {
		panic(err)
	}
	for _, s := range is {
		fmt.Printf("%s\t%s\n", s.GUID, s.Name)
	}
}

func createIsolationSegment(c *cfclient.Client) {
	if len(os.Args) != 3 {
		printHelp()
	}
	is, errCreate := c.CreateIsolationSegment(os.Args[2])
	if errCreate != nil {
		panic(errCreate)
	}
	fmt.Printf("Creates Isolation Segment with the Name %s (GUID: %s)\n", is.Name, is.GUID)
}

func deleteIsolationSegment(c *cfclient.Client) {
	if len(os.Args) != 3 {
		printHelp()
	}
	c.DeleteIsolationSegmentByGUID(os.Args[2])
}

func addToOrg(c *cfclient.Client) {
	if len(os.Args) != 4 {
		printHelp()
	}

	is, errIs := c.GetIsolationSegmentByGUID(os.Args[2])
	if errIs != nil {
		panic(errIs)
	}

	org, errOrg := c.GetOrgByName(os.Args[3])
	if errOrg != nil {
		panic(errOrg)
	}

	if errAddOrg := is.AddOrg(org.Guid); errAddOrg != nil {
		panic(errAddOrg)
	}
}

func printHelp() {
	fmt.Println("Usage: cf-isolation-segments [list | create <name> | delete <guid> | add-to-org <guid> <org_name>")
	os.Exit(1)
}

// sample for testing go-cfclient isolation segment functionality
func main() {

	if len(os.Args) <= 1 {
		printHelp()
	}

	if os.Getenv("CF_API") == "" || os.Getenv("CF_USER") == "" || os.Getenv("CF_PASSWORD") == "" {
		fmt.Println("CF_API, CF_USER, and CF_PASSWORD needs to be set.")
		os.Exit(1)
	}

	c := &cfclient.Config{
		ApiAddress:        os.Getenv("CF_API"),
		Username:          os.Getenv("CF_USER"),
		Password:          os.Getenv("CF_PASSWORD"),
		SkipSslValidation: true,
	}

	client, err := cfclient.NewClient(c)
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "list":
		listIsolationSegments(client)
	case "create":
		createIsolationSegment(client)
	case "delete":
		deleteIsolationSegment(client)
	case "add-to-org":
		addToOrg(client)
	default:
		printHelp()
	}
}
