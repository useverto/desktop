package fs

import (
	"bytes"
	"io"
	"net/http"
	"path"
	"path/filepath"
)

// Walk walks the file tree rooted at root,
// calling walkFn for each file or directory in the tree, including root.
// All errors that arise visiting files and directories are filtered by walkFn.
//
// As with filepath.Walk, if the walkFn returns filepath.SkipDir, then the directory is skipped.
func Walk(hfs http.FileSystem, root string, walkFn filepath.WalkFunc) error {
	dh, err := hfs.Open(root)
	if err != nil {
		return err
	}
	di, err := dh.Stat()
	if err != nil {
		return err
	}
	fis, err := dh.Readdir(-1)
	dh.Close()
	if err = walkFn(root, di, err); err != nil {
		if err == filepath.SkipDir {
			return nil
		}
		return err
	}
	for _, fi := range fis {
		fn := path.Join(root, fi.Name())
		if fi.IsDir() {
			if err = Walk(hfs, fn, walkFn); err != nil {
				if err == filepath.SkipDir {
					continue
				}
				return err
			}
			continue
		}
		if err = walkFn(fn, fi, nil); err != nil {
			if err == filepath.SkipDir {
				continue
			}
			return err
		}
	}
	return nil
}

// ReadFile reads the contents of the file of hfs specified by name.
// Just as ioutil.ReadFile does.
func ReadFile(hfs http.FileSystem, name string) ([]byte, error) {
	fh, err := hfs.Open(name)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, fh)
	fh.Close()
	return buf.Bytes(), err
}
