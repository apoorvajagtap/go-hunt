package config

// regexs = {
//     "albums": r'href=\"\.\/albumarchive\/\d*?\/album\/(.*?)\" jsaction.*?>(?:<.*?>){5}(.*?)<\/div><.*?>(\d*?) ',
//     "photos": r'\],\"(https:\/\/lh\d\.googleusercontent\.com\/.*?)\",\[\"\d{21}\"(?:.*?,){16}\"(.*?)\"',
//     "review_loc_by_id": r'{}\",.*?\[\[null,null,(.*?),(.*?)\]'
// }

var headers = map[string]string {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; rv:68.0) Gecko/20100101 Firefox/68.0",
    "Connection": "Keep-Alive",
}

var headless         = true // if true, it doesn't show the browser while scraping GMaps reviews
var ytb_hunt_always  = true // if true, search the Youtube channel everytime
var gmaps_radius     = 30 // in km. The radius distance to create groups of gmaps reviews.
var gdocs_public_doc = "1jaEEHZL32t1RUN5WuZEnFpqiEPf_APYKrRBG9LhLdvE"  // The public Google Doc to use it as an endpoint, to use Google's Search.
var data_path        = "resources/data.txt"
var browser_waiting_timeout = 120

// Profile pictures options
var write_profile_pic = true
var profile_pics_dir = "profile_pics"

// Cookies
// if true, it will uses the Google Account cookies to request the services,
// and gonna be able to read your personal informations
var gmaps_cookies = false
var calendar_cookies = false
var default_consent_cookie = "YES+FR.fr+V10+BX"


func GetDataPath() string {
	return data_path
}