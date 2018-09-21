package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/joeke80215/dumpcat/config"
	"github.com/joeke80215/dumpcat/pkg"
	"github.com/joeke80215/dumpcat/write"
	"github.com/olivere/elastic"
)

type es struct {
	host string
}

// NewESConn new elasticsearch connect for wirte packet log
func NewESConn() write.Writer {
	return &es{
		host: config.Cfg.GetOutput()["elasticsearch"].Host,
	}
}

func (e *es) Write(p pkg.Packet) error {
	client, err := elastic.NewClient(elastic.SetURL(e.host), elastic.SetSniff(false))
	if err != nil {
		return err
	}
	ctx := context.Background()
	_, err = client.Index().
		Index(fmt.Sprintf("dumpcap-%s", time.Now().Format("2006-01-02"))).
		Type("doc").
		BodyJson(p.GetData()).
		Do(ctx)
	if err != nil {
		return err
	}

	return nil
}
