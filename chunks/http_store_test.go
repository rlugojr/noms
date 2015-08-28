package chunks

import (
	"testing"

	"github.com/attic-labs/noms/Godeps/_workspace/src/github.com/stretchr/testify/suite"
)

func TestHttpStoreTestSuite(t *testing.T) {
	suite.Run(t, &HttpStoreTestSuite{})
}

type HttpStoreTestSuite struct {
	ChunkStoreTestSuite
	server *HttpStoreServer
}

func (suite *HttpStoreTestSuite) SetupTest() {
	suite.store = NewHttpStoreClient("http://localhost:8000")
	suite.server = NewHttpStoreServer(&MemoryStore{}, 8000)
	go suite.server.Run()
}

func (suite *HttpStoreTestSuite) TearDownTest() {
	suite.server.Stop()
}