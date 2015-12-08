package main

import (
	"math/rand"
	"time"
)

func Getalias() string {
	adjs := []string{"anachronistic", "analogous", "analytical", "anatomical", "ancestral", "ancient", "ancillary", "andean", "anecdotal", "anemic", "anesthetized", "angeles-based",
		"anglican", "anglo-american", "anglo-saxon", "angolan", "angry", "anguished", "animated", "annoyed", "annoying", "annual", "anomalous", "anonymous", "answerable", "antagonistic", "antebellum",
		"anthropological", "anti-abortion", "anti-aircraft", "anti-american", "anti-apartheid", "anti-cancer", "anti-castro", "anti-catholic", "anticipatory", "anti-communist", "anti-crime",
		"antidemocratic", "anti-depressant", "anti-discrimination", "anti-drug", "anti-dumping", "anti-foreign", "anti-fraud", "anti-government", "anti-growth", "anti-japanese", "anti-lock",
		"anti-nuclear", "antiquated", "antique", "antiqued", "anti-rejection", "anti-semitic", "anti-smoking",
		"antisocial", "anti-social", "anti-soviet", "anti-takeover", "anti-tax", "anti-terrorism", "antithetical", "antitrust", "anti-trust", "anti-virus", "antiwar", "anxious", "anytactical",
		"fabled", "fabulous", "face-saving", "facetious", "face-to-face", "facial", "fact-finding", "factual", "faint", "faintest", "fair", "fairer", "fairy", "faithful",
		"fake", "fallacious", "fallen", "fallible", "fallow", "false", "famed", "familiar", "family-oriented", "family-owned", "family-planning", "famous", "fanatic", "fanatical",
		"fancier", "fanciful", "fancy", "fantastic", "far", "faraway", "far-fetched", "far-flung", "far-left", "farmed", "far-reaching", "far-right", "farthest",
		"fascinating", "fascist", "fashionable", "fashion-conscious", "fast", "fast-acting", "fast-approaching", "faster", "fastest", "fastest-growing",
		"fast-flowing", "fast-forward", "fast-growing", "fastidious", "fasting", "fast-living",
		"fast-moving", "fast-paced", "fast-rising", "fast-track", "fat", "fatal", "fateful", "fatter", "faultless", "faulty", "faux", "favorable", "favorite",
		"roadside", "roasted", "robotic", "robust", "rocket-propelled", "rock-hard", "rocky", "rogue", "roman", "romanesque", "romantic", "rome-based", "roomy",
		"rosier", "rosy", "rotary", "rote", "rotten", "rough", "rough-and-tumble", "rough-cut", "rougher", "roughest",
		"rough-hewn", "round", "round-table", "round-the-clock", "round-the-world", "round-trip", "routine", "rowdy", "royal", "sleepy", "creepy", "vindictive", "horrible", "blazed", "glazed", "420-blaze-it", "grade-A",
		"unrepentant", "unreported", "unresolved", "unresponsive", "unrestrained", "unrestricted", "unrivaled", "unruly", "unsafe", "unsatisfactory", "unsatisfying", "edgy",
		"unsavory", "unscathed", "unscheduled", "unscientific", "unscripted", "unscrupulous", "unseated", "unsecured", "unseemly", "unseen", "unsentimental", "unserious",
		"unsettled", "unsettling", "unshakable", "unsightly", "unsigned", "unskilled", "unsold", "unsolicited", "unsolved", "unsound", "unspeakable", "unspecified", "unspent",
		"unspoiled", "unstable", "unstinting", "unstoppable", "unstructured", "unsubstantiated", "unsuccessful", "unsuitable", "unsupported", "unsure", "unsuspecting",
		"unsustainable", "unsympathetic", "untamed", "untapped", "untarnished", "untenable", "untested", "unthinkable", "untimely", "untold", "untouchable", "untouched", "untrained", "untreated", "untried", "untrue",
		"untrustworthy", "unturned", "unusable", "unused", "unusual", "unviable", "unwanted", "unwarranted", "unwary", "unwashed", "unwed", "unwelcome", "unwholesome", "unwieldy",
		"unwilling", "unwise", "unwitting", "unworkable", "triumphant", "black", "white", "grey", "transcendant", "greedy", "exalted", "radiant", "iradiated", "disgruntled", "explosive", "stiff", "filthy",
		"unworthy", "unwritten", "unyielding", "tired", "sloppy", "sharp"}
	nouns := []string{"shovel", "dog", "jerry", "bus", "cat", "sandwitch", "turtle", "coat", "wall", "pencil", "pen", "cup", "window", "garage", "closet", "glove", "sausage", "speaker", "zealot", "paladin",
		"car", "boat", "airplane", "sofa", "lightbulb", "genius", "map", "programmer", "fish", "yeti", "dragon", "pipe", "wizard", "noodle", "pancake", "rock", "jug", "pickle", "phone", "pie", "dummy", "bookworm",
		"chair", "tube", "squirtgun", "beercan", "tree", "printer", "executive", "knight", "spoon", "man-serpant", "sword", "Stark", "Lannister", "cone", "cupcake", "candle", "batman", "hamster", "goat", "tax-attorney", "dwarf", "elf",
		"notebook", "bear", "mountain", "backpack", "orc", "kobinationskraftwagen", "kugelschreiber", "hand-shoe", "hard-drive", "zombie", "washcloth", "cake", "doughnut", "beefcake", "deck-of-cards", "spatula", "potato", "jellybean",
		"pokemon", "chair", "sediment", "penguin"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return adjs[r1.Intn(len(adjs))] + "-" + nouns[r1.Intn(len(nouns))]
}
