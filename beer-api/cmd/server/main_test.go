package main

import (
	"context"
	"fmt"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/config"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/test"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_buildHandler(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	cfg, _ := config.Load(*flagConfig, logger)
	buildHandler(logger, db, cfg, nil)

}

func Test_logDBQuery(t *testing.T) {
	logger, entries := log.NewForTest()
	f := logDBQuery(logger)
	f(context.Background(), time.Millisecond*3, "sql", nil, nil)
	if assert.Equal(t, 1, entries.Len()) {
		assert.Equal(t, "DB query successful", entries.All()[0].Message)
	}
	entries.TakeAll()

	f(context.Background(), time.Millisecond*3, "sql", nil, fmt.Errorf("test"))
	if assert.Equal(t, 1, entries.Len()) {
		assert.Equal(t, "DB query error: test", entries.All()[0].Message)
	}
}

func Test_logDBExec(t *testing.T) {
	logger, entries := log.NewForTest()
	f := logDBExec(logger)
	f(context.Background(), time.Millisecond*3, "sql", nil, nil)
	if assert.Equal(t, 1, entries.Len()) {
		assert.Equal(t, "DB execution successful", entries.All()[0].Message)
	}
	entries.TakeAll()

	f(context.Background(), time.Millisecond*3, "sql", nil, fmt.Errorf("test"))
	if assert.Equal(t, 1, entries.Len()) {
		assert.Equal(t, "DB execution error: test", entries.All()[0].Message)
	}
}
