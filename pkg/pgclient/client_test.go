package pgclient_test

import (
	"testing"

	"github.com/ahmadmilzam/go/pkg/pgclient"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	sql := pgclient.New()
	defer sql.Close()

	assert.NotNil(t, sql)
}
