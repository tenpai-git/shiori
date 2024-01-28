package model

import (
	"context"
	"io/fs"
	"os"
	"time"

	"github.com/go-shiori/warc"
)

type BookmarksDomain interface {
	HasEbook(b *BookmarkDTO) bool
	HasArchive(b *BookmarkDTO) bool
	HasThumbnail(b *BookmarkDTO) bool
	GetBookmark(ctx context.Context, id DBID) (*BookmarkDTO, error)
}

type AccountsDomain interface {
	CheckToken(ctx context.Context, userJWT string) (*Account, error)
	GetAccountFromCredentials(ctx context.Context, username, password string) (*Account, error)
	CreateTokenForAccount(account *Account, expiration time.Time) (string, error)
}

type ArchiverDomain interface {
	DownloadBookmarkArchive(book BookmarkDTO) (*BookmarkDTO, error)
	GetBookmarkArchive(book *BookmarkDTO) (*warc.Archive, error)
}

type StorageDomain interface {
	Stat(name string) (fs.FileInfo, error)
	Create(name string) (fs.File, error)
	Open(name string) (fs.File, error)
	Remove(name string) error
	MkDirAll(name string, mode os.FileMode) error
	FileExists(path string) bool
	DirExists(path string) bool
	WriteData(dst string, data []byte) error
	WriteFile(dst string, src *os.File) error
}
