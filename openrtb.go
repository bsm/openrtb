package openrtb

import "encoding/json"

// ContentCategory as defined in section 5.1
type ContentCategory string

// ContentCategory values.
const (
	ContentCategoryArtsEntertainment  = "IAB1"
	ContentCategoryBooksLiterature    = "IAB1-1"
	ContentCategoryCelebrityFanGossip = "IAB1-2"
	ContentCategoryFineArt            = "IAB1-3"
	ContentCategoryHumor              = "IAB1-4"
	ContentCategoryMovies             = "IAB1-5"
	ContentCategoryMusic              = "IAB1-6"
	ContentCategoryTelevision         = "IAB1-7"

	ContentCategoryAutomotive          = "IAB2"
	ContentCategoryAutoParts           = "IAB2-1"
	ContentCategoryAutoRepair          = "IAB2-2"
	ContentCategoryBuyingSellingCars   = "IAB2-3"
	ContentCategoryCarCulture          = "IAB2-4"
	ContentCategoryCertifiedPreOwned   = "IAB2-5"
	ContentCategoryConvertible         = "IAB2-6"
	ContentCategoryCoupe               = "IAB2-7"
	ContentCategoryCrossover           = "IAB2-8"
	ContentCategoryDiesel              = "IAB2-9"
	ContentCategoryElectricVehicle     = "IAB2-10"
	ContentCategoryHatchback           = "IAB2-11"
	ContentCategoryHybrid              = "IAB2-12"
	ContentCategoryLuxury              = "IAB2-13"
	ContentCategoryMiniVan             = "IAB2-14"
	ContentCategoryMororcycles         = "IAB2-15"
	ContentCategoryOffRoadVehicles     = "IAB2-16"
	ContentCategoryPerformanceVehicles = "IAB2-17"
	ContentCategoryPickup              = "IAB2-18"
	ContentCategoryRoadSideAssistance  = "IAB2-19"
	ContentCategorySedan               = "IAB2-20"
	ContentCategoryTrucksAccessories   = "IAB2-21"
	ContentCategoryVintageCars         = "IAB2-22"
	ContentCategoryWagon               = "IAB2-23"

	ContentCategoryBusiness          = "IAB3"
	ContentCategoryAdvertising       = "IAB3-1"
	ContentCategoryAgriculture       = "IAB3-2"
	ContentCategoryBiotechBiomedical = "IAB3-3"
	ContentCategoryBusinessSoftware  = "IAB3-4"
	ContentCategoryConstruction      = "IAB3-5"
	ContentCategoryForestry          = "IAB3-6"
	ContentCategoryGovernment        = "IAB3-7"
	ContentCategoryGreenSolutions    = "IAB3-8"
	ContentCategoryHumanResources    = "IAB3-9"
	ContentCategoryLogistics         = "IAB3-10"
	ContentCategoryMarketing         = "IAB3-11"
	ContentCategoryMetals            = "IAB3-12"

	ContentCategoryCareers             = "IAB4"
	ContentCategoryCareerPlanning      = "IAB4-1"
	ContentCategoryCollege             = "IAB4-2"
	ContentCategoryFinancialAid        = "IAB4-3"
	ContentCategoryJobFairs            = "IAB4-4"
	ContentCategoryJobSearch           = "IAB4-5"
	ContentCategoryResumeWritingAdvice = "IAB4-6"
	ContentCategoryNursing             = "IAB4-7"
	ContentCategoryScholarships        = "IAB4-8"
	ContentCategoryTelecommuting       = "IAB4-9"
	ContentCategoryUSMilitary          = "IAB4-10"
	ContentCategoryCareerAdvice        = "IAB4-11"

	ContentCategoryEducation             = "IAB5"
	ContentCategory12Education           = "IAB5-17"
	ContentCategoryAdultEducation        = "IAB5-2"
	ContentCategoryArtHistory            = "IAB5-3"
	ContentCategoryCollegeAdministration = "IAB5-4"
	ContentCategoryCollegeLife           = "IAB5-5"
	ContentCategoryDistanceLearning      = "IAB5-6"
	ContentCategoryEnglishasa2ndLanguage = "IAB5-7"
	ContentCategoryLanguageLearning      = "IAB5-8"
	ContentCategoryGraduateSchool        = "IAB5-9"
	ContentCategoryHomeschooling         = "IAB5-10"
	ContentCategoryHomeworkStudyTips     = "IAB5-11"
	ContentCategoryK6Educators           = "IAB5-12"
	ContentCategoryPrivateSchool         = "IAB5-13"
	ContentCategorySpecialEducation      = "IAB5-14"
	ContentCategoryStudyingBusiness      = "IAB5-15"

	ContentCategoryFamilyParenting  = "IAB6"
	ContentCategoryAdoption         = "IAB6-1"
	ContentCategoryBabiesToddlers   = "IAB6-2"
	ContentCategoryDaycarePreSchool = "IAB6-3"
	ContentCategoryFamilyInternet   = "IAB6-4"
	ContentCategoryParentingK6Kids  = "IAB6-5"
	ContentCategoryParentingteens   = "IAB6-6"
	ContentCategoryPregnancy        = "IAB6-7"
	ContentCategorySpecialNeedsKids = "IAB6-8"
	ContentCategoryEldercare        = "IAB6-9"

	ContentCategoryHealthFitness          = "IAB7"
	ContentCategoryExercise               = "IAB7-1"
	ContentCategoryADD                    = "IAB7-2"
	ContentCategoryAIDSHIV                = "IAB7-3"
	ContentCategoryAllergies              = "IAB7-4"
	ContentCategoryAlternativeMedicine    = "IAB7-5"
	ContentCategoryArthritis              = "IAB7-6"
	ContentCategoryAsthma                 = "IAB7-7"
	ContentCategoryAutismPDD              = "IAB7-8"
	ContentCategoryBipolarDisorder        = "IAB7-9"
	ContentCategoryBrainTumor             = "IAB7-10"
	ContentCategoryCancer                 = "IAB7-11"
	ContentCategoryCholesterol            = "IAB7-12"
	ContentCategoryChronicFatigueSyndrome = "IAB7-13"
	ContentCategoryChronicPain            = "IAB7-14"
	ContentCategoryColdFlu                = "IAB7-15"
	ContentCategoryDeafness               = "IAB7-16"
	ContentCategoryDentalCare             = "IAB7-17"
	ContentCategoryDepression             = "IAB7-18"
	ContentCategoryDermatology            = "IAB7-19"
	ContentCategoryDiabetes               = "IAB7-20"
	ContentCategoryEpilepsy               = "IAB7-21"
	ContentCategoryGERDAcidReflux         = "IAB7-22"
	ContentCategoryHeadachesMigraines     = "IAB7-23"
	ContentCategoryHeartDisease           = "IAB7-24"
	ContentCategoryHerbsforHealth         = "IAB7-25"
	ContentCategoryHolisticHealing        = "IAB7-26"
	ContentCategoryIBSCrohnsDisease       = "IAB7-27"
	ContentCategoryIncestAbuseSupport     = "IAB7-28"
	ContentCategoryIncontinence           = "IAB7-29"
	ContentCategoryInfertility            = "IAB7-30"
	ContentCategoryMensHealth             = "IAB7-31"
	ContentCategoryNutrition              = "IAB7-32"
	ContentCategoryOrthopedics            = "IAB7-33"
	ContentCategoryPanicAnxietyDisorders  = "IAB7-34"
	ContentCategoryPediatrics             = "IAB7-35"
	ContentCategoryPhysicalTherapy        = "IAB7-36"
	ContentCategoryPsychologyPsychiatry   = "IAB7-37"
	ContentCategorySeniorHealth           = "IAB7-38"
	ContentCategorySexuality              = "IAB7-39"
	ContentCategorySleepDisorders         = "IAB7-40"
	ContentCategorySmokingCessation       = "IAB7-41"
	ContentCategorySubstanceAbuse         = "IAB7-42"
	ContentCategoryThyroidDisease         = "IAB7-43"
	ContentCategoryWeightLoss             = "IAB7-44"
	ContentCategoryWomensHealth           = "IAB7-45"

	ContentCategoryFoodDrink           = "IAB8"
	ContentCategoryAmericanCuisine     = "IAB8-1"
	ContentCategoryBarbecuesGrilling   = "IAB8-2"
	ContentCategoryCajunCreole         = "IAB8-3"
	ContentCategoryChineseCuisine      = "IAB8-4"
	ContentCategoryCocktailsBeer       = "IAB8-5"
	ContentCategoryCoffeeTea           = "IAB8-6"
	ContentCategoryCuisineSpecific     = "IAB8-7"
	ContentCategoryDessertsBaking      = "IAB8-8"
	ContentCategoryDiningOut           = "IAB8-9"
	ContentCategoryFoodAllergies       = "IAB8-10"
	ContentCategoryFrenchCuisine       = "IAB8-11"
	ContentCategoryHealthLowfatCooking = "IAB8-12"
	ContentCategoryItalianCuisine      = "IAB8-13"
	ContentCategoryJapaneseCuisine     = "IAB8-14"
	ContentCategoryMexicanCuisine      = "IAB8-15"
	ContentCategoryVegan               = "IAB8-16"
	ContentCategoryVegetarian          = "IAB8-17"
	ContentCategoryWine                = "IAB8-18"

	ContentCategoryHobbiesInterests   = "IAB9"
	ContentCategoryArtTechnology      = "IAB9-1"
	ContentCategoryArtsCrafts         = "IAB9-2"
	ContentCategoryBeadwork           = "IAB9-3"
	ContentCategoryBirdwatching       = "IAB9-4"
	ContentCategoryBoardGamesPuzzles  = "IAB9-5"
	ContentCategoryCandleSoapMaking   = "IAB9-6"
	ContentCategoryCardGames          = "IAB9-7"
	ContentCategoryChess              = "IAB9-8"
	ContentCategoryCigars             = "IAB9-9"
	ContentCategoryCollecting         = "IAB9-10"
	ContentCategoryComicBooks         = "IAB9-11"
	ContentCategoryDrawingSketching   = "IAB9-12"
	ContentCategoryFreelanceWriting   = "IAB9-13"
	ContentCategoryGenealogy          = "IAB9-14"
	ContentCategoryGettingPublished   = "IAB9-15"
	ContentCategoryGuitar             = "IAB9-16"
	ContentCategoryHomeRecording      = "IAB9-17"
	ContentCategoryInvestorsPatents   = "IAB9-18"
	ContentCategoryJewelryMaking      = "IAB9-19"
	ContentCategoryMagicIllusion      = "IAB9-20"
	ContentCategoryNeedlework         = "IAB9-21"
	ContentCategoryPainting           = "IAB9-22"
	ContentCategoryPhotography        = "IAB9-23"
	ContentCategoryRadio              = "IAB9-24"
	ContentCategoryRoleplayingGames   = "IAB9-25"
	ContentCategorySciFiFantasy       = "IAB9-26"
	ContentCategoryScrapbooking       = "IAB9-27"
	ContentCategoryScreenwriting      = "IAB9-28"
	ContentCategoryStampsCoins        = "IAB9-29"
	ContentCategoryVideoComputerGames = "IAB9-30"
	ContentCategoryWoodworking        = "IAB9-31"

	ContentCategoryHomeGarden             = "IAB10"
	ContentCategoryAppliances             = "IAB10-1"
	ContentCategoryEntertaining           = "IAB10-2"
	ContentCategoryEnvironmentalSafety    = "IAB10-3"
	ContentCategoryGardening              = "IAB10-4"
	ContentCategoryHomeRepair             = "IAB10-5"
	ContentCategoryHomeTheater            = "IAB10-6"
	ContentCategoryInteriorDecorating     = "IAB10-7"
	ContentCategoryLandscaping            = "IAB10-8"
	ContentCategoryRemodelingConstruction = "IAB10-9"

	ContentCategoryLawGovtPolitics       = "IAB11"
	ContentCategoryImmigration           = "IAB11-1"
	ContentCategoryLegalIssues           = "IAB11-2"
	ContentCategoryUSGovernmentResources = "IAB11-3"
	ContentCategoryPolitics              = "IAB11-4"
	ContentCategoryCommentary            = "IAB11-5"

	ContentCategoryNews              = "IAB12"
	ContentCategoryInternationalNews = "IAB12-1"
	ContentCategoryNationalNews      = "IAB12-2"
	ContentCategoryLocalNews         = "IAB12-3"

	ContentCategoryPersonalFinance    = "IAB13"
	ContentCategoryBeginningInvesting = "IAB13-1"
	ContentCategoryCreditDebtLoans    = "IAB13-2"
	ContentCategoryFinancialNews      = "IAB13-3"
	ContentCategoryFinancialPlanning  = "IAB13-4"
	ContentCategoryHedgeFund          = "IAB13-5"
	ContentCategoryInsurance          = "IAB13-6"
	ContentCategoryInvesting          = "IAB13-7"
	ContentCategoryMutualFunds        = "IAB13-8"
	ContentCategoryOptions            = "IAB13-9"
	ContentCategoryRetirementPlanning = "IAB13-10"
	ContentCategoryStocks             = "IAB13-11"
	ContentCategoryTaxPlanning        = "IAB13-12"

	ContentCategorySociety        = "IAB14"
	ContentCategoryDating         = "IAB14-1"
	ContentCategoryDivorceSupport = "IAB14-2"
	ContentCategoryGayLife        = "IAB14-3"
	ContentCategoryMarriage       = "IAB14-4"
	ContentCategorySeniorLiving   = "IAB14-5"
	ContentCategoryTeens          = "IAB14-6"
	ContentCategoryWeddings       = "IAB14-7"
	ContentCategoryEthnicSpecific = "IAB14-8"

	ContentCategoryScience             = "IAB15"
	ContentCategoryAstrology           = "IAB15-1"
	ContentCategoryBiology             = "IAB15-2"
	ContentCategoryChemistry           = "IAB15-3"
	ContentCategoryGeology             = "IAB15-4"
	ContentCategoryParanormalPhenomena = "IAB15-5"
	ContentCategoryPhysics             = "IAB15-6"
	ContentCategorySpaceAstronomy      = "IAB15-7"
	ContentCategoryGeography           = "IAB15-8"
	ContentCategoryBotany              = "IAB15-9"
	ContentCategoryWeather             = "IAB15-10"

	ContentCategoryPets               = "IAB16"
	ContentCategoryAquariums          = "IAB16-1"
	ContentCategoryBirds              = "IAB16-2"
	ContentCategoryCats               = "IAB16-3"
	ContentCategoryDogs               = "IAB16-4"
	ContentCategoryLargeAnimals       = "IAB16-5"
	ContentCategoryReptiles           = "IAB16-6"
	ContentCategoryVeterinaryMedicine = "IAB16-7"

	ContentCategorySports              = "IAB17"
	ContentCategoryAutoRacing          = "IAB17-1"
	ContentCategoryBaseball            = "IAB17-2"
	ContentCategoryBicycling           = "IAB17-3"
	ContentCategoryBodybuilding        = "IAB17-4"
	ContentCategoryBoxing              = "IAB17-5"
	ContentCategoryCanoeingKayaking    = "IAB17-6"
	ContentCategoryCheerleading        = "IAB17-7"
	ContentCategoryClimbing            = "IAB17-8"
	ContentCategoryCricket             = "IAB17-9"
	ContentCategoryFigureSkating       = "IAB17-10"
	ContentCategoryFlyFishing          = "IAB17-11"
	ContentCategoryFootball            = "IAB17-12"
	ContentCategoryFreshwaterFishing   = "IAB17-13"
	ContentCategoryGameFish            = "IAB17-14"
	ContentCategoryGolf                = "IAB17-15"
	ContentCategoryHorseRacing         = "IAB17-16"
	ContentCategoryHorses              = "IAB17-17"
	ContentCategoryHuntingShooting     = "IAB17-18"
	ContentCategoryInlineSkating       = "IAB17-19"
	ContentCategoryMartialArts         = "IAB17-20"
	ContentCategoryMountainBiking      = "IAB17-21"
	ContentCategoryNASCARRacing        = "IAB17-22"
	ContentCategoryOlympics            = "IAB17-23"
	ContentCategoryPaintball           = "IAB17-24"
	ContentCategoryPowerMotorcycles    = "IAB17-25"
	ContentCategoryProBasketball       = "IAB17-26"
	ContentCategoryProIceHockey        = "IAB17-27"
	ContentCategoryRodeo               = "IAB17-28"
	ContentCategoryRugby               = "IAB17-29"
	ContentCategoryRunningJogging      = "IAB17-30"
	ContentCategorySailing             = "IAB17-31"
	ContentCategorySaltwaterFishing    = "IAB17-32"
	ContentCategoryScubaDiving         = "IAB17-33"
	ContentCategorySkateboarding       = "IAB17-34"
	ContentCategorySkiing              = "IAB17-35"
	ContentCategorySnowboarding        = "IAB17-36"
	ContentCategorySurfingBodyboarding = "IAB17-37"
	ContentCategorySwimming            = "IAB17-38"
	ContentCategoryTableTennisPingPong = "IAB17-39"
	ContentCategoryTennis              = "IAB17-40"
	ContentCategoryVolleyball          = "IAB17-41"
	ContentCategoryWalking             = "IAB17-42"
	ContentCategoryWaterskiWakeboard   = "IAB17-43"
	ContentCategoryWorldSoccer         = "IAB17-44"

	ContentCategoryStyleFashion = "IAB18"
	ContentCategoryBeauty       = "IAB18-1"
	ContentCategoryBodyArt      = "IAB18-2"
	ContentCategoryFashion      = "IAB18-3"
	ContentCategoryJewelry      = "IAB18-4"
	ContentCategoryClothing     = "IAB18-5"
	ContentCategoryAccessories  = "IAB18-6"

	ContentCategoryTechnologyComputing   = "IAB19"
	ContentCategoryDGraphics             = "IAB19-13"
	ContentCategoryAnimation             = "IAB19-2"
	ContentCategoryAntivirusSoftware     = "IAB19-3"
	ContentCategoryCC                    = "IAB19-4"
	ContentCategoryCamerasCamcorders     = "IAB19-5"
	ContentCategoryCellPhones            = "IAB19-6"
	ContentCategoryComputerCertification = "IAB19-7"
	ContentCategoryComputerNetworking    = "IAB19-8"
	ContentCategoryComputerPeripherals   = "IAB19-9"
	ContentCategoryComputerReviews       = "IAB19-10"
	ContentCategoryDataCenters           = "IAB19-11"
	ContentCategoryDatabases             = "IAB19-12"
	ContentCategoryDesktopPublishing     = "IAB19-13"
	ContentCategoryDesktopVideo          = "IAB19-14"
	ContentCategoryEmail                 = "IAB19-15"
	ContentCategoryGraphicsSoftware      = "IAB19-16"
	ContentCategoryHomeVideoDVD          = "IAB19-17"
	ContentCategoryInternetTechnology    = "IAB19-18"
	ContentCategoryJava                  = "IAB19-19"
	ContentCategoryJavaScript            = "IAB19-20"
	ContentCategoryMacSupport            = "IAB19-21"
	ContentCategoryMP3MIDI               = "IAB19-22"
	ContentCategoryNetConferencing       = "IAB19-23"
	ContentCategoryNetforBeginners       = "IAB19-24"
	ContentCategoryNetworkSecurity       = "IAB19-25"
	ContentCategoryPalmtopsPDAs          = "IAB19-26"
	ContentCategoryPCSupport             = "IAB19-27"
	ContentCategoryPortable              = "IAB19-28"
	ContentCategoryEntertainment         = "IAB19-29"
	ContentCategorySharewareFreeware     = "IAB19-30"
	ContentCategoryUnix                  = "IAB19-31"
	ContentCategoryVisualBasic           = "IAB19-32"
	ContentCategoryWebClipArt            = "IAB19-33"
	ContentCategoryWebDesignHTML         = "IAB19-34"
	ContentCategoryWebSearch             = "IAB19-35"
	ContentCategoryWindows               = "IAB19-36"

	ContentCategoryTravel               = "IAB20"
	ContentCategoryAdventureTravel      = "IAB20-1"
	ContentCategoryAfrica               = "IAB20-2"
	ContentCategoryAirTravel            = "IAB20-3"
	ContentCategoryAustraliaNewZealand  = "IAB20-4"
	ContentCategoryBedBreakfasts        = "IAB20-5"
	ContentCategoryBudgetTravel         = "IAB20-6"
	ContentCategoryBusinessTravel       = "IAB20-7"
	ContentCategoryByUSLocale           = "IAB20-8"
	ContentCategoryCamping              = "IAB20-9"
	ContentCategoryCanada               = "IAB20-10"
	ContentCategoryCaribbean            = "IAB20-11"
	ContentCategoryCruises              = "IAB20-12"
	ContentCategoryEasternEurope        = "IAB20-13"
	ContentCategoryEurope               = "IAB20-14"
	ContentCategoryFrance               = "IAB20-15"
	ContentCategoryGreece               = "IAB20-16"
	ContentCategoryHoneymoonsGetaways   = "IAB20-17"
	ContentCategoryHotels               = "IAB20-18"
	ContentCategoryItaly                = "IAB20-19"
	ContentCategoryJapan                = "IAB20-20"
	ContentCategoryMexicoCentralAmerica = "IAB20-21"
	ContentCategoryNationalParks        = "IAB20-22"
	ContentCategorySouthAmerica         = "IAB20-23"
	ContentCategorySpas                 = "IAB20-24"
	ContentCategoryThemeParks           = "IAB20-25"
	ContentCategoryTravelingwithKids    = "IAB20-26"
	ContentCategoryUnitedKingdom        = "IAB20-27"

	ContentCategoryRealEstate         = "IAB21"
	ContentCategoryApartments         = "IAB21-1"
	ContentCategoryArchitects         = "IAB21-2"
	ContentCategoryBuyingSellingHomes = "IAB21-3"

	ContentCategoryShopping         = "IAB22"
	ContentCategoryContestsFreebies = "IAB22-1"
	ContentCategoryCouponing        = "IAB22-2"
	ContentCategoryComparison       = "IAB22-3"
	ContentCategoryEngines          = "IAB22-4"

	ContentCategoryReligionSpirituality = "IAB23"
	ContentCategoryAlternativeReligions = "IAB23-1"
	ContentCategoryAtheismAgnosticism   = "IAB23-2"
	ContentCategoryBuddhism             = "IAB23-3"
	ContentCategoryCatholicism          = "IAB23-4"
	ContentCategoryChristianity         = "IAB23-5"
	ContentCategoryHinduism             = "IAB23-6"
	ContentCategoryIslam                = "IAB23-7"
	ContentCategoryJudaism              = "IAB23-8"
	ContentCategoryLatterDaySaints      = "IAB23-9"
	ContentCategoryPaganWiccan          = "IAB23-10"

	ContentCategoryUncategorized = "IAB24"

	ContentCategoryNonStandardContent             = "IAB25"
	ContentCategoryUnmoderatedUGC                 = "IAB25-1"
	ContentCategoryExtremeGraphicExplicitViolence = "IAB25-2"
	ContentCategoryPornography                    = "IAB25-3"
	ContentCategoryProfaneContent                 = "IAB25-4"
	ContentCategoryHateContent                    = "IAB25-5"
	ContentCategoryUnderConstruction              = "IAB25-6"
	ContentCategoryIncentivized                   = "IAB25-7"

	ContentCategoryAnyIllegalContent     = "IAB26"
	ContentCategoryIllegalContent        = "IAB26-1"
	ContentCategoryWarez                 = "IAB26-2"
	ContentCategorySpywareMalware        = "IAB26-3"
	ContentCategoryCopyrightInfringement = "IAB26-4"
)

// BannerType as defined in section 5.2.
type BannerType int

// 5.2 Banner Ad Types
const (
	BannerTypeXHTMLText BannerType = 1
	BannerTypeXHTML     BannerType = 2
	BannerTypeJS        BannerType = 3
	BannerTypeFrame     BannerType = 4
)

// CreativeAttribute as defined in section 5.3.
type CreativeAttribute int

// 5.3 Creative Attributes
const (
	CreativeAttributeAudioAdAutoPlay                 CreativeAttribute = 1
	CreativeAttributeAudioAdUserInitiated            CreativeAttribute = 2
	CreativeAttributeExpandableAuto                  CreativeAttribute = 3
	CreativeAttributeExpandableUserInitiatedClick    CreativeAttribute = 4
	CreativeAttributeExpandableUserInitiatedRollover CreativeAttribute = 5
	CreativeAttributeInBannerVideoAdAutoPlay         CreativeAttribute = 6
	CreativeAttributeInBannerVideoAdUserInitiated    CreativeAttribute = 7
	CreativeAttributePop                             CreativeAttribute = 8
	CreativeAttributeProvocativeOrSuggestiveImagery  CreativeAttribute = 9
	CreativeAttributeExtremeAnimation                CreativeAttribute = 10
	CreativeAttributeSurveys                         CreativeAttribute = 11
	CreativeAttributeTextOnly                        CreativeAttribute = 12
	CreativeAttributeUserInitiated                   CreativeAttribute = 13
	CreativeAttributeWindowsDialogOrAlert            CreativeAttribute = 14
	CreativeAttributeHasAudioWithPlayer              CreativeAttribute = 15
	CreativeAttributeAdProvidesSkipButton            CreativeAttribute = 16
	CreativeAttributeAdobeFlash                      CreativeAttribute = 17
)

// AdPosition as defined in section 5.4.
type AdPosition int

// 5.4 Ad Position
const (
	AdPositionUnknown    AdPosition = 0
	AdPositionAboveFold  AdPosition = 1
	AdPositionBelowFold  AdPosition = 3
	AdPositionHeader     AdPosition = 4
	AdPositionFooter     AdPosition = 5
	AdPositionSidebar    AdPosition = 6
	AdPositionFullscreen AdPosition = 7
)

// ExpDir as defined in section 5.5.
type ExpDir int

// 5.5 Expandable Direction
const (
	ExpDirUnknown    ExpDir = 0
	ExpDirLeft       ExpDir = 1
	ExpDirRight      ExpDir = 2
	ExpDirUp         ExpDir = 3
	ExpDirDown       ExpDir = 4
	ExpDirFullScreen ExpDir = 5
)

// APIFramework as defined in section 5.6.
type APIFramework int

// 5.6 API Frameworks
const (
	APIFrameworkUnknown APIFramework = 0
	APIFrameworkVPAID1  APIFramework = 1
	APIFrameworkVPAID2  APIFramework = 2
	APIFrameworkMRAID1  APIFramework = 3
	APIFrameworkORMMA   APIFramework = 4
	APIFrameworkMRAID2  APIFramework = 5
)

// VideoLinearity as defined in section 5.7.
type VideoLinearity int

// 5.7 Video Linearity
const (
	VideoLinearityUnknown   VideoLinearity = 0
	VideoLinearityLinear    VideoLinearity = 1
	VideoLinearityNonLinear VideoLinearity = 2
)

// Protocol as defined in section 5.8.
type Protocol int

// 5.8 Video and Audio Bid Response Protocols
const (
	ProtocolUnknown       Protocol = 0
	ProtocolVAST1         Protocol = 1
	ProtocolVAST2         Protocol = 2
	ProtocolVAST3         Protocol = 3
	ProtocolVAST1Wrapper  Protocol = 4
	ProtocolVAST2Wrapper  Protocol = 5
	ProtocolVAST3Wrapper  Protocol = 6
	ProtocolVAST4         Protocol = 7
	ProtocolVAST4Wrapper  Protocol = 8
	ProtocolDAAST1        Protocol = 9
	ProtocolDAAST1Wrapper Protocol = 10
)

// VideoPlacement as defined in section 5.9.
type VideoPlacement int

// Video Placement Types
const (
	VideoPlacementUnknown      VideoPlacement = 0
	VideoPlacementInStream     VideoPlacement = 1
	VideoPlacementInBanner     VideoPlacement = 2
	VideoPlacementInArticle    VideoPlacement = 3
	VideoPlacementInFeed       VideoPlacement = 4
	VideoPlacementInterstitial VideoPlacement = 5
)

// VideoPlayback as defined in section 5.10.
type VideoPlayback int

// 5.10 Video Playback Methods
const (
	VideoPlaybackUnknown          VideoPlayback = 0
	VideoPlaybackPageLoadSoundOn  VideoPlayback = 1
	VideoPlaybackPageLoadSoundOff VideoPlayback = 2
	VideoPlaybackClickToPlay      VideoPlayback = 3
	VideoPlaybackMouseOver        VideoPlayback = 4
	VideoPlaybackEnterSoundOn     VideoPlayback = 5
	VideoPlaybackEnterSoundOff    VideoPlayback = 6
)

// StartDelay as defined in section 5.12.
type StartDelay int

// 5.12 Video Start Delay
const (
	StartDelayPreRoll         StartDelay = 0
	StartDelayGenericMidRoll  StartDelay = -1
	StartDelayGenericPostRoll StartDelay = -2
)

// ProductionQuality as defined in section 5.13.
type ProductionQuality int

// 5.13 Video Quality
const (
	ProductionQualityUnknown      ProductionQuality = 0
	ProductionQualityProfessional ProductionQuality = 1
	ProductionQualityProsumer     ProductionQuality = 2
	ProductionQualityUGC          ProductionQuality = 3
)

// CompanionType as defined in section 5.14.
type CompanionType int

// 5.14  Companion Types
const (
	CompanionTypeUnknown CompanionType = 0
	CompanionTypeStatic  CompanionType = 1
	CompanionTypeHTML    CompanionType = 2
	CompanionTypeIFrame  CompanionType = 3
)

// ContentDelivery as defined in section 5.15.
type ContentDelivery int

// 5.15 Content Delivery Methods
const (
	ContentDeliveryUnknown     ContentDelivery = 0
	ContentDeliveryStreaming   ContentDelivery = 1
	ContentDeliveryProgressive ContentDelivery = 2
	ContentDeliveryDownload    ContentDelivery = 3
)

// FeedType as defined in section 5.16.
type FeedType int

// 5.16 Feed Types
const (
	FeedTypeUnknown   FeedType = 0
	FeedTypeMusic     FeedType = 1
	FeedTypeBroadcast FeedType = 2
	FeedTypePodcast   FeedType = 3
)

// VolumeNorm as defined in section 5.17.
type VolumeNorm int

// 5.17 Volume Normalization Modes
const (
	VolumeNormNone     VolumeNorm = 0
	VolumeNormAverage  VolumeNorm = 1
	VolumeNormPeak     VolumeNorm = 2
	VolumeNormLoudness VolumeNorm = 3
	VolumeNormCustom   VolumeNorm = 4
)

// ContentContext as defined in section 5.18.
type ContentContext int

// UnmarshalJSON implements json.Unmarshaler
func (n *ContentContext) UnmarshalJSON(data []byte) (err error) {
	var v int

	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		err = json.Unmarshal(data[1:len(data)-1], &v)
	} else {
		err = json.Unmarshal(data, &v)
	}
	if err != nil {
		return err
	}

	*n = ContentContext(v)
	return nil
}

// 5.18 Content Context
const (
	ContentContextVideo       ContentContext = 1
	ContentContextGame        ContentContext = 2
	ContentContextMusic       ContentContext = 3
	ContentContextApplication ContentContext = 4
	ContentContextText        ContentContext = 5
	ContentContextOther       ContentContext = 6
	ContentContextUnknown     ContentContext = 7
)

// IQGRating as defined in section 5.19.
type IQGRating int

// 5.19 IQG Media Ratings
const (
	IQGRatingUnknown IQGRating = 0
	IQGRatingAll     IQGRating = 1
	IQGRatingOver12  IQGRating = 2
	IQGRatingMature  IQGRating = 3
)

// LocationType as defined in section 5.20.
type LocationType int

// 5.20 Location Type
const (
	LocationTypeUnknown LocationType = 0
	LocationTypeGPS     LocationType = 1
	LocationTypeIP      LocationType = 2
	LocationTypeUser    LocationType = 3
)

// DeviceType as defined in section 5.21.
type DeviceType int

// 5.21 Device Type
const (
	DeviceTypeUnknown   DeviceType = 0
	DeviceTypeMobile    DeviceType = 1
	DeviceTypePC        DeviceType = 2
	DeviceTypeTV        DeviceType = 3
	DeviceTypePhone     DeviceType = 4
	DeviceTypeTablet    DeviceType = 5
	DeviceTypeConnected DeviceType = 6
	DeviceTypeSetTopBox DeviceType = 7
)

// ConnType as defined in section 5.22.
type ConnType int

// 5.22 Connection Type
const (
	ConnTypeUnknown  ConnType = 0
	ConnTypeEthernet ConnType = 1
	ConnTypeWIFI     ConnType = 2
	ConnTypeCell     ConnType = 3
	ConnTypeCell2G   ConnType = 4
	ConnTypeCell3G   ConnType = 5
	ConnTypeCell4G   ConnType = 6
)

// IPLocation as defined in section 5.23.
type IPLocation int

// 5.22 IP Location Services
const (
	IPLocationUnknown     IPLocation = 0
	IPLocationIP2Location IPLocation = 1
	IPLocationNeustar     IPLocation = 2
	IPLocationMaxMind     IPLocation = 3
	IPLocationNetAquity   IPLocation = 4
)

// NBR as defined in section 5.24.
type NBR int

// 5.24 No-Bid Reason Codes
const (
	NBRUnknownError      NBR = 0
	NBRTechnicalError    NBR = 1
	NBRInvalidRequest    NBR = 2
	NBRKnownSpider       NBR = 3
	NBRSuspectedNonHuman NBR = 4
	NBRProxyIP           NBR = 5
	NBRUnsupportedDevice NBR = 6
	NBRBlockedSite       NBR = 7
	NBRUnmatchedUser     NBR = 8
)

/*************************************************************************
 * COMMON OBJECT STRUCTS
 *************************************************************************/

// This object may be useful in the situation where syndicated content contains impressions and
// does not necessarily match the publisher's general content.  The exchange might or might not
// have knowledge of the page where the content is running, as a result of the syndication
// method.  (For example, video impressions embedded in an iframe on an unknown web property
// or device.)
// type Content struct {
// }

// ThirdParty abstract attributes.
type ThirdParty struct {
	ID         string            `json:"id,omitempty"`
	Name       string            `json:"name,omitempty"`
	Categories []ContentCategory `json:"cat,omitempty"` // Array of IAB content categories
	Domain     string            `json:"domain,omitempty"`
	Ext        json.RawMessage   `json:"ext,omitempty"`
}

// Publisher object itself and all of its parameters are optional, so default values are not
// provided. If an optional parameter is not specified, it should be considered unknown.
type Publisher ThirdParty

// Producer is useful when content where the ad is shown is syndicated, and may appear on a
// completely different publisher. The producer object itself and all of its parameters are optional,
// so default values are not provided. If an optional parameter is not specified, it should be
// considered unknown.
type Producer ThirdParty

// Geo object may appear in one or both the Device Object and the User Object.
// This is intentional, since the information may be derived from either a device-oriented source
// (such as IP geo lookup), or by user registration information (for example provided to a publisher
// through a user registration).
type Geo struct {
	Latitude      float64         `json:"lat,omitempty"`           // Latitude from -90 to 90
	Longitude     float64         `json:"lon,omitempty"`           // Longitude from -180 to 180
	Type          LocationType    `json:"type,omitempty"`          // Indicate the source of the geo data
	Accuracy      int             `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
	LastFix       int             `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     IPLocation      `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable
	Country       string          `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 3
	Region        string          `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string          `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string          `json:"metro,omitempty"`
	City          string          `json:"city,omitempty"`
	ZIP           string          `json:"zip,omitempty"`
	UTCOffset     int             `json:"utcoffset,omitempty"` // Local time as the number +/- of minutes from UTC
	Ext           json.RawMessage `json:"ext,omitempty"`
}

// User object contains information known or derived about the human user of the device (i.e., the
// audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
// privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
// frequency capping and retargeting.
type User struct {
	ID          string          `json:"id,omitempty"`         // Unique consumer ID of this user on the exchange
	BuyerID     string          `json:"buyerid,omitempty"`    // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid/buyerid or id is recommended. Valid for OpenRTB 2.3.
	BuyerUID    string          `json:"buyeruid,omitempty"`   // Buyer-specific ID for the user as mapped by the exchange for the buyer. Same as BuyerID but valid for OpenRTB 2.2.
	YearOfBirth int             `json:"yob,omitempty"`        // Year of birth as a 4-digit integer.
	Gender      string          `json:"gender,omitempty"`     // Gender ("M": male, "F" female, "O" Other)
	Keywords    string          `json:"keywords,omitempty"`   // Comma separated list of keywords, interests, or intent
	CustomData  string          `json:"customdata,omitempty"` // Optional feature to pass bidder data that was set in the exchange's cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include "escaped" quotation marks.
	Geo         *Geo            `json:"geo,omitempty"`
	Data        []Data          `json:"data,omitempty"`
	Ext         json.RawMessage `json:"ext,omitempty"`
}

// Data and segment objects together allow additional data about the user to be specified. This data
// may be from multiple sources whether from the exchange itself or third party providers as specified by
// the id field. A bid request can mix data objects from multiple providers. The specific data providers in
// use should be published by the exchange a priori to its bidders.
type Data struct {
	ID      string          `json:"id,omitempty"`
	Name    string          `json:"name,omitempty"`
	Segment []Segment       `json:"segment,omitempty"`
	Ext     json.RawMessage `json:"ext,omitempty"`
}

// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name,omitempty"`
	Value string          `json:"value,omitempty"`
	Ext   json.RawMessage `json:"ext,omitempty"`
}

// Regulations object contains any legal, governmental, or industry regulations that apply to the request. The
// coppa flag signals whether or not the request falls under the United States Federal Trade Commission's
// regulations for the United States Children's Online Privacy Protection Act ("COPPA").
type Regulations struct {
	COPPA int             `json:"coppa,omitempty"` // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	Ext   json.RawMessage `json:"ext,omitempty"`
}

// Format object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
type Format struct {
	Width  int             `json:"w,omitempty"` // Width in device independent pixels (DIPS).
	Height int             `json:"h,omitempty"` //Height in device independent pixels (DIPS).
	Ext    json.RawMessage `json:"ext,omitempty"`
}
