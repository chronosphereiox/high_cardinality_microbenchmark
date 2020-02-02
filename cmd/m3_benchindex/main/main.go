package main

import (
	"flag"
	"os"
	"time"

	"github.com/m3db/m3/src/cmd/services/m3dbnode/config"
	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/persist"
	"github.com/m3db/m3/src/dbnode/persist/fs"
	"github.com/m3db/m3/src/m3ninx/index"
	"github.com/m3db/m3/src/m3ninx/postings"
	"github.com/m3db/m3/src/x/ident"
	"github.com/m3db/m3/src/x/instrument"
	"go.uber.org/zap"
)

var (
	blockSize = 2 * time.Hour
)

func main() {
	var (
		flagDir = flag.String("dir", "", "m3 data directory")
	)

	flag.Parse()

	logger := instrument.NewOptions().Logger()

	if *flagDir == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	var (
		dir = *flagDir
	)
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
	fsOpts := fs.NewOptions().
		SetFilePathPrefix(dir).
		SetWriterBufferSize(cfg.Filesystem.WriteBufferSizeOrDefault()).
		SetDataReaderBufferSize(cfg.Filesystem.DataReadBufferSizeOrDefault()).
		SetInfoReaderBufferSize(cfg.Filesystem.InfoReadBufferSizeOrDefault()).
		SetSeekReaderBufferSize(cfg.Filesystem.SeekReadBufferSizeOrDefault()).
		SetForceIndexSummariesMmapMemory(cfg.Filesystem.ForceIndexSummariesMmapMemoryOrDefault()).
		SetForceBloomFilterMmapMemory(cfg.Filesystem.ForceBloomFilterMmapMemoryOrDefault()).
		SetIndexBloomFilterFalsePositivePercent(cfg.Filesystem.BloomFilterFalsePositivePercentOrDefault())

	infoFiles := fs.ReadIndexInfoFiles(fsOpts.FilePathPrefix(), ns.ID(),
		fsOpts.InfoReaderBufferSize())
	if len(infoFiles) != 1 {
		logger.Fatal("expected single index segment",
			zap.Int("infoFiles", len(infoFiles)))
	}

	infoFile := infoFiles[0]
	segments, err := fs.ReadIndexSegments(fs.ReadIndexSegmentsOptions{
		ReaderOptions: fs.IndexReaderOpenOptions{
			Identifier:  infoFile.ID,
			FileSetType: persist.FileSetFlushType,
		},
		FilesystemOptions: fsOpts,
	})
	if err != nil {
		logger.Fatal("could not load segments", zap.Error(err))
	}
	if len(segments) != 1 {
		logger.Fatal("unexpected number of segments")
	}

	segment := segments[0]
	reader, err := segment.Reader()
	if err != nil {
		logger.Fatal("could not load segment reader", zap.Error(err))
	}

	field := []byte("pod")
	regexp, err := index.CompileRegex([]byte("^abc.*$"))
	if err != nil {
		logger.Fatal("compile regexp error", zap.Error(err))
	}

	var postings postings.List
	sevenDayTwoHourNumBlocks := 7 * 12 // 12 blocks per day
	start := time.Now()
	for i := 0; i < sevenDayTwoHourNumBlocks; i++ {
		postings, err = reader.MatchRegexp(field, regexp)
		if err != nil {
			logger.Fatal("could not get postings", zap.Error(err))
		}
	}
	took := time.Since(start)

	logger.Info("matched time series",
		zap.Int("n", postings.Len()),
		zap.Stringer("took", took))
}
