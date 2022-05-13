package adcom1

// PodSequence identifies the pod sequence field, for use in audio and video content streams with one or more ad pods.
type PodSequence int8

// PodSequence options.
const (
	PodSeqLast  PodSequence = -1 // -1	Last pod in the content stream.
	PodSeqAny   PodSequence = 0  // 0	Any pod in the content stream.
	PodSeqFirst PodSequence = 1  // 1	First pod in the content stream.
)
