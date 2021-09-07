package main

type VerseDetails struct {
	Text      string `json:"text"`
	Reference string `json:"reference"`
	Version   string `json:"version"`
}

type VerseMeta struct {
	Details VerseDetails `json:"details"`
}

type VOTD struct {
	Verse VerseMeta `json:"verse"`
}
