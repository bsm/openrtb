package openrtb

type Inventory struct {
	ID            string     `json:"id,omitempty"` // ID on the exchange
	Name          string     `json:"name,omitempty"`
	Domain        string     `json:"domain,omitempty"`
	Cat           []string   `json:"cat,omitempty"`           // Array of IAB content categories
	SectionCat    []string   `json:"sectioncat,omitempty"`    // Array of IAB content categories for subsection
	PageCat       []string   `json:"pagecat,omitempty"`       // Array of IAB content categories for page
	PrivacyPolicy *int       `json:"privacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
	Publisher     *Publisher `json:"publisher,omitempty"`     // Details about the Publisher
	Content       *Content   `json:"content,omitempty"`       // Details about the Content
	Keywords      string     `json:"keywords,omitempty"`      // Comma separated list of keywords about the site.
	Ext           Extension  `json:"ext,omitempty"`
}

// GetPrivacyPolicy returns the privacy policy value
func (a *Inventory) GetPrivacyPolicy() int {
	if a.PrivacyPolicy != nil {
		return *a.PrivacyPolicy
	}
	return 1
}

// An "app" object should be included if the ad supported content is part of a mobile application
// (as opposed to a mobile website).  A bid request must not contain both an "app" object and a
// "site" object.
type App struct {
	Inventory
	Bundle   string `json:"bundle,omitempty"`   // App bundle or package name
	StoreURL string `json:"storeurl,omitempty"` // App store URL for an installed app
	Ver      string `json:"ver,omitempty"`      // App version
	Paid     int    `json:"paid,omitempty"`     // "1": Paid, "2": Free
}

// A site object should be included if the ad supported content is part of a website (as opposed to
// an application).  A bid request must not contain both a site object and an app object.
type Site struct {
	Inventory
	Page   string `json:"page,omitempty"`   // URL of the page
	Ref    string `json:"ref,omitempty"`    // Referrer URL
	Search string `json:"search,omitempty"` // Search string that caused naviation
	Mobile int    `json:"mobile,omitempty"` // Mobile ("1": site is mobile optimised)
}
