package openrtb

var sptr = func(s string) *string { return &s }
var iptr = func(i int) *int { return &i }

var testFixtures = struct {
	simpleBanner         []byte
	expandableCreative   []byte
	simpleResponse       []byte
	simpleNativeResponse []byte
	simpleLink           Link
	simpleImg            ResponseImg
	simpleTitle          ResponseTitle
	simpleData           ResponseData
	installData          ResponseData
	fullLink             Link
}{
	simpleBanner: []byte(`
{
  "id":"1234534625254",
  "at":2,
  "tmax":120,
  "imp":[
    {
      "id":"1",
      "banner":{
        "w":300,
        "h":250,
        "pos":1,
        "battr":[13]
      }
    }
  ],
  "badv":["company1.com","company2.com"],
  "site":{
    "id":"234563",
    "name":"Site ABCD",
    "domain":"siteabcd.com",
    "cat":["IAB2-1", "IAB2-2"],
    "privacypolicy":1,
    "page":"http://siteabcd.com/page.htm",
    "ref":"http://referringsite.com/referringpage.htm",
    "publisher":{
      "id":"pub12345",
      "name":"Publisher A"
    },
    "content":{
      "keywords":["keyword a","keyword b","keyword c"]
    }
  },
  "device":{
    "ip":"64.124.253.1",
    "ua":"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
    "os":"OS X",
    "flashver":"10.1",
    "js":1
  },
  "user":{
    "id":"45asdf987656789adfad4678rew656789",
    "buyeruid":"5df678asd8987656asdf78987654"
  }
}
`),
	expandableCreative: []byte(`
{
  "id":"1234567893",
  "at":2,
  "tmax":120,
  "imp":[
    {
      "id":"1",
      "iframebuster":[
        "vendor1.com",
        "vendor2.com"
      ],
      "banner":{
        "w":300,
        "h":250,
        "wmin":300,
        "hmin":250,
        "pos":1,
        "battr":[
          13
        ],
        "expdir":[
          2,
          4
        ]
      }
    }
  ],
  "site":{
    "id":"1345135123",
    "name":"Site ABCD",
    "domain":"siteabcd.com",
    "sitecat":[
      "IAB2-1",
      "IAB2-2"
    ],
    "page":"http://siteabcd.com/page.htm",
    "ref":"http://referringsite.com/referringpage.htm",
    "privacypolicy":1,
    "publisher":{
      "id":"pub12345",
      "name":"Publisher A"
    },
    "content":{
      "keyword":[
        "keyword1",
        "keyword2",
        "keyword3"
      ]
    }
  },
  "device":{
    "ip":"64.124.253.1",
    "ua":"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
    "os":"OS X",
    "flashver":"10.1",
    "js":1
  },
  "user":{
    "id":"456789876567897654678987656789",
    "buyeruid":"545678765467876567898765678987654",
    "data":[
      {
        "id":"6",
        "name":"Data Provider 1",
        "segment":[
          {
            "id":"12341318394918",
            "name":"auto intenders"
          },
          {
            "id":"1234131839491234",
            "name":"auto enthusiasts"
          },
          {
            "id":"23423424",
            "name":"data-provider1-age",
            "value":"30-40"
          }
        ]
      }
    ]
  }
}`),
	simpleResponse: []byte(`
{
    "id": "BID-4-ZIMP-4b309eae-504a-4252-a8a8-4c8ceee9791a",
    "seatbid": [
        {
            "bid": [
                {
                    "id": "32a69c6ba388f110487f9d1e63f77b22d86e916b",
                    "impid": "32a69c6ba388f110487f9d1e63f77b22d86e916b",
                    "price": 0.065445,
                    "adid": "529833ce55314b19e8796116",
                    "nurl": "http://ads.com/win/529833ce55314b19e8796116?won=${auction_price}",
                    "adm": "<iframe src=\"foo.bar\"/>" ,
                    "adomain": [],
                    "cid": "529833ce55314b19e8796116",
                    "crid": "529833ce55314b19e8796116_1385706446",
                    "attr": []
                }
            ],
            "seat": "772"
        }
    ],
    "cur": "USD"
}
`),
	simpleNativeResponse: []byte(`
{
    "native": {
        "ver": "1",
        "link": {
            "url": "deeplink://deeplink/url/into/app",
            "fallback": "http: //i.am.a/URL",
            "clicktrackers": [
                "http: //a.com/a",
                "http: //b.com/b"
            ]
        },
        "imptrackers": [
            "http: //a.com/a",
            "http: //b.com/b"
        ],
        "assets": [
            {
                "id": 1,
                "title": {
                    "text": "InstallBOA"
                },
                "link": {
                    "url": "http: //i.am.a/URL"
                }
            },
            {
                "id": 2,
                "data": {
                    "value": "5"
                }
            },
            {
                "id": 3,
                "img": {
                    "url": "http: //cdn.mobad.com/ad.png",
                    "w": 64,
                    "h": 64
                }
            },
            {
                "id": 4,
                "data": {
                    "value": "Install"
                },
                "link": {
                    "url": "http: //i.am.a/URL"
                }
            }
        ]
    }
}
`),
	simpleLink:  Link{Url: sptr("http: //i.am.a/URL")},
	simpleImg:   ResponseImg{Url: sptr("http: //cdn.mobad.com/ad.png"), W: iptr(64), H: iptr(64)},
	simpleTitle: ResponseTitle{Text: sptr("InstallBOA")},
	simpleData:  ResponseData{Value: sptr("5")},
	installData: ResponseData{Value: sptr("Install")},
	fullLink:    Link{Url: sptr("deeplink://deeplink/url/into/app"), Clicktrackers: []string{"http: //a.com/a", "http: //b.com/b"}, Fallback: sptr("http: //i.am.a/URL")},
}
