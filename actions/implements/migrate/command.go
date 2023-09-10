package migrate

import (
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/AH-dark/epay-cli/pkg/utils"
)

func (svc *service) Command() *cli.Command {
	return &cli.Command{
		Name:   "migrate",
		Usage:  "generate sql schema",
		Action: svc.Do,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "database.driver",
				Usage:   "database driver",
				EnvVars: []string{"DATABASE_DRIVER", "DB_DRIVER"},
				Value:   "mysql",
			},
			&cli.StringFlag{
				Name:    "database.host",
				Usage:   "database host",
				EnvVars: []string{"DATABASE_HOST", "DB_HOST"},
				Value:   "localhost",
			},
			&cli.IntFlag{
				Name:    "database.port",
				Usage:   "database port",
				EnvVars: []string{"DATABASE_PORT", "DB_PORT"},
				Value:   3306,
			},
			&cli.StringFlag{
				Name:    "database.name",
				Usage:   "database name",
				EnvVars: []string{"DATABASE_NAME", "DB_NAME"},
				Value:   "epay",
			},
			&cli.StringFlag{
				Name:    "database.user",
				Usage:   "database user",
				EnvVars: []string{"DATABASE_USER", "DB_USER"},
				Value:   "root",
			},
			&cli.StringFlag{
				Name:    "database.password",
				Usage:   "database password",
				EnvVars: []string{"DATABASE_PASSWORD", "DB_PASSWORD"},
				Value:   "",
			},
			&cli.StringFlag{
				Name:    "database.prefix",
				Usage:   "database prefix",
				EnvVars: []string{"DATABASE_PREFIX", "DB_PREFIX"},
				Value:   "pay_",
			},
			&cli.StringFlag{
				Name:    "database.sslmode",
				Usage:   "database sslmode",
				EnvVars: []string{"DATABASE_SSL_MODE", "DB_SSL_MODE"},
				Value:   "disable",
			},
			&cli.StringFlag{
				Name:    "database.charset",
				Usage:   "database charset",
				EnvVars: []string{"DATABASE_CHARSET", "DB_CHARSET"},
				Value:   "utf8mb4",
			},
			&cli.StringFlag{
				Name:    "app.syskey",
				Usage:   "app syskey",
				EnvVars: []string{"APP_SYSKEY", "APP_SYS_KEY"},
				Value:   utils.RandString(32),
			},
			&cli.StringFlag{
				Name:    "app.cronkey",
				Usage:   "app cronkey",
				EnvVars: []string{"APP_CRONKEY", "APP_CRON_KEY"},
				Value:   strconv.Itoa(utils.RandInt(100000, 999999)),
			},
			&cli.StringFlag{
				Name:    "migrate.default_config",
				Usage:   "migrate default config",
				EnvVars: []string{"MIGRATE_DEFAULT_CONFIG", "MIGRATION_DEFAULT_CONFIG"},
				Value:   "./data/default_config.json",
			},
			&cli.StringFlag{
				Name:    "migrate.default_payment_types",
				Usage:   "migrate default payment types",
				EnvVars: []string{"MIGRATE_DEFAULT_PAYMENT_TYPES", "MIGRATION_DEFAULT_PAYMENT_TYPES"},
				Value:   "./data/default_payment_types.json",
			},
		},
	}
}
