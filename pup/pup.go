package pup

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/asgaines/pupsniffer/null"
)

// WoofASCII holds the data for the pup ascii picture
const WoofASCII = `
             __
            /  \        WOOF!
           / ..|\     /
          (_\  |_)
          /  \@'
         /     \
    _   /  \   |
    \\/  \  | _\
     \   /_ || \\_
      \____)|_) \_)
`

// Pup is all the info about a good ole' pup
type Pup struct {
	CompanyID      string      `json:"CompanyID"`      // "993",
	ID             string      `json:"ID"`             // "41727077",
	AnimalName     string      `json:"AnimalName"`     // "Shadow",
	Species        string      `json:"Species"`        // "Dog",
	Sex            string      `json:"Sex"`            // "Male",
	Altered        string      `json:"Altered"`        // "Yes",
	PrimaryBreed   string      `json:"PrimaryBreed"`   // "German Shepherd",
	SecondaryBreed null.String `json:"SecondaryBreed"` // {
	// "0": "\n  "
	// }, OR "Mix",
	PrimaryColor   string      `json:"PrimaryColor"`   // "Fawn",
	SecondaryColor null.String `json:"SecondaryColor"` // {
	// "0": "\n  "
	// }, OR "White"
	Age            string      `json:"Age"`            // "36",
	Size           string      `json:"Size"`           // "L",
	Housetrained   string      `json:"Housetrained"`   // "Unknown",
	Declawed       string      `json:"Declawed"`       // "No",
	Price          string      `json:"Price"`          // "19.00",
	LastIntakeDate string      `json:"LastIntakeDate"` // "2019-05-20 14:30:00",
	Location       string      `json:"Location"`       // "DA",
	Description    null.String `json:"Dsc"`            // {
	// "0": "\n  "
	// }, OR "Shadow is a handsome guy!  Just a bit sensitive, and always alert, Shadow hopes to find a family with a secure back yard and lots of time to play.  Shadow has had conflict with other dogs previously and harassed the cat in his former home so at this time, HSBV recommends an experienced family who keeps him as the only pet.  Shadow was overwhelmed by the young child in his previous home so HSBV will only rehome him to a family with children 10 years of age or older.",
	Photo1 null.String `json:"Photo1"` // "http:\/\/g.petango.com\/photos\/993\/b3a36066-352c-49bd-b495-ffb960123738.jpg",
	Photo2 null.String `json:"Photo2"` // {
	// "0": "\n  "
	// }, OR "http:\/\/g.petango.com\/photos\/993\/7471ab7a-a257-4c80-988e-ae1caaa603c4.jpg",
	Photo3 null.String `json:"Photo3"` // {
	// "0": "\n  "
	// }, OR "http:\/\/g.petango.com\/photos\/993\/ffb2899b-c7c2-48b1-b5a0-4a976e455fdd.jpg",
	OnHold       string            `json:"OnHold"`       // "No",
	SpecialNeeds map[string]string `json:"SpecialNeeds"` // {
	// "0": "\n  "
	// },
	NoDogs map[string]string `json:"NoDogs"` // {
	// "0": "\n  "
	// },
	NoCats map[string]string `json:"NoCats"` // {
	// "0": "\n  "
	// },
	NoKids map[string]string `json:"NoKids"` // {
	// "0": "\n  "
	// },
	BehaviorResult map[string]string `json:"BehaviorResult"` // {
	// "0": "\n  "
	// },
	MemoList         interface{} `json:"MemoList"`         // {},
	Site             string      `json:"Site"`             // "Humane Society of Boulder Valley",
	DateOfSurrender  string      `json:"DateOfSurrender"`  // "2019-05-20 14:30:21",
	TimeInFormerHome null.String `json:"TimeInFormerHome"` // {
	// "0": "\n  "
	// }, OR "1 Years",
	ReasonForSurrender null.String `json:"ReasonForSurrender"` // {
	// "0": "\n  "
	// }, OR "Behavioral",
	PrevEnvironment      string            `json:"PrevEnvironment"`      // "Unknown                       ",
	LivedWithChildren    string            `json:"LivedWithChildren"`    // "No",
	LivedWithAnimals     string            `json:"LivedWithAnimals"`     // "No",
	LivedWithAnimalTypes map[string]string `json:"LivedWithAnimalTypes"` // {
	// "0": "\n  "
	// },
	BodyWeight  string            `json:"BodyWeight"`  // "82 pounds",
	DateOfBirth string            `json:"DateOfBirth"` // "2016-05-20",
	ARN         map[string]string `json:"ARN"`         // {
	// "0": "\n  "
	// },
	VideoID map[string]string `json:"VideoID"` // {
	// "0": "\n  "
	// },
	BehaviorTestList     interface{}       `json:"BehaviorTestList"`     // {},
	Stage                string            `json:"Stage"`                // "Available",
	AnimalType           string            `json:"AnimalType"`           // "Dog",
	AgeGroup             string            `json:"AgeGroup"`             // "Adult",
	WildlifeIntakeInjury map[string]string `json:"WildlifeIntakeInjury"` // {
	// "0": "\n  "
	// },
	WildlifeIntakeCause map[string]string `json:"WildlifeIntakeCause"` // {
	// "0": "\n  "
	// },
	BuddyID     string      `json:"BuddyID"`     // "0",
	Featured    string      `json:"Featured"`    // "No",
	Sublocation string      `json:"Sublocation"` // "07",
	ChipNumber  null.String `json:"ChipNumber"`  // {
	// "0": "\n  "
	// }, OR "981020013624912",
	ColorPattern null.String `json:"ColorPattern"` // {
	// "0": "\n  "
	// }, OR "Tick"
	AdoptionApplicationURL map[string]string `json:"AdoptionApplicationUrl"` // {
	// "0": "\n  "
	// },
	BannerURL string `json:"BannerURL"` // "https:\/\/ws.petango.com\/webservices\/adoptablesearch\/images\/24PW-Web-Services-Graphic-Trial.png"
}

func (p Pup) IsAvailable() (bool, error) {
	if p.OnHold != "No" && p.OnHold != "Yes" {
		return false, fmt.Errorf("Unexpected value for pup on-hold status: %s", p.OnHold)
	}

	return p.OnHold == "No", nil
}

func (p Pup) IsRightAge() (bool, error) {
	months, err := strconv.Atoi(p.Age)
	if err != nil {
		return false, err
	}

	return months >= 36 && months < 84, nil
}

func (p Pup) BarkGreeting() {
	fmt.Println()
	fmt.Println()
	fmt.Println()

	p.barkAscii()

	p.barkName()
	p.barkAge()
	p.barkWeight()
	p.barkBreed()
	p.barkDesc()
	p.barkPics()

	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func (p Pup) barkAscii() {
	fmt.Println(WoofASCII)
}

func (p Pup) barkName() {
	fmt.Printf("My name is %v\n", p.AnimalName)
}

func (p Pup) barkAge() {
	months, err := strconv.Atoi(p.Age)
	if err != nil {
		log.Println(err)
		fmt.Println("They're not sure how old I am...")
		return
	}

	if months < 12 {
		fmt.Printf("I'm %d months old\n", months)
	} else {
		years := months / 12
		remainder := months % 12

		s := fmt.Sprintf("I'm %d year", years)
		if years > 1 {
			s += "s"
		}

		if remainder > 0 {
			s += fmt.Sprintf(" and %d month", remainder)
		}
		if remainder > 1 {
			s += "s"
		}

		fmt.Printf("%s old\n", s)
	}
}

func (p Pup) barkWeight() {
	fmt.Printf("I'm %s\n", p.BodyWeight)
}

func (p Pup) barkBreed() {
	s := fmt.Sprintf("I'm a %s", p.PrimaryBreed)

	if p.SecondaryBreed.Valid {
		s += fmt.Sprintf(" %s", p.SecondaryBreed.Value)
	}

	fmt.Println(s)
}

func (p Pup) barkDesc() {
	if p.Description.Valid {
		fmt.Printf("Here's what people say about me:\n%s\n", p.Description.Value)
	}
}

func (p Pup) barkPics() {
	pics := []string{}

	for _, photo := range []null.String{p.Photo1, p.Photo2, p.Photo3} {
		if photo.Valid {
			picStripped := strings.Replace(photo.Value, "\\", "", -1)
			pics = append(pics, picStripped)
		}
	}

	if len(pics) > 0 {
		fmt.Println("Wanna see pictures of me?")
		for _, pic := range pics {
			fmt.Println(pic)
		}
	}
}
