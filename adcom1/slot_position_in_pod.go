package adcom1

// PodSequence identifies the slot position in pod field, for use in audio and video ad pods.
type SlotPositionInPod int8

// PodSequence options.
const (
	SlotPosLast        SlotPositionInPod = -1 // -1	Last ad in the pod.
	SlotPosAny         SlotPositionInPod = 0  // 0	Any ad in the pod.
	SlotPosFirst       SlotPositionInPod = 1  // 1	First ad in the pod.
	SlotPosFirstOrLast SlotPositionInPod = 2  // 1	First or Last ad in the pod.
)
