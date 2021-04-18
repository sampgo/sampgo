package sampgo

const (
	firstInvalidVehicleModel = iota + 399
	Landstalker
	Bravura
	Buffalo
	Linerunner
	Perennial
	Sentinel
	Dumper
	Firetruck
	Trashmaster
	Stretch
	Manana
	Infernus
	Voodoo
	Pony
	Mule
	Cheetah
	Ambulance
	Leviathan
	Moonbeam
	Esperanto
	Taxi
	Washington
	Bobcat
	MrWhoopee
	BFInjection
	Hunter
	Premier
	Enforcer
	Securicar
	Banshee
	Predator
	Bus
	Rhino
	Barracks
	Hotknife
	ArticleTrailer
	Previon
	Coach
	Cabbie
	Stallion
	Rumpo
	RCBandit
	Romero
	Packer
	Monster
	Admiral
	Squalo
	Seasparrow
	Pizzaboy
	Tram
	ArticleTrailer2
	Turismo
	Speeder
	Reefer
	Tropic
	Flatbed
	Yankee
	Caddy
	Solair
	TopfunVanBerkleysRC
	Skimmer
	PCJ600
	Faggio
	Freeway
	RCBaron
	RCRaider
	Glendale
	Oceanic
	Sanchez
	Sparrow
	Patriot
	Quad
	Coastguard
	Dinghy
	Hermes
	Sabre
	Rustler
	ZR350
	Walton
	Regina
	Comet
	BMX
	Burrito
	Camper
	Marquis
	Baggage
	Dozer
	Maverick
	SANNewsMaverick
	Rancher
	FBIRancher
	Virgo
	Greenwood
	Jetmax
	HotringRacer
	Sandking
	BlistaCompact
	PoliceMaverick
	Boxville
	Benson
	Mesa
	RCGoblin
	HotringRacerA
	HotringRacerB
	BloodringBanger
	RancherLure
	SuperGT
	Elegant
	Journey
	Bike
	MountainBike
	Beagle
	Cropduster
	Stuntplane
	Tanker
	Roadtrain
	Nebula
	Majestic
	Buccaneer
	Shamal
	Hydra
	FCR900
	NRG500
	HPV1000
	CementTruck
	Towtruck
	Fortune
	Cadrona
	FBITruck
	Willard
	Forklift
	Tractor
	CombineHarvester
	Feltzer
	Remington
	Slamvan
	Blade
	FreightTrain
	BrownstreakTrain
	Vortex
	Vincent
	Bullet
	Clover
	Sadler
	FiretruckLA
	Hustler
	Intruder
	Primo
	Cargobob
	Tampa
	Sunrise
	Merit
	UtilityVan
	Nevada
	Yosemite
	Windsor
	MonsterA
	MonsterB
	Uranus
	Jester
	Sultan
	Stratum
	Elegy
	Raindance
	RCTiger
	Flash
	Tahoma
	Savanna
	Bandito
	FreightFlatTrailerTrain
	StreakTrailerTrain
	Kart
	Mower
	Dune
	Sweeper
	Broadway
	Tornado
	AT400
	DFT30
	Huntley
	Stafford
	BF400
	Newsvan
	Tug
	PetrolTrailer
	Emperor
	Wayfarer
	Euros
	Hotdog
	Club
	FreightBoxTrailerTrain
	ArticleTrailer3
	Andromada
	Dodo
	RCCam
	Launch
	PoliceCarLSPD
	PoliceCarSFPD
	PoliceCarLVPD
	PoliceRanger
	Picador
	SWAT
	Alpha
	Phoenix
	GlendaleShit
	SadlerShit
	BaggageTrailerA
	BaggageTrailerB
	TugStairsTrailer
	Boxville2
	FarmTrailer
	UtilityTrailer
	lastInvalidVehicleModel
)

func IsValidVehicleModel(modelid int) bool {
	return modelid > firstInvalidVehicleModel && modelid < lastInvalidVehicleModel
}

func IsValidVehicleModelName(model string) bool {
	if model == "" {
		return false
	}

	vehicleMap := map[string]int{
		"Landstalker":               400,
		"Bravura":                   401,
		"Buffalo":                   402,
		"Linerunner":                403,
		"Perennial":                 404,
		"Sentinel":                  405,
		"Dumper":                    406,
		"Firetruck":                 407,
		"Trashmaster":               408,
		"Stretch":                   409,
		"Manana":                    410,
		"Infernus":                  411,
		"Voodoo":                    412,
		"Pony":                      413,
		"Mule":                      414,
		"Cheetah":                   415,
		"Ambulance":                 416,
		"Leviathan":                 417,
		"Moonbeam":                  418,
		"Esperanto":                 419,
		"Taxi":                      420,
		"Washington":                421,
		"Bobcat":                    422,
		"Mr Whoopee":                423,
		"BF Injection":              424,
		"Hunter":                    425,
		"Premier":                   426,
		"Enforcer":                  427,
		"Securicar":                 428,
		"Banshee":                   429,
		"Predator":                  430,
		"Bus":                       431,
		"Rhino":                     432,
		"Barracks":                  433,
		"Hotknife":                  434,
		"Article Trailer":           435,
		"Previon":                   436,
		"Coach":                     437,
		"Cabbie":                    438,
		"Stallion":                  439,
		"Rumpo":                     440,
		"RC Bandit":                 441,
		"Romero":                    442,
		"Packer":                    443,
		"Monster":                   444,
		"Admiral":                   445,
		"Squalo":                    446,
		"Seasparrow":                447,
		"Pizzaboy":                  448,
		"Tram":                      449,
		"Article Trailer 2":         450,
		"Turismo":                   451,
		"Speeder":                   452,
		"Reefer":                    453,
		"Tropic":                    454,
		"Flatbed":                   455,
		"Yankee":                    456,
		"Caddy":                     457,
		"Solair":                    458,
		"BerkleysRC":                459,
		"Skimmer":                   460,
		"PCJ-600":                   461,
		"Faggio":                    462,
		"Freeway":                   463,
		"RC Baron":                  464,
		"RC Raider":                 465,
		"Glendale":                  466,
		"Oceanic":                   467,
		"Sanchez":                   468,
		"Sparrow":                   469,
		"Patriot":                   470,
		"Quad":                      471,
		"Coastguard":                472,
		"Dinghy":                    473,
		"Hermes":                    474,
		"Sabre":                     475,
		"Rustler":                   476,
		"ZR-350":                    477,
		"Walton":                    478,
		"Regina":                    479,
		"Comet":                     480,
		"BMX":                       481,
		"Burrito":                   482,
		"Camper":                    483,
		"Marquis":                   484,
		"Baggage":                   485,
		"Dozer":                     486,
		"Maverick":                  487,
		"SANews Maverick":           488,
		"Rancher":                   489,
		"FBI Rancher":               490,
		"Virgo":                     491,
		"Greenwood":                 492,
		"Jetmax":                    493,
		"Hotring Racer":             494,
		"Sandking":                  495,
		"Blista Compact":            496,
		"Police Maverick":           497,
		"Boxville":                  498,
		"Benson":                    499,
		"Mesa":                      500,
		"RC Goblin":                 501,
		"Hotring RacerA":            502,
		"Hotring RacerB":            503,
		"Bloodring Banger":          504,
		"Rancher Lure":              505,
		"Super GT":                  506,
		"Elegant":                   507,
		"Journey":                   508,
		"Bike":                      509,
		"Mountain Bike":             510,
		"Beagle":                    511,
		"Cropduster":                512,
		"Stuntplane":                513,
		"Tanker":                    514,
		"Roadtrain":                 515,
		"Nebula":                    516,
		"Majestic":                  517,
		"Buccaneer":                 518,
		"Shamal":                    519,
		"Hydra":                     520,
		"FCR-900":                   521,
		"NRG-500":                   522,
		"HPV1000":                   523,
		"Cement Truck":              524,
		"Towtruck":                  525,
		"Fortune":                   526,
		"Cadrona":                   527,
		"FBI Truck":                 528,
		"Willard":                   529,
		"Forklift":                  530,
		"Tractor":                   531,
		"Combine Harvester":         532,
		"Feltzer":                   533,
		"Remington":                 534,
		"Slamvan":                   535,
		"Blade":                     536,
		"Freight Train":             537,
		"Brownstreak Train":         538,
		"Vortex":                    539,
		"Vincent":                   540,
		"Bullet":                    541,
		"Clover":                    542,
		"Sadler":                    543,
		"Firetruck 2":               544,
		"Hustler":                   545,
		"Intruder":                  546,
		"Primo":                     547,
		"Cargobob":                  548,
		"Tampa":                     549,
		"Sunrise":                   550,
		"Merit":                     551,
		"Utility Van":               552,
		"Nevada":                    553,
		"Yosemite":                  554,
		"Windsor":                   555,
		"Monster A":                 556,
		"Monster B":                 557,
		"Uranus":                    558,
		"Jester":                    559,
		"Sultan":                    560,
		"Stratum":                   561,
		"Elegy":                     562,
		"Raindance":                 563,
		"RC Tiger":                  564,
		"Flash":                     565,
		"Tahoma":                    566,
		"Savanna":                   567,
		"Bandito":                   568,
		"Flat Trailer Train":        569,
		"Streak Trailer Train":      570,
		"Kart":                      571,
		"Mower":                     572,
		"Dune":                      573,
		"Sweeper":                   574,
		"Broadway":                  575,
		"Tornado":                   576,
		"AT-400":                    577,
		"DFT30":                     578,
		"Huntley":                   579,
		"Stafford":                  580,
		"BF400":                     581,
		"Newsvan":                   582,
		"Tug":                       583,
		"Petrol Trailer":            584,
		"Emperor":                   585,
		"Wayfarer":                  586,
		"Euros":                     587,
		"Hotdog":                    588,
		"Club":                      589,
		"Freight Box Trailer Train": 590,
		"Article Trailer 3":         591,
		"Andromada":                 592,
		"Dodo":                      593,
		"RCCam":                     594,
		"Launch":                    595,
		"Police Car LSPD":           596,
		"Police Car SFPD":           597,
		"Police Car LVPD":           598,
		"Police Ranger":             599,
		"Picador":                   600,
		"SWAT":                      601,
		"Alpha":                     602,
		"Phoenix":                   603,
		"Glendale Shit":             604,
		"SadlerShit":                605,
		"Baggage Trailer A":         606,
		"Baggage Trailer B":         607,
		"TugStairs Trailer":         608,
		"Boxville 2":                609,
		"Farm Trailer":              610,
		"Utility Trailer":           611,
	}

	_, ok := vehicleMap[model]
	return ok
}

func GetVehicleModelName(modelid int) string {
	switch modelid {
	case Landstalker:
		return "Landstalker"
	case Bravura:
		return "Bravura"
	case Buffalo:
		return "Buffalo"
	case Linerunner:
		return "Linerunner"
	case Perennial:
		return "Perennial"
	case Sentinel:
		return "Sentinel"
	case Dumper:
		return "Dumper"
	case Firetruck:
		return "Firetruck"
	case Trashmaster:
		return "Trashmaster"
	case Stretch:
		return "Stretch"
	case Manana:
		return "Manana"
	case Infernus:
		return "Infernus"
	case Voodoo:
		return "Voodoo"
	case Pony:
		return "Pony"
	case Mule:
		return "Mule"
	case Cheetah:
		return "Cheetah"
	case Ambulance:
		return "Ambulance"
	case Leviathan:
		return "Leviathan"
	case Moonbeam:
		return "Moonbeam"
	case Esperanto:
		return "Esperanto"
	case Taxi:
		return "Taxi"
	case Washington:
		return "Washington"
	case Bobcat:
		return "Bobcat"
	case MrWhoopee:
		return "Mr Whoopee"
	case BFInjection:
		return "BF Injection"
	case Hunter:
		return "Hunter"
	case Premier:
		return "Premier"
	case Enforcer:
		return "Enforcer"
	case Securicar:
		return "Securicar"
	case Banshee:
		return "Banshee"
	case Predator:
		return "Predator"
	case Bus:
		return "Bus"
	case Rhino:
		return "Rhino"
	case Barracks:
		return "Barracks"
	case Hotknife:
		return "Hotknife"
	case ArticleTrailer:
		return "Article Trailer"
	case Previon:
		return "Previon"
	case Coach:
		return "Coach"
	case Cabbie:
		return "Cabbie"
	case Stallion:
		return "Stallion"
	case Rumpo:
		return "Rumpo"
	case RCBandit:
		return "RC Bandit"
	case Romero:
		return "Romero"
	case Packer:
		return "Packer"
	case Monster:
		return "Monster"
	case Admiral:
		return "Admiral"
	case Squalo:
		return "Squalo"
	case Seasparrow:
		return "Seasparrow"
	case Pizzaboy:
		return "Pizzaboy"
	case Tram:
		return "Tram"
	case ArticleTrailer2:
		return "Article Trailer 2"
	case Turismo:
		return "Turismo"
	case Speeder:
		return "Speeder"
	case Reefer:
		return "Reefer"
	case Tropic:
		return "Tropic"
	case Flatbed:
		return "Flatbed"
	case Yankee:
		return "Yankee"
	case Caddy:
		return "Caddy"
	case Solair:
		return "Solair"
	case TopfunVanBerkleysRC:
		return "Topfun Van (Berkley's RC)"
	case Skimmer:
		return "Skimmer"
	case PCJ600:
		return "PCJ-600"
	case Faggio:
		return "Faggio"
	case Freeway:
		return "Freeway"
	case RCBaron:
		return "RC Baron"
	case RCRaider:
		return "RC Raider"
	case Glendale:
		return "Glendale"
	case Oceanic:
		return "Oceanic"
	case Sanchez:
		return "Sanchez"
	case Sparrow:
		return "Sparrow"
	case Patriot:
		return "Patriot"
	case Quad:
		return "Quad"
	case Coastguard:
		return "Coastguard"
	case Dinghy:
		return "Dinghy"
	case Hermes:
		return "Hermes"
	case Sabre:
		return "Sabre"
	case Rustler:
		return "Rustler"
	case ZR350:
		return "ZR-350"
	case Walton:
		return "Walton"
	case Regina:
		return "Regina"
	case Comet:
		return "Comet"
	case BMX:
		return "BMX"
	case Burrito:
		return "Burrito"
	case Camper:
		return "Camper"
	case Marquis:
		return "Marquis"
	case Baggage:
		return "Baggage"
	case Dozer:
		return "Dozer"
	case Maverick:
		return "Maverick"
	case SANNewsMaverick:
		return "SAN News Maverick"
	case Rancher:
		return "Rancher"
	case FBIRancher:
		return "FBI Rancher"
	case Virgo:
		return "Virgo"
	case Greenwood:
		return "Greenwood"
	case Jetmax:
		return "Jetmax"
	case HotringRacer:
		return "Hotring Racer"
	case Sandking:
		return "Sandking"
	case BlistaCompact:
		return "Blista Compact"
	case PoliceMaverick:
		return "Police Maverick"
	case Boxville:
		return "Boxville"
	case Benson:
		return "Benson"
	case Mesa:
		return "Mesa"
	case RCGoblin:
		return "RC Goblin"
	case HotringRacerA:
		return "Hotring Racer \"A\""
	case HotringRacerB:
		return "Hotring Racer \"B\""
	case BloodringBanger:
		return "Bloodring Banger"
	case RancherLure:
		return "Rancher Lure"
	case SuperGT:
		return "Super GT"
	case Elegant:
		return "Elegant"
	case Journey:
		return "Journey"
	case Bike:
		return "Bike"
	case MountainBike:
		return "Mountain Bike"
	case Beagle:
		return "Beagle"
	case Cropduster:
		return "Cropduster"
	case Stuntplane:
		return "Stuntplane"
	case Tanker:
		return "Tanker"
	case Roadtrain:
		return "Roadtrain"
	case Nebula:
		return "Nebula"
	case Majestic:
		return "Majestic"
	case Buccaneer:
		return "Buccaneer"
	case Shamal:
		return "Shamal"
	case Hydra:
		return "Hydra"
	case FCR900:
		return "FCR-900"
	case NRG500:
		return "NRG-500"
	case HPV1000:
		return "HPV1000"
	case CementTruck:
		return "Cement Truck"
	case Towtruck:
		return "Towtruck"
	case Fortune:
		return "Fortune"
	case Cadrona:
		return "Cadrona"
	case FBITruck:
		return "FBI Truck"
	case Willard:
		return "Willard"
	case Forklift:
		return "Forklift"
	case Tractor:
		return "Tractor"
	case CombineHarvester:
		return "Combine Harvester"
	case Feltzer:
		return "Feltzer"
	case Remington:
		return "Remington"
	case Slamvan:
		return "Slamvan"
	case Blade:
		return "Blade"
	case FreightTrain:
		return "Freight (Train)"
	case BrownstreakTrain:
		return "Brownstreak (Train)"
	case Vortex:
		return "Vortex"
	case Vincent:
		return "Vincent"
	case Bullet:
		return "Bullet"
	case Clover:
		return "Clover"
	case Sadler:
		return "Sadler"
	case FiretruckLA:
		return "Firetruck LA"
	case Hustler:
		return "Hustler"
	case Intruder:
		return "Intruder"
	case Primo:
		return "Primo"
	case Cargobob:
		return "Cargobob"
	case Tampa:
		return "Tampa"
	case Sunrise:
		return "Sunrise"
	case Merit:
		return "Merit"
	case UtilityVan:
		return "Utility Van"
	case Nevada:
		return "Nevada"
	case Yosemite:
		return "Yosemite"
	case Windsor:
		return "Windsor"
	case MonsterA:
		return "Monster \"A\""
	case MonsterB:
		return "Monster \"B\""
	case Uranus:
		return "Uranus"
	case Jester:
		return "Jester"
	case Sultan:
		return "Sultan"
	case Stratum:
		return "Stratum"
	case Elegy:
		return "Elegy"
	case Raindance:
		return "Raindance"
	case RCTiger:
		return "RC Tiger"
	case Flash:
		return "Flash"
	case Tahoma:
		return "Tahoma"
	case Savanna:
		return "Savanna"
	case Bandito:
		return "Bandito"
	case FreightFlatTrailerTrain:
		return "Freight Flat Trailer (Train)"
	case StreakTrailerTrain:
		return "Streak Trailer (Train)"
	case Kart:
		return "Kart"
	case Mower:
		return "Mower"
	case Dune:
		return "Dune"
	case Sweeper:
		return "Sweeper"
	case Broadway:
		return "Broadway"
	case Tornado:
		return "Tornado"
	case AT400:
		return "AT400"
	case DFT30:
		return "DFT-30"
	case Huntley:
		return "Huntley"
	case Stafford:
		return "Stafford"
	case BF400:
		return "BF-400"
	case Newsvan:
		return "Newsvan"
	case Tug:
		return "Tug"
	case PetrolTrailer:
		return "Petrol Trailer"
	case Emperor:
		return "Emperor"
	case Wayfarer:
		return "Wayfarer"
	case Euros:
		return "Euros"
	case Hotdog:
		return "Hotdog"
	case Club:
		return "Club"
	case FreightBoxTrailerTrain:
		return "Freight Box Trailer (Train)"
	case ArticleTrailer3:
		return "Article Trailer 3"
	case Andromada:
		return "Andromada"
	case Dodo:
		return "Dodo"
	case RCCam:
		return "RC Cam"
	case Launch:
		return "Launch"
	case PoliceCarLSPD:
		return "Police Car (LSPD)"
	case PoliceCarSFPD:
		return "Police Car (SFPD)"
	case PoliceCarLVPD:
		return "Police Car (LVPD)"
	case PoliceRanger:
		return "Police Ranger"
	case Picador:
		return "Picador"
	case SWAT:
		return "S.W.A.T."
	case Alpha:
		return "Alpha"
	case Phoenix:
		return "Phoenix"
	case GlendaleShit:
		return "Glendale Shit"
	case SadlerShit:
		return "Sadler Shit"
	case BaggageTrailerA:
		return "Baggage Trailer \"A\""
	case BaggageTrailerB:
		return "Baggage Trailer \"B\""
	case TugStairsTrailer:
		return "Tug Stairs Trailer"
	case Boxville2:
		return "Boxville"
	case FarmTrailer:
		return "Farm Trailer"
	case UtilityTrailer:
		return "Utility Trailer"
	default:
		return "invalid vehicle model"
	}
}
