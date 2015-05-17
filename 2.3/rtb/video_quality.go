package rtb

// 5.11 Video Quality
//
// The following table lists the options for the video quality. These values are defined by the IAB –
// http://www.iab.net/media/file/long-form-video-final.pdf.
const (
	VideoQualityUnknown                uint8 = 0 // 0 Unknown
	VideoQualityProfessionallyProduced uint8 = 1 // 1 Professionally Produced
	VideoQualityProsumer               uint8 = 2 // 2 Prosumer
	VideoQualityUserGenerated          uint8 = 3 // 3 User Generated (UGC)
)
