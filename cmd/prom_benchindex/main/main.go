package main

import (
	"flag"
	"os"
	"time"

	kitlogzap "github.com/go-kit/kit/log/zap"
	"github.com/m3db/m3/src/x/instrument"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/tsdb/chunkenc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	var (
		flagBlockDir = flag.String("block-dir", "", "block directory")
	)

	flag.Parse()

	logger := instrument.NewOptions().Logger()

	if *flagBlockDir == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	var (
		blockDir = *flagBlockDir
	)
	kitLogger := kitlogzap.NewZapSugarLogger(logger, zapcore.InfoLevel)
	block, err := tsdb.OpenBlock(kitLogger, blockDir, chunkenc.NewPool())
	if err != nil {
		logger.Fatal("could not open block", zap.Error(err))
	}

	index, err := block.Index()
	if err != nil {
		logger.Fatal("could not open block index", zap.Error(err))
	}

	matcher, err := labels.NewMatcher(labels.MatchRegexp,
		"hostname", "^host_[0-9]+9$")
	if err != nil {
		logger.Fatal("could not create matcher", zap.Error(err))
	}

	start := time.Now()
	postings, err := tsdb.PostingsForMatchers(index, matcher)
	took := time.Since(start)
	if err != nil {
		logger.Fatal("could not retrieve postings", zap.Error(err))
	}

	n := 0
	for postings.Next() {
		n++
	}
	if err := postings.Err(); err != nil {
		logger.Fatal("iterate postings error", zap.Error(err))
	}

	logger.Info("matched time series",
		zap.Int("n", n),
		zap.Stringer("took", took))
}
