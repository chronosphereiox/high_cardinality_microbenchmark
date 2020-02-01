package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof" // pprof: for debug listen server if configured
	"os"
	"time"

	"github.com/chronosphereiox/high_cardinality_microbenchmark/pkg/generator"

	"github.com/m3db/m3/src/cmd/services/m3dbnode/config"
	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/persist"
	"github.com/m3db/m3/src/dbnode/persist/fs"
	"github.com/m3db/m3/src/dbnode/runtime"
	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/m3ninx/index/segment/builder"
	"github.com/m3db/m3/src/x/ident"
	"github.com/m3db/m3/src/x/instrument"
	"go.uber.org/zap"
)

var (
	blockSize = 2 * time.Hour
)

func main() {
	var (
		flagCardinality = flag.Int("cardinality", 5000000, "cardinality to generate")
		flagDir         = flag.String("dir", "/tmp", "directory for output")
	)

	flag.Parse()

	logger := instrument.NewOptions().Logger()

	if *flagCardinality <= 0 || *flagDir == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	srv := httptest.NewServer(http.DefaultServeMux)
	logger.Info("test server with pprof", zap.String("url", srv.URL))

	start := time.Now().Truncate(blockSize).Add(-1 * blockSize)
	timeNowFn := func() time.Time { return start }

	gen := generator.NewHostsSimulator(10000, start,
		generator.HostsSimulatorOptions{TimeNowFn: timeNowFn})

	idxOpts := namespace.NewIndexOptions().
		SetEnabled(true).
		SetBlockSize(blockSize)
	nsOpts := namespace.NewOptions().SetIndexOptions(idxOpts)
	ns, err := namespace.NewMetadata(ident.StringID("data"),
		nsOpts)
	if err != nil {
		logger.Fatal("unable to create namespace metadata", zap.Error(err))
	}

	cfg := config.DBConfiguration{}

	runtimeOptsMgr := runtime.NewOptionsManager()
	runtimeOpts := runtime.NewOptions().
		SetFlushIndexBlockNumSegments(1)
	if err := runtimeOptsMgr.Update(runtimeOpts); err != nil {
		logger.Fatal("could not update runtime options", zap.Error(err))
	}

	fsOpts := fs.NewOptions().
		SetFilePathPrefix(m3Dir).
		SetRuntimeOptionsManager(runtimeOptsMgr).
		SetWriterBufferSize(cfg.Filesystem.WriteBufferSizeOrDefault()).
		SetDataReaderBufferSize(cfg.Filesystem.DataReadBufferSizeOrDefault()).
		SetInfoReaderBufferSize(cfg.Filesystem.InfoReadBufferSizeOrDefault()).
		SetSeekReaderBufferSize(cfg.Filesystem.SeekReadBufferSizeOrDefault()).
		SetForceIndexSummariesMmapMemory(cfg.Filesystem.ForceIndexSummariesMmapMemoryOrDefault()).
		SetForceBloomFilterMmapMemory(cfg.Filesystem.ForceBloomFilterMmapMemoryOrDefault()).
		SetIndexBloomFilterFalsePositivePercent(cfg.Filesystem.BloomFilterFalsePositivePercentOrDefault())

	persistMgr, err := fs.NewPersistManager(fsOpts)
	if err != nil {
		logger.Fatal("unable to create persist manager", zap.Error(err))
	}

	flush, err := persistMgr.StartIndexPersist()
	if err != nil {
		logger.Fatal("unable to create persist manager", zap.Error(err))
	}

	noUUIDFn := func() ([]byte, error) { return nil, fmt.Errorf("ID prebuilt") }
	builderOpts := builder.NewOptions().SetNewUUIDFn(noUUIDFn)
	builder, err := builder.NewBuilderFromDocuments(builderOpts)
	if err != nil {
		logger.Fatal("unable to create builder", zap.Error(err))
	}

	_, err = builder.Insert(doc.Document{
		ID: []byte("foo"),
		Fields: []doc.Field{
			{Name: []byte("name"), Value: []byte("value")},
		}})
	if err != nil {
		logger.Fatal("unable to insert document", zap.Error(err))
	}

	preparedPersist, err := flush.PrepareIndex(persist.IndexPrepareOptions{
		NamespaceMetadata: ns,
		BlockStart:        opts.start,
		FileSetType:       persist.FileSetFlushType,
		Shards:            map[uint32]struct{}{0: struct{}{}},
	})
	if err != nil {
		logger.Fatal("unable to start persist", zap.Error(err))
	}

	if err := preparedPersist.Persist(builder); err != nil {
		logger.Fatal("unable to persist", zap.Error(err))
	}

	if _, err := preparedPersist.Close(); err != nil {
		logger.Fatal("unable to close persist", zap.Error(err))
	}
}
