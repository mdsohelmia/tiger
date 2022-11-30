package models

type Plugin struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	IconLink    string `json:"icon_link"`
	BannerLink  string `json:"banner_link"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Provider    string `json:"provider"`
	IsPaid      bool   `json:"is_paid"`
}

type LibraryPlugin struct {
	LibraryId string `json:"library_id"`
	PluginId  string `json:"plugin_id"`
	Enabled   bool   `json:"enabled"`
}

func (plugin *Plugin) GetPlugin() {}

/*

	{
		"name" :"New Lib",
		"plugins":[
			{
				"name" : "Video Importer"
				"is_paid":true
				"enabled":true
			}
		]
	}

*/
