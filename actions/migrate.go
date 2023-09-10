package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"

	"github.com/AH-dark/epay-cli/database/conn"
	"github.com/AH-dark/epay-cli/database/model"
	"github.com/AH-dark/epay-cli/pkg/log"
	"github.com/AH-dark/epay-cli/pkg/utils"
)

func MigrateCommand() *cli.Command {
	return &cli.Command{
		Name:   "migrate",
		Usage:  "generate sql schema",
		Action: MigrateAction,
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

func getDefaultDatabaseConfig(c *cli.Context) (map[string]string, error) {
	var data map[string]string

	logger := log.Log(c.Context)

	path := utils.GetAbsolutePath(c.String("migrate.default_config"))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.WithError(err).Error("default config file not exists")
		return data, nil
	}

	logger = logger.WithField("path", path)

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		logger.WithError(err).Error("open default config file failed")
		return data, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			logger.WithError(err).Warn("close default config file failed")
		}
	}(f)

	if err := json.NewDecoder(f).Decode(&data); err != nil {
		logger.WithError(err).Error("decode default config file failed")
		return data, err
	}

	return data, nil
}

func getDefaultPaymentTypes(c *cli.Context) ([]model.Type, error) {
	var data []model.Type

	logger := log.Log(c.Context)

	path := utils.GetAbsolutePath(c.String("migrate.default_payment_types"))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.WithError(err).Error("default payment types file not exists")
		return data, nil
	}

	logger = logger.WithField("path", path)

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		logger.WithError(err).Error("open default payment types file failed")
		return data, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			logger.WithError(err).Warn("close default payment types file failed")
		}
	}(f)

	if err := json.NewDecoder(f).Decode(&data); err != nil {
		logger.WithError(err).Error("decode default payment types file failed")
		return data, err
	}

	return data, nil
}

var defaultGroup = model.Group{
	GID:  0,
	Name: "默认分组",
	Info: `{"1":{"type":"","channel":"-1","rate":""},"2":{"type":"","channel":"-1","rate":""},"3":{"type":"","channel":"-1","rate":""}}`,
}

func createDatabaseConfig(c *cli.Context, configs map[string]string) error {
	db := conn.GetDB(c)
	ctx := c.Context

	for k, v := range configs {
		logger := log.Log(ctx).WithField("key", k).WithField("val", v)
		logger.Debug("create default config")

		// check if config exists
		if err := db.Model(&model.Config{}).
			Where("k = ?", k).
			First(&model.Config{}).
			Error; err == nil {
			logger.Debug("config already exists")
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.WithError(err).Error("check default config failed")
			return err
		}

		// create config
		if err := db.
			Model(&model.Config{}).
			Create(&model.Config{
				Key: k,
				Val: v,
			}).Error; err != nil {
			logger.WithError(err).Error("create default config failed")
			return err
		}
	}

	return nil
}

func MigrateAction(c *cli.Context) error {
	ctx := c.Context
	db := conn.GetDB(c)

	log.Log(ctx).Info("auto migrate database tables")
	if err := db.AutoMigrate(
		&model.AlipayRisk{},
		&model.Anounce{},
		&model.Batch{},
		&model.Channel{},
		&model.Config{},
		&model.Domain{},
		&model.Group{},
		&model.Log{},
		&model.Order{},
		&model.Plugin{},
		&model.Record{},
		&model.Regcode{},
		&model.Risk{},
		&model.Roll{},
		&model.Settle{},
		&model.Type{},
		&model.User{},
		&model.Weixin{},
	); err != nil {
		log.Log(ctx).WithError(err).Error("auto migrate database tables failed")
		return err
	}

	// alter auto increment if user table is empty
	if err := db.Model(&model.User{}).First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Log(ctx).Info("alter auto increment")
		if err := db.Model(&model.User{}).
			Exec(fmt.Sprintf(
				"alter table `%s` AUTO_INCREMENT = 1000",
				db.NamingStrategy.TableName("user"),
			)).Error; err != nil {
			log.Log(ctx).WithError(err).Error("alter auto increment failed")
			return err
		}
	}

	log.Log(ctx).Info("get default config")
	defaultDatabaseConfig, err := getDefaultDatabaseConfig(c)
	if err != nil {
		log.Log(ctx).WithError(err).Error("get default config failed")
		return err
	}

	log.Log(ctx).Info("create default config")
	if err := createDatabaseConfig(c, defaultDatabaseConfig); err != nil {
		log.Log(ctx).WithError(err).Panic("create default config failed")
		return err
	}

	log.Log(ctx).Info("init app config")
	initData := map[string]string{
		"syskey":  c.String("app.syskey"),
		"build":   time.Now().Format("2006-01-02"),
		"cronkey": c.String("app.cronkey"),
	}
	if err := createDatabaseConfig(c, initData); err != nil {
		log.Log(ctx).WithError(err).Panic("create app init config failed")
		return err
	}

	log.Log(ctx).Info("create default payment types")
	defaultPaymentTypes, err := getDefaultPaymentTypes(c)
	if err != nil {
		log.Log(ctx).WithError(err).Error("get default payment types failed")
		return err
	}

	for _, t := range defaultPaymentTypes {
		if err := db.Model(&model.Type{}).
			Where("name = ?", t.Name).
			First(&model.Type{}).Error; err == nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Log(ctx).WithError(err).Error("check default payment type failed")
			return err
		}

		if err := db.Model(&model.Type{}).Create(&t).Error; err != nil {
			log.Log(ctx).WithError(err).Error("create default payment type failed")
			return err
		}
	}

	log.Log(ctx).Info("create default group")
	if err := db.Model(&model.Group{}).
		Where("gid = ?", defaultGroup.GID).
		First(&model.Group{}).Error; err == nil {
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Log(ctx).WithError(err).Error("check default group failed")
		return err
	} else if err := db.Model(&model.Group{}).Create(&defaultGroup).Error; err != nil {
		log.Log(ctx).WithError(err).Error("create default group failed")
		return err
	}

	fmt.Printf(`
System Key: %s
Build Time: %s
Cron Key: %s

Admin Username: admin
Admin Password: 123456
`, initData["syskey"], initData["build"], initData["cronkey"])

	return nil
}
