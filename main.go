package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/asgaines/pupsniffr/pupservice"
	"github.com/asgaines/pupsniffr/pupservice/fetcher"
	"github.com/asgaines/pupsniffr/utils"
)

func init() {
	flag.Usage = func() {
		fmt.Println(`
         ____                         _ ________
        / __ \__  ______  _________  (_) __/ __/____
       / /_/ / / / / __ \/ ___/ __ \/ / /_/ /_/ ___/
      / ____/ /_/ / /_/ (__  ) / / / / __/ __/ /
     /_/    \__,_/ .___/____/_/ /_/_/_/ /_/ /_/
                /_/                                      

                                    __               _ ______   __
                          ___  ___ / /_    ___ ___  (_/ _/ _/__/ /
                _ _ _    / _ \/ -_/ __/   (_-</ _ \/ / _/ _/ _  /
               (_|_|_)   \_, /\__/\__/   /___/_//_/_/_//_/ \_,_/
                        /___/                                       
		`)
		fmt.Println("Welcome to pupsniffr, a way to know about new pups at the Boulder Humane Society!")
		fmt.Println()
		fmt.Fprintln(flag.CommandLine.Output(), "Usage of pupsniffr:")
		flag.PrintDefaults()
	}
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var kennelPath string
	flag.StringVar(&kennelPath, "kennel", fmt.Sprintf("%s/.config/pupsniffr/kennel", usr.HomeDir), "Path to kennel (where previous searches are stored for comparison)")
	flag.Parse()

	if err := os.MkdirAll(kennelPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	pupsvc := pupservice.New(
		fetcher.NewFetcher(),
		kennelPath,
	)

	pupIDs, err := pupsvc.FetchPupIDs(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if len(pupIDs) == 0 {
		fmt.Printf("There aren't any pups at the Boulder Humane Society right now...")
		os.Exit(0)
	}

	fmt.Printf("%d %s at the Boulder Humane Society right now\n", len(pupIDs), utils.Pluralize("pup", "pups", len(pupIDs)))

	newPupIDs, err := pupsvc.SniffOutNew(pupIDs)
	if err != nil {
		log.Fatal(err)
	}

	if len(newPupIDs) == 0 {
		fmt.Println("Looks like you've already seen all the pups currently at the center. Check back soon!")
		os.Exit(0)
	}

	if err := pupsvc.KennelPups(pupIDs); err != nil {
		log.Fatal(err)
	}

	newPups, err := pupsvc.FetchPups(ctx, newPupIDs)
	if err != nil {
		log.Fatal(err)
	}

	filteredPups, err := pupsvc.FilterPups(newPups)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d/%d new %s %s your criteria!\n", len(filteredPups), len(newPups), utils.Pluralize("pup", "pups", len(filteredPups)), utils.Pluralize("meets", "meet", len(filteredPups)))

	for _, pup := range filteredPups {
		pup.BarkGreeting()
	}
}
