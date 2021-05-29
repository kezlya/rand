package rand

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var  FlickrAPI  = "https://api.flickr.com/services/feeds/photos_public.gne"

func Image(keyword string, order uint) (string ,error){
	log.Println("Random from Flicker: ", keyword)
	type flickerImage struct {
		Media struct {
			Url string `json:"m"`
		} `json:"media"`
	}
	type flickerFeed struct {
		Images []flickerImage `json:"items"`
	}
	var fImg flickerFeed

	resp, err := http.Get(FlickrAPI+"?format=json&tags=" + keyword)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	_ ,err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "",err
	}

	respStr := strings.TrimPrefix(string(buf.Bytes()), "jsonFlickrFeed(")
	respStr = strings.TrimSuffix(respStr, ")")

	err = json.Unmarshal([]byte(respStr), &fImg)
	if err != nil {
		return "", err
	}

	return fImg.Images[order].Media.Url, nil
}

func Noun() string {
	nouns := []string{"world", "house", "place", "group", "party", "money", "point", "state", "night", "water", "thing", "order", "power", "times", "court", "level", "child", "south", "staff", "woman", "north", "sense", "death", "range", "table", "trade", "study", "other", "price", "class", "union", "value", "paper", "right", "voice", "stage", "light", "march", "board", "month", "music", "field", "award", "issue", "basis", "front", "heart", "force", "model", "means", "space", "peter", "hotel", "floor", "style", "miles", "needs", "sales", "event", "press", "doubt", "blood", "sound", "title", "glass", "earth", "river", "whole", "piece", "mouth", "radio", "peace", "start", "share", "works", "truth", "media", "smith", "stone", "queen", "stock", "horse", "plant", "visit", "scale", "image", "trust", "chair", "cause", "speed", "crime", "pound", "henry", "match", "scene", "stuff", "japan", "claim", "video", "trial", "time", "year", "work", "life", "part", "case", "fact", "area", "head", "hand", "john", "side", "home", "days", "week", "room", "road", "form", "face", "sort", "body", "name", "book", "view", "door", "line", "city", "kind", "idea", "west", "mind", "land", "care", "back", "rate", "word", "food", "team", "role", "town", "bank", "need", "east", "type", "date", "wife", "club", "lord", "king", "cost", "ways", "girl", "game", "love", "news", "rest", "hair", "bill", "feet", "fire", "size", "term", "plan", "hall", "list", "loss", "wall", "paul", "army", "unit", "park", "hour", "test", "arms", "look", "deal", "help", "page", "risk", "fish", "film", "shop", "site", "mark", "lady", "task", "sale", "lack", "post", "firm", "show", "baby", "base", "miss", "past", "cash", "rule", "turn", "duty", "ball", "way", "day", "man", "end", "use", "lot", "war", "car", "law", "bit", "god", "job", "act", "age", "sir", "air", "tax", "art", "may", "bed", "top", "boy", "son", "sea", "cup", "sun", "set", "new", "oil", "eye", "arm", "box", "sex", "mum", "aid", "tea", "dog", "bar", "gas", "dad", "pp.", "tom", "bus", "bag", "leg", "aim", "key", "sky", "row", "run", "pay", "van", "bob", "sum", "ice", "joe", "ref", "map", "cat", "guy", "lee", "pub", "gun", "gap", "fun", "no.", "bay", "ken", "bid", "hat", "fee", "joy", "cut", "sam", "ben", "ear", "san", "net", "red", "egg", "ban", "vat", "fox", "win", "pop", "era", "ray", "pen", "tip", "tin", "ann", "pot", "pat", "cap", "lad", "mud", "pan", "tie", "fat", "kit"}
	return nouns[Int(len(nouns))]
}

func Adj() string {
	adjectives := []string{"little", "social", "second", "public", "likely", "better", "common", "single", "former", "recent", "strong", "higher", "simple", "modern", "normal", "soviet", "direct", "useful", "german", "future", "senior", "annual", "latter", "active", "middle", "united", "sexual", "actual", "latest", "famous", "formal", "proper", "unable", "lovely", "fourth", "female", "mental", "double", "afraid", "bright", "bloody", "narrow", "entire", "severe", "unique", "guilty", "yellow", "longer", "golden", "sudden", "living", "global", "silent", "bigger", "secret", "visual", "wooden", "stupid", "stable", "honest", "slight", "remote", "gentle", "junior", "broken", "smooth", "pretty", "fellow", "square", "steady", "bitter", "ethnic", "weekly", "random", "modest", "asleep", "clever", "liable", "ruling", "mutual", "nearby", "urgent", "marked", "superb", "strict", "closer", "marine", "retail", "closed", "unfair", "flying", "hungry", "subtle", "secure", "decent", "bottom", "lesser", "casual", "finest", "lonely", "other", "first", "great", "local", "small", "right", "large", "young", "early", "major", "clear", "black", "whole", "third", "white", "short", "human", "royal", "wrong", "legal", "final", "close", "total", "prime", "happy", "lower", "sorry", "basic", "aware", "ready", "green", "heavy", "extra", "civil", "later", "chief", "usual", "front", "fresh", "joint", "alone", "worse", "rural", "light", "equal", "quiet", "quick", "daily", "urban", "upper", "moral", "vital", "empty", "brief", "broad", "worst", "clean", "ideal", "minor", "thick", "inner", "grand", "prior", "roman", "funny", "brown", "sharp", "alive", "angry", "lucky", "cheap", "tired", "armed", "fixed", "red", "rapid", "fifth", "solid", "rough", "sweet", "given", "tough", "awful", "proud", "plain", "still", "round", "silly", "above", "blind", "dirty", "sixth", "loose", "level", "outer", "gross", "acute", "valid", "tight", "exact", "world", "house", "place", "group", "party", "money", "point", "state", "night", "water", "thing", "order", "power", "times", "court", "level", "child", "south", "staff", "woman", "north", "sense", "death", "range", "table", "trade", "study", "other", "price", "class", "union", "value", "paper", "right", "voice", "stage", "light", "march", "board", "month", "music", "field", "award", "issue", "basis", "front", "heart", "force", "model", "means", "space", "peter", "hotel", "floor", "style", "miles", "needs", "sales", "event", "press", "doubt", "blood", "sound", "title", "glass", "earth", "river", "whole", "piece", "mouth", "radio", "peace", "start", "share", "works", "truth", "media", "smith", "stone", "queen", "stock", "horse", "plant", "visit", "scale", "image", "trust", "chair", "cause", "speed", "crime", "pound", "henry", "match", "scene", "stuff", "japan", "claim", "video", "trial"}
	return adjectives[Int(len(adjectives))]
}

func String(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[Int(len(charset))]
	}
	return string(b)
}

func Color() string {
	colors := []string{"0048BA","B0BF1A","7CB9E8","C0E8D5","B284BE","72A0C1","EDEAE0","F0F8FF","C46210","EFDECD","E52B50","9F2B68","F19CBB","AB274F","D3212D","3B7A57","FFBF00","FF7E00","9966CC","3DDC84","CD9575","665D1E","915C83","841B2D","FAEBD7","008000","8DB600","FBCEB1","00FFFF","7FFFD4","D0FF14","4B5320","8F9779","E9D66B","B2BEB5","87A96B","FF9966","A52A2A","FDEE00","568203","007FFF","F0FFFF","89CFF0","A1CAF1","F4C2C2","FEFEFA","FF91AF","FAE7B5","DA1884","7C0A02","848482","BCD4E6","9F8170","F5F5DC","2E5894","9C2542","FFE4C4","3D2B1F","967117","CAE00D","BFFF00","FE6F5E","BF4F51","000000","3D0C02","1B1811","3B2F2F","54626F","3B3C36","BFAFB2","FFEBCD","A57164","318CE7","ACE5EE","FAF0BE","660000","0000FF","1F75FE","0093AF","0087BD","0018A8","333399","0247FE","A2A2D0","6699CC","0D98BA","064E40","5DADEC","126180","8A2BE2","7366BD","4D1A7F","5072A7","3C69E7","DE5D83","79443B","E3DAC9","006A4E","87413F","CB4154","66FF00","D891EF","C32148","1974D2","FFAA1D","FF55A3","FB607F","004225","CD7F32","88540B","AF6E4D","1B4D3E","7BB661","FFC680","800020","DEB887","A17A74","CC5500","E97451","8A3324","BD33A4","702963","536872","5F9EA0","A9B2C3","91A3B0","006B3C","ED872D","E30022","FFF600","A67B5B","4B3621","A3C1AD","C19A6B","EFBBCC","FFFF99","FFEF00","FF0800","E4717A","00BFFF","592720","C41E3A","00CC99","960018","D70040","FFA6C9","B31B1B","56A0D3","ED9121","00563F","703642","C95A49","ACE1AF","007BA7","2F847C","B2FFFF","246BCE","DE3163","007BA7","2A52BE","6D9BC3","1DACD6","007AA5","E03C31","F7E7CE","F1DDCF","36454F","232B2B","E68FAC","DFFF00","7FFF00","FFB7C5","954535","E23D28","DE6FA1","A8516E","AA381E","856088","FFB200","7B3F00","D2691E","58111A","FFA700","98817B","E34234","CD607E","E4D00A","9FA91F","7F1734","0047AB","D2691E","6F4E37","B9D9EB","F88379","8C92AC","B87333","DA8A67","AD6F69","CB6D51","996666","FF3800","FF7F50","F88379","893F45","FBEC5D","B31B1B","6495ED","FFF8DC","2E2D88","FFF8E7","81613C","FFBCD9","FFFDD0","DC143C","9E1B32","A7D8DE","F5F5F5","00FFFF","00B7EB","58427C","FFD300","F56FA1","666699","654321","5D3954","26428B","008B8B","536878","B8860B","013220","006400","1A2421","BDB76B","483C32","534B4F","543D37","8B008B","4A5D23","556B2F","FF8C00","9932CC","03C03C","301934","8B0000","E9967A","8FBC8F","3C1414","8CBED6","483D8B","2F4F4F","177245","00CED1","9400D3","00703C","555555","DA3287","FAD6A5","B94E48","004B49","FF1493","FF9933","00BFFF","4A646C","7E5E60","1560BD","2243B6","C19A6B","EDC9AF","696969","1E90FF","D71868","967117","00009C","EFDFBB","E1A95F","555D50","C2B280","1B1B1B","614051","F0EAD6","1034A6","16161D","7DF9FF","00FF00","6F00FF","CCFF00","BF00FF","8F00FF","50C878","6C3082","1B4D3E","B48395","AB4B52","CC474B","563C5C","00FF40","96C8A2","C19A6B","801818","B53389","DE5285","F400A1","E5AA70","4D5D53","4F7942","6C541E","FF5470","B22222","CE2029","E95C4B","E25822","EEDC82","A2006D","FFFAF0","15F4EE","5FA777","014421","228B22","A67B5B","856D4D","0072BB","FD3F92","86608E","9EFD38","D473D4","FD6C9E","C72C48","F64A8A","77B5FE","8806CE","E936A7","FF00FF","C154C1","CC397B","C74375","E48400","87421F"}
return colors[Int(len(colors))]
}

func Int(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}