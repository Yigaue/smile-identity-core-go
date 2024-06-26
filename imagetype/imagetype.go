package imagetype

import "errors"

const (
	SELFIE_FILE           = 0
	ID_CARD_FILE          = 1
	SELFIE_IMAGE_STRING   = 2
	ID_CARD_IMAGE_STRING  = 3
	LIVENESS_IMAGE_FILE   = 4
	ID_CARD_BACK_FILE     = 5
	LIVENESS_IMAGE_STRING = 6
	ID_CARD_BACK_STRING   = 7
)

var (
	ErrInvalidImagetype       = errors.New("invalid imagetype")
)
