package model

import (
	"gopkg.in/mgo.v2/bson"
)

//PageVars struct
type PageVars struct {
	Title  string `json:"title"`
	Tipo   string `json:"tipo"`
	Val    string `json:"val"`
	User   string `json:"user"`
	Estado string `json:"Estado"`
}
type Estado struct {
	Estado string `json:"Estado"`
	Valor  string `json:"valor"`
}
type U struct {
	User []User `json:"user"`
}
type L struct {
	Link []Link `json:"link"`
}
type H struct {
	Hist []Hist `json:"link"`
}

type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	User     string        `bson:"user" json:"user"`
	Password string        `bson:"password" json:"password"`
}

type Link struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	User  string        `bson:"user" json:"user"`
	Link  string        `bson:"link" json:"link"`
	Theme string        `bson:"theme" json:"theme"`
	Title string        `json:"title"`
}
type Hist struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	User  string        `bson:"user" json:"user"`
	Link  string        `bson:"link" json:"link"`
	Theme string        `bson:"theme" json:"theme"`
	Hours string        `json:"hours"`
	Date  string        `json:"date"`
}

/**************************************
*	Api Content
*********************************************/
type Api struct {
	Title            string `json:"title"`
	Content          string `json:"content"`
	Date_publisheds  string `json:"date_published"`
	Lead_image_urls  string `json:"lead_image_url"`
	Dek              string `json:"dek"`
	Urls             string `json:"url"`
	Dominio          string `json:"dominio"`
	Extracto         string `json:"extracto"`
	Word_counts      int    `json:"word_count"`
	Direccion        string `json:"dirección"`
	Total_pagess     int    `json:"total_pages"`
	Rendering_pagess int    `json:"rendering_pages"`
	Next_page_urls   bool   `json:"next_page_url"`
}

/**************************************
*	Search bing
*********************************************/
type SearchBing struct {
	Fav       L      `json:"fav"`
	Favoritos string `json:"favoritos"`
	Kind      string `json:"kind"`
	URL       struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		PreviousPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"previousPage"`
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title  string `json:"title"`
		Facets [][]struct {
			Label       string `json:"label"`
			Anchor      string `json:"anchor"`
			LabelWithOp string `json:"label_with_op"`
		} `json:"facets"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`

	Items []Items `json:"items"`
	Error struct {
		Errors []struct {
			Domain       string `json:"domain"`
			Reason       string `json:"reason"`
			Message      string `json:"message"`
			ExtendedHelp string `json:"extendedHelp"`
		} `json:"errors"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Items struct {
	Kind             string `json:"kind"`
	Title            string `json:"title"`
	HTMLTitle        string `json:"htmlTitle"`
	Link             string `json:"link"`
	DisplayLink      string `json:"displayLink"`
	Snippet          string `json:"snippet"`
	HTMLSnippet      string `json:"htmlSnippet"`
	CacheID          string `json:"cacheId,omitempty"`
	FormattedURL     string `json:"formattedUrl"`
	HTMLFormattedURL string `json:"htmlFormattedUrl"`

	Content     string `json:"content"`
	ContentType string `json:"contentTipe"`
	Ocurrencias int    `json:"ocurrencia"`

	Pagemap struct {
		CseThumbnail []struct {
			Width  string `json:"width"`
			Height string `json:"height"`
			Src    string `json:"src"`
		} `json:"cse_thumbnail"`
		Metatags []struct {
			Title                string `json:"title"`
			ThemeColor           string `json:"theme-color"`
			OgSiteName           string `json:"og:site_name"`
			OgURL                string `json:"og:url"`
			OgTitle              string `json:"og:title"`
			OgImage              string `json:"og:image"`
			OgDescription        string `json:"og:description"`
			AlIosAppStoreID      string `json:"al:ios:app_store_id"`
			AlIosAppName         string `json:"al:ios:app_name"`
			AlIosURL             string `json:"al:ios:url"`
			AlAndroidURL         string `json:"al:android:url"`
			AlAndroidAppName     string `json:"al:android:app_name"`
			AlAndroidPackage     string `json:"al:android:package"`
			AlWebURL             string `json:"al:web:url"`
			OgType               string `json:"og:type"`
			OgVideoURL           string `json:"og:video:url"`
			OgVideoSecureURL     string `json:"og:video:secure_url"`
			OgVideoType          string `json:"og:video:type"`
			OgVideoWidth         string `json:"og:video:width"`
			OgVideoHeight        string `json:"og:video:height"`
			OgVideoTag           string `json:"og:video:tag"`
			FbAppID              string `json:"fb:app_id"`
			TwitterCard          string `json:"twitter:card"`
			TwitterSite          string `json:"twitter:site"`
			TwitterURL           string `json:"twitter:url"`
			TwitterTitle         string `json:"twitter:title"`
			TwitterDescription   string `json:"twitter:description"`
			TwitterImage         string `json:"twitter:image"`
			TwitterAppNameIphone string `json:"twitter:app:name:iphone"`
			TwitterAppIDIphone   string `json:"twitter:app:id:iphone"`
			TwitterAppNameIpad   string `json:"twitter:app:name:ipad"`
			TwitterAppIDIpad     string `json:"twitter:app:id:ipad"`
		} `json:"metatags"`
		Videoobject []struct {
			URL              string `json:"url"`
			Name             string `json:"name"`
			Description      string `json:"description"`
			Paid             string `json:"paid"`
			Channelid        string `json:"channelid"`
			Videoid          string `json:"videoid"`
			Duration         string `json:"duration"`
			Unlisted         string `json:"unlisted"`
			Thumbnailurl     string `json:"thumbnailurl"`
			Embedurl         string `json:"embedurl"`
			Playertype       string `json:"playertype"`
			Width            string `json:"width"`
			Height           string `json:"height"`
			Isfamilyfriendly string `json:"isfamilyfriendly"`
			Regionsallowed   string `json:"regionsallowed"`
			Interactioncount string `json:"interactioncount"`
			Datepublished    string `json:"datepublished"`
			Genre            string `json:"genre"`
		} `json:"videoobject"`
		Imageobject []struct {
			URL    string `json:"url"`
			Width  string `json:"width"`
			Height string `json:"height"`
		} `json:"imageobject"`
		Person []struct {
			URL string `json:"url"`
		} `json:"person"`
		CseImage []struct {
			Src string `json:"src"`
		} `json:"cse_image"`
	} `json:"pagemap"`
}
