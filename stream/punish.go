package stream

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"gitlab.com/Startail/Systera-API/systerapb"
)

// PublishPunish - Publish Punish
func PublishPunish(target string, data *systerapb.PunishEntry) {
	c := pool.Get()
	defer c.Close()

	d := &systerapb.PunishEntryStream{
		Type:  systerapb.PunishEntryStream_PUNISH,
		Entry: data,
	}
	serialized, _ := json.Marshal(&d)
	logrus.Debugln(d)

	_, err := c.Do("PUBLISH", "systera.punish."+target, string(serialized))
	if err != nil {
		logrus.WithError(err).Errorf("[Publish] Failed Publish Punishment")
	}
}
