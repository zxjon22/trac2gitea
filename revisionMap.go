package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var revisionRegexp = regexp.MustCompile(`^r\d+`)

// readRevisionMap reads the revision map (SVN revisions -> Git commits map)
// from the provided file. Returns nil if no file was provided
func readRevisionMap(mapFile string) (map[string]string, error) {
	if mapFile == "" {
		return nil, nil
	}

	fd, err := os.Open(mapFile)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	revisionMap := make(map[string]string)
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		revisionMapLine := scanner.Text()
		if revisionMapLine == "" {
			continue
		}

		equalsPos := strings.LastIndex(revisionMapLine, "=")
		if equalsPos == -1 {
			return nil, fmt.Errorf("badly formatted revision map file %s: found line %s", mapFile, revisionMapLine)
		}

		gitRef := strings.Trim(revisionMapLine[0:equalsPos], " ")
		svnRevision := strings.Trim(revisionMapLine[equalsPos+1:], " ")
		svnRevision = revisionRegexp.FindString(svnRevision)

		if svnRevision != "" {
			if _, found := revisionMap[svnRevision]; found == true {
				return nil, fmt.Errorf("Revision map file %s: contains multiple entries for %s", mapFile, svnRevision)
			}
			revisionMap[svnRevision] = gitRef
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return revisionMap, nil
}
