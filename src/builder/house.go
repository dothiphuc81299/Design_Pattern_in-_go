package builder

// House ..
type House struct {
	windowType string
	doorType   string
	floor      int
}

// GetWindowType ..
func (h *House) GetWindowType() string {
	return h.windowType
}

// GetDoorType ..
func (h *House) GetDoorType() string {
	return h.doorType
}

// GetFloor
func (h *House) GetFloor() int {
	return h.floor
}
