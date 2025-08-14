package utils

import (
	"strings"
	"time"

	"github.com/wushiling50/aster/config"
	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/gen/contribution"
	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/pkg/constants"
)

// GetMysqlDSN 会拼接 Mysql 的 DSN
func GetMysqlDSN(mysql config.MysqlConf) string {
	dsn := strings.Join([]string{
		mysql.Username, ":", mysql.Password,
		"@tcp(", mysql.Addr, ")/",
		mysql.Database, "?charset=" + mysql.Charset + "&parseTime=true",
	}, "")

	return dsn
}

func GetTextFromProfile(profile *developer.Developer) string {
	text := "|Developer Profile Start|" + "|Name:" +
		profile.Name + "|Bio:" + profile.Bio + "|TwitterUsername:" + profile.TwitterUsername +
		"|Company:" + profile.Company + "|Location:" + profile.Location + "|Developer Profile End|"

	return text
}

func GetTextFromContribution(contributions []*contribution.Contribution) string {
	limitCharacterCount := 1000
	text := "|Contribution Start|"

	for _, theContribution := range contributions {
		text += theContribution.Content
		if len(text) > limitCharacterCount {
			break
		}
	}

	text = strings.ReplaceAll(text, "\n", " ")
	if len(text) > limitCharacterCount {
		text = text[:limitCharacterCount]
	}

	text += "|Contribution End|"

	return text
}

func GetTextFromLanguages(languages *analysis.Languages) string {
	text := "|Language Usage Start|" + languages.Languages + "|Language Usage End|"

	return text
}

func CheckIfDataExpired(lastUpdate time.Time) bool {
	return time.Since(lastUpdate) > constants.DataExpiredTime
}

func CustomizeCheckIfDataExpired(lastUpdate time.Time, duration time.Duration) bool {
	return time.Since(lastUpdate) > duration
}
