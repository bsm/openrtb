package openrtb

import (
	"bytes"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
  var subject *Request

  BeforeEach(func() { subject = new(Request) })

  Describe("ParseRequest()", func() {
    req, err := ParseRequest(bytes.NewBuffer(expandableCreative))
    Expect(err).NotTo(HaveOccurred())
    Expect(req).To(BeAssignableToTypeOf(&Request{}))
  })

  Describe("ParseRequestBytes()", func() {
    It("should return blank request with defaults when blank", func() {
      req, err := ParseRequestBytes([]byte("{}"))
      Expect(err).NotTo(HaveOccurred())
      Expect(req).To(BeAssignableToTypeOf(&Request{}))
    })

    It("should return accordingly when with simple banner", func() {
      req, err := ParseRequestBytes(simpleBanner)
      Expect(err).NotTo(HaveOccurred())
      Expect(req).To(BeAssignableToTypeOf(&Request{}))

      Expect(*req.At).To(Equal(2))
      Expect(*req.Id).To(Equal("1234534625254"))

      Expect(len(req.Badv)).To(Equal(2))
      Expect(len(req.Bcat)).To(Equal(0))
      Expect(len(req.Imp)).To(Equal(1))

      Expect(*req.Imp[0].Banner.W).To(Equal(300))
      Expect(*req.Imp[0].Banner.H).To(Equal(250))
      Expect(*req.Site.Name).To(Equal("Site ABCD"))
      Expect(*req.Site.Publisher.Name).To(Equal("Publisher A"))
      Expect(*req.Device.Ip).To(Equal("64.124.253.1"))
      Expect(*req.User.Buyeruid).To(Equal("5df678asd8987656asdf78987654"))
    })

    It("should return accordingly when with expandable creatives", func(){
      req, err := ParseRequestBytes(expandableCreative)
      Expect(err).NotTo(HaveOccurred())
      Expect(req).To(BeAssignableToTypeOf(&Request{}))

      Expect(*req.At).To(Equal(2))
      Expect(*req.Tmax).To(Equal(120))
      Expect(req.Imp[0].Banner.Expdir).To(Equal([]int{2,4}))
      Expect(*req.Site.Privacypolicy).To(Equal(1))
      Expect(*req.Device.Flashver).To(Equal("10.1"))
      Expect(*req.User.Data[0].Id).To(Equal("6"))
      Expect(*req.User.Data[0].Segment[2].Id).To(Equal("23423424"))

      Expect(len(req.User.Data)).To(Equal(1))
      Expect(len(req.User.Data[0].Segment)).To(Equal(3))
    })
  })

  Describe("Valid()", func() {

    var (
      site        *Site
      app         *App
      impression  *Impression
      banner      *Banner
    )

    BeforeEach(func() {
      site        = new(Site)
      app         = new(App)
      impression  = new(Impression)
      banner      = new(Banner)
    })

    It("should return error messages when attributes missing", func(){
      ok, err := subject.Valid()
      Expect(err.Error()).To(Equal("openrtb parse: request ID missing"))

      subject.SetId("RAND_ID") // With ID
      ok, err = subject.Valid()
      Expect(err.Error()).To(Equal("openrtb parse: no impressions"))

      subject.SetSite(site)    // With Site
      ok, err = subject.Valid()
      Expect(err.Error()).To(Equal("openrtb parse: no impressions"))

      subject.SetApp(app)      // With App
      ok, err = subject.Valid()
      Expect(err.Error()).To(Equal("openrtb parse: no impressions"))

      // With Impression
      impression.SetId("IMPID").SetBanner(banner).WithDefaults()
      subject.Imp = []Impression{*impression}
      ok, err = subject.Valid()
      Expect(err.Error()).To(Equal("openrtb parse: request has site and app"))

      // with valid attrs
      subject.App = nil
      ok, err = subject.Valid()
      Expect(err).NotTo(HaveOccurred())
      Expect(ok).To(BeTrue())
    })

  })

  Describe("WithDefaults()", func() {

    var (
      site        *Site
      app         *App
      device      *Device
      impression  *Impression
      banner      *Banner
      video       *Video
    )

    BeforeEach(func() {
      site        = new(Site)
      app         = new(App)
      device      = new(Device)
      impression  = new(Impression)
      banner      = new(Banner)
      video       = new(Video)

      impression.SetBanner(banner).SetVideo(video)
      subject.Site   = site
      subject.App    = app
      subject.Device = device
      subject.Imp    = []Impression{*impression}
    })

    It("should return blank request with default values", func(){
       request := subject.WithDefaults()
       Expect(*request.At).To(Equal(2))
       Expect(*request.App.Privacypolicy).To(Equal(0))
       Expect(*request.App.Paid).To(Equal(0))
       Expect(*request.Site.Privacypolicy).To(Equal(0))
       Expect(*request.Device.Dnt).To(Equal(0))
       Expect(*request.Device.Js).To(Equal(0))
       Expect(*request.Device.Connectiontype).To(Equal(CONN_TYPE_UNKNOWN))
       Expect(*request.Device.Devicetype).To(Equal(DEVICE_TYPE_UNKNOWN))
       Expect(*request.Imp[0].Instl).To(Equal(0))
       Expect(*request.Imp[0].Bidfloor).To(Equal(float32(0)))
       Expect(*request.Imp[0].Bidfloorcur).To(Equal("USD"))
       Expect(*request.Imp[0].Banner.Topframe).To(Equal(0))
       Expect(*request.Imp[0].Banner.Pos).To(Equal(AD_POS_UNKNOWN))
       Expect(*request.Imp[0].Video.Sequence).To(Equal(1))
       Expect(*request.Imp[0].Video.Boxingallowed).To(Equal(1))
       Expect(*request.Imp[0].Video.Pos).To(Equal(AD_POS_UNKNOWN))
    })
  })

})

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
