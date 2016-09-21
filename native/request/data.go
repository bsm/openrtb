package request

import "encoding/json"

type DataTypeID int

const (
	DataTypeSponsored      DataTypeID = 1  // Sponsored By message where response should contain the brand name of the sponsor
	DataTypeDesc           DataTypeID = 2  // Descriptive text associated with the product or service being advertised
	DataTypeRating         DataTypeID = 3  // Rating of the product being offered to the user
	DataTypeLikes          DataTypeID = 4  // Number of social ratings or “likes” of the product being offered to the user
	DataTypeDownloads      DataTypeID = 5  // Number downloads/installs of this product
	DataTypePrice          DataTypeID = 6  // Price for product / app / in-app purchase. Value should include currency symbol in localised format
	DataTypeSalePrice      DataTypeID = 7  // Sale price that can be used together with price to indicate a discounted price compared to a regular price. Value should include currency symbol in localised format
	DataTypePhone          DataTypeID = 8  // Phone number
	DataTypeAddress        DataTypeID = 9  // Address
	DataTypeDescAdditional DataTypeID = 10 // Additional descriptive text associated with the product or service being advertised
	DataTypeDisplayURL     DataTypeID = 11 // Display URL for the text ad. To be used when sponsoring entity doesn’t own the content. IE sponsored by BRAND on SITE (where SITE is transmitted in this field)
	DataTypeCTADesc        DataTypeID = 12 // CTA description - descriptive text describing a ‘call to action’ button for the destination URL

)

type Data struct {
	TypeID DataTypeID      `json:"type"` // Type ID of the element supported by the publisher. The publisher can display this information in an appropriate format
	Length int             `json:"len"`  // Maximum length of the text in the element’s response
	Ext    json.RawMessage `json:"ext,omitempty"`
}
