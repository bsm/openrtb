package openrtb

// An "app" object should be included if the ad supported content is part of a mobile application
// (as opposed to a mobile website).  A bid request must not contain both an "app" object and a
// "site" object.
type App struct {
	Id            *string    `json:"id,omitempty"`   // App ID on the exchange
	Name          *string    `json:"name,omitempty"` // App name
	Domain        *string    `json:"domain,omitempty"`
	Cat           []string   `json:"cat,omitempty"`           // Array of IAB content categories
	Sectioncat    []string   `json:"sectioncat,omitempty"`    // Array of IAB content categories for subsection
	Pagecat       []string   `json:"pagecat,omitempty"`       // Array of IAB content categories for page
	Ver           *string    `json:"ver,omitempty"`           // App version
	Bundle        *string    `json:"bundle,omitempty"`        // App bundle or package name
	Privacypolicy *int       `json:"privacypolicy,omitempty"` // Default: 1 ("1": site has a privacy policy)
	Paid          *int       `json:"paid,omitempty"`          // "1": Paid, "2": Free
	Publisher     *Publisher `json:"publisher,omitempty"`
	// Content       Content
	Keywords *string    `json:"keywords,omitempty"`
	Storeurl *string    `json:"storeurl,omitempty"` // App store URL for an installed app
	Ext      Extensions `json:"ext,omitempty"`
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
