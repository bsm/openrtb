package openrtb

// A site object should be included if the ad supported content is part of a website (as opposed to
// an application).  A bid request must not contain both a site object and an app object.
type Site struct {
	Id            *string // Site ID on the exchange
	Name          *string // Site name
	Domain        *string
	Cat           []string // Array of IAB content categories
	Sectioncat    []string // Array of IAB content categories for subsection
	Pagecat       []string // Array of IAB content categories for page
	Page          *string  // URL of the page
	Privacypolicy *int     // Default: 1 ("1": site has a privacy policy)
	Ref           *string  // Referrer URL
	Search        *string  // Search string that caused naviation
	Publisher     *Publisher
	// Content       Content
	Keywords []string
	Ext      Extensions
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
