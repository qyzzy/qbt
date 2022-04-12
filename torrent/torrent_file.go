package torrent

const SHALEN int = 20

type rawInfo struct {
	Name        string
	Length      int
	Pieces      string
	PieceLength int
}

type rawFile struct {
	Announce string
	Info     rawInfo
}

type TorrentFile struct {
	Announce string
	InfoSHA  [SHALEN]byte
	FileName string
	FileLen  int
	PieceLen int
	PieceSHA [][SHALEN]byte
}
