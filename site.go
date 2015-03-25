package openrtb

// A site object should be included if the ad supported content is part of a website (as opposed to
// an application).  A bid request must not contain both a site object and an app object.
type Site struct {
	Id            *string    `json:"id,omitempty"`   // Site ID on the exchange
	Name          *string    `json:"name,omitempty"` // Site name
	Domain        *string    `json:"domain,omitempty"`
	Cat           []string   `json:"cat,omitempty"`           // Array of IAB content categories
	Sectioncat    []string   `json:"sectioncat,omitempty"`    // Array of IAB content categories for subsection
	Pagecat       []string   `json:"pagecat,omitempty"`       // Array of IAB content categories for page
	Page          *string    `json:"page,omitempty"`          // URL of the page
	Privacypolicy *int       `json:"privacypolicy,omitempty"` // Default: 1 ("1": site has a privacy policy)
	Ref           *string    `json:"ref,omitempty"`           // Referrer URL
	Search        *string    `json:"search,omitempty"`        // Search string that caused naviation
	Mobile        *int       `json:"mobile,omitempty"`        // Mobile ("1": site is mobile optimised)
	Publisher     *Publisher `json:"publisher,omitempty"`
	// Content       Content
	Keywords *string    `json:"keywords,omitempty"`
	Ext      Extensions `json:"ext,omitempty"`
}

// Returns the privacy policy status, with default fallback
func (s *Site) IsPrivacyPolicy() bool {
	if s.Privacypolicy != nil {
		return *s.Privacypolicy == 1
	}
	return false
}

// Applies defaults
func (s *Site) WithDefaults() *Site {
	if s.Privacypolicy == nil {
		s.Privacypolicy = new(int)
		*s.Privacypolicy = 0
	}
	return s
}
