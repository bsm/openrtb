package openrtb

import "encoding/json"

// ContentCategory as defined in section 5.1
type ContentCategory string

// ContentCategory values.
const (
	ContentCategoryArtsEntertainment  ContentCategory = "IAB1"
	ContentCategoryBooksLiterature    ContentCategory = "IAB1-1"
	ContentCategoryCelebrityFanGossip ContentCategory = "IAB1-2"
	ContentCategoryFineArt            ContentCategory = "IAB1-3"
	ContentCategoryHumor              ContentCategory = "IAB1-4"
	ContentCategoryMovies             ContentCategory = "IAB1-5"
	ContentCategoryMusic              ContentCategory = "IAB1-6"
	ContentCategoryTelevision         ContentCategory = "IAB1-7"

	ContentCategoryAutomotive          ContentCategory = "IAB2"
	ContentCategoryAutoParts           ContentCategory = "IAB2-1"
	ContentCategoryAutoRepair          ContentCategory = "IAB2-2"
	ContentCategoryBuyingSellingCars   ContentCategory = "IAB2-3"
	ContentCategoryCarCulture          ContentCategory = "IAB2-4"
	ContentCategoryCertifiedPreOwned   ContentCategory = "IAB2-5"
	ContentCategoryConvertible         ContentCategory = "IAB2-6"
	ContentCategoryCoupe               ContentCategory = "IAB2-7"
	ContentCategoryCrossover           ContentCategory = "IAB2-8"
	ContentCategoryDiesel              ContentCategory = "IAB2-9"
	ContentCategoryElectricVehicle     ContentCategory = "IAB2-10"
	ContentCategoryHatchback           ContentCategory = "IAB2-11"
	ContentCategoryHybrid              ContentCategory = "IAB2-12"
	ContentCategoryLuxury              ContentCategory = "IAB2-13"
	ContentCategoryMiniVan             ContentCategory = "IAB2-14"
	ContentCategoryMororcycles         ContentCategory = "IAB2-15"
	ContentCategoryOffRoadVehicles     ContentCategory = "IAB2-16"
	ContentCategoryPerformanceVehicles ContentCategory = "IAB2-17"
	ContentCategoryPickup              ContentCategory = "IAB2-18"
	ContentCategoryRoadSideAssistance  ContentCategory = "IAB2-19"
	ContentCategorySedan               ContentCategory = "IAB2-20"
	ContentCategoryTrucksAccessories   ContentCategory = "IAB2-21"
	ContentCategoryVintageCars         ContentCategory = "IAB2-22"
	ContentCategoryWagon               ContentCategory = "IAB2-23"

	ContentCategoryBusiness          ContentCategory = "IAB3"
	ContentCategoryAdvertising       ContentCategory = "IAB3-1"
	ContentCategoryAgriculture       ContentCategory = "IAB3-2"
	ContentCategoryBiotechBiomedical ContentCategory = "IAB3-3"
	ContentCategoryBusinessSoftware  ContentCategory = "IAB3-4"
	ContentCategoryConstruction      ContentCategory = "IAB3-5"
	ContentCategoryForestry          ContentCategory = "IAB3-6"
	ContentCategoryGovernment        ContentCategory = "IAB3-7"
	ContentCategoryGreenSolutions    ContentCategory = "IAB3-8"
	ContentCategoryHumanResources    ContentCategory = "IAB3-9"
	ContentCategoryLogistics         ContentCategory = "IAB3-10"
	ContentCategoryMarketing         ContentCategory = "IAB3-11"
	ContentCategoryMetals            ContentCategory = "IAB3-12"

	ContentCategoryCareers             ContentCategory = "IAB4"
	ContentCategoryCareerPlanning      ContentCategory = "IAB4-1"
	ContentCategoryCollege             ContentCategory = "IAB4-2"
	ContentCategoryFinancialAid        ContentCategory = "IAB4-3"
	ContentCategoryJobFairs            ContentCategory = "IAB4-4"
	ContentCategoryJobSearch           ContentCategory = "IAB4-5"
	ContentCategoryResumeWritingAdvice ContentCategory = "IAB4-6"
	ContentCategoryNursing             ContentCategory = "IAB4-7"
	ContentCategoryScholarships        ContentCategory = "IAB4-8"
	ContentCategoryTelecommuting       ContentCategory = "IAB4-9"
	ContentCategoryUSMilitary          ContentCategory = "IAB4-10"
	ContentCategoryCareerAdvice        ContentCategory = "IAB4-11"

	ContentCategoryEducation             ContentCategory = "IAB5"
	ContentCategory12Education           ContentCategory = "IAB5-1"
	ContentCategoryAdultEducation        ContentCategory = "IAB5-2"
	ContentCategoryArtHistory            ContentCategory = "IAB5-3"
	ContentCategoryCollegeAdministration ContentCategory = "IAB5-4"
	ContentCategoryCollegeLife           ContentCategory = "IAB5-5"
	ContentCategoryDistanceLearning      ContentCategory = "IAB5-6"
	ContentCategoryEnglishasa2ndLanguage ContentCategory = "IAB5-7"
	ContentCategoryLanguageLearning      ContentCategory = "IAB5-8"
	ContentCategoryGraduateSchool        ContentCategory = "IAB5-9"
	ContentCategoryHomeschooling         ContentCategory = "IAB5-10"
	ContentCategoryHomeworkStudyTips     ContentCategory = "IAB5-11"
	ContentCategoryK6Educators           ContentCategory = "IAB5-12"
	ContentCategoryPrivateSchool         ContentCategory = "IAB5-13"
	ContentCategorySpecialEducation      ContentCategory = "IAB5-14"
	ContentCategoryStudyingBusiness      ContentCategory = "IAB5-15"

	ContentCategoryFamilyParenting  ContentCategory = "IAB6"
	ContentCategoryAdoption         ContentCategory = "IAB6-1"
	ContentCategoryBabiesToddlers   ContentCategory = "IAB6-2"
	ContentCategoryDaycarePreSchool ContentCategory = "IAB6-3"
	ContentCategoryFamilyInternet   ContentCategory = "IAB6-4"
	ContentCategoryParentingK6Kids  ContentCategory = "IAB6-5"
	ContentCategoryParentingteens   ContentCategory = "IAB6-6"
	ContentCategoryPregnancy        ContentCategory = "IAB6-7"
	ContentCategorySpecialNeedsKids ContentCategory = "IAB6-8"
	ContentCategoryEldercare        ContentCategory = "IAB6-9"

	ContentCategoryHealthFitness          ContentCategory = "IAB7"
	ContentCategoryExercise               ContentCategory = "IAB7-1"
	ContentCategoryADD                    ContentCategory = "IAB7-2"
	ContentCategoryAIDSHIV                ContentCategory = "IAB7-3"
	ContentCategoryAllergies              ContentCategory = "IAB7-4"
	ContentCategoryAlternativeMedicine    ContentCategory = "IAB7-5"
	ContentCategoryArthritis              ContentCategory = "IAB7-6"
	ContentCategoryAsthma                 ContentCategory = "IAB7-7"
	ContentCategoryAutismPDD              ContentCategory = "IAB7-8"
	ContentCategoryBipolarDisorder        ContentCategory = "IAB7-9"
	ContentCategoryBrainTumor             ContentCategory = "IAB7-10"
	ContentCategoryCancer                 ContentCategory = "IAB7-11"
	ContentCategoryCholesterol            ContentCategory = "IAB7-12"
	ContentCategoryChronicFatigueSyndrome ContentCategory = "IAB7-13"
	ContentCategoryChronicPain            ContentCategory = "IAB7-14"
	ContentCategoryColdFlu                ContentCategory = "IAB7-15"
	ContentCategoryDeafness               ContentCategory = "IAB7-16"
	ContentCategoryDentalCare             ContentCategory = "IAB7-17"
	ContentCategoryDepression             ContentCategory = "IAB7-18"
	ContentCategoryDermatology            ContentCategory = "IAB7-19"
	ContentCategoryDiabetes               ContentCategory = "IAB7-20"
	ContentCategoryEpilepsy               ContentCategory = "IAB7-21"
	ContentCategoryGERDAcidReflux         ContentCategory = "IAB7-22"
	ContentCategoryHeadachesMigraines     ContentCategory = "IAB7-23"
	ContentCategoryHeartDisease           ContentCategory = "IAB7-24"
	ContentCategoryHerbsforHealth         ContentCategory = "IAB7-25"
	ContentCategoryHolisticHealing        ContentCategory = "IAB7-26"
	ContentCategoryIBSCrohnsDisease       ContentCategory = "IAB7-27"
	ContentCategoryIncestAbuseSupport     ContentCategory = "IAB7-28"
	ContentCategoryIncontinence           ContentCategory = "IAB7-29"
	ContentCategoryInfertility            ContentCategory = "IAB7-30"
	ContentCategoryMensHealth             ContentCategory = "IAB7-31"
	ContentCategoryNutrition              ContentCategory = "IAB7-32"
	ContentCategoryOrthopedics            ContentCategory = "IAB7-33"
	ContentCategoryPanicAnxietyDisorders  ContentCategory = "IAB7-34"
	ContentCategoryPediatrics             ContentCategory = "IAB7-35"
	ContentCategoryPhysicalTherapy        ContentCategory = "IAB7-36"
	ContentCategoryPsychologyPsychiatry   ContentCategory = "IAB7-37"
	ContentCategorySeniorHealth           ContentCategory = "IAB7-38"
	ContentCategorySexuality              ContentCategory = "IAB7-39"
	ContentCategorySleepDisorders         ContentCategory = "IAB7-40"
	ContentCategorySmokingCessation       ContentCategory = "IAB7-41"
	ContentCategorySubstanceAbuse         ContentCategory = "IAB7-42"
	ContentCategoryThyroidDisease         ContentCategory = "IAB7-43"
	ContentCategoryWeightLoss             ContentCategory = "IAB7-44"
	ContentCategoryWomensHealth           ContentCategory = "IAB7-45"

	ContentCategoryFoodDrink           ContentCategory = "IAB8"
	ContentCategoryAmericanCuisine     ContentCategory = "IAB8-1"
	ContentCategoryBarbecuesGrilling   ContentCategory = "IAB8-2"
	ContentCategoryCajunCreole         ContentCategory = "IAB8-3"
	ContentCategoryChineseCuisine      ContentCategory = "IAB8-4"
	ContentCategoryCocktailsBeer       ContentCategory = "IAB8-5"
	ContentCategoryCoffeeTea           ContentCategory = "IAB8-6"
	ContentCategoryCuisineSpecific     ContentCategory = "IAB8-7"
	ContentCategoryDessertsBaking      ContentCategory = "IAB8-8"
	ContentCategoryDiningOut           ContentCategory = "IAB8-9"
	ContentCategoryFoodAllergies       ContentCategory = "IAB8-10"
	ContentCategoryFrenchCuisine       ContentCategory = "IAB8-11"
	ContentCategoryHealthLowfatCooking ContentCategory = "IAB8-12"
	ContentCategoryItalianCuisine      ContentCategory = "IAB8-13"
	ContentCategoryJapaneseCuisine     ContentCategory = "IAB8-14"
	ContentCategoryMexicanCuisine      ContentCategory = "IAB8-15"
	ContentCategoryVegan               ContentCategory = "IAB8-16"
	ContentCategoryVegetarian          ContentCategory = "IAB8-17"
	ContentCategoryWine                ContentCategory = "IAB8-18"

	ContentCategoryHobbiesInterests   ContentCategory = "IAB9"
	ContentCategoryArtTechnology      ContentCategory = "IAB9-1"
	ContentCategoryArtsCrafts         ContentCategory = "IAB9-2"
	ContentCategoryBeadwork           ContentCategory = "IAB9-3"
	ContentCategoryBirdwatching       ContentCategory = "IAB9-4"
	ContentCategoryBoardGamesPuzzles  ContentCategory = "IAB9-5"
	ContentCategoryCandleSoapMaking   ContentCategory = "IAB9-6"
	ContentCategoryCardGames          ContentCategory = "IAB9-7"
	ContentCategoryChess              ContentCategory = "IAB9-8"
	ContentCategoryCigars             ContentCategory = "IAB9-9"
	ContentCategoryCollecting         ContentCategory = "IAB9-10"
	ContentCategoryComicBooks         ContentCategory = "IAB9-11"
	ContentCategoryDrawingSketching   ContentCategory = "IAB9-12"
	ContentCategoryFreelanceWriting   ContentCategory = "IAB9-13"
	ContentCategoryGenealogy          ContentCategory = "IAB9-14"
	ContentCategoryGettingPublished   ContentCategory = "IAB9-15"
	ContentCategoryGuitar             ContentCategory = "IAB9-16"
	ContentCategoryHomeRecording      ContentCategory = "IAB9-17"
	ContentCategoryInvestorsPatents   ContentCategory = "IAB9-18"
	ContentCategoryJewelryMaking      ContentCategory = "IAB9-19"
	ContentCategoryMagicIllusion      ContentCategory = "IAB9-20"
	ContentCategoryNeedlework         ContentCategory = "IAB9-21"
	ContentCategoryPainting           ContentCategory = "IAB9-22"
	ContentCategoryPhotography        ContentCategory = "IAB9-23"
	ContentCategoryRadio              ContentCategory = "IAB9-24"
	ContentCategoryRoleplayingGames   ContentCategory = "IAB9-25"
	ContentCategorySciFiFantasy       ContentCategory = "IAB9-26"
	ContentCategoryScrapbooking       ContentCategory = "IAB9-27"
	ContentCategoryScreenwriting      ContentCategory = "IAB9-28"
	ContentCategoryStampsCoins        ContentCategory = "IAB9-29"
	ContentCategoryVideoComputerGames ContentCategory = "IAB9-30"
	ContentCategoryWoodworking        ContentCategory = "IAB9-31"

	ContentCategoryHomeGarden             ContentCategory = "IAB10"
	ContentCategoryAppliances             ContentCategory = "IAB10-1"
	ContentCategoryEntertaining           ContentCategory = "IAB10-2"
	ContentCategoryEnvironmentalSafety    ContentCategory = "IAB10-3"
	ContentCategoryGardening              ContentCategory = "IAB10-4"
	ContentCategoryHomeRepair             ContentCategory = "IAB10-5"
	ContentCategoryHomeTheater            ContentCategory = "IAB10-6"
	ContentCategoryInteriorDecorating     ContentCategory = "IAB10-7"
	ContentCategoryLandscaping            ContentCategory = "IAB10-8"
	ContentCategoryRemodelingConstruction ContentCategory = "IAB10-9"

	ContentCategoryLawGovtPolitics       ContentCategory = "IAB11"
	ContentCategoryImmigration           ContentCategory = "IAB11-1"
	ContentCategoryLegalIssues           ContentCategory = "IAB11-2"
	ContentCategoryUSGovernmentResources ContentCategory = "IAB11-3"
	ContentCategoryPolitics              ContentCategory = "IAB11-4"
	ContentCategoryCommentary            ContentCategory = "IAB11-5"

	ContentCategoryNews              ContentCategory = "IAB12"
	ContentCategoryInternationalNews ContentCategory = "IAB12-1"
	ContentCategoryNationalNews      ContentCategory = "IAB12-2"
	ContentCategoryLocalNews         ContentCategory = "IAB12-3"

	ContentCategoryPersonalFinance    ContentCategory = "IAB13"
	ContentCategoryBeginningInvesting ContentCategory = "IAB13-1"
	ContentCategoryCreditDebtLoans    ContentCategory = "IAB13-2"
	ContentCategoryFinancialNews      ContentCategory = "IAB13-3"
	ContentCategoryFinancialPlanning  ContentCategory = "IAB13-4"
	ContentCategoryHedgeFund          ContentCategory = "IAB13-5"
	ContentCategoryInsurance          ContentCategory = "IAB13-6"
	ContentCategoryInvesting          ContentCategory = "IAB13-7"
	ContentCategoryMutualFunds        ContentCategory = "IAB13-8"
	ContentCategoryOptions            ContentCategory = "IAB13-9"
	ContentCategoryRetirementPlanning ContentCategory = "IAB13-10"
	ContentCategoryStocks             ContentCategory = "IAB13-11"
	ContentCategoryTaxPlanning        ContentCategory = "IAB13-12"

	ContentCategorySociety        ContentCategory = "IAB14"
	ContentCategoryDating         ContentCategory = "IAB14-1"
	ContentCategoryDivorceSupport ContentCategory = "IAB14-2"
	ContentCategoryGayLife        ContentCategory = "IAB14-3"
	ContentCategoryMarriage       ContentCategory = "IAB14-4"
	ContentCategorySeniorLiving   ContentCategory = "IAB14-5"
	ContentCategoryTeens          ContentCategory = "IAB14-6"
	ContentCategoryWeddings       ContentCategory = "IAB14-7"
	ContentCategoryEthnicSpecific ContentCategory = "IAB14-8"

	ContentCategoryScience             ContentCategory = "IAB15"
	ContentCategoryAstrology           ContentCategory = "IAB15-1"
	ContentCategoryBiology             ContentCategory = "IAB15-2"
	ContentCategoryChemistry           ContentCategory = "IAB15-3"
	ContentCategoryGeology             ContentCategory = "IAB15-4"
	ContentCategoryParanormalPhenomena ContentCategory = "IAB15-5"
	ContentCategoryPhysics             ContentCategory = "IAB15-6"
	ContentCategorySpaceAstronomy      ContentCategory = "IAB15-7"
	ContentCategoryGeography           ContentCategory = "IAB15-8"
	ContentCategoryBotany              ContentCategory = "IAB15-9"
	ContentCategoryWeather             ContentCategory = "IAB15-10"

	ContentCategoryPets               ContentCategory = "IAB16"
	ContentCategoryAquariums          ContentCategory = "IAB16-1"
	ContentCategoryBirds              ContentCategory = "IAB16-2"
	ContentCategoryCats               ContentCategory = "IAB16-3"
	ContentCategoryDogs               ContentCategory = "IAB16-4"
	ContentCategoryLargeAnimals       ContentCategory = "IAB16-5"
	ContentCategoryReptiles           ContentCategory = "IAB16-6"
	ContentCategoryVeterinaryMedicine ContentCategory = "IAB16-7"

	ContentCategorySports              ContentCategory = "IAB17"
	ContentCategoryAutoRacing          ContentCategory = "IAB17-1"
	ContentCategoryBaseball            ContentCategory = "IAB17-2"
	ContentCategoryBicycling           ContentCategory = "IAB17-3"
	ContentCategoryBodybuilding        ContentCategory = "IAB17-4"
	ContentCategoryBoxing              ContentCategory = "IAB17-5"
	ContentCategoryCanoeingKayaking    ContentCategory = "IAB17-6"
	ContentCategoryCheerleading        ContentCategory = "IAB17-7"
	ContentCategoryClimbing            ContentCategory = "IAB17-8"
	ContentCategoryCricket             ContentCategory = "IAB17-9"
	ContentCategoryFigureSkating       ContentCategory = "IAB17-10"
	ContentCategoryFlyFishing          ContentCategory = "IAB17-11"
	ContentCategoryFootball            ContentCategory = "IAB17-12"
	ContentCategoryFreshwaterFishing   ContentCategory = "IAB17-13"
	ContentCategoryGameFish            ContentCategory = "IAB17-14"
	ContentCategoryGolf                ContentCategory = "IAB17-15"
	ContentCategoryHorseRacing         ContentCategory = "IAB17-16"
	ContentCategoryHorses              ContentCategory = "IAB17-17"
	ContentCategoryHuntingShooting     ContentCategory = "IAB17-18"
	ContentCategoryInlineSkating       ContentCategory = "IAB17-19"
	ContentCategoryMartialArts         ContentCategory = "IAB17-20"
	ContentCategoryMountainBiking      ContentCategory = "IAB17-21"
	ContentCategoryNASCARRacing        ContentCategory = "IAB17-22"
	ContentCategoryOlympics            ContentCategory = "IAB17-23"
	ContentCategoryPaintball           ContentCategory = "IAB17-24"
	ContentCategoryPowerMotorcycles    ContentCategory = "IAB17-25"
	ContentCategoryProBasketball       ContentCategory = "IAB17-26"
	ContentCategoryProIceHockey        ContentCategory = "IAB17-27"
	ContentCategoryRodeo               ContentCategory = "IAB17-28"
	ContentCategoryRugby               ContentCategory = "IAB17-29"
	ContentCategoryRunningJogging      ContentCategory = "IAB17-30"
	ContentCategorySailing             ContentCategory = "IAB17-31"
	ContentCategorySaltwaterFishing    ContentCategory = "IAB17-32"
	ContentCategoryScubaDiving         ContentCategory = "IAB17-33"
	ContentCategorySkateboarding       ContentCategory = "IAB17-34"
	ContentCategorySkiing              ContentCategory = "IAB17-35"
	ContentCategorySnowboarding        ContentCategory = "IAB17-36"
	ContentCategorySurfingBodyboarding ContentCategory = "IAB17-37"
	ContentCategorySwimming            ContentCategory = "IAB17-38"
	ContentCategoryTableTennisPingPong ContentCategory = "IAB17-39"
	ContentCategoryTennis              ContentCategory = "IAB17-40"
	ContentCategoryVolleyball          ContentCategory = "IAB17-41"
	ContentCategoryWalking             ContentCategory = "IAB17-42"
	ContentCategoryWaterskiWakeboard   ContentCategory = "IAB17-43"
	ContentCategoryWorldSoccer         ContentCategory = "IAB17-44"

	ContentCategoryStyleFashion ContentCategory = "IAB18"
	ContentCategoryBeauty       ContentCategory = "IAB18-1"
	ContentCategoryBodyArt      ContentCategory = "IAB18-2"
	ContentCategoryFashion      ContentCategory = "IAB18-3"
	ContentCategoryJewelry      ContentCategory = "IAB18-4"
	ContentCategoryClothing     ContentCategory = "IAB18-5"
	ContentCategoryAccessories  ContentCategory = "IAB18-6"

	ContentCategoryTechnologyComputing   ContentCategory = "IAB19"
	ContentCategoryDGraphics             ContentCategory = "IAB19-1"
	ContentCategoryAnimation             ContentCategory = "IAB19-2"
	ContentCategoryAntivirusSoftware     ContentCategory = "IAB19-3"
	ContentCategoryCC                    ContentCategory = "IAB19-4"
	ContentCategoryCamerasCamcorders     ContentCategory = "IAB19-5"
	ContentCategoryCellPhones            ContentCategory = "IAB19-6"
	ContentCategoryComputerCertification ContentCategory = "IAB19-7"
	ContentCategoryComputerNetworking    ContentCategory = "IAB19-8"
	ContentCategoryComputerPeripherals   ContentCategory = "IAB19-9"
	ContentCategoryComputerReviews       ContentCategory = "IAB19-10"
	ContentCategoryDataCenters           ContentCategory = "IAB19-11"
	ContentCategoryDatabases             ContentCategory = "IAB19-12"
	ContentCategoryDesktopPublishing     ContentCategory = "IAB19-13"
	ContentCategoryDesktopVideo          ContentCategory = "IAB19-14"
	ContentCategoryEmail                 ContentCategory = "IAB19-15"
	ContentCategoryGraphicsSoftware      ContentCategory = "IAB19-16"
	ContentCategoryHomeVideoDVD          ContentCategory = "IAB19-17"
	ContentCategoryInternetTechnology    ContentCategory = "IAB19-18"
	ContentCategoryJava                  ContentCategory = "IAB19-19"
	ContentCategoryJavaScript            ContentCategory = "IAB19-20"
	ContentCategoryMacSupport            ContentCategory = "IAB19-21"
	ContentCategoryMP3MIDI               ContentCategory = "IAB19-22"
	ContentCategoryNetConferencing       ContentCategory = "IAB19-23"
	ContentCategoryNetforBeginners       ContentCategory = "IAB19-24"
	ContentCategoryNetworkSecurity       ContentCategory = "IAB19-25"
	ContentCategoryPalmtopsPDAs          ContentCategory = "IAB19-26"
	ContentCategoryPCSupport             ContentCategory = "IAB19-27"
	ContentCategoryPortable              ContentCategory = "IAB19-28"
	ContentCategoryEntertainment         ContentCategory = "IAB19-29"
	ContentCategorySharewareFreeware     ContentCategory = "IAB19-30"
	ContentCategoryUnix                  ContentCategory = "IAB19-31"
	ContentCategoryVisualBasic           ContentCategory = "IAB19-32"
	ContentCategoryWebClipArt            ContentCategory = "IAB19-33"
	ContentCategoryWebDesignHTML         ContentCategory = "IAB19-34"
	ContentCategoryWebSearch             ContentCategory = "IAB19-35"
	ContentCategoryWindows               ContentCategory = "IAB19-36"

	ContentCategoryTravel               ContentCategory = "IAB20"
	ContentCategoryAdventureTravel      ContentCategory = "IAB20-1"
	ContentCategoryAfrica               ContentCategory = "IAB20-2"
	ContentCategoryAirTravel            ContentCategory = "IAB20-3"
	ContentCategoryAustraliaNewZealand  ContentCategory = "IAB20-4"
	ContentCategoryBedBreakfasts        ContentCategory = "IAB20-5"
	ContentCategoryBudgetTravel         ContentCategory = "IAB20-6"
	ContentCategoryBusinessTravel       ContentCategory = "IAB20-7"
	ContentCategoryByUSLocale           ContentCategory = "IAB20-8"
	ContentCategoryCamping              ContentCategory = "IAB20-9"
	ContentCategoryCanada               ContentCategory = "IAB20-10"
	ContentCategoryCaribbean            ContentCategory = "IAB20-11"
	ContentCategoryCruises              ContentCategory = "IAB20-12"
	ContentCategoryEasternEurope        ContentCategory = "IAB20-13"
	ContentCategoryEurope               ContentCategory = "IAB20-14"
	ContentCategoryFrance               ContentCategory = "IAB20-15"
	ContentCategoryGreece               ContentCategory = "IAB20-16"
	ContentCategoryHoneymoonsGetaways   ContentCategory = "IAB20-17"
	ContentCategoryHotels               ContentCategory = "IAB20-18"
	ContentCategoryItaly                ContentCategory = "IAB20-19"
	ContentCategoryJapan                ContentCategory = "IAB20-20"
	ContentCategoryMexicoCentralAmerica ContentCategory = "IAB20-21"
	ContentCategoryNationalParks        ContentCategory = "IAB20-22"
	ContentCategorySouthAmerica         ContentCategory = "IAB20-23"
	ContentCategorySpas                 ContentCategory = "IAB20-24"
	ContentCategoryThemeParks           ContentCategory = "IAB20-25"
	ContentCategoryTravelingwithKids    ContentCategory = "IAB20-26"
	ContentCategoryUnitedKingdom        ContentCategory = "IAB20-27"

	ContentCategoryRealEstate         ContentCategory = "IAB21"
	ContentCategoryApartments         ContentCategory = "IAB21-1"
	ContentCategoryArchitects         ContentCategory = "IAB21-2"
	ContentCategoryBuyingSellingHomes ContentCategory = "IAB21-3"

	ContentCategoryShopping         ContentCategory = "IAB22"
	ContentCategoryContestsFreebies ContentCategory = "IAB22-1"
	ContentCategoryCouponing        ContentCategory = "IAB22-2"
	ContentCategoryComparison       ContentCategory = "IAB22-3"
	ContentCategoryEngines          ContentCategory = "IAB22-4"

	ContentCategoryReligionSpirituality ContentCategory = "IAB23"
	ContentCategoryAlternativeReligions ContentCategory = "IAB23-1"
	ContentCategoryAtheismAgnosticism   ContentCategory = "IAB23-2"
	ContentCategoryBuddhism             ContentCategory = "IAB23-3"
	ContentCategoryCatholicism          ContentCategory = "IAB23-4"
	ContentCategoryChristianity         ContentCategory = "IAB23-5"
	ContentCategoryHinduism             ContentCategory = "IAB23-6"
	ContentCategoryIslam                ContentCategory = "IAB23-7"
	ContentCategoryJudaism              ContentCategory = "IAB23-8"
	ContentCategoryLatterDaySaints      ContentCategory = "IAB23-9"
	ContentCategoryPaganWiccan          ContentCategory = "IAB23-10"

	ContentCategoryUncategorized ContentCategory = "IAB24"

	ContentCategoryNonStandardContent             ContentCategory = "IAB25"
	ContentCategoryUnmoderatedUGC                 ContentCategory = "IAB25-1"
	ContentCategoryExtremeGraphicExplicitViolence ContentCategory = "IAB25-2"
	ContentCategoryPornography                    ContentCategory = "IAB25-3"
	ContentCategoryProfaneContent                 ContentCategory = "IAB25-4"
	ContentCategoryHateContent                    ContentCategory = "IAB25-5"
	ContentCategoryUnderConstruction              ContentCategory = "IAB25-6"
	ContentCategoryIncentivized                   ContentCategory = "IAB25-7"

	ContentCategoryAnyIllegalContent     ContentCategory = "IAB26"
	ContentCategoryIllegalContent        ContentCategory = "IAB26-1"
	ContentCategoryWarez                 ContentCategory = "IAB26-2"
	ContentCategorySpywareMalware        ContentCategory = "IAB26-3"
	ContentCategoryCopyrightInfringement ContentCategory = "IAB26-4"
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

// APIFramework as defined in Adcom1.0.
type APIFramework int

// API Frameworks options as defined in Adcom1.0.
const (
	APIFrameworkUnknown APIFramework = 0
	APIFrameworkVPAID1  APIFramework = 1
	APIFrameworkVPAID2  APIFramework = 2
	APIFrameworkMRAID1  APIFramework = 3
	APIFrameworkORMMA   APIFramework = 4
	APIFrameworkMRAID2  APIFramework = 5
	APIFrameworkMRAID3  APIFramework = 6
	APIFrameworkOMID    APIFramework = 7
	APIFrameworkSIMID1  APIFramework = 8
	APIFrameworkSIMID11 APIFramework = 9
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
	DeviceTypeOOH       DeviceType = 8
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
// It is recommended that either the w/h pair or the wratio/hratio/wmin set (i.e., for Flex Ads) be specified.
type Format struct {
	Width       int             `json:"w,omitempty"`       // Width in device independent pixels (DIPS).
	Height      int             `json:"h,omitempty"`       // Height in device independent pixels (DIPS).
	WidthRatio  int             `json:"wratio,omitempty"`  // Relative width when expressing size as a ratio.
	HeightRatio int             `json:"hration,omitempty"` // Relative height when expressing size as a ratio.
	WidthMin    int             `json:"wmin,omitempty"`    // The minimum width in device independent pixels (DIPS) at which the ad will be displayed the size is expressed as a ratio.
	Ext         json.RawMessage `json:"ext,omitempty"`
}

// PodSequence identifies the pod sequence field, for use in video content streams with one or more ad pods as defined in Adcom1.0
type PodSequence int

// PodSequence options as defined in Adcom1.0
const (
	PodSeqLast  PodSequence = -1 // Last pod in the content stream.
	PodSeqAny   PodSequence = 0  // Any pod in the content stream.
	PodSeqFirst PodSequence = 1  // First pod in the content stream.
)

// SlotPositionInPod identifies the slot position in pod field, for use in video ad pods as defined in Adcom1.0
type SlotPositionInPod int

// SlotPositionInPod options as defined in Adcom1.0
const (
	SlotPosLast        SlotPositionInPod = -1 // Last ad in the pod.
	SlotPosAny         SlotPositionInPod = 0  // Any ad in the pod.
	SlotPosFirst       SlotPositionInPod = 1  // First ad in the pod.
	SlotPosFirstOrLast SlotPositionInPod = 2  // First or Last ad in the pod.
)

// Type of the creative markup so that it can properly be associated with the right sub-object of the BidRequest.Imp.
type MarkupType int

// MarkupType available options
const (
	MarkupUnknown MarkupType = 0
	MarkupBanner  MarkupType = 1
	MarkupVideo   MarkupType = 2
	MarkupAudio   MarkupType = 3
	MarkupNative  MarkupType = 4
)

// CategoryTaxonomy identifies the taxonomy in effect when content categories as defined in Adcom1.0.
type CategoryTaxonomy int

// CategoryTaxonomy options as defined in Adcom1.0
const (
	CategoryTaxonomyIABContent1   CategoryTaxonomy = 1 // 1	IAB Content Category Taxonomy 1.0.
	CategoryTaxonomyIABContent2   CategoryTaxonomy = 2 // 2	IAB Content Category Taxonomy 2: www.iab.com/guidelines/taxonomy
	CategoryTaxonomyIABProduct1   CategoryTaxonomy = 3 // 3	IAB Ad Product Taxonomy 1.0.
	CategoryTaxonomyIABAudience11 CategoryTaxonomy = 4 // 4	IAB Audience Taxonomy 1.1.
	CategoryTaxonomyIABContent21  CategoryTaxonomy = 5 // 5	IAB Content Category Taxonomy 2.1.
	CategoryTaxonomyIABContent22  CategoryTaxonomy = 6 // 6	IAB Content Category Taxonomy 2.2
)
