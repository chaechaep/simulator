package log

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"path/filepath"
)

type logLevelHook struct {
	level logrus.Level
	logrus.Hook
}

func setupFileHook(dirPath string, level logrus.Level, componentName string) {
	fullPath := filepath.Join(dirPath, componentName+".log")

	if err := os.MkdirAll(path.Dir(fullPath), os.ModePerm); err != nil {
		log.Fatalf("Couldn't create directory %s: %q", path.Dir(fullPath), err)
	}

	Log.AddHook(&logLevelHook{
		level,
		NewFSHook(
			fullPath,
			&Formatter{
				TimestampFormat: "2006-01-02 15:04:05.999999",
				NoColors:        true,
				NoFieldsColors:  true,
			},
			&DailyRotationSchedule{GZip: true},
		),
	})

}

func SetupHookLogging(path string, level string, processName string) {
	Log.SetReportCaller(true)
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		Log.Fatalf("Unrecognised logging level %s: %q", level, err)
	}

	if Log.GetLevel() < logLevel {
		Log.SetLevel(logLevel)
	}
	setupFileHook(path, logLevel, processName)

}

var Log = *logrus.New()

func Init(filePath, level, processName string) {
	Log.SetFormatter(&Formatter{
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05.999999",
	})
	SetupHookLogging(filePath, level, processName)

	m := Log.WithField("component", "main")
	m.WithFields(logrus.Fields{
		"path":        filePath,
		"level":       level,
		"processName": processName,
	}).Info("create logging")

	Log.Info("start log...")
}
