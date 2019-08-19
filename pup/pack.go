package pup

import "strconv"

// Pack is the collection of all listed pups
type Pack []packPup

// GetPupIDs extracts the pups' IDs to be used for a fetching of more detailed,
// per-pup info
func (p *Pack) GetPupIDs() ([]int, error) {
	pupIDs := make([]int, 0, len(*p))

	for _, pup := range *p {
		pupID, err := strconv.Atoi(pup.KPup.ID)
		if err != nil {
			return nil, err
		}

		pupIDs = append(pupIDs, pupID)
	}

	return pupIDs, nil
}

type packPup struct {
	KPup kpup `json:"adoptableSearch"`
}

type kpup struct {
	ID             string      `json:"ID"`             // "31396035"
	Name           string      `json:"Name"`           // "Tallulah Pearl"
	Species        string      `json:"Species"`        // "Dog"
	Sex            string      `json:"Sex"`            // "Female"
	PrimaryBreed   string      `json:"PrimaryBreed"`   // "Terrier, Americal Pit Bull"
	SecondaryBreed interface{} `json:"SecondaryBreed"` // "Mix"
	//SecondaryBreed string   `json:"SecondaryBreed"` // "Mix"
	SN interface{} `json:"SN"` // "Spayed" OR [
	//"\n      "
	//],
	Age          string   `json:"Age"`          // "41",
	Photo        string   `json:"Photo"`        // "http:\/\/g.petango.com\/photos\/993\/e1c6acfd-6aee-46bb-a9e7-a6f1d6533f50_TN1.jpg",
	Location     string   `json:"Location"`     // "DA",
	OnHold       string   `json:"OnHold"`       // "No",
	SpecialNeeds []string `json:"SpecialNeeds"` // [
	//"\n      "
	//],
	NoDogs []string `json:"NoDogs"` // [
	//"\n      "
	//],
	NoCats []string `json:"NoCats"` // [
	//"\n      "
	//],
	NoKids []string `json:"NoKids"` // [
	//"\n      "
	//],
	MemoList []string `json:"MemoList"` // [],
	ARN      []string `json:"ARN"`      // [
	//"\n      "
	//],
	BehaviorTestList []string `json:"BehaviorTestList"` // [
	//"\n      "
	//],
	Stage                string   `json:"Stage"`                // "Available",
	AnimalType           string   `json:"AnimalType"`           // "Dog",
	AgeGroup             string   `json:"AgeGroup"`             // "Adult",
	WildlifeIntakeInjury []string `json:"WildlifeIntakeInjury"` // [
	//"\n      "
	//],
	WildlifeIntakeCase []string `json:"WildlifeIntakeCause"` // [
	//"\n      "
	//],
	BuddyID     string      `json:"BuddyID"`     // "0",
	Featured    string      `json:"Featured"`    // "No",
	Sublocation string      `json:"Sublocation"` // "05",
	ChipNumber  interface{} `json:"ChipNumber"`  // "982000406116374",
	//ChipNumber     string `json:"ChipNumber"`     // "982000406116374",
	FreshnessStamp string `json:"FreshnessStamp"` // "06\/16\/2019 08:42:18pm MDT"
}
