package config

// BaseURL is the base path to the Boulder Humane Society's API
const BaseURL = "https://www.boulderhumane.org/wp-content/plugins/Petpoint-Webservices-2018/"

// PackURL is the base path to the endpoint serving the collection of pups
const PackURL = BaseURL + "pullanimals.php?type=dog"

// PupURL is the base path to the endpoint serving the details for a single pup
const PupURL = BaseURL + "viewanimal.php?id=%d"
