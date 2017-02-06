package repo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDependency(t *testing.T) {
	dependencies, error := NewDependencies();
	assert.NotNil(t, dependencies);
	assert.Nil(t, error);

	assert.NotZero(t, len(dependencies.Repositories));

	// ファイルに書き込める
	assert.Nil(t, dependencies.ToFile("dependencies.json"));

	// ファイルから読み込める
	load, error := FromFile("dependencies.json");

	assert.NotZero(t, len(load.Repositories));
	assert.Nil(t, error);
}
