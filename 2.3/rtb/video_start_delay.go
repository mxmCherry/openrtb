package rtb

// 5.10 Video Start Delay
//
// The following table lists the various options for the video start delay. If the start delay value is greater
// than 0, then the position is mid-roll and the value indicates the start delay.
const (
	// > 0 Mid-Roll (value indicates start delay in second)
	VideoStartDelayPreRoll         int64 = 0  //  0 Pre-Roll
	VideoStartDelayGenericMidRoll  int64 = -1 // -1 Generic Mid-Roll
	VideoStartDelayGenericPostRoll int64 = -2 // -2 Generic Post-Roll
)
