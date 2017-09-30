package pacchetto

const (
	acSubPath      = "steamapps/common/assettocorsa"
	contentSubPath = "content"
	serverSubPath  = "server"
	tracksSubPath  = "tracks"
	carsSubPath    = "cars"
	weatherSubPath = "weather"
	outputPrefix   = "assetto-corsa-server"
	tempPrefix     = ".pacchetto"
)

var contentSubPaths = [...]string{
	tracksSubPath, carsSubPath, weatherSubPath,
}

var windowsDriveLetters = [...]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}
