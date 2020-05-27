package pup

import (
	"fmt"
	"html"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/asgaines/pupsniffer/null"
	"github.com/asgaines/pupsniffer/utils"
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
	OnHold       string      `json:"OnHold"`       // "No",
	SpecialNeeds null.String `json:"SpecialNeeds"` // {
	// "0": "\n  "
	// },
	NoDogs null.String `json:"NoDogs"` // {
	// "0": "\n  "
	// },
	NoCats null.String `json:"NoCats"` // {
	// "0": "\n  "
	// },
	NoKids null.String `json:"NoKids"` // {
	// "0": "\n  "
	// },
	BehaviorResult null.String `json:"BehaviorResult"` // {
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
	PrevEnvironment      string      `json:"PrevEnvironment"`      // "Unknown                       ",
	LivedWithChildren    string      `json:"LivedWithChildren"`    // "No",
	LivedWithAnimals     string      `json:"LivedWithAnimals"`     // "No",
	LivedWithAnimalTypes null.String `json:"LivedWithAnimalTypes"` // {
	// "0": "\n  "
	// },
	BodyWeight  string      `json:"BodyWeight"`  // "82 pounds",
	DateOfBirth string      `json:"DateOfBirth"` // "2016-05-20",
	ARN         null.String `json:"ARN"`         // {
	// "0": "\n  "
	// },
	VideoID null.String `json:"VideoID"` // {
	// "0": "\n  "
	// },
	BehaviorTestList     interface{} `json:"BehaviorTestList"`     // {},
	Stage                string      `json:"Stage"`                // "Available",
	AnimalType           string      `json:"AnimalType"`           // "Dog",
	AgeGroup             string      `json:"AgeGroup"`             // "Adult",
	WildlifeIntakeInjury null.String `json:"WildlifeIntakeInjury"` // {
	// "0": "\n  "
	// },
	WildlifeIntakeCause null.String `json:"WildlifeIntakeCause"` // {
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
	AdoptionApplicationURL null.String `json:"AdoptionApplicationUrl"` // {
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

func (p Pup) IsOlderAndBigger() (bool, error) {
	months, err := p.AgeInMonths()
	if err != nil {
		return false, err
	}

	weight, err := p.getWeight()
	if err != nil {
		return false, err
	}

	return months >= 36 && months < 84 && weight >= 35, nil
}

func (p Pup) AgeInMonths() (int, error) {
	months, err := strconv.Atoi(p.Age)
	return months, err
}

func (p Pup) IsPuppy() (bool, error) {
	months, err := p.AgeInMonths()
	if err != nil {
		return false, err
	}

	return months <= 12, nil
}

func (p Pup) BarkGreeting() {
	fmt.Println()
	fmt.Println()
	fmt.Println()

	p.barkASCII()

	p.barkName()
	fmt.Printf("I'm %s old\n", p.BarkAge())
	p.barkSex()
	p.barkWeight()
	fmt.Printf("I'm a %s", p.BarkBreed())
	p.barkDesc()
	p.barkReasonForSurrender()
	p.barkSpecialNeeds()
	pics := p.BarkPics()

	if len(pics) > 0 {
		fmt.Println("Wanna see pictures of me?")
		for _, pic := range pics {
			fmt.Println(pic)
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func (p Pup) barkASCII() {
	fmt.Println(WoofASCII)
}

func (p Pup) barkName() {
	fmt.Printf("My name is %v\n", p.AnimalName)
}

func (p Pup) BarkAge() string {
	months, err := strconv.Atoi(p.Age)
	if err != nil {
		log.Println(err)
		return "They're not sure how old I am..."
	}

	if months < 12 {
		return fmt.Sprintf("%d months", months)
	}

	years := months / 12
	remainder := months % 12

	s := fmt.Sprintf("%d year", years)
	if years > 1 {
		s += "s"
	}

	if remainder > 0 {
		s += fmt.Sprintf(" and %d month", remainder)
	}
	if remainder > 1 {
		s += "s"
	}

	return s
}

func (p Pup) barkSex() {
	fmt.Printf("I'm a %s\n", p.Sex)
}

func (p Pup) barkWeight() {
	fmt.Printf("I'm %s\n", p.BodyWeight)
}

func (p Pup) getWeight() (int, error) {
	weightParts := strings.Split(p.BodyWeight, " ")
	if len(weightParts) < 1 {
		return 0, fmt.Errorf("unexpected weight: %s", p.BodyWeight)
	}

	weight, err := strconv.Atoi(weightParts[0])
	if err != nil {
		return 0, fmt.Errorf("expected int for first part of weight: %s. err: %s", p.BodyWeight, err)
	}

	return weight, nil
}

func (p Pup) BarkBreed() string {
	s := fmt.Sprintf("%s", p.PrimaryBreed)

	if p.SecondaryBreed.Valid {
		s += fmt.Sprintf(" %s", p.SecondaryBreed.Value)
	}

	return s
}

func (p Pup) BarkTimeSince() (string, error) {
	intakeTime, err := time.Parse("2006-01-02 15:04:05", p.LastIntakeDate)
	if err != nil {
		return "", err
	}

	duration := time.Since(intakeTime)

	days := duration.Hours() / 24
	if days > 0 {
		return fmt.Sprintf("%d %s", int(days), utils.Pluralize("day", "days", int(days))), nil
	}

	if duration.Hours() > 0 {
		return fmt.Sprintf("%d %s", int(duration.Hours()), utils.Pluralize("hour", "hours", int(duration.Hours()))), nil
	}

	if duration.Minutes() > 0 {
		return fmt.Sprintf("%d %s", int(duration.Minutes()), utils.Pluralize("minute", "minutes", int(duration.Minutes()))), nil
	}

	return fmt.Sprintf("%d %s", int(duration.Seconds()), utils.Pluralize("second", "seconds", int(duration.Seconds()))), nil
}

func (p Pup) barkDesc() {
	if p.Description.Valid {
		fmt.Printf("Here's what people say about me: \"%s\n\"", p.Description.Value)
	}
}

func (p Pup) BarkDesc() string {
	return html.UnescapeString(p.Description.Value)
}

func (p Pup) barkReasonForSurrender() {
	if p.ReasonForSurrender.Valid {
		fmt.Printf("Here's why I'm here: %s\n", p.ReasonForSurrender.Value)
	}
}

func (p Pup) barkSpecialNeeds() {
	if p.SpecialNeeds.Valid {
		fmt.Printf("My special needs: %s\n", p.SpecialNeeds.Value)
	}
}

func (p Pup) BarkDeclawed() (bool, error) {
	if p.Declawed == "Yes" {
		return true, nil
	}

	if p.Declawed == "No" {
		return false, nil
	}

	return false, fmt.Errorf("Unexpected value for Declawed: %s", p.Declawed)
}

func (p Pup) BarkPics() []string {
	pics := []string{}

	for _, photo := range []null.String{p.Photo1, p.Photo2, p.Photo3} {
		if photo.Valid {
			picStripped := strings.Replace(photo.Value, "\\", "", -1)
			pics = append(pics, picStripped)
		}
	}

	return pics
}
