package actions

import (
	"context"
	"foursquare.com/airfog/cmd/util"
	"foursquare.com/airfog/internal/airflow"
)

type ApiCtx struct {
	Api        *airflow.APIClient
	ApiConfig  Configuration
	ApiContext context.Context
}

func AirflowApi() ApiCtx {
	airfogConfig := loadConfig()
	conf := airflow.NewConfiguration()
	conf.Host = airfogConfig.Host
	conf.Scheme = airfogConfig.Scheme

	util.PrintValue("Airflow:", airfogConfig.Scheme+"://"+airfogConfig.Host)

	api := airflow.NewAPIClient(conf)

	cred := airflow.BasicAuth{
		UserName: airfogConfig.Username,
		Password: airfogConfig.Password,
	}

	var ctx ApiCtx
	ctx.Api = api
	ctx.ApiConfig = airfogConfig
	ctx.ApiContext = context.WithValue(context.Background(), airflow.ContextBasicAuth, cred)
	return ctx
}
