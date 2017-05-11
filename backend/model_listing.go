package main

import (
	"gopkg.in/mgo.v2/bson"
)

type listing struct {
	ID                  bson.ObjectId `bson:"_id" json:"id"`
	ListingID           int           `json:"listing_id"`
	State               string        `json:"state"`
	UserID              int           `json:"user_id"`
	CategoryID          int           `json:"category_id"`
	Title               string        `json:"title"`
	Description         string        `json:"description"`
	CreationTsz         int           `json:"creation_tsz"`
	EndingTsz           int           `json:"ending_tsz"`
	OriginalCreationTsz int           `json:"original_creation_tsz"`
	LastModifiedTsz     int           `json:"last_modified_tsz"`
	Price               string        `json:"price"`
	CurrencyCode        string        `json:"currency_code"`
	Quantity            int           `json:"quantity"`
	ShopSectionID       int           `json:"shop_section_id"`
	StateTsz            int           `json:"state_tsz"`
	URL                 string        `json:"url"`
	Views               int           `json:"views"`
	NumFavorers         int           `json:"num_favorers"`
	ProcessingMin       int           `json:"processing_min"`
	ProcessingMax       int           `json:"processing_max"`
	WhoMade             string        `json:"who_made"`
	IsSupply            string        `json:"is_supply"`
	WhenMade            string        `json:"when_made"`
	ItemDimensionsUnit  string        `json:"item_dimensions_unit"`
	IsPrivate           bool          `json:"is_private"`
	FileData            string        `json:"file_data"`
	Language            string        `json:"language"`
	HasVariations       bool          `json:"has_variations"`
	TaxonomyID          int           `json:"taxonomy_id"`
}

type offering struct {
	OfferingID int `json:"offering_id"`
	Price      struct {
		Amount                 int    `json:"amount"`
		Divisor                int    `json:"divisor"`
		CurrencyCode           string `json:"currency_code"`
		CurrencyFormattedShort string `json:"currency_formatted_short"`
		CurrencyFormattedLong  string `json:"currency_formatted_long"`
		CurrencyFormattedRaw   string `json:"currency_formatted_raw"`
	} `json:"price"`
	Quantity int `json:"quantity"`
}

type product struct {
	ID             bson.ObjectId `bson:"_id"`
	Reference      bson.ObjectId `bson:"reference"`
	ProductID      int           `json:"product_id"`
	PropertyValues []struct {
		PropertyID   int      `json:"property_id"`
		PropertyName string   `json:"property_name"`
		ScaleID      int      `json:"scale_id"`
		ScaleName    string   `json:"scale_name"`
		ValueIds     []int64  `json:"value_ids"`
		Values       []string `json:"values"`
	} `json:"property_values"`
	Offerings []offering `json:"offerings"`
}

// getListingInventoryResponse represents the response from etsy/v2/listings/:listing_id/inventory
type getListingInventoryResponse struct {
	Count   int `json:"count"`
	Results struct {
		Products []product `json:"products"`
	} `json:"results"`
	Type       string `json:"type"`
	Pagination struct {
		EffectiveLimit  int `json:"effective_limit"`
		EffectiveOffset int `json:"effective_offset"`
		NextOffset      int `json:"next_offset"`
		EffectivePage   int `json:"effective_page"`
		NextPage        int `json:"next_page"`
	} `json:"pagination"`
}

// getActiveListingsResponse is the etsy response for active listings.
type getActiveListingsResponse struct {
	Count      int       `json:"count"`
	Results    []listing `json:"results"`
	Type       string    `json:"type"`
	Pagination struct {
		EffectiveLimit  int `json:"effective_limit"`
		EffectiveOffset int `json:"effective_offset"`
		NextOffset      int `json:"next_offset"`
		EffectivePage   int `json:"effective_page"`
		NextPage        int `json:"next_page"`
	} `json:"pagination"`
}

// ListingImage is represents a listings image data.
type listingImage struct {
	ID              bson.ObjectId `bson:"_id"`
	Reference       bson.ObjectId `bson:"reference"`
	ListingImageID  int           `json:"listing_image_id"`
	HexCode         string        `json:"hex_code"`
	Red             int           `json:"red"`
	Green           int           `json:"green"`
	Blue            int           `json:"blue"`
	Hue             int           `json:"hue"`
	Saturation      int           `json:"saturation"`
	Brightness      int           `json:"brightness"`
	IsBlackAndWhite bool          `json:"is_black_and_white"`
	CreationTsz     int           `json:"creation_tsz"`
	ListingID       int           `json:"listing_id"`
	Rank            int           `json:"rank"`
	URL75X75        string        `json:"url_75x75"`
	URL170X135      string        `json:"url_170x135"`
	URL570XN        string        `json:"url_570xN"`
	URLFullxfull    string        `json:"url_fullxfull"`
	FullHeight      int           `json:"full_height"`
	FullWidth       int           `json:"full_width"`
}

// GetListingImagesResponse is the response from etsy/v2/listings/:listing_id/images
type getListingImagesResponse struct {
	Count      int            `json:"count"`
	Results    []listingImage `json:"results"`
	Type       string         `json:"type"`
	Pagination struct {
		EffectiveLimit  int `json:"effective_limit"`
		EffectiveOffset int `json:"effective_offset"`
		NextOffset      int `json:"next_offset"`
		EffectivePage   int `json:"effective_page"`
		NextPage        int `json:"next_page"`
	} `json:"pagination"`
}
