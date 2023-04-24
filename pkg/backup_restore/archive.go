package backup_restore

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
)

func archiveMapToTarGz(data map[string][]byte) (io.Reader, int, error) {
	var buf bytes.Buffer

	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	for filename, content := range data {
		hdr := &tar.Header{
			Name: filename,
			Mode: 0600,
			Size: int64(len(content)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return nil, 0, err
		}

		if _, err := tw.Write(content); err != nil {
			return nil, 0, err
		}
	}

	if err := tw.Close(); err != nil {
		return nil, 0, err
	}

	if err := gw.Close(); err != nil {
		return nil, 0, err
	}

	return &buf, buf.Len(), nil
}

func unarchiveTarGzToMap(backup []byte) (map[string][]byte, error) {
	gzr, err := gzip.NewReader(bytes.NewReader(backup))
	if err != nil {
		return nil, err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	data := make(map[string][]byte)
	for {
		header, err := tr.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if header.Typeflag == tar.TypeReg {
			content, err := ioutil.ReadAll(tr)
			if err != nil {
				return nil, err
			}
			data[header.Name] = content
		}
	}

	return data, nil
}
