package openrtb

// An "app" object should be included if the ad supported content is part of a mobile application
// (as opposed to a mobile website).  A bid request must not contain both an "app" object and a
// "site" object.
type App struct {
	Id            *string // App ID on the exchange
	Name          *string // App name
	Domain        *string
	Cat           []string // Array of IAB content categories
	Sectioncat    []string // Array of IAB content categories for subsection
	Pagecat       []string // Array of IAB content categories for page
	Ver           *string  // App version
	Bundle        *string  // App bundle or package name
	Privacypolicy *int     // Default: 1 ("1": site has a privacy policy)
	Paid          *int     // "1": Paid, "2": Free
	Publisher     *Publisher
	// Content       Content
	Keywords []string
	Storeurl *string // App store URL for an installed app
	Ext      Extensions
}

// Returns the privacy policy status, with default fallback
func (a *App) IsPrivacyPolicy() bool {
	if a.Privacypolicy != nil {
		return *a.Privacypolicy == 1
	}
	return false
}

// Returns the paid status, with default fallback
func (a *App) IsPaid() bool {
	if a.Paid != nil {
		return *a.Paid == 1
	}
	return false
}

// Applies defaults
func (a *App) WithDefaults() *App {
	if a.Privacypolicy == nil {
		a.Privacypolicy = new(int)
		*a.Privacypolicy = 0
	}
	if a.Paid == nil {
		a.Paid = new(int)
		*a.Paid = 0
	}
	return a
}
