package main

// GetActiveListingResponse is the etsy response for active listings.
// Refactor, we don't need all of this.
type GetActiveListingResponse struct {
	Count   int `json:"count"`
	Results []struct {
		ListingID           int      `json:"listing_id"`
		State               string   `json:"state"`
		UserID              int      `json:"user_id"`
		CategoryID          int      `json:"category_id"`
		Title               string   `json:"title"`
		Description         string   `json:"description"`
		CreationTsz         int      `json:"creation_tsz"`
		EndingTsz           int      `json:"ending_tsz"`
		OriginalCreationTsz int      `json:"original_creation_tsz"`
		LastModifiedTsz     int      `json:"last_modified_tsz"`
		Price               string   `json:"price"`
		CurrencyCode        string   `json:"currency_code"`
		Quantity            int      `json:"quantity"`
		Tags                []string `json:"tags"`
		CategoryPath        []string `json:"category_path"`
		CategoryPathIds     []int    `json:"category_path_ids"`
		Materials           []string `json:"materials"`
		ShopSectionID       int      `json:"shop_section_id"`
		StateTsz            int      `json:"state_tsz"`
		URL                 string   `json:"url"`
		Views               int      `json:"views"`
		NumFavorers         int      `json:"num_favorers"`
		ProcessingMin       int      `json:"processing_min"`
		ProcessingMax       int      `json:"processing_max"`
		WhoMade             string   `json:"who_made"`
		IsSupply            string   `json:"is_supply"`
		WhenMade            string   `json:"when_made"`
		ItemDimensionsUnit  string   `json:"item_dimensions_unit"`
		IsPrivate           bool     `json:"is_private"`
		NonTaxable          bool     `json:"non_taxable"`
		IsCustomizable      bool     `json:"is_customizable"`
		IsDigital           bool     `json:"is_digital"`
		FileData            string   `json:"file_data"`
		Language            string   `json:"language"`
		HasVariations       bool     `json:"has_variations"`
		TaxonomyID          int      `json:"taxonomy_id"`
		TaxonomyPath        []string `json:"taxonomy_path"`
		UsedManufacturer    bool     `json:"used_manufacturer"`
	} `json:"results"`
	Params struct {
		Limit             string `json:"limit"`
		Offset            string `json:"offset"`
		ShopID            string `json:"shop_id"`
		SortOn            string `json:"sort_on"`
		SortOrder         string `json:"sort_order"`
		ColorAccuracy     int    `json:"color_accuracy"`
		TranslateKeywords string `json:"translate_keywords"`
		IncludePrivate    int    `json:"include_private"`
	} `json:"params"`
	Type       string `json:"type"`
	Pagination struct {
		EffectiveLimit  int `json:"effective_limit"`
		EffectiveOffset int `json:"effective_offset"`
		NextOffset      int `json:"next_offset"`
		EffectivePage   int `json:"effective_page"`
		NextPage        int `json:"next_page"`
	} `json:"pagination"`
}
