package fileinfo

import (
	"database/sql/driver"
	"errors"
)

type FileKind string

func (me FileKind) Value() (driver.Value, error) {
	return driver.Value(string(me)), nil
}
func (me *FileKind) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		*me = FileKind(src.(string))
		return nil
	case []byte:
		*me = FileKind(string(src.([]byte)))
		return nil
	default:
		return errors.New("Incompatible type for FileKind (string should be string or []byte)")
	}
}

const PictureKind = FileKind("picture")
const SoundKind = FileKind("sound")
const MovieKind = FileKind("movie")
const DocKind = FileKind("doc")
const OtherKind = FileKind("other")

type FileType string

func (me FileType) Value() (driver.Value, error) {
	return driver.Value(string(me)), nil
}
func (me *FileType) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		*me = FileType(src.(string))
		return nil
	case []byte:
		*me = FileType(string(src.([]byte)))
		return nil
	default:
		return errors.New("Incompatible type for FileType (string should be string or []byte)")
	}
}

const JpegType = FileType("jpeg")
const PngType = FileType("png")
const GifType = FileType("gif")
const TiffType = FileType("tiff")
const Jpeg2KType = FileType("jp2")
const PhotoCdType = FileType("pcd")

const WavType = FileType("wav")
const OggType = FileType("ogg")
const Mp3Type = FileType("mp3")
const M4aType = FileType("m4a")

const Mp4Type = FileType("mp4")
const MpgType = FileType("mpg")
const MovType = FileType("mov")
const AviType = FileType("avi")
const WmvType = FileType("wmv")
const FlvType = FileType("flv")
const SwfType = FileType("swf")
const M4vType = FileType("m4v")

const PdfType = FileType("pdf")

const OtherType = FileType("other")

var _typeKind map[FileType]FileKind
var _extensionType map[string]FileType

func init() {
	_typeKind = make(map[FileType]FileKind)
	_typeKind[JpegType] = PictureKind
	_typeKind[PngType] = PictureKind
	_typeKind[GifType] = PictureKind
	_typeKind[TiffType] = PictureKind
	_typeKind[Jpeg2KType] = PictureKind
	_typeKind[PhotoCdType] = PictureKind

	_typeKind[WavType] = SoundKind
	_typeKind[OggType] = SoundKind
	_typeKind[Mp3Type] = SoundKind
	_typeKind[M4aType] = SoundKind

	_typeKind[MovType] = MovieKind
	_typeKind[AviType] = MovieKind
	_typeKind[Mp4Type] = MovieKind
	_typeKind[MpgType] = MovieKind
	_typeKind[FlvType] = MovieKind
	_typeKind[SwfType] = MovieKind
	_typeKind[M4vType] = MovieKind

	_typeKind[PdfType] = DocKind

	_typeKind[OtherType] = OtherKind

	_extensionType = make(map[string]FileType)
	// Each FileType has at least its own value as a known file extension:
	for ftype, _ := range _typeKind {
		_extensionType[string(ftype)] = ftype
	}
	// Add on some known alternates
	_extensionType["jpg"] = JpegType
	_extensionType["jif"] = JpegType
	_extensionType["jfif"] = JpegType
	_extensionType["tif"] = TiffType
	_extensionType["jpx"] = Jpeg2KType
	_extensionType["j2k"] = Jpeg2KType
	_extensionType["j2c"] = Jpeg2KType
	_extensionType["mpeg"] = MpgType
}

func FileTypeForExt(ext string) FileType {
	ftype, ok := _extensionType[ext]
	if !ok {
		return OtherType
	}
	return ftype
}

func FileKindForExt(ext string) FileKind {
	fkind, ok := _typeKind[FileTypeForExt(ext)]
	if !ok {
		return OtherKind
	}
	return fkind
}
