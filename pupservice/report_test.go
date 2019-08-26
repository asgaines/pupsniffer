package pupservice

import (
	"os"
	"testing"

	"github.com/asgaines/pupsniffer/null"
	"github.com/asgaines/pupsniffer/pup"
)

func TestPupReport(t *testing.T) {
	pupsvc := newMockPupSvc()

	cases := []struct {
		pups []*pup.Pup
		err  error
	}{
		{
			pups: []*pup.Pup{
				&pup.Pup{
					CompanyID:      "993",
					ID:             "42399012",
					AnimalName:     "Dewey",
					Species:        "Dog",
					Sex:            "Male",
					Altered:        "Yes",
					PrimaryBreed:   "Retriever, Labrador",
					SecondaryBreed: null.String{Valid: false},
					PrimaryColor:   "Black",
					SecondaryColor: null.String{Valid: false},
					Age:            "36",
					Size:           "L",
					Housetrained:   "Unknown",
					Declawed:       "No",
					Price:          "99.00",
					LastIntakeDate: "2019-08-10 18:59:00",
					Location:       "DA",
					Description: null.String{
						Value: "Dewey is a handsome boy!  He is athletic and will enjoy living with an active family.  Dewey can be fussy when interacting with some other dogs and would likely appreciate being the center of his family&#39;s attention.  He is recommended to live without other dogs at this time.",
						Valid: true,
					},
					Photo1: null.String{
						Value: "http://g.petango.com/photos/993/c5273d77-c1da-444e-a75b-71710fa46b2e.jpg",
						Valid: true,
					},
					Photo2: null.String{
						Value: "http://g.petango.com/photos/993/f37d1e51-038c-4e7c-88f7-ed4347094404.jpg",
						Valid: true,
					},
					Photo3: null.String{
						Value: "http://g.petango.com/photos/993/5d0e21a7-c026-41e4-873a-124f232155c7.jpg",
						Valid: true,
					},
					OnHold:               "No",
					SpecialNeeds:         null.String{Valid: false},
					NoDogs:               null.String{Valid: false},
					NoCats:               null.String{Valid: false},
					NoKids:               null.String{Valid: false},
					BehaviorResult:       null.String{Valid: false},
					MemoList:             struct{}{},
					Site:                 "Humane Society of Boulder Valley",
					TimeInFormerHome:     null.String{Valid: false},
					ReasonForSurrender:   null.String{Valid: false},
					PrevEnvironment:      "Unknown                       ",
					LivedWithChildren:    "No",
					LivedWithAnimals:     "No",
					LivedWithAnimalTypes: null.String{Valid: false},
					BodyWeight:           "63 pounds",
					DateOfBirth:          "2016-08-06",
					ARN:                  null.String{Valid: false},
					VideoID:              null.String{Valid: false},
					BehaviorTestList:     struct{}{},
					Stage:                "Available",
					AnimalType:           "Dog",
					AgeGroup:             "Adult",
					WildlifeIntakeInjury: null.String{Valid: false},
					WildlifeIntakeCause:  null.String{Valid: false},
					BuddyID:              "0",
					Featured:             "No",
					Sublocation:          "03",
					ChipNumber: null.String{
						Value: "982126055775478",
						Valid: true,
					},
					ColorPattern:           null.String{Valid: false},
					AdoptionApplicationURL: null.String{Valid: false},
					BannerURL:              "https://ws.petango.com/webservices/adoptablesearch/images/24PW-Web-Services-Graphic-Trial.png",
				},
				&pup.Pup{
					CompanyID:      "993",
					ID:             "42431385",
					AnimalName:     "Lenny",
					Species:        "Dog",
					Sex:            "Male",
					Altered:        "Yes",
					PrimaryBreed:   "German Shepherd",
					SecondaryBreed: null.String{Value: "Mix"},
					PrimaryColor:   "Black",
					SecondaryColor: null.String{Value: "Tan"},
					Age:            "40",
					Size:           "M",
					Housetrained:   "Unknown",
					Declawed:       "No",
					Price:          "79.00",
					LastIntakeDate: "2019-08-08 16:26:00",
					Location:       "DA",
					Description:    null.String{Value: "Lenny is athletic and will enjoy exploring Colorado with an active guardian!   Lenny has been stressed in the shelter and will benefit from a quiet transition into his new home.  HSBV will only rehome him to a family with children 10 years of age or older."},
					Photo1: null.String{
						Value: "http://g.petango.com/photos/993/ed69a9ae-c7d6-4fd9-9073-9ee14925ab6c.jpg",
						Valid: true,
					},
					Photo2: null.String{
						Value: "http://g.petango.com/photos/993/ea1edda7-3e2c-4a1e-af1b-26adb1c06831.jpg",
						Valid: true,
					},
					Photo3: null.String{
						Value: "http://g.petango.com/photos/993/f0522702-2421-481b-b2dc-0146f332181f.jpg",
						Valid: true,
					},
					OnHold:                 "No",
					SpecialNeeds:           null.String{Valid: false},
					NoDogs:                 null.String{Valid: false},
					NoCats:                 null.String{Valid: false},
					NoKids:                 null.String{Valid: false},
					BehaviorResult:         null.String{Valid: false},
					MemoList:               struct{}{},
					Site:                   "Humane Society of Boulder Valley",
					TimeInFormerHome:       null.String{Valid: false},
					ReasonForSurrender:     null.String{Valid: false},
					PrevEnvironment:        "Unknown                       ",
					LivedWithChildren:      "No",
					LivedWithAnimals:       "No",
					LivedWithAnimalTypes:   null.String{Valid: false},
					BodyWeight:             "60 pounds",
					DateOfBirth:            "2016-04-04",
					ARN:                    null.String{Valid: false},
					VideoID:                null.String{Valid: false},
					BehaviorTestList:       struct{}{},
					Stage:                  "Available",
					AnimalType:             "Dog",
					AgeGroup:               "Adult",
					WildlifeIntakeInjury:   null.String{Valid: false},
					WildlifeIntakeCause:    null.String{Valid: false},
					BuddyID:                "0",
					Featured:               "No",
					Sublocation:            "19",
					ChipNumber:             null.String{Value: "981020027306983"},
					ColorPattern:           null.String{Valid: false},
					AdoptionApplicationURL: null.String{Valid: false},
					BannerURL:              "https://ws.petango.com/webservices/adoptablesearch/images/24PW-Web-Services-Graphic-Trial.png",
				},
			},
			err: nil,
		},
	}

	for _, c := range cases {
		f, err := os.OpenFile("../testpups.html", os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		if err := pupsvc.PupReport(len(c.pups), c.pups, f); err != c.err {
			t.Errorf("Expected err %s, got %s\n", err, c.err)
		}
	}
}
