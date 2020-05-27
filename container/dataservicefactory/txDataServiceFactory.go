package dataservicefactory

import (
	"github.com/pkg/errors"

	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice/txdataservice"
	"github.com/jfeng45/servicetmpl/tool/gdbc"
)

// txDataServiceFactory is a empty receiver for Build method
type txDataServiceFactory struct{}

func (tdsf *txDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("txDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(gdbc.SqlGdbc)
	tdm := txdataservice.TxDataSql{ds}
	logger.Log.Debug("udm:", tdm.DB)
	return &tdm, nil

}
