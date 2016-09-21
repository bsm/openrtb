package request

import "encoding/json"

type Data_Type int

const (
	Data_Sponsored      Data_Type = 1  // Sponsored By message where response should contain the brand name of the sponsor
	Data_Desc           Data_Type = 2  // Descriptive text associated with the product or service being advertised
	Data_Rating         Data_Type = 3  // Rating of the product being offered to the user
	Data_Likes          Data_Type = 4  // Number of social ratings or “likes” of the product being offered to the user
	Data_Downloads      Data_Type = 5  // Number downloads/installs of this product
	Data_Price          Data_Type = 6  // Price for product / app / in-app purchase. Value should include currency symbol in localised format
	Data_SalePrice      Data_Type = 7  // Sale price that can be used together with price to indicate a discounted price compared to a regular price. Value should include currency symbol in localised format
	Data_Phone          Data_Type = 8  // Phone number
	Data_Address        Data_Type = 9  // Address
	Data_DescAdditional Data_Type = 10 // Additional descriptive text associated with the product or service being advertised
	Data_DisplayURL     Data_Type = 11 // Display URL for the text ad. To be used when sponsoring entity doesn’t own the content. IE sponsored by BRAND on SITE (where SITE is transmitted in this field)
	Data_CTADesc        Data_Type = 12 // CTA description - descriptive text describing a ‘call to action’ button for the destination URL

)

type Data struct {
	Type   Data_Type       `json:"type"` // Type ID of the element supported by the publisher. The publisher can display this information in an appropriate format
	Length int             `json:"len"`  // Maximum length of the text in the element’s response
	Ext    json.RawMessage `json:"ext,omitempty"`
}
