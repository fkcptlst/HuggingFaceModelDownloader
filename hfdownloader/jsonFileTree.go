package hfdownloader

type hfmodel []struct {
	Type string `json:"type"`
	Oid  string `json:"oid"`
	Size int    `json:"size"`
	Path string `json:"path"`
	Lfs  struct {
		Oid         string `json:"oid"` // in lfs, oid is sha256 of the file
		Size        int64  `json:"size"`
		PointerSize int    `json:"pointerSize"`
	} `json:"lfs,omitempty"`
}