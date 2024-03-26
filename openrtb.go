package openrtb

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

// 5.2 Banner Ad Types
const (
	BannerTypeXHTMLText = 1
	BannerTypeXHTML     = 2
	BannerTypeJS        = 3
	BannerTypeFrame     = 4
)

// 5.3 Creative Attributes
const (
	CreativeAttributeAudioAdAutoPlay                 = 1
	CreativeAttributeAudioAdUserInitiated            = 2
	CreativeAttributeExpandableAuto                  = 3
	CreativeAttributeExpandableUserInitiatedClick    = 4
	CreativeAttributeExpandableUserInitiatedRollover = 5
	CreativeAttributeInBannerVideoAdAutoPlay         = 6
	CreativeAttributeInBannerVideoAdUserInitiated    = 7
	CreativeAttributePop                             = 8
	CreativeAttributeProvocativeOrSuggestiveImagery  = 9
	CreativeAttributeExtremeAnimation                = 10
	CreativeAttributeSurveys                         = 11
	CreativeAttributeTextOnly                        = 12
	CreativeAttributeUserInitiated                   = 13
	CreativeAttributeWindowsDialogOrAlert            = 14
	CreativeAttributeHasAudioWithPlayer              = 15
	CreativeAttributeAdProvidesSkipButton            = 16
	CreativeAttributeAdobeFlash                      = 17
)

// 5.4 Ad Position
const (
	AdPosUnknown    = 0
	AdPosAboveFold  = 1
	AdPosBelowFold  = 3
	AdPosHeader     = 4
	AdPosFooter     = 5
	AdPosSidebar    = 6
	AdPosFullscreen = 7
)

// 5.5 Expandable Direction
const (
	ExpDirUnknown    = 0
	ExpDirLeft       = 1
	ExpDirRight      = 2
	ExpDirUp         = 3
	ExpDirDown       = 4
	ExpDirFullScreen = 5
)

// 5.6 API Frameworks
const (
	APIFrameworkUnknown = 0
	APIFrameworkVPAID1  = 1
	APIFrameworkVPAID2  = 2
	APIFrameworkMRAID1  = 3
	APIFrameworkORMMA   = 4
	APIFrameworkMRAID2  = 5
	APIFrameworkMRAID3  = 6
	APIFrameworkOMID1   = 7
	APIFrameworkSIMID1  = 8
	APIFrameworkSIMID11 = 9
)

// 5.7 Video Linearity
const (
	VideoLinearityUnknown   = 0
	VideoLinearityLinear    = 1
	VideoLinearityNonLinear = 2
)

// 5.8 Video and Audio Bid Response Protocols
const (
	ProtocolUnknown            = 0
	VideoProtoVAST1            = 1
	VideoProtoVAST2            = 2
	VideoProtoVAST3            = 3
	VideoProtoVAST1Wrapper     = 4
	VideoProtoVAST2Wrapper     = 5
	VideoProtoVAST3Wrapper     = 6
	VideoProtoVAST4            = 7
	VideoProtoVAST4Wrapper     = 8
	AudioProtocolDAAST1        = 9
	AudioProtocolDAAST1Wrapper = 10
)

// 5.9 Video Placement Types
const (
	VideoPlacementUnknown      = 0
	VideoPlacementInStream     = 1
	VideoPlacementInBanner     = 2
	VideoPlacementInArticle    = 3
	VideoPlacementInFeed       = 4
	VideoPlacementInterstitial = 5
)

// 5.10 Video Playback Methods
const (
	VideoPlaybackUnknown       = 0
	VideoPlaybackAutoSoundOn   = 1
	VideoPlaybackAutoSoundOff  = 2
	VideoPlaybackClickToPlay   = 3
	VideoPlaybackMouseOver     = 4
	VideoPlaybackEnterSoundOn  = 5
	VideoPlaybackEnterSoundOff = 6
)

// 5.12 Video Start Delay
const (
	VideoStartDelayPreRoll         = 0
	VideoStartDelayGenericMidRoll  = -1
	VideoStartDelayGenericPostRoll = -2
)

// 5.13 Video Quality
const (
	VideoQualityUnknown      = 0
	VideoQualityProfessional = 1
	VideoQualityProsumer     = 2
	VideoQualityUGC          = 3
)

// 5.14 Companion Types
const (
	VASTCompanionUnknown = 0
	VASTCompanionStatic  = 1
	VASTCompanionHTML    = 2
	VASTCompanionIFrame  = 3
)

// 5.15 Content Delivery Methods
const (
	ContentDeliveryUnknown     = 0
	ContentDeliveryStreaming   = 1
	ContentDeliveryProgressive = 2
	ContentDeliveryDownload    = 3
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

// 5.18 Content Context
const (
	ContextVideo       = 1
	ContextGame        = 2
	ContextMusic       = 3
	ContextApplication = 4
	ContextText        = 5
	ContextOther       = 6
	ContextUnknown     = 7
)

// 5.19 IQG Media Ratings
const (
	QAGUnknown = 0
	QAGAll     = 1
	QAGOver12  = 2
	QAGMature  = 3
)

// 5.20 Location Type
const (
	LocationTypeUnknown = 0
	LocationTypeGPS     = 1
	LocationTypeIP      = 2
	LocationTypeUser    = 3
)

// 5.21 Device Type
const (
	DeviceTypeUnknown   = 0
	DeviceTypeMobile    = 1
	DeviceTypePC        = 2
	DeviceTypeTV        = 3
	DeviceTypePhone     = 4
	DeviceTypeTablet    = 5
	DeviceTypeConnected = 6
	DeviceTypeSetTopBox = 7
	DeviceTypeOOH       = 8
)

// 5.22 Connection Type
const (
	ConnTypeUnknown  = 0
	ConnTypeEthernet = 1
	ConnTypeWIFI     = 2
	ConnTypeCell     = 3
	ConnTypeCell2G   = 4
	ConnTypeCell3G   = 5
	ConnTypeCell4G   = 6
	ConnTypeCell5G   = 7
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

// 5.24 No-Bid Reason Codes
const (
	NBRUnknownError      = 0
	NBRTechnicalError    = 1
	NBRInvalidRequest    = 2
	NBRKnownSpider       = 3
	NBRSuspectedNonHuman = 4
	NBRProxyIP           = 5
	NBRUnsupportedDevice = 6
	NBRBlockedSite       = 7
	NBRUnmatchedUser     = 8
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
	ID     string    `json:"id,omitempty"`
	Name   string    `json:"name,omitempty"`
	Cat    []string  `json:"cat,omitempty"` // Array of IAB content categories
	Domain string    `json:"domain,omitempty"`
	Ext    Extension `json:"ext,omitempty"`
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
	Lat           float64   `json:"lat,omitempty"`           // Latitude from -90 to 90
	Lon           float64   `json:"lon,omitempty"`           // Longitude from -180 to 180
	Type          int       `json:"type,omitempty"`          // Indicate the source of the geo data
	Accuracy      int       `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
	LastFix       int       `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     int       `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable
	Country       string    `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 3
	Region        string    `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string    `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string    `json:"metro,omitempty"`
	City          string    `json:"city,omitempty"`
	Zip           string    `json:"zip,omitempty"`
	UTCOffset     int       `json:"utcoffset,omitempty"` // Local time as the number +/- of minutes from UTC
	Ext           Extension `json:"ext,omitempty"`
}

// User object contains information known or derived about the human user of the device (i.e., the
// audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
// privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
// frequency capping and retargeting.
type User struct {
	ID         string        `json:"id,omitempty"`         // Unique consumer ID of this user on the exchange
	BuyerID    string        `json:"buyerid,omitempty"`    // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid/buyerid or id is recommended. Valid for OpenRTB 2.3.
	BuyerUID   string        `json:"buyeruid,omitempty"`   // Buyer-specific ID for the user as mapped by the exchange for the buyer. Same as BuyerID but valid for OpenRTB 2.2.
	YOB        int           `json:"yob,omitempty"`        // Year of birth as a 4-digit integer.
	Gender     string        `json:"gender,omitempty"`     // Gender ("M": male, "F" female, "O" Other)
	Keywords   string        `json:"keywords,omitempty"`   // Comma separated list of keywords, interests, or intent
	CustomData string        `json:"customdata,omitempty"` // Optional feature to pass bidder data that was set in the exchange's cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include "escaped" quotation marks.
	Geo        *Geo          `json:"geo,omitempty"`
	Data       []Data        `json:"data,omitempty"`
	Ext        UserExtension `json:"ext,omitempty"`
}

// Data and segment objects together allow additional data about the user to be specified. This data
// may be from multiple sources whether from the exchange itself or third party providers as specified by
// the id field. A bid request can mix data objects from multiple providers. The specific data providers in
// use should be published by the exchange a priori to its bidders.
type Data struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Segment []Segment `json:"segment,omitempty"`
	Ext     Extension `json:"ext,omitempty"`
}

// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	ID    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Value string    `json:"value,omitempty"`
	Ext   Extension `json:"ext,omitempty"`
}

// Regulations object contains any legal, governmental, or industry regulations that apply to the request. The
// coppa flag signals whether or not the request falls under the United States Federal Trade Commission's
// regulations for the United States Children's Online Privacy Protection Act ("COPPA").
type Regulations struct {
	Coppa     int          `json:"coppa,omitempty"`      // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	GDPR      int          `json:"gdpr,omitempty"`       // Flag that indicates whether or not the request is subject to GDPR regulations 0 = No, 1 = Yes, omission indicates Unknown.
	USPrivacy string       `json:"us_privacy,omitempty"` // Communicates signals regarding consumer privacy under US privacy regulation.
	Ext       RegExtension `json:"ext,omitempty"`
}

// Format object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
// It is recommended that either the w/h pair or the wratio/hratio/wmin set (i.e., for Flex Ads) be specified.
type Format struct {
	W           int       `json:"w,omitempty"`       // Width in device independent pixels (DIPS).
	H           int       `json:"h,omitempty"`       // Height in device independent pixels (DIPS).
	WidthRatio  int       `json:"wratio,omitempty"`  // Relative width when expressing size as a ratio.
	HeightRatio int       `json:"hration,omitempty"` // Relative height when expressing size as a ratio.
	WidthMin    int       `json:"wmin,omitempty"`    // The minimum width in device independent pixels (DIPS) at which the ad will be displayed the size is expressed as a ratio.
	Ext         Extension `json:"ext,omitempty"`
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

// VideoPlcmt represents the the various types of video placements in accordance with updated IAB Digital Video Guidelines.
type VideoPlcmt int

// Types of video placements derived largely from the IAB Digital Video Guidelines.
const (
	VideoPlcmtInstream            VideoPlcmt = 1
	VideoPlcmtAccompanyingContent VideoPlcmt = 2
	VideoPlcmtInterstitial        VideoPlcmt = 3
	VideoPlcmtNoContent           VideoPlcmt = 4
)

// ChannelEntity describes the network or channel an ad will be displayed on. (Reffer Section 3.2.23 and 3.2.24 OpenRTB_2.6)
type ChannelEntity struct {
	ID     string    `json:"id,omitempty"`
	Name   string    `json:"name,omitempty"`
	Domain string    `json:"domain,omitempty"`
	Ext    Extension `json:"ext,omitempty"`
}

// RegExtension Extension object for Regulations
type RegExtension struct {
	GDPR      int    `json:"gdpr,omitempty"`
	LGPD      bool   `json:"lgpd,omitempty"`
	PIPL      bool   `json:"pipl,omitempty"`
	USPrivacy string `json:"us_privacy,omitempty"`
}

// UserExtension Extension object for User
type UserExtension struct {
	Consent      string  `json:"consent,omitempty"`
	SdkData      SdkData `json:"sdkdata,omitempty"`
	AppStartTime int64   `json:"appStartTime,omitempty"`
	SessionID    string  `json:"sessionId,omitempty"`
}

// SdkData object for UserExtension. Required for direct demand header bidding integration
type SdkData struct {
	Seq    int32 `json:"seq"`    // Number of times Header Bidding same prevaluation token has been requested by Mediation SDK
	Loads  int32 `json:"loads"`  // Number of times same prevaluation token has been used to load AdM
	Starts int32 `json:"starts"` // Number of times starts has been registered for the same prevaluation token
}
