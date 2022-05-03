package views

type SlashCommand struct {
	Token          string `schema:"token"`
	TeamID         string `schema:"team_id"`
	TeamDomain     string `schema:"team_domain"`
	EnterpriseID   string `schema:"enterprise_id,omitempty"`
	EnterpriseName string `schema:"enterprise_name,omitempty"`
	ChannelID      string `schema:"channel_id"`
	ChannelName    string `schema:"channel_name"`
	UserID         string `schema:"user_id"`
	UserName       string `schema:"user_name"`
	Command        string `schema:"command"`
	Text           string `schema:"text"`
	ResponseURL    string `schema:"response_url"`
	TriggerID      string `schema:"trigger_id"`
	APIAppID       string `schema:"api_app_id"`
}
