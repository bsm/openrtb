package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRequest_BlankUnvalidated(t *testing.T) {
	req, err := ParseRequest([]byte("{}"), false)
	assert.Nil(t, err)
	assert.IsType(t, &Request{}, req)
}

func TestParseRequest_BlankValidated(t *testing.T) {
	req, err := ParseRequest([]byte("{}"), true)
	assert.Nil(t, req)
	assert.Equal(t, err.Error(), "openrtb parse: request ID missing")
}

func TestParseRequest_SimpleBanner(t *testing.T) {
	req, err := ParseRequest(simpleBanner, true)
	assert.Nil(t, err)
	assert.IsType(t, &Request{}, req)

	assert.Equal(t, *req.At, 2)
	assert.Equal(t, *req.Id, "1234534625254")
	assert.Equal(t, len(req.Badv), 2)
	assert.Equal(t, len(req.Bcat), 0)

	assert.Equal(t, len(req.Imp), 1)
	assert.Equal(t, *req.Imp[0].Banner.W, 300)
	assert.Equal(t, *req.Imp[0].Banner.H, 250)

	assert.Equal(t, *req.Site.Name, "Site ABCD")
	assert.Equal(t, *req.Site.Publisher.Name, "Publisher A")
	assert.Equal(t, *req.Device.Ip, "64.124.253.1")
	assert.Equal(t, *req.User.Buyeruid, "5df678asd8987656asdf78987654")
}

func TestParseExpandableCreative(t *testing.T) {
	req, err := ParseRequest(expandableCreative, true)
	assert.Nil(t, err)
	assert.IsType(t, &Request{}, req)

	assert.Equal(t, *req.At, 2)
	assert.Equal(t, *req.Tmax, 120)
	assert.Equal(t, req.Imp[0].Banner.Expdir, []int{2, 4})
	assert.Equal(t, *req.Site.Privacypolicy, 1)
	assert.Equal(t, *req.Device.Flashver, "10.1")
	assert.Equal(t, len(req.User.Data), 1)
	assert.Equal(t, *req.User.Data[0].Id, "6")
	assert.Equal(t, len(req.User.Data[0].Segment), 3)
	assert.Equal(t, *req.User.Data[0].Segment[2].Id, "23423424")
}

var simpleBanner []byte = []byte(`
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
`)

var expandableCreative []byte = []byte(`
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
}`)
