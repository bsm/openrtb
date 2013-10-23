package openrtb

// The "device" object provides information pertaining to the device including its hardware,
// platform, location, and carrier. This device can refer to a mobile handset, a desktop computer,
// set top box or other digital device.
type Device struct {
	Dnt            *int    // "1": Do not track
	Ua             *string // User agent
	Ip             *string // IPv4
	Geo            *Geo
	Didsha1        *string // SHA1 hashed device ID
	Didmd5         *string // MD5 hashed device ID
	Dpidsha1       *string // SHA1 hashed platform device ID
	Dpidmd5        *string // MD5 hashed platform device ID
	Ipv6           *string // IPv6
	Carrier        *string // Carrier or ISP derived from the IP address
	Language       *string // Browser language
	Make           *string // Device make
	Model          *string // Device model
	Os             *string // Device OS
	Osv            *string // Device OS version
	Js             *int    // Javascript status ("0": Disabled, "1": Enabled)
	Connectiontype *int
	Devicetype     *int
	Flashver       *string // Flash version
	Ext            map[string]string
}

// Returns the DNT status, with default fallback
func (d *Device) IsDnt() bool {
	if d.Dnt != nil {
		return *d.Dnt == 1
	}
	return false
}

// Returns the JS status, with default fallback
func (d *Device) IsJs() bool {
	if d.Js != nil {
		return *d.Js == 1
	}
	return false
}

// Returns the connection type, with default fallback
func (d *Device) ConnectionType() int {
	if d.Connectiontype != nil {
		return *d.Connectiontype
	}
	return CONN_TYPE_UNKNOWN
}

// Returns the connection type, with default fallback
func (d *Device) DeviceType() int {
	if d.Devicetype != nil {
		return *d.Devicetype
	}
	return DEVICE_TYPE_UNKNOWN
}

// Applies defaults
func (d *Device) WithDefaults() *Device {
	if d.Dnt == nil {
		d.Dnt = new(int)
		*d.Dnt = 0
	}
	if d.Js == nil {
		d.Js = new(int)
		*d.Js = 0
	}
	if d.Connectiontype == nil {
		d.Connectiontype = new(int)
		*d.Connectiontype = CONN_TYPE_UNKNOWN
	}
	if d.Devicetype == nil {
		d.Devicetype = new(int)
		*d.Devicetype = DEVICE_TYPE_UNKNOWN
	}
	return d
}
