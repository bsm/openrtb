package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Device", func() {
	var subject *Device

	BeforeEach(func() {
		subject = new(Device)
	})

	It("should have accessors", func() {
		Expect(subject.IsDnt()).To(BeFalse())
		Expect(subject.IsJs()).To(BeFalse())
		Expect(subject.ConnectionType()).To(Equal(0))
		Expect(subject.DeviceType()).To(Equal(0))
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(*subject.Dnt).To(Equal(0))
		Expect(*subject.Js).To(Equal(0))
		Expect(*subject.Connectiontype).To(Equal(CONN_TYPE_UNKNOWN))
		Expect(*subject.Connectiontype).To(Equal(DEVICE_TYPE_UNKNOWN))
	})
})
