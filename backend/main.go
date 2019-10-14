package main

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"strconv"
)

type Quote struct {
	Text   string `json:"text"`
	Tag    string `json:"tag"`
	Author string `json:"author"`
}

type Quotes []Quote

const (
	LIFE         = 1100
	LIKE         = 1139
	DIE          = 1106
	MARRIAGE     = 1109
	HAPPY        = 1108
	MONEY        = 1111
	LOVE         = 1107
	FRIEND       = 1104
	WOMAN        = 1113
	TIME         = 1103
	BROKEN_HEAET = 1116
	SELF         = 1102
	YOUTH        = 1115
	MAN          = 1112
	SOCIETY      = 1114
	LONELY       = 1118
	DESTINY      = 1105
	TRUTH        = 1133
	MISFORTUNE   = 1117
	FREEDOM      = 1129
	EDUCATION    = 1101
	POSITIVE     = 1128
	MANAGEMENT   = 1147
	CHALLENGE    = 1137
	EFFORT       = 1126
	FUTURE       = 1120
)

var (
	tags = map[int]string{
		LIFE:         "人生",
		LIKE:         "恋愛",
		DIE:          "死",
		MARRIAGE:     "結婚",
		HAPPY:        "幸福",
		MONEY:        "金",
		LOVE:         "愛",
		FRIEND:       "友達",
		WOMAN:        "女",
		TIME:         "時間",
		BROKEN_HEAET: "失恋",
		SELF:         "自己",
		YOUTH:        "青春",
		MAN:          "男",
		SOCIETY:      "処世",
		LONELY:       "孤独",
		DESTINY:      "運命",
		TRUTH:        "真実",
		MISFORTUNE:   "不幸",
		FREEDOM:      "自由",
		EDUCATION:    "教育",
		POSITIVE:     "ポジティブ",
		MANAGEMENT:   "マネジメント",
		CHALLENGE:    "挑戦",
		EFFORT:       "努力",
		FUTURE:       "未来",
	}
)

func getQuotes(tagId int, lastPage int) Quotes {
	var quotes Quotes
	for i := 1; i <= lastPage; i++ {
		doc, err := goquery.NewDocument("http://www.meigensyu.com/tags/view/" + strconv.Itoa(tagId) + "/page" + strconv.Itoa(i) + ".html")
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".meigenbox").Each(func(_ int, selection *goquery.Selection) {
			var text = selection.Find(".text").Text()
			var author = selection.Find(".link li").First().Text()
			var tag = tags[tagId]

			quotes = append(quotes, Quote{
				Text:   text,
				Tag:    tag,
				Author: author,
			})
		})
	}

	return quotes
}

func outputJson(fileName string, tagId int, lastPage int) {
	quotes := getQuotes(tagId, lastPage)
	jsonBytes, err := json.Marshal(quotes)
	if err != nil {
		log.Fatal(err)
	}

	out := new(bytes.Buffer)
	err = json.Indent(out, jsonBytes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(out)
	err = ioutil.WriteFile("../client/static/data/"+fileName, buf, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 人生
	outputJson("life.json", LIFE, 13)
	// 恋愛
	outputJson("like.json", LIKE, 12)
	// 死
	outputJson("die.json", DIE, 7)
	// 結婚
	outputJson("marriage.json", MARRIAGE, 7)
	// 幸福
	outputJson("happy.json", HAPPY, 6)
	// 金
	outputJson("money.json", MONEY, 5)
	// 愛
	outputJson("happy.json", LOVE, 5)
	// 友情
	outputJson("friend.json", FRIEND, 4)
	// 女
	outputJson("woman.json", WOMAN, 4)
	// 時間
	outputJson("time.json", TIME, 3)
	// 失恋
	outputJson("broken_heart.json", BROKEN_HEAET, 3)
	// 自己
	outputJson("self.json", SELF, 3)
	// 青春
	outputJson("youth.json", YOUTH, 12)
	// 男
	outputJson("man.json", MAN, 3)
	// 処世
	outputJson("society.json", SOCIETY, 3)
	// 孤独
	outputJson("lonely.json", LONELY, 2)
	// 運命
	outputJson("destiny.json", DESTINY, 2)
	// 真実
	outputJson("truth.json", TRUTH, 2)
	// 不幸
	outputJson("misfortune.json", MISFORTUNE, 2)
	// 自由
	outputJson("freedom.json", FREEDOM, 2)
	// 教育
	outputJson("education.json", EDUCATION, 2)
	// ポジティブ
	outputJson("positive.json", POSITIVE, 1)
	// マネージメント
	outputJson("management.json", MANAGEMENT, 1)
	// 挑戦
	outputJson("challenge.json", CHALLENGE, 1)
	// 努力
	outputJson("effort.json", EFFORT, 1)
	// 未来
	outputJson("future.json", FUTURE, 1)
}
