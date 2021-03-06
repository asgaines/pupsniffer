package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/asgaines/pupsniffer/pupservice"
	"github.com/asgaines/pupsniffer/pupservice/fetcher"
	"github.com/asgaines/pupsniffer/utils"
)

func init() {
	flag.Usage = func() {
		fmt.Println(`
         ____                         _ ________
        / __ \__  ______  _________  (_) __/ __/__  ____
       / /_/ / / / / __ \/ ___/ __ \/ / /_/ /_/ _ \/ ___/
      / ____/ /_/ / /_/ (__  ) / / / / __/ __/  __/ /
     /_/    \__,_/ .___/____/_/ /_/_/_/ /_/ /\___/_/
                /_/

                                    __               _ ______       __
                          ___  ___ / /_    ___ ___  (_/ _/ _/__ ___/ /
                _ _ _    / _ \/ -_/ __/   (_-</ _ \/ / _/ _/ -_/ _  /
               (_|_|_)   \_, /\__/\__/   /___/_//_/_/_//_/ \__/\_,_/
                        /___/                                       
		`)
		fmt.Println("Welcome to pupsniffer, a way to know about new pups at the Boulder Humane Society!")
		fmt.Println()
		fmt.Fprintln(flag.CommandLine.Output(), "Usage of pupsniffer:")
		flag.PrintDefaults()
	}
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var kennelPath string
	var filterPups bool
	var newOnly bool
	var email bool

	flag.StringVar(&kennelPath, "kennel", fmt.Sprintf("%s/.config/pupsniffer/kennel", usr.HomeDir), "Path to kennel (where previous searches are stored for comparison)")
	flag.BoolVar(&filterPups, "filter", true, "Should pups be filtered?")
	flag.BoolVar(&newOnly, "newonly", true, "Only show pups not before seen?")
	flag.BoolVar(&email, "email", false, "Send report to email")

	flag.Parse()

	log.Printf("filtering: %v, new only: %v, email: %v", filterPups, newOnly, email)

	if err := os.MkdirAll(kennelPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	pupsvc := pupservice.New(
		fetcher.NewFetcher(),
		kennelPath,
		"./static",
	)

	pupIDs, err := pupsvc.FetchPupIDs(ctx)
	if err != nil {
		log.Fatal(err)
	}
	numTotalPups := len(pupIDs)

	if numTotalPups == 0 {
		log.Println("There aren't any pups at the Boulder Humane Society right now...")
		os.Exit(0)
	}

	log.Printf("%d %s at the Boulder Humane Society right now\n", len(pupIDs), utils.Pluralize("pup", "pups", len(pupIDs)))

	if newOnly {
		newPupIDs, err := pupsvc.SniffOutNew(pupIDs)
		if err != nil {
			log.Fatal(err)
		}

		if len(newPupIDs) == 0 {
			log.Println("Looks like you've already seen all the pups currently at the center. Check back soon!")
			os.Exit(0)
		}

		if err := pupsvc.KennelPups(pupIDs); err != nil {
			log.Fatal(err)
		}

		pupIDs = newPupIDs
	} else {
		if err := pupsvc.KennelPups(pupIDs); err != nil {
			log.Fatal(err)
		}
	}

	pups, err := pupsvc.FetchPups(ctx, pupIDs)
	if err != nil {
		log.Fatal(err)
	}

	if filterPups {
		filteredPups, err := pupsvc.FilterPups(pups)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%d/%d new %s %s your criteria!\n", len(filteredPups), len(pups), utils.Pluralize("pup", "pups", len(filteredPups)), utils.Pluralize("meets", "meet", len(filteredPups)))

		pups = filteredPups
	}

	buf := &bytes.Buffer{}

	if err := pupsvc.PupReport(numTotalPups, pups, buf); err != nil {
		log.Fatal(err)
	}

	if len(pups) > 0 && email {
		recipients := []string{
			"andrew.s.gaines@gmail.com",
			"kelley36@gmail.com",
			"susannahgaines@gmail.com",
			"ajneumann@yahoo.com",
		}
		if err := pupsvc.Mailman(buf, recipients); err != nil {
			log.Fatal(err)
		}
	}
}
